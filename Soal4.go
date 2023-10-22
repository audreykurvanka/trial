//mendeklarasikan paket untuk menjalankan 
package main

//import "fmt" untuk input dan output
import "fmt"

//membuat fungsi utama untuk menjalankan kode
func main() {
	//mendeklarasikan variabel n dan hasil
	var n int

	//menginput n (angka yang akan dijadikan faktorial)
	print("n : ")
	fmt.Scan(&n)

	//memberikan nilai awal 
	hasil := 1
	//menghitung faktorial menggunakan loop sebanyak n 
	for i := 1; i <= n; i++ {
		//mendeklarasikan rumus untuk mengalikan sebanyak n
		hasil *= i
	}

	//output berupa hasil 
	fmt.Printf("%d", hasil)
}
