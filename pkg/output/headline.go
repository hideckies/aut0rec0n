package output

import (
	"fmt"

	"github.com/fatih/color"
)

// Print the headline of result
func Headline(title string) {
	color.Cyan(TMPL_BAR_SINGLE_M)
	fmt.Printf("%s %s\n", color.GreenString("+"), title)
	color.Cyan(TMPL_BAR_SINGLE_M)
}
