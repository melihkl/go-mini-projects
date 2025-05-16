package prompt

import (
	"bufio"
	"fmt"
	"json-user-register-cli/models"
	"os"
	"strconv"
	"strings"
)

func GetUserInput() models.User {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Adiniz: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Yasiniz: ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Printf("Gecersiz Yas")
	}

	return models.User{
		Name: name,
		Age:  age,
	}
}
