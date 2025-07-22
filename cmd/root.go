package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"specCon18/bubblewand/render"
)

// CLI flag variables
var (
	modName        string
	packageName    string
	programVersion string
	programDesc    string
	outputDir      string // NEW: output directory flag
)

// rootCmd renders templates using CLI flags
var rootCmd = &cobra.Command{
	Use:   "bubblewand",
    	Short: "A tool to generate a go project template for building a terminal application with bubbletea + cobra + viper + log",
	Run: func(cmd *cobra.Command, args []string) {
		// Fill ProgramData from CLI input
		data := render.ProgramData{
			ModName:        modName,
			PackageName:    packageName,
			ProgramVersion: programVersion,
			ProgramDesc:    programDesc,
		}

		// Render templates to the specified output directory
		if err := render.RenderTemplates(data, outputDir); err != nil {
			log.Fatalf("rendering failed: %v", err)
		}
	},
}

// Execute starts the CLI application
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Register flags
	rootCmd.Flags().StringVar(&modName, "mod-name", "", "Module name (e.g. github.com/user/app)")
	rootCmd.Flags().StringVar(&packageName, "package-name", "", "Package name")
	rootCmd.Flags().StringVar(&programVersion, "program-version", "", "Program version")
	rootCmd.Flags().StringVar(&programDesc, "program-desc", "", "Program description")
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", "output", "Output directory for rendered files")

	// Mark required
	rootCmd.MarkFlagRequired("mod-name")
	rootCmd.MarkFlagRequired("package-name")
	rootCmd.MarkFlagRequired("program-version")
	rootCmd.MarkFlagRequired("program-desc")
}

