package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Estructura para representar un punto en el plano 2D
type Point struct {
	X, Y float64
}

// Función para calcular la distancia euclidiana entre dos puntos
func distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Función para leer el archivo y devolver una lista de puntos
func readPoints(filename string) ([]Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "NODE_COORD_SECTION") {
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "EOF" {
			break
		}
		fields := strings.Fields(line)
		x, _ := strconv.ParseFloat(fields[1], 64)
		y, _ := strconv.ParseFloat(fields[2], 64)
		points = append(points, Point{x, y})
	}

	return points, nil
}

// Función para encontrar el punto más cercano a un punto en la lista de puntos
func findNearestPoint(p Point, points []Point, visited []bool) int {
	minDist := math.Inf(1)
	nearestIndex := -1

	for i, point := range points {
		if !visited[i] {
			dist := distance(p, point)
			if dist < minDist {
				minDist = dist
				nearestIndex = i
			}
		}
	}

	return nearestIndex
}

// Función para resolver el TSP utilizando el algoritmo de inserción más cercana
func solveTSP(points []Point) []int {
	n := len(points)
	tour := make([]int, n)
	visited := make([]bool, n)
	start := 0
	tour[0] = start
	visited[start] = true

	for i := 1; i < n; i++ {
		prevPoint := points[tour[i-1]]
		nearestIndex := findNearestPoint(prevPoint, points, visited)
		tour[i] = nearestIndex
		visited[nearestIndex] = true
	}

	// Agregar el punto de origen al final del recorrido para cerrar el ciclo
	tour = append(tour, start)

	return tour
}

func totalDistance(points []Point, tour []int) float64 {
	total := 0.0
	n := len(tour)

	for i := 0; i < n; i++ {
		from := points[tour[i]]
		to := points[tour[(i+1)%n]] // Para cerrar el ciclo
		total += distance(from, to)
	}

	return total
}

func main() {
	filename := "dj38.tsp"
	points, err := readPoints(filename)

	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	tour := solveTSP(points)
	total := totalDistance(points, tour)

	fmt.Println("Tour óptimo:", tour)
	fmt.Println("Total del tour:", total)
}
