package main

import "fmt"

var slc_1 []int

func main() {
	// burada vermiş oldugu 3 kapasite değeri appendlerde kapasiteyi doldurduk 3 oldu array 4 ü ekleyince kapasitesine tanımladığın değer kadar arttrır. kapasitesi default olarak 2 dir bir şey belirtmezsek ve 2nin katları olarak arttıyor tabi
	slc_2 := make([]int, 0, 3)
	//sıfır boyutlu oldugundan aşağıdaki kod hata verir neden çünkü sıfırıncı elemanı yok buna ekleme yapılması gerkeir c# list.add gibi
	// slc_2[0] = 2

	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)

	slc_2 = append(slc_2, 1)

	slc_1 = append(slc_1, 2)

	fmt.Println(slc_1, slc_2)
	fmt.Printf("slc_1 len:%d cap:%d \n", len(slc_1), cap(slc_1))
	fmt.Printf("slc_2 len:%d cap:%d \n", len(slc_2), cap(slc_2))

}
