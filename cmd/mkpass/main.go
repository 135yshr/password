//nolint:forbidigo // Printf is used for command line interface
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/135yshr/password"
	"github.com/135yshr/password/policy"
)

const (
	defaultGenerateCount = 1
	defaultLength        = 12
)

var (
	version string
	commit  string
	date    string
	builtBy string
)

var (
	length int
	all    bool
	upper  bool
	lower  bool
	number bool
	symbol bool
	custom string
	help   bool
	ver    bool
)

func initFlags() {
	flag.IntVar(&length, "length", defaultLength, "length of password")
	flag.BoolVar(&all, "all", true, "use all characters")
	flag.BoolVar(&all, "a", true, "use all characters")
	flag.BoolVar(&upper, "upper", false, "use uppercase")
	flag.BoolVar(&upper, "u", false, "use uppercase")
	flag.BoolVar(&lower, "lower", false, "use lowercase")
	flag.BoolVar(&lower, "l", false, "use lowercase")
	flag.BoolVar(&number, "number", false, "use numbers")
	flag.BoolVar(&number, "n", false, "use numbers")
	flag.BoolVar(&symbol, "symbol", false, "use symbols")
	flag.BoolVar(&symbol, "s", false, "use symbols")
	flag.StringVar(&custom, "custom", "", "use custom characters")
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&ver, "version", false, "show version")
	flag.BoolVar(&ver, "v", false, "show version")
}

func updateFlagsBasedOnConditions() {
	if upper || lower || number || symbol {
		all = false
	}

	if custom != "" {
		all = false
		upper = false
		lower = false
		number = false
		symbol = false
	}
}

func main() {
	initFlags()
	flag.Parse()
	updateFlagsBasedOnConditions()

	// show help
	if help {
		printVersion()
		flag.Usage()
		os.Exit(0)
	}

	// show version
	if ver {
		printVersion()
		os.Exit(0)
	}

	count := defaultGenerateCount

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

func printVersion() {
	fmt.Printf("%s - generate passwords\n", os.Args[0])
	fmt.Printf("https://github.com/135yshr/password\n\n")
	fmt.Printf("Version:\t%s\n", version)
	fmt.Printf("Commit:\t\t%s\n", commit)
	fmt.Printf("BuildDate:\t%s\n", date)
	fmt.Printf("BuiltBy:\t%s\n", builtBy)
	fmt.Printf("GoVersion:\t%s\n", runtime.Version())
	fmt.Printf("GoOS:\t\t%s\n", runtime.GOOS)
	fmt.Printf("GoArch:\t\t%s\n", runtime.GOARCH)
	fmt.Println("\nMIT License (c) 2024 135yshr")
}

func createPolicies() []policy.Policy {
	policies := make([]policy.Policy, 0, 4) //nolint:mnd // 4 is the number of policies
	policies = createPolicy(policies, upper, policy.WithUppercase)
	policies = createPolicy(policies, lower, policy.WithLowercase)
	policies = createPolicy(policies, number, policy.WithNumbers)
	policies = createPolicy(policies, symbol, policy.WithSymbols)

	if custom != "" {
		policies = append(policies, policy.WithCustomString([]rune(custom)))
	}

	return policies
}

func createPolicy(policies []policy.Policy, create bool, policy policy.Policy) []policy.Policy {
	if create {
		return append(policies, policy)
	}

	return policies
}
