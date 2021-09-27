package error

type SourceError struct {
	name string
}

func (s SourceError) Error() string { return s.name }

var (
	SourceNotSupportedError = SourceError{name: "source not supported"}
)
