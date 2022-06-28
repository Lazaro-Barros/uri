package main

import (
	"fmt"
	"time"
)

func main() {

	var entrada int64
	var anos, meses, dias int64 = 0, 0, 0
	fmt.Scan(&entrada)
	start := time.Now()

	// anos
	if entrada >= 365 {
		if entrada%365 == 0 {
			anos = entrada / 365
			entrada = 0
		} else {
			rdi := entrada % 365
			anos = 365 / (entrada - rdi)
			entrada = entrada % 365
		}
	}
	// meses
	if entrada >= 30 {
		if entrada%30 == 0 {
			meses = entrada / 30
			entrada = 0
		} else {
			rdi := entrada % 30
			meses = 30 / (entrada - rdi)
			entrada = entrada % 30
		}
	}

	dias = entrada

	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", anos, meses, dias)
	fmt.Printf("main, execution time %s\n", time.Since(start))

}
