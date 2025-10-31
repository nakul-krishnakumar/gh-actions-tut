package adder


type Adder struct {
	A int
	B int
	Result int
}

func New(a, b int) *Adder {
	result := a + b
	return &Adder{
		A: a,
		B: b,
		Result: result,
	}
}

func (A *Adder) GetResult() int {
	return A.Result
}