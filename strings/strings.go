package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "lorem ipsum dolor sit amet"
	//:5 0 İLE 5 arasında ki karakterleri ver
	str_1 := str[:5]
	//toplam uzunluğundan 4 karakter geri gel ameti alır sadece
	str_2 := str[len(str)-4:]
	str_3 := fmt.Sprintf("%s ipsum dolor sit %s", str_1, str_2)

	if strings.EqualFold(str_1, "L0rEM") {
		fmt.Println("str_1  equal to L0rEM")
	}

	if strings.HasPrefix(str, "lorem") {
		fmt.Println("string is started lorem")
	}

	fmt.Println(str)
	fmt.Println(str_1)
	fmt.Println(str_2)
	fmt.Println(str_3)
	fmt.Println(strings.ToUpper(str))

}
