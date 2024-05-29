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
	Generate(length int) Password
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

// Generate generates a password.
func (g generator) Generate(length int) Password {
	password := g.generate(length)
	for !g.IsValidate(password) {
		password = g.generate(length)
	}

	return password
}

// IsValidate returns true if the password is valid.
func (g generator) IsValidate(password Password) bool {
	return regexp.
		MustCompile(g.regexString()).
		MatchString(password.String())
}

func (g generator) generate(length int) Password {
	buf := make([]rune, length)

	for i := range buf {
		buf[i] = g.letters[rand.N(g.size)] //nolint:gosec // use rand.N here because it is not safe to use rand.Intn
	}

	return Password(buf)
}

func (g generator) regexString() string {
	return "^[" + regexp.QuoteMeta(string(g.letters)) + "]+$"
}
