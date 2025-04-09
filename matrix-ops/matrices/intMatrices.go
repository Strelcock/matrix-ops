package matrices

type Integer struct {
	Value int
}

func (i Integer) GetData() any {
	return i.Value
}
