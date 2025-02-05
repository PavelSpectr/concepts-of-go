package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats читает значение float64 из каждой строки файла.
// Имя файла с данными передается в аргументе.
// Функция возвращает массив чисел и любую обнаруженную ошибку.
func GetFloats(filename string) ([]float64, error) {
	var numbers []float64 ////Объявление возвращаемого массива.

	file, err := os.Open(filename) //Открывает файл с переданным именем.
	//Если при открытии файла произошла ошибка, вернуть ее.
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64) //Строка, прочитанная из файла, преобразуется в float64.
		//Если при преобразовании значения в число произошла ошибка, вернуть ее.
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	//Если при закрытии файла произошла ошибка, вернуть ее.
	if err != nil {
		return nil, err
	}

	//Если при сканировании файла произошла ошибка, вернуть ее.
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil //Если выполнение дошло до этой точки, значит, ошибок не было, поэтому программа возвращает массив чисел и значение ошибки «nil».
}
