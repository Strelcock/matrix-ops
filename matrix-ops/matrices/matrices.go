package matrices

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errWrongSize = errors.New("wrong size")

type Matrix struct {
	Values [][]Data
}

func (m *Matrix) IsEmpty() bool {
	return len(m.Values) == 0
}

// INT OR BOOL
type Data interface {
	GetData() any
}

func GetInput() ([][]Data, error) {
	newScanner := bufio.NewScanner(os.Stdin)
	var matrixData [][]Data

	for newScanner.Scan() {
		line := newScanner.Text()

		if line == "" {
			break
		}

		stringNumbers := strings.Fields(line)
		var row []Data

		for _, v := range stringNumbers {

			if number, err := strconv.Atoi(v); err == nil {
				row = append(row, Integer{Value: number})
			} else if val, err := strconv.ParseBool(v); err == nil {
				row = append(row, Boolean{Value: val})
			} else {
				return nil, errors.New("cannot parse string")
			}
		}

		matrixData = append(matrixData, row)
	}

	return matrixData, nil
}

func NewMatrix(numbers [][]Data) (Matrix, error) {
	checkList := make(map[int]struct{})
	for _, v := range numbers {
		checkList[len(v)] = struct{}{}
	}
	if len(checkList) != 1 {
		return Matrix{}, errWrongSize
	}
	return Matrix{
		Values: numbers,
	}, nil
}

func PrintMatrix(m Matrix) {
	for i := range m.Values {
		for j := range m.Values[i] {
			fmt.Printf("%v ", m.Values[i][j])
		}
		fmt.Println()
	}
}
