package shared

// Output struct
type Output struct {
	Result interface{}
	Error  error
}

type OutputV1 struct {
	Result interface{}
	Errors interface{}
	Error  error
	Code   int
}
