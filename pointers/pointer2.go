package main

func main() {
	a := 5
	println("a'nın değeri:", a)
	println("a'nın adresi:", &a)
	arttir(a)
	println("a'nın değeri:", a)
	//ampersant ile göndermek zorrundayız. adresi göndeririyoruz
	arttir2(&a)
	println("a'nın değeri:", a)

	//b ye a nın adresini veriyoruz
	b := &a
	//c ye de sende b nin adresini tut diyoruz
	c := b

	*c = 11

	println(*b)

}

func arttir(numb int) {
	numb++
	println("a +'nın değeri:", numb)
	println("a + 'nın adresi:", &numb)
}

// * ile pointerın işarettiği adresteki değeri almış oluyoruz.
func arttir2(numb *int) {
	*numb++
	println("a +2'nın değeri:", *numb)
	println("a +2 'nın adresi:", numb)
}
