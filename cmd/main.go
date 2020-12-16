package main

import (
	"AIS/internal/collaborate"
	"AIS/internal/data"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//data.GenerateUserData()
	_, users, matrix := data.LoadUserItem()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите как вас зовут: ")
	U, _ := reader.ReadString('\n')
	U = strings.Replace(U, "\n", "", -1)

	userIndex, userError := data.FindUser(users, U)
	if userError != nil {
		fmt.Println("Мы вас не нашли :( ")
		os.Exit(2)
	}
	fmt.Printf("Привет, %s!\n", U)

	fmt.Println("Введите название статьи: ")
	A, _ := reader.ReadString('\n')
	A = strings.Replace(A, "\n", "", -1)

	articles := data.Load()
	a, arError := data.Find(articles, A)
	if arError != nil {
		fmt.Println("Такая статья не найдена :(  ")
		os.Exit(2)
	}

	fmt.Printf("Статья А: %+v \n\n", a)
	collaborate.GetBest(5, userIndex, articles, matrix)

}
