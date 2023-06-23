package progress

import (
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
)

type ProgressBar *progressbar.ProgressBar

func NewProgressBar(length int, description string) ProgressBar {
	bar := progressbar.NewOptions(
		length,
		progressbar.OptionSetWriter(os.Stdout),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription(fmt.Sprintf("[cyan][reset] %s", description)),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionClearOnFinish())
	return bar
}
