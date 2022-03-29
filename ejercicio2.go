package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Categories []Category
type Category struct {
	ID   string `json:"id"`
	Name string `json: "name"`
}

func main() {
	GetCategories("MLA")
	if err != nil {

	}
	fmt.Println("Las categorias son")
}

func GetCategories(siteID string) (Categories, error) {

	response, err := http.Get              //completar
	bytes := iotil.ReadAll(response.Bytes) //completar
	var Categories Categories
	json.Unmarshal(bytes, &cats)

	return cats, nil

}
