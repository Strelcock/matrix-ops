package menu

import "fmt"

func describe() {
	fmt.Print(
		"To define a matrix just type data, separeted by space\n" +
			"Inputting matrix must look like matrix! Every row must be the same size as every col\n" +
			"When you are done press enter\n",
	)
}
