package menu

import (
	"bufio"
	"fmt"
	"labs/matrix-ops/matrices"
	"log"
	"os"
)

func DefineMatrix(dest *matrices.Matrix, tmp matrices.Matrix, n int) bool {

	scanner := bufio.NewScanner(os.Stdin)

	if !tmp.IsEmpty() {
		fmt.Printf("Matrix %d is already defined, do you want to redefine it? [Y/N]\n", n)
		scanner.Scan()
		in := scanner.Text()
		if in != "Y" {
			return true
		}
	}

	input, err := matrices.GetInput()
	if err != nil {
		log.Fatalf("GetInput: %s", err)
	}

	*dest, err = matrices.NewMatrix(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matrix %d:\n", n)
	matrices.PrintMatrix(*dest)
	return false
}
