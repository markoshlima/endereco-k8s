package main

import (
	s "endereco/src/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/enderecos", s.GetEnderecos)
	router.HandleFunc("/enderecos/{id}", s.GetEndereco)
	router.HandleFunc("/enderecos/cliente/{id}", s.GetEnderecoByCliente)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Rest API Endereco v1.0s")
	s.LoadEnderecos()
	handleRequests()
}
