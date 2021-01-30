package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/main/config"
	"todo/main/controller"
)


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	controller.HandleToDo()
	config.MyRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", config.MyRouter))
}

func main() {
	config.InitDb()
	config.InitRouter()
	handleRequests()
}