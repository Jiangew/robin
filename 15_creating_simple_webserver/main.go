package main

import(
	"fmt"
    "log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	http.HandleFunc("/", echoString)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/increment", incrementCounter)

	// log.Fatal(http.ListenAndServe(":8081", nil))

	// Serving Content over HTTPS
	log.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil))
}