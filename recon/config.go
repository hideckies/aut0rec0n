package recon

type Config struct {
	Host string

	Script []string

	Color     bool
	OutputDir string
	NoOutput  bool
	Quiet     bool
	Verbose   bool
}
