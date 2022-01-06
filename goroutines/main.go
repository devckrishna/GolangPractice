package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var signals = []string{"test"}
var mut sync.Mutex

func main() {
	website := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://github.com",
		"https://google.com",
	}

	for _, web := range website {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oops")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%s the status code is %d \n", endpoint, res.StatusCode)
	}
}
