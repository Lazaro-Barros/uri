package main

import "fmt"

func main() {
	var tempoGasto, velocidadeMedia int64
	fmt.Scan(&tempoGasto, &velocidadeMedia)
	distancia := tempoGasto * velocidadeMedia
	fmt.Printf("%.3f\n", float64(distancia)/12)
}
