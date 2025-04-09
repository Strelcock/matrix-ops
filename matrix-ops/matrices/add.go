package matrices

import (
	"errors"
	"fmt"
	"reflect"
)

func AddMatrix(data1, data2 *Matrix) (*Matrix, error) {
	if len(data1.Values) != len(data2.Values) || len(data1.Values[0]) != len(data2.Values[0]) {
		return nil, errWrongSize
	}

	sum := make([][]Data, len(data1.Values))
	for i := range len(data1.Values) {
		for j := range len(data1.Values[0]) {
			val, err := adder(data1.Values[i][j], data2.Values[i][j])
			if err != nil {
				return nil, fmt.Errorf("adder: %s", err)
			}

			sum[i] = append(sum[i], val)

		}
	}
	return &Matrix{
		Values: sum,
	}, nil
}

func adder(a, b Data) (Data, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil, errors.New("нельзя складывать Int и bool")
	}
	switch a.GetData().(type) {
	case int:
		return Integer{
			Value: (a.GetData()).(int) + (b.GetData()).(int),
		}, nil
	case bool:
		return Boolean{
			Value: (a.GetData()).(bool) || (b.GetData()).(bool),
		}, nil
	default:
		return nil, errors.New("что то пошло не так")
	}

}
