package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devckrishna/mongoapi/router"
)

func main() {
	fmt.Println("Mongo API----")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Listening to 3000")

}
