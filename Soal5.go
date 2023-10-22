//mendeklarasikan paket untuk menjalankan
package main

//import "fmt" untuk input dan output
import "fmt"

//membuat fungsi utama untuk menjalankan kode
func main(){
	//mendeklarasikan variabel n berupa integer
	var n int

	//input bilangan dalam variabel n
	print("bilangan : ")
	fmt.Scan(&n)

	//string untuk memberikan informasi sebelum mencetak hasilnya
	fmt.Println("faktor-faktor dari", n, "adalah")

	//loop untuk menentukan faktor bilangan
	for i := 1; i <= n; i++ {
		//output berupa hasil boolean faktor dari angka 1 sampai n
		fmt.Printf("%d = %t\n", i, n%i == 0)
		//%d = placeholder untuk memasukkan variabel berupa integer
		//%t\n = placeholder untuk memasukkan variabel berupa boolean dengan kondisi n%i == 0
	}
}