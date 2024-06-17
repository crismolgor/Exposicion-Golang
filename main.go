package main

import (
	"fmt"
	"os"
)

var path = "GruposEstudiantes.txt"

func main() {
	crearArchivo()
	escribeArchivo()
}

func crearArchivo() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("Archivo creado exitosamente", path)
}

func escribeArchivo() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("GRUPOS PRESENTACION DE LENGUAJES \n")
	if existeError(err) {
		return
	}
	_, err = file.WriteString("GRUPO Golang: \n + Veronica \n + David \n + Cristian \n + Luna Gil")
	if existeError(err) {
		return
	}

	err = file.Sync()
	if existeError(err) {
		return
	}

	fmt.Println("Archivo actualizado existosamente.")
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
