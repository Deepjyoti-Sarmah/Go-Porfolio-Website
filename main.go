package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DeepjyotiSarmah/portfolio/routes"
)

func Execute() error {
	routes.Route()
	fmt.Println("Strating the server at: 4000")
	err := http.ListenAndServe(":4000", nil)
	return err
}

func main() {
	if err := Execute(); err != nil {
		log.Fatal(err)
	}
}
