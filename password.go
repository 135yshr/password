package password

// Password is a password type
type Password string

// String returns the password string
func (p Password) String() string {
	return string(p)
}

type Generator interface {
	Generate(length int) string
}

type generator struct {
}

func New() Generator {
	return generator{}
}

func (g generator) Generate(length int) string {
	return "aaaaa"
}
