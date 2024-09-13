package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/yagizklc/from-scratch-server/app/handlers"
	"github.com/yagizklc/from-scratch-server/app/pkg"
)

const (
	DEFAULT_PORT = "8080"
	HOST         = "localhost"
)

func main() {
	PORT := flag.String("port", DEFAULT_PORT, "help message for flag n")
	HOST := flag.String("host", HOST, "help message for flag n")
	flag.Parse()

	config, err := pkg.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)
	mux.Handle("/pokedex", handlers.NewPokemonHandler(config))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *HOST, *PORT), mux))
}

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
}
