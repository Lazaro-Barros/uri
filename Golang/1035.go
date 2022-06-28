package main

import "fmt"

func main() {
	var a, b, c, d int64
	fmt.Scanln(&a, &b, &c, &d)
	bIsBiggerThanCAndDIsBiggerThanA := (b > c && d > a)
	sumOfCAndDIsBiggerThanSumOfAAndB := ((c + d) > (a + b))
	AreCAndDPositive := c >= 0 && d >= 0
	isApair := (a % 2) == 0
	if bIsBiggerThanCAndDIsBiggerThanA && sumOfCAndDIsBiggerThanSumOfAAndB && AreCAndDPositive && isApair {
		fmt.Println("Valores aceitos")
		return
	}
	fmt.Println("Valores nao aceitos")
}
