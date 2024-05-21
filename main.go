package main

import (
	"fmt"
	"sync"

	"github.com/Xavier2920093/SegundoParcial/LectorInstancias"
	"github.com/Xavier2920093/SegundoParcial/TSP"
)

func main() {
	PuntoCanal := make(chan []LectorInstancias.Nodo, 1)
	CanalVecino := make(chan LectorInstancias.Resultado, 1)
	CanalInsercion := make(chan LectorInstancias.Resultado, 1)
	CanalVecindario := make(chan LectorInstancias.Resultado, 1)

	var wg sync.WaitGroup

	go func() {
		defer close(PuntoCanal)
		nodos := LectorInstancias.LeerNodos("dj38.tsp")
		PuntoCanal <- nodos
	}()

	IndiceNodos := <-PuntoCanal

	wg.Add(1)
	go func() {
		defer wg.Done()
		distanciasPrim, distanciasSec := TSP.Calculo(IndiceNodos)
		fmt.Println("Distancias calculadas por búsqueda de vecindario:", append(distanciasPrim, distanciasSec...))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rutaVecinoMasCercano, distanciaTotalVecinoMasCercano := TSP.VecinoMasCercano(IndiceNodos)
		fmt.Println("Ruta utilizando Vecino Más Cercano:", rutaVecinoMasCercano)
		fmt.Println("Distancia total utilizando Vecino Más Cercano:", distanciaTotalVecinoMasCercano)
		Resve := LectorInstancias.CrearResultado(rutaVecinoMasCercano, distanciaTotalVecinoMasCercano)
		CanalVecino <- *Resve
		close(CanalVecino)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rutaInsercionMasCercana, distanciaTotalInsercionMasCercana := TSP.InsercionMasCercana(IndiceNodos)
		fmt.Println("Ruta utilizada Inserción Más Cercana:", rutaInsercionMasCercana)
		fmt.Println("Distancia total utilizando Inserción Más Cercana:", distanciaTotalInsercionMasCercana)
		Resin := LectorInstancias.CrearResultado(rutaInsercionMasCercana, distanciaTotalInsercionMasCercana)
		CanalInsercion <- *Resin
		close(CanalInsercion)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rutaVecindario, distanciaTotalVecindario := TSP.BusquedaVecindario(IndiceNodos)
		fmt.Println("Ruta utilizaDA Búsqueda de Vecindario:", rutaVecindario)
		fmt.Println("Distancia total utilizando Búsqueda de Vecindario:", distanciaTotalVecindario)
		ResVecindario := LectorInstancias.CrearResultado(rutaVecindario, distanciaTotalVecindario)
		CanalVecindario <- *ResVecindario
		close(CanalVecindario)
	}()

	wg.Wait()

	fmt.Println(<-CanalVecino)
	fmt.Println(<-CanalInsercion)
	fmt.Println(<-CanalVecindario)
}
