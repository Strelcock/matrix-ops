package menu

import (
	"bufio"
	"fmt"
	"labs/matrix-ops/matrices"
	"log"
	"os"
	"strings"
)

func Greet() {
	fmt.Println("\nHere you can add and multiply matrices!")
	fmt.Println("Choose what you want to do: ")
	fmt.Print(
		"1. Define first matrix\n" +
			"2. Define second matrix\n" +
			"3. Add\n" +
			"4. Multiply\n" +
			"5. Quit\n",
	)
}

func Menu() {
	var m1, m2 matrices.Matrix
	for {
		Greet()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		switch strings.TrimSpace(input) {
		case "1":
			describe()
			if DefineMatrix(&m1, m1, 1) {
				continue
			}

		case "2":
			describe()
			if DefineMatrix(&m2, m2, 2) {
				continue
			}

		case "3":
			sum, err := matrices.AddMatrix(&m1, &m2)
			if err != nil {
				log.Fatalf("add: %s", err)
			}
			fmt.Println("Matrix 1 + Matrix 2:")
			matrices.PrintMatrix(*sum)
			fmt.Println("Enter anythnig to continue...")
			scanner.Scan()

		case "4":
			mult, err := matrices.MultiplyMatrix(&m1, &m2)
			if err != nil {
				log.Fatalf("mult: %s", err)
			}
			fmt.Println("Matrix 1 * Matrix 2:")
			matrices.PrintMatrix(*mult)
			fmt.Println("Enter anythnig to continue...")
			scanner.Scan()

		default:
			os.Exit(0)
		}

	}
}
