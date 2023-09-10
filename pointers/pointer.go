// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // stringe nil atanabilir bir hale getirdik . referans type çevirip
// // & işareti pointerin adresini verir
// type String *string

// func main() {
// 	var rstr String
// 	var pstr string

// 	fmt.Println(rstr) //point eden adresi basrıcak

// 	fmt.Println(pstr)

// 	pstr = "go emre"
// 	rstr = &pstr

// 	fmt.Println(rstr)
// 	fmt.Println(*rstr) //pointte değeri verir

// 	fmt.Println(pstr)

// 	pstr = "değeri değiştiridm"
// 	fmt.Println(*rstr)
// 	fmt.Println(pstr)

// 	*rstr = "değeri değiştirdim 2 "

// 	fmt.Println(*rstr)
// 	fmt.Println(pstr)

// 	//method parametre gönderirken değerin kopyasını gönderir. bu yuzden replaceStr olanda yeni bir adres oluşur o string için
// 	replaceStr2(rstr)
// 	fmt.Println(*rstr)
// 	fmt.Println(pstr)

// }

// func replaceStr(s string) {
// 	s = strings.Replace(s, "değeri", "value", 1)
// }
// func replaceStr2(s String) {
// 	*s = strings.Replace(*s, "değeri", "value", 1)
// }
