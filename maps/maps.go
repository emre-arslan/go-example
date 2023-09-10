package main

import "fmt"

//burada initialize ediliyor ama referansı yok bellekte kullanım olarak
var map_1 map[int]string

func main() {
	//buarada yukarıda initialize edilmiş mapi tanımlıyoruz
	map_1 = make(map[int]string)

	map_2 := make(map[int]int)
	//virgulden sonra ki parametre boyuta karşılık gelmektedir.
	map_3 := make(map[int]int, 3)

	map_1[0] = "1"
	map_2[0] = 2
	map_3[0] = 3
	map_3[1] = 3
	map_3[2] = 3
	map_3[3] = 3

	fmt.Println(map_1, map_2, map_3)
	fmt.Printf("map_1 len:%d \n ", len(map_1))
	fmt.Printf("map_2 len:%d \n ", len(map_2))
	fmt.Printf("map_3 len:%d \n ", len(map_3))

}
