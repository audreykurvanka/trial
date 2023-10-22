//mendeklarasikan paket untuk menjalankan kode GO
package main

//import "fmt" untuk input dan output
import "fmt"

//membuat fungsi utama untuk menjalankan kode
func main() {
	//mendeklarasikan variabel n berupa integer
	var n int

	//input jumlah persegi dalam variabel n
	fmt.Print("banyak persegi :")
	fmt.Scan(&n)

	//loop untuk mengulang perhitungan persegi sebanyak n
	for i := 0; i < n; i++ {
		//mendeklarasikan variabel sisi berupa integer
		var sisi int

		//input sisi dalam variabel sisi
		fmt.Scan(&sisi)

		//mendeklarasikan rumus luas dan keliling persegi
		luas := sisi * sisi
		keliling := 4 * sisi

		//output berupa hasil perhitungan rumus luas dan keliling persegi
		fmt.Printf("luas dan keliling adalah %d dan %d", luas, keliling)
	}
}
