package main

import (
	"net/http"
	"server/utils"
	"log"
)

func main() {
	mux := http.NewServeMux()



	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	log.Println("Servidor escuchando en puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))

	
}
