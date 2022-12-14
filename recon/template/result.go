package template

import (
	"fmt"

	"github.com/fatih/color"
)

func CreateResultWithTemplate(title string, contents string, colorized bool) []string {
	results := []string{}

	result := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", BarDoubleM, title, BarDoubleM, contents, BarDoubleM)
	resultColor := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", BarDoubleMColor, color.HiCyanString(title), BarDoubleMColor, contents, BarDoubleMColor)

	results = append(results, result)
	results = append(results, resultColor)

	return results
}
