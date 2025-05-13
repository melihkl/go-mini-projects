package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task/task"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- GÖREV TAKİP ---")
		fmt.Println("1. Görev Ekle")
		fmt.Println("2. Görevleri Listele")
		fmt.Println("3. Görev Sil")
		fmt.Println("4. Görev Güncelle")
		fmt.Println("5. Çıkış Yap")
		fmt.Print("Seçiminiz: ")

		secim, _ := reader.ReadString('\n')
		secim = strings.TrimSpace(secim)

		switch secim {
		case "1":
			fmt.Print("Başlık: ")
			title, _ := reader.ReadString('\n')

			fmt.Print("Açıklama: ")
			description, _ := reader.ReadString('\n')

			fmt.Print("Durum: ")
			status, _ := reader.ReadString('\n')

			title = strings.TrimSpace(title)
			description = strings.TrimSpace(description)
			status = strings.TrimSpace(status)

			task.AddTask(title, description, status)

		case "2":
			task.ListTask()

		case "3":
			fmt.Print("Silmek için görev ID giriniz\n")
			idInput, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idInput))

			if err != nil {
				fmt.Printf("Bu ID için görev kaydı bulunmamaktadır!")
				continue
			}

			if !task.Exists(id) {
				fmt.Printf("ID %d ile görev bulunamadı.\n", id)
				continue
			}
			task.DeleteTask(id)

		case "4":
			fmt.Print("Görev ID giriniz : ")
			idInput, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idInput))

			if err != nil {
				fmt.Printf("Bu ID için görev kaydı bulunmamaktadır!")
				continue
			}

			if !task.Exists(id) {
				fmt.Printf("ID %d ile görev bulunamadı.\n", id)
				continue
			}

			fmt.Print("Güncel Başlık: ")
			title, _ := reader.ReadString('\n')

			fmt.Print("Güncel Açıklama: ")
			description, _ := reader.ReadString('\n')

			fmt.Print("Güncel Durum: ")
			status, _ := reader.ReadString('\n')

			title = strings.TrimSpace(title)
			description = strings.TrimSpace(description)
			status = strings.TrimSpace(status)

			task.UpdateTask(id, title, description, status)

		case "5":
			fmt.Printf("Çıkış yaptınız!")
			return
		}
	}
}
