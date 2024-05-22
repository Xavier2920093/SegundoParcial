package main

import (
	"fmt"
	"sync"

	"github.com/Xavier2920093/SegundoParcial/LectorInstancias"
	"github.com/Xavier2920093/SegundoParcial/TSP"
)

func main() {
	PuntoCanal := make(chan []LectorInstancias.Punto, 1)
	CanalVecino := make(chan LectorInstancias.TipoResultado, 1)
	CanalInsercion := make(chan LectorInstancias.TipoResultado, 1)
	CanalVecindario := make(chan LectorInstancias.TipoResultado, 1)

	var wg sync.WaitGroup

	go func() {
		defer close(PuntoCanal)
		nodos := LectorInstancias.LecturaPuntos("dj38.tsp")
		PuntoCanal <- nodos
	}()

	IndiceNodos := <-PuntoCanal

	wg.Add(1)
	go func() {
		defer wg.Done()
		rutaVecindario, distanciaTotalVecindario := TSP.BusquedaVecindario(IndiceNodos)
		fmt.Println("Ruta  Búsqueda de Vecindario:", rutaVecindario)
		fmt.Println("Distancia total  Búsqueda de Vecindario:", distanciaTotalVecindario)
		ResVecindario := LectorInstancias.Resultado(rutaVecindario, distanciaTotalVecindario)
		CanalVecindario <- *ResVecindario
		close(CanalVecindario)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rutaInsercionMasCercana, distanciaTotalInsercionMasCercana := TSP.InsercionMasCercana(IndiceNodos)
		fmt.Println("Ruta  Inserciion cercana:", rutaInsercionMasCercana)
		fmt.Println("Distancia   insercion mas cercana:", distanciaTotalInsercionMasCercana)
		Resin := LectorInstancias.Resultado(rutaInsercionMasCercana, distanciaTotalInsercionMasCercana)
		CanalInsercion <- *Resin
		close(CanalInsercion)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rutaVecinoMasCercano, distanciaTotalVecinoMasCercano := TSP.VecinoMasCercano(IndiceNodos)
		fmt.Println("Ruta  vecino mas cercano:", rutaVecinoMasCercano)
		fmt.Println("Distancia total  vecino mas cercano:", distanciaTotalVecinoMasCercano)
		Resve := LectorInstancias.Resultado(rutaVecinoMasCercano, distanciaTotalVecinoMasCercano)
		CanalVecino <- *Resve
		close(CanalVecino)
	}()

	fmt.Println(<-CanalVecino)
	fmt.Print("")
	fmt.Println(<-CanalVecindario)
	fmt.Print("")
	fmt.Println(<-CanalInsercion)

}
