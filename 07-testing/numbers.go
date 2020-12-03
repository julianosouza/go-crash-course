package numbers

// Sum returns the sum of all the arguments passed in
// notice that here we're using a variadic argument. This means
// that the var `args` can accept any number of ints.
func Sum(args ...int) int {
	total := 0

	for _, v := range args {
		total += v
	}

	return total
}
