//nolint:forbidigo // Printf is used for command line interface
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/135yshr/password"
)

const defaultLength = 12

var (
	length int
	upper  bool
	lower  bool
	number bool
	symbol bool
	help   bool
)

func main() {
	flag.IntVar(&length, "length", defaultLength, "length of password")
	flag.BoolVar(&upper, "upper", true, "use uppercase")
	flag.BoolVar(&lower, "lower", true, "use lowercase")
	flag.BoolVar(&number, "number", true, "use numbers")
	flag.BoolVar(&symbol, "symbol", true, "use symbols")
	flag.BoolVar(&help, "help", false, "show help")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	count := 1

	if flag.NArg() > 0 {
		cn, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			fmt.Println("invalid count:", flag.Arg(0))
			flag.Usage()
			os.Exit(1)
		}

		count = cn
	}

	policies := createPolicies()
	gen := password.New(policies...)

	for range count {
		fmt.Println(gen.Generate(length))
	}
}

func createPolicies() []password.Policy {
	policies := make([]password.Policy, 0, 4) //nolint:mnd // 4 is the number of policies
	policies = createPolicy(policies, upper, password.WithUppercase)
	policies = createPolicy(policies, lower, password.WithLowercase)
	policies = createPolicy(policies, number, password.WithNumbers)
	policies = createPolicy(policies, symbol, password.WithSymbols)

	return policies
}

func createPolicy(policies []password.Policy, create bool, policy password.Policy) []password.Policy {
	if create {
		return append(policies, policy)
	}

	return policies
}
