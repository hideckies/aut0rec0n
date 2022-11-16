package cmd

type Flag struct {
	Host string

	// Recon scripts
	Script []string

	Color     bool
	OutputDir string
	Verbose   bool
}
