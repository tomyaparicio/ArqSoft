package main

import("fmt"
        "errors"
	     "net/http"
		 "ioutil"
		 "encoding/json")

		    http://api.mercadolibre.com/sites/MLA/search?q=Motorola
			type Categories struct {

			}

func main(){
	 GetCategories("MLA")

}
			

func GetCategories(siteID string) (Categories, error){

 response := http.Get//completar
bytes := iotil.ReadAll(response.Bytes)//completar
var Categories Categories
json.Unmarshal(bytes, &cats)

}		 