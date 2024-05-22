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

func DistanciaEuclidiana(nodo0, nodo1 lectorinstancias.Punto) float64 {
	DX := nodo0.Posicionx - nodo1.Posicionx
	DY := nodo0.Posiciony - nodo1.Posiciony
	return math.Sqrt(DX*DX + DY*DY)
}
func VecinoMasCercano(nodos []lectorinstancias.Punto) ([]lectorinstancias.Distancia, float64) {
	if len(nodos) == 0 {
		return nil, 0
	}

	visitados := make(map[string]bool)
	var ruta []lectorinstancias.Distancia
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
			ruta = append(ruta, lectorinstancias.Distancia{
				PuntoInicial: nodoActual.Nombre,
				PuntoFinal:   nodoMasCercano.Nombre,
				Distancia:    minDistancia,
			})
			totalDistancia += minDistancia
			nodoActual = nodoMasCercano
			visitados[nodoActual.Nombre] = true
		}
	}

	// Regresar al nodo inicial para completar el ciclo
	ruta = append(ruta, lectorinstancias.Distancia{
		PuntoInicial: nodoActual.Nombre,
		PuntoFinal:   nodos[0].Nombre,
		Distancia:    DistanciaEuclidiana(nodoActual, nodos[0]),
	})
	totalDistancia += DistanciaEuclidiana(nodoActual, nodos[0])

	return ruta, totalDistancia
}

func InsercionMasCercana(nodos []lectorinstancias.Punto) ([]lectorinstancias.Distancia, float64) {
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
			PuntoInicial: nodos[0].Nombre,
			PuntoFinal:   nodos[1].Nombre,
			Distancia:    DistanciaEuclidiana(nodos[0], nodos[1]),
		})
		ruta = append(ruta, lectorinstancias.Distancia{
			PuntoInicial: nodos[1].Nombre,
			PuntoFinal:   nodos[0].Nombre,
			Distancia:    DistanciaEuclidiana(nodos[1], nodos[0]),
		})
		totalDistancia = 2 * DistanciaEuclidiana(nodos[0], nodos[1])
	}

	// Inserción más cercana
	for len(visitados) < len(nodos) {
		nodoMasCercano := lectorinstancias.Punto{}
		minIncremento := math.MaxFloat64
		posicion := 0

		for _, nodo := range nodos {
			if !visitados[nodo.Nombre] {
				for i := 0; i < len(ruta); i++ {
					nodoI := encontrarNodo(ruta[i].PuntoInicial, nodos)
					nodoF := encontrarNodo(ruta[i].PuntoFinal, nodos)
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
		if (nodoMasCercano != lectorinstancias.Punto{}) {
			nodoI := encontrarNodo(ruta[posicion].PuntoInicial, nodos)
			nodoF := encontrarNodo(ruta[posicion].PuntoFinal, nodos)
			ruta = append(ruta[:posicion+1], ruta[posicion:]...) // Hacer espacio para la nueva distancia
			ruta[posicion] = lectorinstancias.Distancia{
				PuntoInicial: nodoI.Nombre,
				PuntoFinal:   nodoMasCercano.Nombre,
				Distancia:    DistanciaEuclidiana(nodoI, nodoMasCercano),
			}
			ruta[posicion+1] = lectorinstancias.Distancia{
				PuntoInicial: nodoMasCercano.Nombre,
				PuntoFinal:   nodoF.Nombre,
				Distancia:    DistanciaEuclidiana(nodoMasCercano, nodoF),
			}
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
	return lectorinstancias.Punto{}
}

func calcularDistancias(nodos []lectorinstancias.Punto, distancias *[]lectorinstancias.Distancia, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(nodos)-1; i++ {
		for j := i + 1; j < len(nodos); j++ {
			distancia := DistanciaEuclidiana(nodos[i], nodos[j])
			*distancias = append(*distancias, lectorinstancias.Distancia{
				PuntoInicial: nodos[i].Nombre,
				PuntoFinal:   nodos[j].Nombre,
				Distancia:    distancia,
			})
		}
	}
}

func Calculo(IndiceNodos []lectorinstancias.Punto) ([]lectorinstancias.Distancia, []lectorinstancias.Distancia) {
	rand.Seed(time.Now().UnixNano())
	if len(IndiceNodos) == 0 {
		return nil, nil
	}
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

// Implementación del algoritmo de 2-opt para mejorar la ruta en la búsqueda de vecindario
func DosOpt(ruta []lectorinstancias.Punto) ([]lectorinstancias.Distancia, float64) {
	mejorado := true
	distanciaTotal := calcularDistanciaTotal(ruta)

	for mejorado {
		mejorado = false
		for i := 1; i < len(ruta)-2; i++ {
			for j := i + 1; j < len(ruta)-1; j++ {
				nuevaRuta := intercambiarRutas(ruta, i, j)
				nuevaDistancia := calcularDistanciaTotal(nuevaRuta)

				if nuevaDistancia < distanciaTotal {
					ruta = nuevaRuta
					distanciaTotal = nuevaDistancia
					mejorado = true
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

func intercambiarRutas(ruta []lectorinstancias.Punto, i, j int) []lectorinstancias.Punto {
	nuevaRuta := make([]lectorinstancias.Punto, len(ruta))
	copy(nuevaRuta, ruta)

	for k := 0; k < (j-i)/2; k++ {
		nuevaRuta[i+k], nuevaRuta[j-k] = nuevaRuta[j-k], nuevaRuta[i+k]
	}

	return nuevaRuta
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
