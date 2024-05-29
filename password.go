package password

import (
	"math/rand/v2"
	"regexp"

	"github.com/135yshr/password/policy"
)

// Password is a password type.
type Password string

// String returns the password string.
func (p Password) String() string {
	return string(p)
}

// Generator is a password generator.
type Generator interface {
	Generate(length int) string
}

type generator struct {
	letters  []rune
	size     int
	policies []policy.Policy
}

// New returns a new generator.
func New(policies ...policy.Policy) Generator {
	if len(policies) == 0 {
		policies = []policy.Policy{policy.WithDefault}
	}

	letters := []rune{}
	for _, policy := range policies {
		letters = append(letters, policy.Letters()...)
	}

	return generator{
		letters:  letters,
		size:     len(letters),
		policies: policies,
	}
}

func (g generator) Generate(length int) string {
	password := g.generate(length)
	for !g.IsValidate(password) {
		password = g.generate(length)
	}

	return password
}

func (g generator) IsValidate(password string) bool {
	return regexp.MustCompile("^[" + regexp.QuoteMeta(string(g.letters)) + "]+$").MatchString(password)
}

func (g generator) generate(length int) string {
	buf := make([]rune, length)

	for i := range buf {
		buf[i] = g.letters[rand.N(g.size)] //nolint:gosec // use rand.N here because it is not safe to use rand.Intn
	}

	return string(buf)
}
