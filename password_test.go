package password_test

import (
	"regexp"
	"testing"

	"github.com/135yshr/password"
	"github.com/stretchr/testify/require"
)

func TestPasswordGeneratorGenerate(t *testing.T) {
	t.Parallel()

	type args struct {
		length   int
		policies []password.Policy
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
				policies: []password.Policy{
					password.WithLowercase,
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
				policies: []password.Policy{
					password.WithUppercase,
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
				policies: []password.Policy{
					password.WithLowercase,
					password.WithUppercase,
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
				policies: []password.Policy{
					password.WithLowercase,
					password.WithUppercase,
					password.WithNumbers,
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
				policies: []password.Policy{
					password.WithLowercase,
					password.WithUppercase,
					password.WithNumbers,
					password.WithSymbols,
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
				policies: []password.Policy{
					password.WithLowercase,
					password.WithMinLength(5),
					password.WithMaxLength(10),
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
				policies: []password.Policy{
					password.WithLowercase,
					password.WithMinLength(5),
					password.WithMaxLength(10),
				},
			},
			want: want{
				length:  6,
				pattern: "^[a-z]{5,10}$",
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
