package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// - benim json çıktımda olmasın denildiği zaman kullanılır
// taglari json çıksına basıyor.
// omitempty null değerleri getirmez
type Kullanici struct {
	Ad      string      `json:"name"`
	Soyad   string      `json:"surname"`
	Takipci []Kullanici `json:"followers,omitempty"`
	// Begeni []struct {
	// 	Tarih time.Time
	// }
	// Begeni Begen
}

type Begen struct {
	Tarih time.Time
}

func main() {
	k := Kullanici{
		Ad:    "GO",
		Soyad: "Emre",
		Takipci: []Kullanici{
			{
				Ad:    "takipçi 1",
				Soyad: "1",
			},
			{
				Ad:    "takipçi 2",
				Soyad: "2",
			},
			{
				Ad:    "takipçi 3",
				Soyad: "3",
			},
		},
	}
	//arr Marshall modeli byte array çeviriyor
	arr, _ := json.Marshal(k)
	//string byte array verircen byte arrayi stringe çevirmek için yeterli oluyor.
	fmt.Println(string(arr))
	fmt.Println(k)
}
