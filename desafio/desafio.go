package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
	"time"
)

func main() {
	canal := make(chan string, 20)

	fmt.Println("Iniciando o processo")
	start := time.Now()

	tituloDoSite(canal,
		"https://www.google.com", "https://www.youtube.com", "https://www.mercadolivre.com.br/",
		"https://www.verdazzo.com.br", "https://www.casadocodigo.com.br/", "https://developers.mercadoenvios.com",
		"https://thedevconf.com/pt", "https://www.postgresql.org/", "https://www.mongodb.com/", "https://www.netflix.com/br/title/80057281",
		"https://jovemnerd.com.br", "https://www.a12.com",
	)

	for titulo := range canal {
		fmt.Println(titulo)
	}

	end := float64(time.Since(start) / time.Millisecond)

	fmt.Printf("Processo executado em %vms", end)
}

func tituloDoSite(canal chan string, urls ...string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			canal <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	wg.Wait()
	close(canal)
}