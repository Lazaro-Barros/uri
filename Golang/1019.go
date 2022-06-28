package main

import "fmt"

func main() {
	var entrada, horas, minutos, segundos int64
	fmt.Scan(&entrada)
	for entrada >= 3600 {
		horas += 1
		entrada -= 3600
	}
	for entrada >= 60 {
		minutos += 1
		entrada -= 60
	}
	for entrada != 0 {
		segundos += 1
		entrada -= 1
	}
	fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)
}
