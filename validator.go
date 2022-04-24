package markednotes

type Validator interface {
	Struct(any) ([]ErrorField, bool)
}
