package main

import (
	"fmt"

	"github.com/Xavier2920093/SegundoParcial/LectorInstancias"
	"github.com/Xavier2920093/SegundoParcial/TSP"
)

func main() {
	PuntoCanal := make(chan []LectorInstancias.Punto, 1)
	CanalVecino := make(chan []LectorInstancias.Punto, 1)
	CanalInsercion := make(chan LectorInstancias.TipoResultado, 1)
	CanalVecindario := make(chan LectorInstancias.TipoResultado, 1)

	go func() {
		defer close(PuntoCanal)
		nodos := LectorInstancias.LecturaPuntos("dj38.tsp")
		PuntoCanal <- nodos
	}()

	IndiceNodos := <-PuntoCanal

	go func() {
		//defer wg.Done()
		rutaVecinoMasCercano, distanciaTotalVecinoMasCercano := TSP.VecinoMasCercano(IndiceNodos)
		fmt.Println("\nRuta  vecino mas cercano:\n", rutaVecinoMasCercano)
		fmt.Println("\nDistancia total  vecino mas cercano:\n", distanciaTotalVecinoMasCercano)
		//Resve := LectorInstancias.Resultado(rutaVecinoMasCercano, distanciaTotalVecinoMasCercano)
		CanalVecino <- rutaVecinoMasCercano
		close(CanalVecino)
	}()

	go func() {
		//defer wg.Done()
		rutaInsercionMasCercana, distanciaTotalInsercionMasCercana := TSP.InsercionMasCercana(IndiceNodos)
		fmt.Println("\nRuta  Insercion cercana:\n", rutaInsercionMasCercana)
		fmt.Println("\nDistancia   insercion mas cercana:\n", distanciaTotalInsercionMasCercana)
		Resin := LectorInstancias.Resultado(rutaInsercionMasCercana, distanciaTotalInsercionMasCercana)
		CanalInsercion <- *Resin
		close(CanalInsercion)
	}()

	go func() {
		//defer wg.Done()
		rutaVMC, _ := TSP.VecinoMasCercano(IndiceNodos)
		rutaVecindario, distanciaTotalVecindario := TSP.BusquedaVecindario(rutaVMC)
		fmt.Println("\nRuta  Búsqueda de Vecindario:\n", rutaVecindario)
		fmt.Println("\nDistancia total  Búsqueda de Vecindario:\n", distanciaTotalVecindario)
		ResVecindario := LectorInstancias.Resultado(rutaVecindario, distanciaTotalVecindario)
		CanalVecindario <- *ResVecindario
		close(CanalVecindario)
	}()

	fmt.Println(<-CanalVecino)
	fmt.Print("-----------------------------------------------------------------------------------------------\n")
	fmt.Println(<-CanalInsercion)
}
