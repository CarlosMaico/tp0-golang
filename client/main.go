package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	// validar que la config este cargada correctamente

	if globals.ClientConfig == nil {
		log.Fatal("No se pudo cargar la configuración")
	}
	// loggeamos el valor de la config
	log.Printf("Configuración cargada: IP=%s, Puerto=%d, Mensaje=%s",
		globals.ClientConfig.Ip,
		globals.ClientConfig.Puerto,
		globals.ClientConfig.Mensaje,
	)
	//PAra leer de consola
	//utils.LeerConsola()

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él


	// enviar un mensaje al servidor con el valor de la config

	// leer de la consola el mensaje
	// utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
}
