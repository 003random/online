package main

import (
	"fmt"
	"os"
	"net/http"
	"bufio"
	"time"
	"crypto/tls"
)

func main() {
	var urls []string

	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] Usage: cat urls.txt | online", err)
		os.Exit(3)
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "[!] Usage: cat urls.txt | online", err)
			os.Exit(3)
		}
	}

	for _, url := range urls {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		timeout := time.Duration(1 * time.Second)
		client := http.Client{
			Timeout: timeout,
		}
		_, err := client.Get(url)
		if err == nil {
    			fmt.Println(url)
		}
	}
}
