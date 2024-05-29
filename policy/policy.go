package policy

import "regexp"

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	symbols   = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

var (
	// WithDefault is a default policy.
	WithDefault = NewPolicy(
		[]rune(lowercase+uppercase+numbers+symbols),
		regexpValidator{regexpString: regexp.QuoteMeta("a-zA-Z0-9!#$%&'\"()*+,-./:;<=>?@[\\]^_`{|}~")})

	// WithLowercase is a policy with lowercase.
	WithLowercase = NewPolicy([]rune(lowercase), regexpValidator{regexpString: "a-z"})

	// WithUppercase is a policy with uppercase.
	WithUppercase = NewPolicy([]rune(uppercase), regexpValidator{regexpString: "A-Z"})

	// WithNumbers is a policy with numbers.
	WithNumbers = NewPolicy([]rune(numbers), regexpValidator{regexpString: "0-9"})

	// WithSymbols is a policy with symbols.
	WithSymbols = NewPolicy(
		[]rune(symbols),
		regexpValidator{regexpString: regexp.QuoteMeta("a-zA-Z0-9!#$%&'\"()*+,-./:;<=>?@[\\]^_`{|}~")})

	// WithMinLength is a policy with minimum length.
	WithMinLength = func(min int) Policy {
		return NewPolicy([]rune{}, minLengthValidator{minLength: min})
	}

	// WithMaxLength is a policy with maximum length.
	WithMaxLength = func(max int) Policy {
		return NewPolicy([]rune{}, maxLengthValidator{maxLength: max})
	}

	// WithCustomString is a policy with custom string.
	WithCustomString = func(letters []rune) Policy {
		return NewPolicy(letters,
			regexpValidator{regexpString: regexp.QuoteMeta(string(letters))})
	}
)

// Validator is a validator.
type Validator interface {
	IsValid(password string) bool
}

type regexpValidator struct {
	regexpString string
}

// IsValid returns true if the password is valid.
func (v regexpValidator) IsValid(password string) bool {
	return regexp.MustCompile(`^[` + v.regexpString + `]+$`).MatchString(password)
}

type minLengthValidator struct {
	minLength int
}

// IsValid returns true if the password is valid.
func (v minLengthValidator) IsValid(password string) bool {
	return len(password) >= v.minLength
}

type maxLengthValidator struct {
	maxLength int
}

// IsValid returns true if the password is valid.
func (v maxLengthValidator) IsValid(password string) bool {
	return len(password) <= v.maxLength
}

// Policy is a password policy.
type Policy interface {
	Validator
	Letters() []rune
}

type policy struct {
	letters   []rune
	validator Validator
}

// Letters returns the letters.
func (p *policy) Letters() []rune {
	return p.letters
}

// IsValid returns true if the password is valid.
func (p *policy) IsValid(password string) bool {
	return p.validator.IsValid(password)
}

// NewPolicy returns a new policy.
func NewPolicy(letters []rune, validator Validator) Policy {
	return &policy{
		letters:   letters,
		validator: validator,
	}
}
