package cmd

type Flag struct {
	Host string

	// Recon scripts
	Script          []string
	PrintScriptList bool

	Color     bool
	OutputDir string
	NoOutput  bool
	Quiet     bool
	Verbose   bool

	PrintVersion bool
}
