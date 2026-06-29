package main

import (
	"fmt"
)

func faktorial(n int) int {			//fungsi untuk menghitung faktorial dari n
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 1; i <= n; i++ {			//mengalikan hasil dengan i untuk setiap i dari 1 sampai n
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {			//fungsi untuk menghitung permutasi dari n dan r
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))		//fungsi untuk menghitung kombinasi dari n dan r dengan rumus n! / (r! * (n-r)!)
}

func main() {			//input nilai a, b, c, d
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutasi(a, c), kombinasi(a, c))		//menghitung dan mencetak hasil permutasi dan kombinasi untuk a dan c
	fmt.Println(permutasi(b, d), kombinasi(b, d))		//menghitung dan mencetak hasil permutasi dan kombinasi untuk a dan c, serta b dan d
}