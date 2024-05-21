package LectorInstancias

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Distancia struct {
	PuntoInicial string
	PuntoFinal   string
	Distancia    float64
}

type TipoResultado struct {
	Ruta      []Distancia
	Distancia float64
}
type Punto struct {
	Nombre    string
	Posicionx float64
	Posiciony float64
	Pasado    bool
}

func Resultado(Ruta []Distancia, Distancia float64) *TipoResultado {
	rsult := &TipoResultado{
		Ruta:      Ruta,
		Distancia: Distancia,
	}

	return rsult

}

func CrearPuntos(nombre string, X, Y float64) *Punto {
	Resultado := &Punto{
		Nombre:    nombre,
		Posicionx: X,
		Posiciony: Y,
		Pasado:    false,
	}

	return Resultado

}

func LecturaPuntos(NombreArchivo string) []Punto {

	var ColeccionNodos []Punto

	file, err := os.Open(NombreArchivo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linea := scanner.Text()
		campos := strings.Split(linea, " ")
		Z := campos[0]
		X, _ := strconv.ParseFloat(campos[1], 32)
		Y, _ := strconv.ParseFloat(campos[2], 32)

		ColeccionNodos = append(ColeccionNodos, *CrearPuntos(Z, X, Y))

	}
	return ColeccionNodos
}
