package pepcli

// holds the configuration for all flags

import "github.com/urfave/cli/v2"

var Flags = []cli.Flag{
	&cli.StringSliceFlag{
		Name:        "files",
		Aliases:     []string{"f"},
		Usage:       "Specifiy fasta files",
		DefaultText: "data.fasta",
	},
	&cli.Float64Flag{
		Name:        "start",
		Aliases:     []string{"s"},
		Usage:       "Specify start of mass range",
		Required:    true,
		Destination: &minRange,
	},
	&cli.Float64Flag{
		Name:        "end",
		Aliases:     []string{"e"},
		Usage:       "Specify end of mass range",
		Required:    true,
		Destination: &maxRange,
	},
	&cli.BoolFlag{
		Name:     "wcomments",
		Aliases:  []string{"wc"},
		Usage:    "Use this if you want the mass as a comment between seq ID and peptide seq",
		Value:    false,
		Required: false,
	},
}
