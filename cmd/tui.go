package cmd

import (
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"specCon18/bubblewand/internal/render"
	"specCon18/bubblewand/internal/logger"
)

// tuiCmd renders templates interactively via a form
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "an interactive TUI for generating a go module template for terminal applications using bubbletea + cobra + viper + log",
	Run: func(cmd *cobra.Command, args []string) {
		var data render.ProgramData

		// Build the input form
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title("Module Name").Placeholder("e.g. github.com/you/project").Value(&data.ModName),
				huh.NewInput().Title("Package Name").Placeholder("e.g. myapp").Value(&data.PackageName),
				huh.NewInput().Title("Program Version").Placeholder("e.g. 1.0.0").Value(&data.ProgramVersion),
				huh.NewInput().Title("Program Description").Placeholder("Describe your program").Value(&data.ProgramDesc),
				huh.NewInput().Title("Output Directory").Placeholder("Where to output your templated project").Value(&data.OutputDir),
			),
		)

		// Run the form and exit on cancel
		if err := form.Run(); err != nil {
			logger.Log.Fatalf("form cancelled or failed: %v", err)
		}

		// Render templates with user input
		if err := render.RenderTemplates(data, data.OutputDir, verbose); err != nil {
			logger.Log.Fatalf("rendering failed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}

