//mendeklarasikan paket untuk menjalankan kode GO
package main

//import "fmt" untuk input dan output
import "fmt"

//membuat fungsi utama untuk menjalankan kode
func main(){
	//mendeklarasikan  variabel n berupa integer
	var n int
	//mendeklarasikan variabel hasil
	var s string

	//input banyaknya string untuk diulangi
	print("n :")
	fmt.Scan(&n)

	//input string yang akan diulangi
	print("string :")
	fmt.Scan(&s)

	//loop untuk print string sebanyak n
	for i := 0; i < n; i++ {
		//output berupa string
		fmt.Println(s)
	}
}
