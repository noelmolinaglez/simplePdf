package interfaces

type SimpleDoc interface {
	GetHeaders() map[string]string
	GetValues() map[string]string
}
