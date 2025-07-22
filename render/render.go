package render

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ProgramData holds user-supplied template values
type ProgramData struct {
	ModName        string
	PackageName    string
	ProgramVersion string
	ProgramDesc    string
}

// RenderTemplates renders all .tmpl files from the templates/ directory into outputDir
func RenderTemplates(data ProgramData, outputDir string) error {
	return filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".tmpl") {
			return nil
		}

		// Create relative output path (preserving subdirs)
		relPath, err := filepath.Rel("templates", path)
		if err != nil {
			return err
		}
		outputPath := filepath.Join(outputDir, strings.TrimSuffix(relPath, ".tmpl"))

		// Ensure parent directories exist
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return err
		}

		// Parse and execute template
		tmpl, err := template.ParseFiles(path)
		if err != nil {
			return err
		}

		outFile, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		log.Printf("Rendering %s → %s\n", path, outputPath)
		return tmpl.Execute(outFile, data)
	})
}

