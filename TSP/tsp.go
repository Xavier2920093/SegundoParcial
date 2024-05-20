package TSP

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	lectorinstancias "github.com/Xavier2920093/SegundoParcial/LectorInstancias"
	//aca iria el uso del modulo lectorinstancias
)

// Calcula la distancia euclidiana entre dos nodos
func DistanciaEuclidiana(nodo1, nodo2 lectorinstancias.Nodo) float64 {
	deltaX := nodo1.CoorX - nodo2.CoorX
	deltaY := nodo1.CoorY - nodo2.CoorY
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

// VecinoMásCercano calcula la ruta óptima utilizando el algoritmo del vecino más cercano
func VecinoMasCercano(nodos []lectorinstancias.Nodo) ([]lectorinstancias.Distancia, float64) {
	if len(nodos) == 0 {
		return nil, 0
	}

	visitados := make(map[string]bool)
	var ruta []lectorinstancias.Distancia
	totalDistancia := 0.0

	nodoActual := nodos[0]
	visitados[nodoActual.Nombre] = true

	for len(visitados) < len(nodos) {
		nodoMasCercano := lectorinstancias.Nodo{}
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

		if (nodoMasCercano != lectorinstancias.Nodo{}) {
			ruta = append(ruta, lectorinstancias.Distancia{
				NodoI:     nodoActual.Nombre,
				NodoFinal: nodoMasCercano.Nombre,
				Distancia: minDistancia,
			})
			totalDistancia += minDistancia
			nodoActual = nodoMasCercano
			visitados[nodoActual.Nombre] = true
		}
	}

	// Regresar al nodo inicial para completar el ciclo
	ruta = append(ruta, lectorinstancias.Distancia{
		NodoI:     nodoActual.Nombre,
		NodoFinal: nodos[0].Nombre,
		Distancia: DistanciaEuclidiana(nodoActual, nodos[0]),
	})
	totalDistancia += DistanciaEuclidiana(nodoActual, nodos[0])

	return ruta, totalDistancia
}

// InsercionMasCercana calcula la ruta óptima utilizando el algoritmo de inserción más cercana
func InsercionMasCercana(nodos []lectorinstancias.Nodo) ([]lectorinstancias.Distancia, float64) {
	if len(nodos) == 0 {
		return nil, 0
	}

	// Empezamos con un ciclo que incluye el primer nodo
	var ruta []lectorinstancias.Distancia
	totalDistancia := 0.0

	visitados := make(map[string]bool)
	visitados[nodos[0].Nombre] = true

	if len(nodos) > 1 {
		visitados[nodos[1].Nombre] = true
		ruta = append(ruta, lectorinstancias.Distancia{
			NodoI:     nodos[0].Nombre,
			NodoFinal: nodos[1].Nombre,
			Distancia: DistanciaEuclidiana(nodos[0], nodos[1]),
		})
		ruta = append(ruta, lectorinstancias.Distancia{
			NodoI:     nodos[1].Nombre,
			NodoFinal: nodos[0].Nombre,
			Distancia: DistanciaEuclidiana(nodos[1], nodos[0]),
		})
		totalDistancia = 2 * DistanciaEuclidiana(nodos[0], nodos[1])
	}

	// Inserción más cercana
	for len(visitados) < len(nodos) {
		nodoMasCercano := lectorinstancias.Nodo{}
		minIncremento := math.MaxFloat64
		posicion := 0

		// Encontrar el nodo no visitado más cercano y la mejor posición para insertarlo
		for _, nodo := range nodos {
			if !visitados[nodo.Nombre] {
				for i := 0; i < len(ruta); i++ {
					nodoI := encontrarNodo(ruta[i].NodoI, nodos)
					nodoF := encontrarNodo(ruta[i].NodoFinal, nodos)
					incremento := DistanciaEuclidiana(nodoI, nodo) + DistanciaEuclidiana(nodo, nodoF) - ruta[i].Distancia
					if incremento < minIncremento {
						nodoMasCercano = nodo
						minIncremento = incremento
						posicion = i
					}
				}
			}
		}

		// Insertar el nodo en la posición encontrada
		if (nodoMasCercano != lectorinstancias.Nodo{}) {
			nodoI := encontrarNodo(ruta[posicion].NodoI, nodos)
			nodoF := encontrarNodo(ruta[posicion].NodoFinal, nodos)
			ruta = append(ruta[:posicion+1], ruta[posicion:]...) // Hacer espacio para la nueva distancia
			ruta[posicion] = lectorinstancias.Distancia{
				NodoI:     nodoI.Nombre,
				NodoFinal: nodoMasCercano.Nombre,
				Distancia: DistanciaEuclidiana(nodoI, nodoMasCercano),
			}
			ruta[posicion+1] = lectorinstancias.Distancia{
				NodoI:     nodoMasCercano.Nombre,
				NodoFinal: nodoF.Nombre,
				Distancia: DistanciaEuclidiana(nodoMasCercano, nodoF),
			}
			totalDistancia += minIncremento
			visitados[nodoMasCercano.Nombre] = true
		}
	}

	return ruta, totalDistancia
}

// Encuentra un nodo por su nombre
func encontrarNodo(nombre string, nodos []lectorinstancias.Nodo) lectorinstancias.Nodo {
	for _, nodo := range nodos {
		if nodo.Nombre == nombre {
			return nodo
		}
	}
	return lectorinstancias.Nodo{}
}

// Calcula las distancias entre los nodos en la lista dada
func calcularDistancias(nodos []lectorinstancias.Nodo, distancias *[]lectorinstancias.Distancia, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(nodos)-1; i++ {
		for j := i + 1; j < len(nodos); j++ {
			distancia := DistanciaEuclidiana(nodos[i], nodos[j])
			*distancias = append(*distancias, lectorinstancias.Distancia{
				NodoI:     nodos[i].Nombre,
				NodoFinal: nodos[j].Nombre,
				Distancia: distancia,
			})
		}
	}
}

func Calculo(IndiceNodos []lectorinstancias.Nodo) ([]lectorinstancias.Distancia, []lectorinstancias.Distancia) {
	rand.Seed(time.Now().UnixNano())
	IndiceAleatorio := rand.Intn(len(IndiceNodos))
	prim := IndiceNodos[:IndiceAleatorio]
	Sec := IndiceNodos[IndiceAleatorio-1:]

	var distanciasPrim []lectorinstancias.Distancia
	var distanciasSec []lectorinstancias.Distancia
	var wg sync.WaitGroup

	wg.Add(2)
	fmt.Println(IndiceNodos[IndiceAleatorio])

	go calcularDistancias(prim, &distanciasPrim, &wg)
	go calcularDistancias(Sec, &distanciasSec, &wg)

	wg.Wait()

	return distanciasPrim, distanciasSec

}

func BusquedaVecindario(nodos []lectorinstancias.Nodo) ([]lectorinstancias.Distancia, float64) {
	// Inicializar la ruta como una permutación de los nodos
	ruta := make([]lectorinstancias.Nodo, len(nodos))
	copy(ruta, nodos)

	// Calcular la distancia total inicial
	distanciaTotal := calcularDistanciaTotal(ruta)

	// Variable para indicar si se realizó un intercambio en la iteración anterior
	intercambio := true

	// Realizar iteraciones hasta que no se realice ningún intercambio en una iteración completa
	for intercambio {
		intercambio = false
		for i := 0; i < len(ruta)-1; i++ {
			for j := i + 1; j < len(ruta); j++ {
				// Aplicar el intercambio
				ruta[i], ruta[j] = ruta[j], ruta[i]

				// Calcular la nueva distancia total
				nuevaDistancia := calcularDistanciaTotal(ruta)

				// Si la nueva distancia es menor, se acepta el intercambio
				if nuevaDistancia < distanciaTotal {
					distanciaTotal = nuevaDistancia
					intercambio = true
				} else {
					// Si no es menor, se deshace el intercambio
					ruta[i], ruta[j] = ruta[j], ruta[i]
				}
			}
		}
	}

	// Construir la lista de distancias basada en la ruta final
	var distancias []lectorinstancias.Distancia
	for i := 0; i < len(ruta)-1; i++ {
		distancia := lectorinstancias.Distancia{
			NodoI:     ruta[i].Nombre,
			NodoFinal: ruta[i+1].Nombre,
			Distancia: DistanciaEuclidiana(ruta[i], ruta[i+1]),
		}
		distancias = append(distancias, distancia)
	}

	// Agregar la distancia desde el último nodo hasta el primero para cerrar el ciclo
	distanciaFinal := lectorinstancias.Distancia{
		NodoI:     ruta[len(ruta)-1].Nombre,
		NodoFinal: ruta[0].Nombre,
		Distancia: DistanciaEuclidiana(ruta[len(ruta)-1], ruta[0]),
	}
	distancias = append(distancias, distanciaFinal)

	return distancias, distanciaTotal
}

func calcularDistanciaTotal(ruta []lectorinstancias.Nodo) float64 {
	distanciaTotal := 0.0
	for i := 0; i < len(ruta)-1; i++ {
		distanciaTotal += DistanciaEuclidiana(ruta[i], ruta[i+1])
	}
	// Agregar la distancia desde el último nodo hasta el primero para cerrar el ciclo
	distanciaTotal += DistanciaEuclidiana(ruta[len(ruta)-1], ruta[0])
	return distanciaTotal
}
