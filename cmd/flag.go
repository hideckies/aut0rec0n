package cmd

type Flag struct {
	Host string

	// Recon scripts
	Script          []string
	PrintScriptList bool

	Color     bool
	OutputDir string
	Quiet     bool
	Verbose   bool

	PrintVersion bool
}
