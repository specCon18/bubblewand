package render

import (
        "os"
        "path/filepath"
        "strings"
        "text/template"
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

// RenderTemplates renders all .tmpl files from the templates/ directory into outputDir
func RenderTemplates(data ProgramData, outputDir string, verbose bool) error {
        var renderedFiles int

        err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
                if err != nil {
                        return err
                }
                if info.IsDir() || !strings.HasSuffix(info.Name(), ".tmpl") {
                        return nil
                }

                relPath, err := filepath.Rel("templates", path)
                if err != nil {
                        return err
                }
                outputPath := filepath.Join(outputDir, strings.TrimSuffix(relPath, ".tmpl"))

                if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
                        return err
                }

                tmpl, err := template.ParseFiles(path)
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

