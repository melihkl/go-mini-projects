package main

import (
	"account/account"
	"flag"
	"fmt"
)

func main() {

	islem := flag.String("islem", "", "Yapilacak islem: yatir, cek, toplam bakiye")
	miktar := flag.Float64("miktar", 0, "yatirilacak ya da cekilecek miktar")
	kisi := flag.String("kisi", "", "hesap sahibi")

	flag.Parse()

	//hesap olusturma islemi

	hesap := account.Hesap{Kisi: *kisi, Bakiye: 0.0}

	switch *islem {
	case "yatir":
		if *miktar <= 0 {
			fmt.Printf("yatirilacak miktar pozitif olmali")
			return
		}

		err := hesap.Yatir(*miktar)

		if err != nil {
			fmt.Println("Hata", err)
		}

	case "cek":
		if *miktar <= 0 {
			fmt.Printf("cekilecek miktar pozitif olmali")
			return
		}

		err := hesap.Cek(*miktar)

		if err != nil {
			fmt.Println("Hata", err)
		}
	case "toplamBakiye":
		hesap.ToplamBakiye()
	}

}
