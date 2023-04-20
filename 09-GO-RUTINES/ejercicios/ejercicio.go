package main

import (
	//"fmt"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Persona struct {
	Id        int
	Nombre    string
	Direccion string
}

var personas = []Persona{
	{Id: 1, Nombre: "Gerardo", Direccion: "CDMX"},
}

var (
	mu sync.Mutex
)

func main() {

	//http.HandleFunc("/data", dataHandler)
	//log.Fatal(http.ListenAndServe("8080", nil))

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		http.HandleFunc("/data", getData)
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	/*go func() {
		defer wg.Done()
		http.HandleFunc("/data", postData)
		log.Fatal(http.ListenAndServe(":8001", nil))
	}()

	go func() {
		defer wg.Done()
		http.HandleFunc("/data", deleteData)
		log.Fatal(http.ListenAndServe(":8002", nil))
	}()*/

	wg.Wait()
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
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
	defer mu.Unlock()
	var data Persona
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	personas = append(personas, data)
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(data)
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	defer mu.Unlock()
	var data Persona
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()

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
