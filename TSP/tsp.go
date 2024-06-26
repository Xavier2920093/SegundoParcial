package TSP

import (
	"math"

	lectorinstancias "github.com/Xavier2920093/SegundoParcial/LectorInstancias"
	//aca iria el uso del modulo lectorinstancias
)

func DistanciaEuclidiana(nodo0, nodo1 lectorinstancias.Punto) float64 {
	DX := nodo0.Posicionx - nodo1.Posicionx
	DY := nodo0.Posiciony - nodo1.Posiciony
	return math.Sqrt(DX*DX + DY*DY)
}
func VecinoMasCercano(nodos []lectorinstancias.Punto) ([]lectorinstancias.Punto, float64) {
	if len(nodos) == 0 {
		return nil, 0
	}

	visitados := make(map[string]bool)
	var ruta []lectorinstancias.Punto
	totalDistancia := 0.0

	nodoActual := nodos[0]
	visitados[nodoActual.Nombre] = true

	for len(visitados) < len(nodos) {
		nodoMasCercano := lectorinstancias.Punto{}
		minDistancia := math.MaxFloat64

		for _, nodo := range nodos {
			if !visitados[nodo.Nombre] {
				dist := DistanciaEuclidiana(nodoActual, nodo)
				if dist < minDistancia {
					nodoMasCercano = nodo
					minDistancia = dist
				}
			}
		}

		if (nodoMasCercano != lectorinstancias.Punto{}) {
			ruta = append(ruta, nodoActual)
			totalDistancia += minDistancia
			nodoActual = nodoMasCercano
			visitados[nodoActual.Nombre] = true
		}
	}

	// Agregar el nodo inicial al final para completar el ciclo
	ruta = append(ruta, nodos[0])
	totalDistancia += DistanciaEuclidiana(nodoActual, nodos[0])

	return ruta, totalDistancia
}

func InsercionMasCercana(nodos []lectorinstancias.Punto) ([]lectorinstancias.Punto, float64) {
	if len(nodos) == 0 {
		return nil, 0
	}

	// Empezamos con un ciclo que incluye el primer nodo
	var ruta []lectorinstancias.Punto
	totalDistancia := 0.0

	visitados := make(map[string]bool)
	visitados[nodos[0].Nombre] = true

	if len(nodos) > 1 {
		visitados[nodos[1].Nombre] = true
		ruta = append(ruta, nodos[0])
		ruta = append(ruta, nodos[1])
		totalDistancia = DistanciaEuclidiana(nodos[0], nodos[1])
	}

	// Inserción más cercana
	for len(visitados) < len(nodos) {
		nodoMasCercano := lectorinstancias.Punto{}
		minIncremento := math.MaxFloat64
		posicion := 0

		for _, nodo := range nodos {
			if !visitados[nodo.Nombre] {
				for i := 0; i < len(ruta); i++ {
					nodoI := encontrarNodo(ruta[i].Nombre, nodos)
					nodoF := encontrarNodo(ruta[(i+1)%len(ruta)].Nombre, nodos)
					distAnterior := DistanciaEuclidiana(nodoI, ruta[i])
					distSiguiente := DistanciaEuclidiana(ruta[(i+1)%len(ruta)], nodoF)
					incremento := DistanciaEuclidiana(nodoI, nodo) +
						DistanciaEuclidiana(nodo, nodoF) - distAnterior - distSiguiente
					if incremento < minIncremento {
						nodoMasCercano = nodo
						minIncremento = incremento
						posicion = i
					}
				}
			}
		}

		// Insertar el nodo en la posición encontrada
		if (nodoMasCercano != lectorinstancias.Punto{}) {
			ruta = append(ruta[:posicion+1], append([]lectorinstancias.Punto{nodoMasCercano}, ruta[posicion+1:]...)...)
			totalDistancia += minIncremento
			visitados[nodoMasCercano.Nombre] = true
		}
	}

	return ruta, totalDistancia

}

func encontrarNodo(nombre string, nodos []lectorinstancias.Punto) lectorinstancias.Punto {
	for _, nodo := range nodos {
		if nodo.Nombre == nombre {
			return nodo
		}
	}
	return lectorinstancias.Punto{} // Retornar un punto vacío si no se encuentra el nodo
}

func BusquedaVecindario(nodos []lectorinstancias.Punto) ([]lectorinstancias.Distancia, float64) {
	ruta := make([]lectorinstancias.Punto, len(nodos))
	copy(ruta, nodos)
	distanciaTotal := calcularDistanciaTotal(ruta)
	intercambio := true
	for intercambio {
		intercambio = false
		for i := 0; i < len(ruta)-1; i++ {
			for j := i + 1; j < len(ruta); j++ {
				ruta[i], ruta[j] = ruta[j], ruta[i] // Intercambio de ruta

				nuevaDistancia := calcularDistanciaTotal(ruta)

				if nuevaDistancia < distanciaTotal {
					distanciaTotal = nuevaDistancia
					intercambio = true
				} else {
					ruta[i], ruta[j] = ruta[j], ruta[i]
				}
			}
		}
	}
	var distancias []lectorinstancias.Distancia
	for i := 0; i < len(ruta)-1; i++ {
		distancia := lectorinstancias.Distancia{
			PuntoInicial: ruta[i].Nombre,
			PuntoFinal:   ruta[i+1].Nombre,
			Distancia:    DistanciaEuclidiana(ruta[i], ruta[i+1]),
		}
		distancias = append(distancias, distancia)
	}
	distanciaFinal := lectorinstancias.Distancia{
		PuntoInicial: ruta[len(ruta)-1].Nombre,
		PuntoFinal:   ruta[0].Nombre,
		Distancia:    DistanciaEuclidiana(ruta[len(ruta)-1], ruta[0]),
	}
	distancias = append(distancias, distanciaFinal)

	return distancias, distanciaTotal
}
func calcularDistanciaTotal(ruta []lectorinstancias.Punto) float64 {
	distanciaTotal := 0.0
	for i := 0; i < len(ruta)-1; i++ {
		distanciaTotal += DistanciaEuclidiana(ruta[i], ruta[i+1])
	}

	distanciaTotal += DistanciaEuclidiana(ruta[len(ruta)-1], ruta[0])
	return distanciaTotal
}
