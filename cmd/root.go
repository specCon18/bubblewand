package cmd

import (
	"github.com/spf13/cobra"
	"specCon18/bubblewand/internal/render"
	"specCon18/bubblewand/internal/logger"
	"github.com/charmbracelet/log"

)

// CLI flag variables
var (
	modName        string
	packageName    string
	programVersion string
	programDesc    string
	outputDir      string 
	logLevel       string
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
			logger.Log.Fatalf("rendering failed: %v",err)
		}
	},
}



func init() {
	// Register flags
	rootCmd.Flags().StringVar(&modName, "mod-name", "", "Module name (e.g. github.com/user/app)")
	rootCmd.Flags().StringVar(&packageName, "package-name", "", "Package name")
	rootCmd.Flags().StringVar(&programVersion, "program-version", "", "Program version")
	rootCmd.Flags().StringVar(&programDesc, "program-desc", "", "Program description")
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", "output", "Output directory for rendered files")
        rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Log level (debug, info, warn, error)")
	// Mark required
	rootCmd.MarkFlagRequired("mod-name")
	rootCmd.MarkFlagRequired("package-name")
	rootCmd.MarkFlagRequired("program-version")
	rootCmd.MarkFlagRequired("program-desc")
}

func initLogging() {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		logger.Log.Warn("Invalid log level; defaulting to info", "input", logLevel)
		level = log.InfoLevel
	}
	logger.Log.SetLevel(level)
}

// Execute starts the CLI application
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
