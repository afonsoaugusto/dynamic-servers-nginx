package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

var serverInformation = getNetworkInformation()

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func getServerInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p = map[string]string{"Server Information": "OK",
		"Version": "1.0.0",
		"IP":      r.Host,
		"Port":    r.URL.Port(),
		"Addr":    serverInformation[1].String(),
	}
	log.Println(p)
	json.NewEncoder(w).Encode(p)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/info", getServerInformation)
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)

}

func getNetworkInformation() []net.Addr {
	addrs, _ := net.InterfaceAddrs()
	fmt.Printf("%v\n", addrs)
	for _, addr := range addrs {
		fmt.Println("IPv4: ", addr)
	}
	return addrs
}
