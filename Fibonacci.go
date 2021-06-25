package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var resultados []int

func inicio() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Leonardo Bonacci: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		text = strings.TrimSuffix(text, "\n")
		indice, err := strconv.ParseInt(text, 10, 32)
		if err != nil {
			fmt.Println(err)
		} else {
			resultados = append(resultados, 0)
			resultados = append(resultados, 1)

			i := int(indice)
			fmt.Println(fibonacci(i))
			fmt.Println(resultados)
		}

	}
}

func fibonacci(i int) int {
	if i <= len(resultados)-1 {
		return resultados[i]
	}
	a := fibonacci(i-1) + fibonacci(i-2)
	resultados = append(resultados, a)
	return resultados[i]
}
