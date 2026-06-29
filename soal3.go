package main

import "fmt"

func main() {
	var cx1, cy1, r1 int		//variabel untuk menyimpan koordinat pusat dan jari-jari lingkaran 1
	var cx2, cy2, r2 int		//variabel untuk menyimpan koordinat pusat dan jari-jari lingkaran 2
	var x, y int		//variabel untuk menyimpan koordinat titik yang akan diperiksa

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	d1 := (x-cx1)*(x-cx1) + (y-cy1)*(y-cy1)		//menghitung jarak kuadrat dari titik ke pusat lingkaran 1
	d2 := (x-cx2)*(x-cx2) + (y-cy2)*(y-cy2)		//menghitung jarak kuadrat dari titik ke pusat lingkaran 2

	dalam1 := d1 <= r1*r1		//memeriksa apakah titik berada di dalam lingkaran 1 dengan membandingkan jarak kuadrat dengan jari-jari kuadrat			
	dalam2 := d2 <= r2*r2		//memeriksa apakah titik berada di dalam lingkaran 2 dengan membandingkan jarak kuadrat dengan jari-jari kuadrat

	if dalam1 && dalam2 {				//memeriksa apakah titik berada di dalam kedua lingkaran
		fmt.Println("Titik di dalam lingkaran 1 dan 2")		//jika benar, mencetak bahwa titik berada di dalam kedua lingkaran	
	} else if dalam1 {										//memeriksa apakah titik berada di dalam lingkaran 1	
		fmt.Println("Titik di dalam lingkaran 1")		//jika benar, mencetak bahwa titik berada di dalam lingkaran 1
	} else if dalam2 {									//memeriksa apakah titik berada di dalam lingkaran 2
		fmt.Println("Titik di dalam lingkaran 2")			//jika benar, mencetak bahwa titik berada di dalam lingkaran 2
	} else {												//	jika titik tidak berada di dalam kedua lingkaran
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}