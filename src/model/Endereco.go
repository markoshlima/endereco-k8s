package model

type Endereco struct {
	EnderecoID int    `json:"enderecoId"`
	Rua        string `json:"rua"`
	Numero     int    `json:"numero"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	Estado     string `json:"estado"`
	Pais       string `json:"pais"`
	ClienteID  int    `json:"clienteId"`
}
