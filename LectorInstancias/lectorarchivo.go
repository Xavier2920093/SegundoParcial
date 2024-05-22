package LectorInstancias

import (
	"bufio"
	"log"
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
		log.Fatal(err)
	}
	defer file.Close()

	encontrado := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()

		if encontrado {
			campos := strings.Fields(linea)
			if len(campos) != 3 {
				continue
			}
			Z := campos[0]
			X, err := strconv.ParseFloat(campos[1], 64)
			if err != nil {
				log.Fatal(err)
			}
			Y, err := strconv.ParseFloat(campos[2], 64)
			if err != nil {
				log.Fatal(err)
			}
			ColeccionNodos = append(ColeccionNodos, Punto{Nombre: Z, Posicionx: X, Posiciony: Y})
		}

		if strings.TrimSpace(linea) == "NODE_COORD_SECTION" {
			encontrado = true
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ColeccionNodos
}
