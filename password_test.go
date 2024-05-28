package password_test

import (
	"testing"

	"github.com/135yshr/password"
	"github.com/stretchr/testify/require"
)

func TestPasswordGeneratorGenerate(t *testing.T) {
	t.Parallel()

	type args struct {
		length int
	}
	type want struct {
		length  int
		pattern string
	}
	tests := map[string]struct {
		args args
		want want
	}{
		"5 characters string of lowercase": {
			args: args{
				length: 5,
			},
			want: want{
				length:  5,
				pattern: "[a-z]{5}",
			},
		},
	}
	for tn, tt := range tests {
		t.Run(tn, func(t *testing.T) {
			t.Parallel()

			sut := password.New()
			require.NotNil(t, sut)

			actual := sut.Generate(tt.args.length)
			t.Log("actual:", actual)

			t.Run("should return a string of the specified length", func(t *testing.T) {
				require.Len(t, actual, tt.want.length)
			})
			t.Run("should return a string of the specified pattern", func(t *testing.T) {
				require.Regexp(t, tt.want.pattern, actual)
			})
		})
	}
}
