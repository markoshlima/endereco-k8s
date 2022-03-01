package service

import (
	"encoding/json"
	e "endereco/src/model"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

func LoadEnderecos() {
	Enderecos = []e.Endereco{
		e.Endereco{EnderecoID: 1, Rua: "Francisco Morato", Numero: 1254, Bairro: "Vila Prudente", Cidade: "São Paulo", Estado: "SP", Pais: "BR", ClienteID: 1},
		e.Endereco{EnderecoID: 2, Rua: "Av. Brasil", Numero: 455, Bairro: "Centro", Cidade: "Rio de Janeiro", Estado: "RJ", Pais: "BR", ClienteID: 1},
		e.Endereco{EnderecoID: 3, Rua: "Chico Julio", Numero: 54, Bairro: "Centro", Cidade: "Manaus", Estado: "AM", Pais: "BR", ClienteID: 1},
		e.Endereco{EnderecoID: 4, Rua: "Otávio Martins", Numero: 987, Bairro: "Vila Marina", Cidade: "Rio Grande do Sul", Estado: "RS", Pais: "BR", ClienteID: 2},
		e.Endereco{EnderecoID: 5, Rua: "Antonio Bandeiras", Numero: 1257, Bairro: "Jardim Brasil", Cidade: "Canela", Estado: "RS", Pais: "BR", ClienteID: 3},
	}
}

var Enderecos []e.Endereco

// obtem enderecos pelo cliente
func GetEnderecoByCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: getEnderecoByCliente")
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	//declare array
	var response []e.Endereco
	for _, endereco := range Enderecos {
		if endereco.ClienteID == key {
			response = append(response, endereco)
		}
	}
	json.NewEncoder(w).Encode(response)
}

// obtem todos os enderecos
func GetEnderecos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: getEnderecos")
	json.NewEncoder(w).Encode(Enderecos)
}

//obtem o endereco pelo ID
func GetEndereco(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: getEndereco")
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for _, endereco := range Enderecos {
		if endereco.EnderecoID == key {
			json.NewEncoder(w).Encode(endereco)
		}
	}
}
