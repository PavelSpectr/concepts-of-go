package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt") //Файл данных открывается для чтения.
	//Если при открытии файла произошла ошибка, сообщить о ней и завершить работу.
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file) ////Для файла создается новое значение Scanner.
	//Цикл выполняется до того, как будет достигнут конец файла, а scanner.Scan вернет false.
	for scanner.Scan() { //Читает строку из файла.
		fmt.Println(scanner.Text()) //Выводит строку.
	}

	err = file.Close() //Закрывает файл для освобождения ресурсов.
	//Если при закрытии файла произошла ошибка, сообщить о ней и завершить работу.
	if err != nil {
		log.Fatal(err)
	}

	//Если при сканировании файла произошла ошибка, сообщить о ней и завершить работу.
	if scanner.Err() != nil {
		log.Fatal(scanner.Err)
	}
}
