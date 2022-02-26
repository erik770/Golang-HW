package uniq

import "flag"

type Flags struct {
	c bool
	d bool
	u bool
	i bool
	f int
	s int
}

func myParseFlags() (input, output string, flags Flags) {
	flag.BoolVar(&flags.c, "c", false, "Count number of repeats")
	flag.BoolVar(&flags.d, "d", false, "Only duplicate strings")
	flag.BoolVar(&flags.u, "u", false, "Only unique strings")
	flag.BoolVar(&flags.i, "i", false, "Ignore register")
	flag.IntVar(&flags.f, "f", -1, "Skip first num_fields")
	flag.IntVar(&flags.s, "s", -1, "Skip first num_chars in sting")
	flag.Parse()
	input = flag.Arg(0)
	output = flag.Arg(1)
	return input, output, flags
}
