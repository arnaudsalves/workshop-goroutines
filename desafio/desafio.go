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
	fmt.Println("Iniciando o processo")
	start := time.Now()

	ch := tituloDoSite(
		"https://www.google.com", "https://www.youtube.com", "https://www.mercadolivre.com.br/",
		"https://www.verdazzo.com.br", "https://www.casadocodigo.com.br/", "https://developers.mercadoenvios.com",
		"https://thedevconf.com/pt", "https://www.postgresql.org/", "https://www.mongodb.com/", "https://www.netflix.com/br/title/80057281",
		"https://jovemnerd.com.br", "https://www.a12.com",
	)
	fmt.Println("\nTitulos encontrados")
	fmt.Println("------------------------------------")
	for titulo := range ch {
		fmt.Println(titulo)
	}
	fmt.Println("------------------------------------")

	end := float64(time.Since(start) / time.Millisecond)

	fmt.Printf("Processo executado em %vms", end)
}

func tituloDoSite(urls ...string) <- chan string {
	var wg sync.WaitGroup
	canal := make(chan string, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Printf("Buscando titulo do site %s\n", url)
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			canal <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	wg.Wait()
	close(canal)
	return canal
}