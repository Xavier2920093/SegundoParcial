# SegundoParcial

# SegundoParcial

Para ejecutar el taller propuesto para el parcial, es decir el vecino mas cercano e insercion mas cercana
realizando busqueda de vencindarios a ambas respuestas basta con llamar la funcion de esta manera

TSP.Tspwithchannels(file)

Por file, nos referimos al archivo TSP con las ubicaciones de los puntos un ejemplo podria ser

TSP.Tspwithchannels("dj38.tsp")

en caso de querer probar cada una de las funciones por separado como el vecino mas cercano, insercion mas
cercana o busqueda de vecindarios seria de la siguiente manera

// Leyendo el archivo
	go func() {
		defer close(PuntoCanal)
		nodos := LectorInstancias.LecturaPuntos(file)
		PuntoCanal <- nodos
	}()

	IndiceNodos := <-PuntoCanal

// funcion del vecino mas cercano
 rutaVecinoMasCercano, distanciaTotalVecinoMasCercano := VecinoMasCercano(IndiceNodos)

 // funcion de insercion mas cercana
 rutaInsercionMasCercana, distanciaTotalInsercionMasCercana := InsercionMasCercana(IndiceNodos)

 // funcion de vecindario
 rutaVecindarioVecino, distanciaTotalVecindarioVecino := BusquedaVecindario(rutaVecinoMasCercano)
