package render

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"specCon18/bubblewand/embed"
	"specCon18/bubblewand/internal/logger"
)

// ProgramData holds user-supplied template values
type ProgramData struct {
	ModName        string
	PackageName    string
	ProgramVersion string
	ProgramDesc    string
	OutputDir      string
}

// RenderTemplates renders embedded .tmpl files into outputDir
func RenderTemplates(data ProgramData, outputDir string, verbose bool) error {
	var renderedFiles int

	err := fs.WalkDir(embed.Templates, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".tmpl") {
			return nil
		}

		// Get relative path inside templates/
		relPath := strings.TrimPrefix(path, "templates/")
		outputPath := filepath.Join(outputDir, strings.TrimSuffix(relPath, ".tmpl"))

		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return err
		}

		tmplBytes, err := embed.Templates.ReadFile(path)
		if err != nil {
			return err
		}

		tmpl, err := template.New(d.Name()).Parse(string(tmplBytes))
		if err != nil {
			return err
		}

		outFile, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		if verbose {
			logger.Log.Infof("Rendering %s → %s", path, outputPath)
		}

		renderedFiles++
		return tmpl.Execute(outFile, data)
	})

	if err != nil {
		return err
	}

	if !verbose && renderedFiles > 0 {
		logger.Log.Info("Rendering templates")
	}

	return nil
}

