package matrices

import (
	"errors"
	"fmt"
	"reflect"
)

func MultiplyMatrix(data1, data2 *Matrix) (*Matrix, error) {
	if len(data1.Values[0]) != len(data2.Values) {
		return nil, errWrongSize
	}
	tmp := make([]Data, len(data2.Values[0]))
	var res [][]Data

	for range len(data1.Values) {
		res = append(res, tmp)
	}

	for i := range data1.Values {
		for j := range data2.Values[0] {
			results := []Data{}
			for k := range data1.Values[i] {
				val, err := multiplier(data1.Values[i][k], data2.Values[k][j])
				if err != nil {
					return nil, fmt.Errorf("multiplier %s", err)
				}

				results = append(results, val)
			}

			sum := results[0]
			for k := 1; k < len(results); k++ {
				sum, _ = adder(sum, results[k])
				_, err := adder(sum, results[k])
				if err != nil {
					return nil, fmt.Errorf("adder %s", err)
				}
			}
			res[i][j] = sum
		}
	}
	return &Matrix{
		Values: res,
	}, nil
}

func multiplier(a, b Data) (Data, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil, errors.New("нельзя умножить Int на bool")
	}
	switch a.GetData().(type) {
	case int:
		return Integer{
			Value: (a.GetData()).(int) * (b.GetData()).(int),
		}, nil
	case bool:
		return Boolean{
			Value: (a.GetData()).(bool) && (b.GetData()).(bool),
		}, nil
	default:
		return nil, errors.New("что то пошло не так")
	}
}
