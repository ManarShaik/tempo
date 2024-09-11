package FlagsCollection

import(
	"flag"
)
func JustifyFlag()string{
	align := flag.String("align", "left", "printing color")
	return *align
}