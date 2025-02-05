package average

import (
	"concepts-of-go/datafile"
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt") //Загружает файл data.txt, разбирает содержащиеся в нем числа и сохраняет массив.
	//Если произошла ошибка, программа сообщает о ней и завершает работу.
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average %0.2f\n", sum/sampleCount)
}
