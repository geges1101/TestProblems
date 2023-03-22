package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//Дана строка (возможно, пустая), состоящая из букв A-Z: AAAABBBCCXYZDDDDEEEFFFAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBB
//Нужно написать функцию RLE, которая на выходе даст строку вида: A4B3C2XYZD4E3F3A6B28
//И сгенерирует ошибку, если на вход пришла невалидная строка.
//Пояснения: Если символ встречается 1 раз, он остается без изменений; Если символ повторяется более 1 раза, к нему добавляется количество повторений.

func compressString(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	var result strings.Builder
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		currentChar := runes[i]
		count := 1

		for j := i + 1; j < len(runes) && runes[j] == currentChar; j++ {
			count++
			i++
		}

		if currentChar < 'A' || currentChar > 'Z' {
			return "", errors.New("невалидная строка")
		}

		if count == 1 {
			result.WriteRune(currentChar)
		} else {
			result.WriteRune(currentChar)
			result.WriteString(strconv.Itoa(count))
		}
	}

	return result.String(), nil
}

func main() {
	input := "AAAABBBCCXYZDDDDEEEFFFAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBB"
	output, err := compressString(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}
