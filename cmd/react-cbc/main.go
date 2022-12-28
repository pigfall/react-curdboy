package main

import (
	"fmt"
	"github.com/pigfall/react-curdboy/mod/cbc"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var entSchemaSrcPath string
	var outputPath string
	cmd := cobra.Command{
		Use: "react-cbc",
		RunE: func(cmd *cobra.Command, args []string) error {
			g, err := (cbc.Factory{}).NewGenerator(entSchemaSrcPath, outputPath)
			if err != nil {
				return err
			}
			return g.Generate()
		},
		Example: `react-cbc --schemaSrcPath=${ENT_PROJECT}/ent --outputPath=${REACT_PROJECT}`,
	}
	cmd.PersistentFlags().StringVar(&entSchemaSrcPath, "schemaSrcPath", "", "ent schema directory path")
	cmd.MarkPersistentFlagRequired("schemaSrcPath")
	cmd.PersistentFlags().StringVar(&outputPath, "outputPath", "", "which diretory to save th generated files")
	cmd.MarkPersistentFlagRequired("outputPath")

	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed: %v", err)
	}
}
