package main

import (
	"AIS/internal/contentbase"
	"AIS/internal/data"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//data.GenerateUserData()
	_, users, _ := data.LoadUserItem()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите как вас зовут: ")
	U, _ := reader.ReadString('\n')
	U = strings.Replace(U, "\n", "", -1)

	_, userError := data.FindUser(users, U)
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
	contentbase.GetBest(5, a, articles)

}
