package account

import (
	"errors"
	"fmt"
)

type Hesap struct {
	Kisi   string
	Bakiye float64
}

func (h *Hesap) Yatir(miktar float64) error {
	if miktar <= 0 {
		return errors.New("yatirilacak miktar pozitif olmali")
	}

	h.Bakiye += miktar
	fmt.Printf("%.2f miktar yatirildi. Yeni bakiye: %.2f", miktar, h.Bakiye)
	return nil
}

func (h *Hesap) Cek(miktar float64) error {
	if miktar <= 0 {
		return errors.New("yatirilacak miktar pozitif olmali")
	}

	if miktar > h.Bakiye {
		return errors.New("yetersiz Bakiye")
	}

	h.Bakiye -= miktar
	fmt.Printf("%.2f miktar cekildi. Yeni Bakiye: %.2f", miktar, h.Bakiye)
	return nil
}

func (h Hesap) ToplamBakiye() {
	fmt.Printf("Kisi %s, Toplam Bakiye: %.2f", h.Kisi, h.Bakiye)
}

