package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func main() {
	fmt.Println("Iniciando o processo")
	start := time.Now()

	titulos := tituloDoSite(
		"https://www.google.com", "https://www.youtube.com", "https://www.mercadolivre.com.br/",
		"https://www.verdazzo.com.br", "https://www.casadocodigo.com.br/", "https://developers.mercadoenvios.com",
		"https://thedevconf.com/pt", "https://www.postgresql.org/", "https://www.mongodb.com/", "https://www.netflix.com/br/title/80057281",
		"https://jovemnerd.com.br", "https://www.a12.com",
	)

	for i, titulo := range titulos {
		fmt.Printf("%d - %s\n", i, titulo)
	}

	end := float64(time.Since(start) / time.Millisecond)

	fmt.Printf("Processo executado em %vms", end)
}

func tituloDoSite(urls ...string) []string {
	result := make([]string, 0, len(urls))
	for _, url := range urls {
		resp, _ := http.Get(url)
		html, _ := ioutil.ReadAll(resp.Body)

		r, _ := regexp.Compile("<title>(.*?)<\\/title>")
		result = append(result, r.FindStringSubmatch(string(html))[1])
	}
	return result
}
