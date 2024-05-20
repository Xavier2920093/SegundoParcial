package main

import (
	"fmt"
	"sync"

	"github.com/Xavier2920093/SegundoParcial/LectorInstancias"
	"github.com/Xavier2920093/SegundoParcial/TSP"
)

func main() {
	CanalNodo := make(chan []LectorInstancias.Nodo, 1)
	CanalVecino := make(chan LectorInstancias.Resultado, 1)
	CanalInsercion := make(chan LectorInstancias.Resultado, 1)
	CanalVecindario := make(chan LectorInstancias.Resultado, 1)

	var wg sync.WaitGroup

	// Iniciar una goroutine para leer los nodos y enviarlos al canal
	go func() {
		defer close(CanalNodo)
		nodos := LectorInstancias.LeerNodos("dj38.tsp")
		CanalNodo <- nodos
	}()

	// Recibir los nodos del canal
	IndiceNodos := <-CanalNodo

	// Calcular la ruta óptima utilizando Vecino Más Cercano
	wg.Add(1)
	go func() {
		defer wg.Done()
		distanciasPrim, distanciasSec := TSP.Calculo(IndiceNodos)
		fmt.Println("Distancias calculadas por búsqueda de vecindario:", append(distanciasPrim, distanciasSec...))
	}()

	// Calcular la ruta óptima utilizando Vecino Más Cercano
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

	// Calcular la ruta óptima utilizando Inserción Más Cercana
	wg.Add(1)
	go func() {
		defer wg.Done()
		rutaInsercionMasCercana, distanciaTotalInsercionMasCercana := TSP.InsercionMasCercana(IndiceNodos)
		fmt.Println("Ruta utilizando Inserción Más Cercana:", rutaInsercionMasCercana)
		fmt.Println("Distancia total utilizando Inserción Más Cercana:", distanciaTotalInsercionMasCercana)
		Resin := LectorInstancias.CrearResultado(rutaInsercionMasCercana, distanciaTotalInsercionMasCercana)
		CanalInsercion <- *Resin
		close(CanalInsercion)
	}()

	// Implementar la búsqueda de vecindario
	wg.Add(1)
	go func() {
		defer wg.Done()
		rutaVecindario, distanciaTotalVecindario := TSP.BusquedaVecindario(IndiceNodos)
		fmt.Println("Ruta utilizando Búsqueda de Vecindario:", rutaVecindario)
		fmt.Println("Distancia total utilizando Búsqueda de Vecindario:", distanciaTotalVecindario)
		ResVecindario := LectorInstancias.CrearResultado(rutaVecindario, distanciaTotalVecindario)
		CanalVecindario <- *ResVecindario
		close(CanalVecindario)
	}()

	// Esperar a que todas las goroutines terminen
	wg.Wait()

	// Imprimir los resultados finales
	fmt.Println(<-CanalVecino)
	fmt.Println(<-CanalInsercion)
	fmt.Println(<-CanalVecindario)
}
