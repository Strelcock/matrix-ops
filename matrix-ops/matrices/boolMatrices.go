package matrices

type Boolean struct {
	Value bool
}

func (b Boolean) GetData() any {
	return b.Value
}
