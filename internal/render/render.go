package render

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"specCon18/bubblewand/embed"           // Import embedded template files
	"specCon18/bubblewand/internal/logger" // Import logger for logging output
)

// ProgramData holds user-supplied values for template substitution.
type ProgramData struct {
	ModName        string // Module name
	PackageName    string // Go package name
	ProgramVersion string // Version string
	ProgramDesc    string // Description of the program
	OutputDir      string // Target output directory for rendered files
}

// RenderTemplates renders embedded .tmpl files into outputDir
func RenderTemplates(data ProgramData, outputDir string, verbose bool) error {
	var renderedFiles int // Count of successfully rendered files

	// Walk through the embedded templates filesystem starting at "templates"
	err := fs.WalkDir(embed.Templates, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// Log and return error encountered while walking the filesystem
			logger.Log.Errorf("Error walking path %q: %v", path, err)
			return err
		}

		// Skip directories and files that do not end with ".tmpl"
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".tmpl") {
			return nil
		}

		// Strip the "templates/" prefix from path and remove ".tmpl" extension
		relPath := strings.TrimPrefix(path, "templates/")
		outputPath := filepath.Join(outputDir, strings.TrimSuffix(relPath, ".tmpl"))

		// Ensure the parent directory for the output file exists
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			logger.Log.Errorf("Failed to create directory for %s: %v", outputPath, err)
			return err
		}

		// Read the embedded template file
		tmplBytes, err := embed.Templates.ReadFile(path)
		if err != nil {
			logger.Log.Errorf("Failed to read template %s: %v", path, err)
			return err
		}

		// Parse the template content
		tmpl, err := template.New(d.Name()).Parse(string(tmplBytes))
		if err != nil {
			logger.Log.Errorf("Failed to parse template %s: %v", path, err)
			return err
		}

		// Create the output file for writing the rendered content
		outFile, err := os.Create(outputPath)
		if err != nil {
			logger.Log.Errorf("Failed to create output file %s: %v", outputPath, err)
			return err
		}
		defer outFile.Close()

		// Log the render operation if verbose mode is enabled
		if verbose {
			logger.Log.Infof("Rendering %s → %s", path, outputPath)
		}

		// Execute the template using the provided data and write to file
		if err := tmpl.Execute(outFile, data); err != nil {
			logger.Log.Errorf("Failed to execute template %s: %v", path, err)
			return err
		}

		renderedFiles++
		return nil
	})

	// If any error occurred during the walk/render process, log and return it
	if err != nil {
		logger.Log.Errorf("Template rendering failed: %v", err)
		return err
	}

	// If not verbose and at least one file was rendered, print a summary log
	if !verbose && renderedFiles > 0 {
		logger.Log.Info("Rendering templates")
	}

	return nil // Success
}

