package main

import "fmt"

var variable_1 string
var variable_2 = "deger2"

//Go da unused kontrolü var kullanımayan değişken oldugu durumlarda hata verir bunu burdan kaldır diyor. paketlerde dahil bu duruma
func main() {
	variable_1 = "Emre"
	variable_3 := "var değişkensiz değer atama işlemi"

	fmt.Println(variable_1, variable_2, variable_3)

}
