package robtest

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Launch() {
	tlsCert, _ := tls.LoadX509KeyPair("../../assets/tls/test.crt", "../../assets/tls/test.key")

	router := mux.NewRouter()

	server := &http.Server{
		Addr:    "0.0.0.0:9000",
		Handler: router,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
		},
	}
	log.Println("Launching web server: https://0.0.0.0:9000")
	log.Fatal(server.ListenAndServeTLS("", ""))
}
