package output

import (
	"fmt"

	progressbar "github.com/schollz/progressbar/v3"
)

type ProgressBar progressbar.ProgressBar

func NewProgressBar(max int, desc string) progressbar.ProgressBar {
	return *progressbar.NewOptions(max,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(10),
		progressbar.OptionShowCount(),
		progressbar.OptionSetDescription(fmt.Sprintf("[cyan]%s[reset]", desc)),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionClearOnFinish())
}
