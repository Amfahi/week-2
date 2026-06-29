package main

import "fmt"

func f(x int, hasil *int) {		//hasil pointer untuk menyimpan hasil dari fungsi f
	*hasil = x * x
}

func g(x int, hasil *int) {		//hasil pointer untuk menyimpan hasil dari fungsi g
	*hasil = x - 2
}

func h(x int, hasil *int) {		//hasil  pointer untuk menyimpan hasil dari fungsi h
	*hasil = x + 1
}

func main() {					//input nilai a, b, c
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	var x, y, z int				//variabel untuk menyimpan hasil sementara dari fungsi f, g, h

	// f(g(h(a)))				//h(a) -> g(h(a)) -> f(g(h(a)))
	h(a, &x)
	g(x, &y)
	f(y, &z)
	fmt.Println(z)				

	// g(h(f(b)))		// f(b) -> h(f(b)) -> g(h(f(b)))
	f(b, &x)
	h(x, &y)
	g(y, &z)
	fmt.Println(z)

	// h(f(g(c)))		// g(c) -> f(g(c)) -> h(f(g(c)))
	g(c, &x)
	f(x, &y)
	h(y, &z)
	fmt.Println(z)
}