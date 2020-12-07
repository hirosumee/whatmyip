package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

func main() {
	log.Println("http server is starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IPAddress := r.Header.Get("X-Real-Ip")
		if IPAddress == "" {
			IPAddress = r.Header.Get("X-Forwarded-For")
		}
		if IPAddress == "" {
			IPAddress, _, _ = net.SplitHostPort(r.RemoteAddr)
		}
		w.Header().Add("Content-Type", "application/json")
		bs, _ := json.Marshal(map[string]interface{}{"ip": IPAddress})
		_, _ = w.Write(bs)
	}))
	if err != nil {
		log.Fatal(err)
	}
}
