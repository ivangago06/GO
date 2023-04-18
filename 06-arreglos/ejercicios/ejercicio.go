package main

import (
	//"fmt"
	"encoding/json"
	"log"
	"net/http"
)

type Persona struct {
	Id        int
	Nombre    string
	Direccion string
}

var personas = []Persona{
	{Id: 1, Nombre: "Gerardo", Direccion: "CDMX"},
}

func main() {

	http.HandleFunc("/data", dataHan)
	log.Fatal(http.ListenAndServe("8080", nil))

	/*var personas []Persona

	var nombre, direccion string

	fmt.Println("Nombre del usuario")
	fmt.Scanln(&nombre)

	fmt.Println("Dirección")
	fmt.Scanln(&direccion)

	personas = append(personas, Persona{nombre: nombre, direccion: direccion})

	for i, pepersona := range personas {
		fmt.Println(i+i, "-", "Nombre: ", pepersona.nombre, "Dirección: ", pepersona.direccion)
	}*/

}

func dataHan(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getData(w, r)
	case "POST":
		postData(w, r)
	case "DELETE":
		deleteData(w, r)
	default:
		http.Error(w, "No existe este método", http.StatusMethodNotAllowed)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(personas)
}

func postData(w http.ResponseWriter, r *http.Request) {
	var data Persona
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	personas = append(personas, data)
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(data)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	var data Persona
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, d := range personas {
		if d.Id == data.Id {
			personas = append(personas[:i], personas[i+1:]...)
			w.Header().Set("Content-Type", "appication/json")
			json.NewEncoder(w).Encode(data)
			return
		}
	}

	http.Error(w, "Indice no encontrado", http.StatusNotFound)

}
