package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/claytonssmint/load-tester/internal/tester"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 100, "Número total de requisições")
	concurrency := flag.Int("concurrency", 10, "Número de requisições concorrentes")
	flag.Parse()

	if *url == "" {
		log.Fatal("A URL é obrigatória. Use a flag -url para informar a URL do serviço a ser testado")
	}

	startTime := time.Now()
	results := tester.RunLoadTest(*url, *requests, *concurrency)
	duration := time.Since(startTime)

	fmt.Println("Relatório de Teste de Carga:")
	fmt.Printf("URL: %s\n", *url)
	fmt.Printf("Requests Totais: %d\n", *requests)
	fmt.Printf("Concorrência: %d\n", *concurrency)
	fmt.Printf("Tempo Total: %v\n", duration)
	fmt.Printf("Status 200: %d\n", results.Status200)
	fmt.Printf("Outros Status:\n")
	for status, count := range results.OtherStatus {
		fmt.Printf("  %d: %d\n", status, count)
	}

}
