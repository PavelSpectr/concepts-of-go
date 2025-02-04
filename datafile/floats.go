package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats читает значение float64 из каждой строки файла.
// Имя файла с данными передается в аргументе.
// Функция возвращает массив чисел и любую обнаруженную ошибку.
func GetFloats(filename string) ([3]float64, error) {
	var numbers [3]float64 ////Объявление возвращаемого массива.

	file, err := os.Open(filename) //Открывает файл с переданным именем.
	//Если при открытии файла произошла ошибка, вернуть ее.
	if err != nil {
		return numbers, err
	}

	i := 0 //Переменная для хранения индекса, по которому должно выполняться присваивание.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) //Строка, прочитанная из файла, преобразуется в float64.
		//Если при преобразовании значения в число произошла ошибка, вернуть ее.
		if err != nil {
			return numbers, err
		}
		i++
	}

	err = file.Close()
	//Если при закрытии файла произошла ошибка, вернуть ее.
	if err != nil {
		return numbers, err
	}

	//Если при сканировании файла произошла ошибка, вернуть ее.
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	return numbers, nil //Если выполнение дошло до этой точки, значит, ошибок не было, поэтому программа возвращает массив чисел и значение ошибки «nil».
}
