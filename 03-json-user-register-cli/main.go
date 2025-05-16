package main

import (
	"fmt"
	"json-user-register-cli/prompt"
	"json-user-register-cli/storage"
	"os"
)

func main() {
	for {
		fmt.Println("\n--- User Registry CLI ---")
		fmt.Println("1. Add User")
		fmt.Println("2. List Users")
		fmt.Println("3. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			user := prompt.GetUserInput()
			err := storage.SaveUser(user)

			if err != nil {
				fmt.Print("kullanici eklenirken hata olustu", err)
			}
		case 2:
			users, err := storage.LoadUsers()
			if err != nil {
				fmt.Println("kullanici listesi yuklenirken hata olustu", err)
			} else {
				fmt.Println("\n--- Kullanici Listesi ---")
				for i, user := range users {
					fmt.Printf("%d. %s (%d)\n", i+1, user.Name, user.Age)
				}
			}

		case 3:
			os.Exit(0)
		}
	}
}
