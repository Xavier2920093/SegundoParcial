package TSP

import (
	"fmt"

	"github.com/Xavier2920093/SegundoParcial/LectorInstancias"
)

func Tspwithchannels(file string) {
	PuntoCanal := make(chan []LectorInstancias.Punto, 1)
	CanalVecino := make(chan []LectorInstancias.Punto, 1)
	CanalInsercion := make(chan []LectorInstancias.Punto, 1)
	CanalVecindarioVecino := make(chan LectorInstancias.TipoResultado, 1)
	CanalVecindarioInsercion := make(chan LectorInstancias.TipoResultado, 1)

	go func() {
		defer close(PuntoCanal)
		nodos := LectorInstancias.LecturaPuntos(file)
		PuntoCanal <- nodos
	}()

	IndiceNodos := <-PuntoCanal

	go func() {
		rutaVecinoMasCercano, distanciaTotalVecinoMasCercano := VecinoMasCercano(IndiceNodos)
		fmt.Println("\nRuta Vecino más cercano:\n", rutaVecinoMasCercano)
		fmt.Println("\nDistancia total Vecino más cercano:\n", distanciaTotalVecinoMasCercano)
		CanalVecino <- rutaVecinoMasCercano
		close(CanalVecino)
	}()

	go func() {
		rutaInsercionMasCercana, distanciaTotalInsercionMasCercana := InsercionMasCercana(IndiceNodos)
		fmt.Println("\nRuta Inserción más cercana:\n", rutaInsercionMasCercana)
		fmt.Println("\nDistancia total Inserción más cercana:\n", distanciaTotalInsercionMasCercana)
		CanalInsercion <- rutaInsercionMasCercana
		close(CanalInsercion)
	}()

	go func() {
		rutaVecinoMasCercano := <-CanalVecino
		rutaVecindarioVecino, distanciaTotalVecindarioVecino := BusquedaVecindario(rutaVecinoMasCercano)
		fmt.Println("\nRuta Búsqueda de Vecindario (a partir de Vecino más cercano):\n", rutaVecindarioVecino)
		fmt.Println("\nDistancia total Búsqueda de Vecindario (Vecino más cercano):\n", distanciaTotalVecindarioVecino)
		ResVecindarioVecino := LectorInstancias.Resultado(rutaVecindarioVecino, distanciaTotalVecindarioVecino)
		CanalVecindarioVecino <- *ResVecindarioVecino
		close(CanalVecindarioVecino)
	}()

	go func() {
		rutaInsercionMasCercana := <-CanalInsercion
		rutaVecindarioInsercion, distanciaTotalVecindarioInsercion := BusquedaVecindario(rutaInsercionMasCercana)
		fmt.Println("\nRuta Búsqueda de Vecindario (a partir de Inserción más cercana):\n", rutaVecindarioInsercion)
		fmt.Println("\nDistancia total Búsqueda de Vecindario (Inserción más cercana):\n", distanciaTotalVecindarioInsercion)
		ResVecindarioInsercion := LectorInstancias.Resultado(rutaVecindarioInsercion, distanciaTotalVecindarioInsercion)
		CanalVecindarioInsercion <- *ResVecindarioInsercion
		close(CanalVecindarioInsercion)
	}()

	// Imprimir resultados de vecino más cercano seguido de búsqueda de vecindario
	fmt.Println(<-CanalVecindarioVecino)
	fmt.Print("-----------------------------------------------------------------------------------------------\n")

	// Imprimir resultados de inserción más cercana seguida de búsqueda de vecindario
	fmt.Println(<-CanalVecindarioInsercion)
}
