package password_test

import (
	"regexp"
	"testing"

	"github.com/135yshr/password"
	"github.com/135yshr/password/policy"
	"github.com/stretchr/testify/require"
)

func TestPasswordGeneratorGenerate(t *testing.T) {
	t.Parallel()

	type args struct {
		length   int
		policies []policy.Policy
	}

	type want struct {
		length  int
		pattern string
	}

	tests := map[string]struct {
		args args
		want want
	}{
		"5 characters string of not set": {
			args: args{
				length:   5,
				policies: nil,
			},
			want: want{
				length:  5,
				pattern: "^[" + regexp.QuoteMeta("a-zA-Z0-9!#$%&'\"()*+,-./:;<=>?@[\\]^_`{|}~") + "]{5}$",
			},
		},
		"5 characters string of lowercase": {
			args: args{
				length: 5,
				policies: []policy.Policy{
					policy.WithLowercase,
				},
			},
			want: want{
				length:  5,
				pattern: "[a-z]{5}",
			},
		},
		"5 characters string of uppercase": {
			args: args{
				length: 5,
				policies: []policy.Policy{
					policy.WithUppercase,
				},
			},
			want: want{
				length:  5,
				pattern: "[A-Z]{5}",
			},
		},
		"5 characters string of lowercase and uppercase": {
			args: args{
				length: 5,
				policies: []policy.Policy{
					policy.WithLowercase,
					policy.WithUppercase,
				},
			},
			want: want{
				length:  5,
				pattern: "[a-zA-Z]{5}",
			},
		},
		"5 characters string of lowercase and uppercase and numbers": {
			args: args{
				length: 5,
				policies: []policy.Policy{
					policy.WithLowercase,
					policy.WithUppercase,
					policy.WithNumbers,
				},
			},
			want: want{
				length:  5,
				pattern: "[a-zA-Z0-9]{5}",
			},
		},
		"30 characters string of lowercase and uppercase and numbers and symbols": {
			args: args{
				length: 30,
				policies: []policy.Policy{
					policy.WithLowercase,
					policy.WithUppercase,
					policy.WithNumbers,
					policy.WithSymbols,
				},
			},
			want: want{
				length:  30,
				pattern: "^[" + regexp.QuoteMeta("a-zA-Z0-9!#$%&'\"()*+,-./:;<=>?@[\\]^_`{|}~") + "]{30}$",
			},
		},
		"5 characters string of lowercase and min length 5 and max length 10": {
			args: args{
				length: 5,
				policies: []policy.Policy{
					policy.WithLowercase,
					policy.WithMinLength(5),
					policy.WithMaxLength(10),
				},
			},
			want: want{
				length:  5,
				pattern: "^[" + regexp.QuoteMeta("a-zA-Z0-9!#$%&'\"()*+,-./:;<=>?@[\\]^_`{|}~") + "]{5,10}$",
			},
		},
		"6 characters string of lowercase and min length 5 and max length 10": {
			args: args{
				length: 6,
				policies: []policy.Policy{
					policy.WithLowercase,
					policy.WithMinLength(5),
					policy.WithMaxLength(10),
				},
			},
			want: want{
				length:  6,
				pattern: "^[a-z]{5,10}$",
			},
		},
		"custom characters string": {
			args: args{
				length: 5,
				policies: []policy.Policy{
					policy.WithCustomString([]rune("abcdef0123456789")),
				},
			},
			want: want{
				length:  5,
				pattern: "^[abcdef0123456789]{5}$",
			},
		},
	}

	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			sut := password.New(tt.args.policies...)
			require.NotNil(t, sut)

			actual := sut.Generate(tt.args.length)

			t.Run("should return a string of the specified length", func(t *testing.T) {
				require.Len(t, actual, tt.want.length)
			})

			t.Run("should return a string of the specified pattern", func(t *testing.T) {
				require.Regexp(t, tt.want.pattern, actual)
			})
		})
	}
}
