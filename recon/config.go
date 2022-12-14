package recon

type Config struct {
	Host string

	Script []string

	Color     bool
	OutputDir string
	Quiet     bool
	Verbose   bool
}
