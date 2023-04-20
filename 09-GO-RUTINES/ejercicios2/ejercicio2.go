package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", slowHandler)
	fmt.Println("Servidor Arriba")
	http.ListenAndServe(":8080", nil)
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	starTime := time.Now()
	rand.Seed(time.Now().UnixNano())
	reandomNu := rand.Intn(10) + 1

	time.Sleep(time.Duration(reandomNu) * time.Second)

	fmt.Fprintf(w, "Peticion procesada en %v segundos\n", time.Since(starTime).Seconds())
}


func main() {
	canal := make(chan int)

	go contar(canal)

	for valor := range canal {
		fmt.Println(valor)
	}
}

func contar(canal chan int) {

	for i := 1; i < 10; i++ {
		canal <- i
	}
	close(canal)
}
*/

func main() {

	data := make([]int, 100000)

	for i := range data {
		data[i] = rand.Intn(10)
	}

	blockTamano := 100000

	var wg sync.WaitGroup

	resultado := make(chan int)

	for i := 0; i < len(data); i += blockTamano {
		wg.Add(1)
		procesarInfo(data[i:i+blockTamano], &wg, resultado)
	}

	go func() {
		wg.Wait()
		close(resultado)
	}()

	sumaTotal := 0

	for sum := range resultado {
		sumaTotal = +sum
	}

	fmt.Println(sumaTotal)

}

func procesarInfo(block []int, wg *sync.WaitGroup, resultado chan<- int) {
	wg.Done()

	time.Sleep(time.Duration(rand.Intn(1000) * int(time.Hour.Milliseconds())))
	sum := 0
	for _, value := range block {
		sum = +value
	}
	resultado <- sum
}
