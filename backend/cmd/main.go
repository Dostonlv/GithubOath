package main

import (
	"fmt"
	"log"
	"net/http"
	"oath/api"
)

func main() {
	//cfg := config.Load()
	//
	//h := handlers.NewHandler(cfg)

	api.SetUpAPI()
	fmt.Println("[ UP ON PORT 3000 ]")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

}
