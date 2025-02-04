package greeting //Принадлежит не пакету «main», а пакету «greeting»!

import "fmt"

func Hello() { //Первые буквы в верхнем регистре, чтобы функции экспортировались!
	fmt.Println("Hello!")
}

func Hi() { //Первые буквы в верхнем регистре, чтобы функции экспортировались!
	fmt.Println("Hi!")
}
