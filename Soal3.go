// mendeklarasikan paket untuk menjalankan kode pada GO
package main

//import "fmt" untuk input dan output
import "fmt"

//membuat fungsi utama untuk menjalankan kode
func main() {
	//mendeklarasikan variabel hari, total jam, dan rata-rata berupa float
	var hari, totalJam, rataRata float32

	//input jumlah hari dalam variabel hari
	fmt.Print("hari :")
	fmt.Scan(&hari)

	//loop untuk menjumlahkan jam sebanyak hari
	for i := 0; i < int(hari); i++ {
		//mendeklarasikan variabel jam berupa float
		var jam float32

		//input jam ngoding dalam variabel jam
		fmt.Scan(&jam)

		//rumus menjumlahkan jam-jam yang telah diinput
		totalJam += jam
	}

	//mendeklarasikan rumus rata-rata waktu yang dihabiskan untuk berlatih ngoding dalam sehari
	rataRata = totalJam / hari

	//output berupa hasil perhitungan rata-rata
	fmt.Print("rata-rata :", rataRata)
}
