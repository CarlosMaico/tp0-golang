package utils

import (
	"bufio"
	"bytes"
	"client/globals"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func IniciarConfiguracion(filePath string) *globals.Config {
	var config *globals.Config
	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

func LeerConsola() {
	// Leer de la consola
	reader := bufio.NewReader(os.Stdin)
	for {
		log.Print("Ingrese un mensaje (enter para salir):")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error leyendo de consola: %s", err.Error())
			continue
		}

		if text == "\n" {
			log.Println("Entrada vacía detectada. Terminando programa.")
			break
		}

		log.Printf("Mensaje ingresado: %s", text)
	}
}

func GenerarYEnviarPaquete() {
	var lineas []string
	reader := bufio.NewReader(os.Stdin)

	log.Println("Ingrese líneas de texto para el paquete (ENTER vacío para finalizar):")

	for {
		log.Print("> ")
		texto, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error leyendo de consola: %s", err.Error())
			continue
		}

		// Eliminar salto de línea final
		texto = texto[:len(texto)-1]

		if texto == "" {
			break
		}

		lineas = append(lineas, texto)
	}

	if len(lineas) == 0 {
		log.Println("No se ingresaron líneas. No se envió ningún paquete.")
		return
	}

	paquete := Paquete{
		Valores: lineas,
	}

	log.Printf("Paquete a enviar: %+v", paquete)
	EnviarPaquete(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, paquete)
	//paquete := Paquete{}
	// Leemos y cargamos el paquete

	//log.Printf("paqute a enviar: %+v", paquete)
	// Enviamos el paqute
}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	mensaje := Mensaje{Mensaje: mensajeTxt}
	body, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("error codificando mensaje: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensaje a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	body, err := json.Marshal(paquete)
	if err != nil {
		log.Printf("error codificando mensajes: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensajes a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func ConfigurarLogger() {
	logFile, err := os.OpenFile("tp0.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //para setear lo que paso con fecha y hora
}
