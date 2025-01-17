package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Build all HTML content into the _gen directory
	err := buildContent()
	if err != nil {
		log.Fatalf("Error building content: %v", err)
	}

	log.Println("Build completed successfully!")
}

func buildContent() error {
	// Create output directory if it doesn't exist
	err := os.MkdirAll("_gen", os.ModePerm)
	if err != nil {
		return err
	}

	// Walk the directory to process all .html files
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the _templates folder and non-HTML files
		if info.IsDir() || !strings.HasSuffix(path, ".html") || strings.HasPrefix(path, "_templates") || strings.HasPrefix(path, "_gen") {
			return nil
		}

		// Generate HTML for each .html file found
		return generateHTML(path)
	})

	if err != nil {
		return err
	}

	return nil
}

// generateHTML will execute the template and write the output to the _gen directory
func generateHTML(filename string) error {
	// Parse the base template (_templates/baseof.html) and the content template (e.g., index.html)
	baseTemplatePath := filepath.Join("_templates", "baseof.html")
	tmplFiles := []string{baseTemplatePath, filename}

	// Collect all partials in the _templates/partials directory
	partialsDir := filepath.Join("_templates", "partials")
	err := filepath.Walk(partialsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only add .html files from the partials folder
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			tmplFiles = append(tmplFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Parse all templates (base + content + partials)
	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		return err
	}

	// Determine the output file path in the _gen folder
	outputFilePath := strings.TrimPrefix(filename, ".")      // Remove leading dot from the path
	outputFilePath = strings.TrimPrefix(outputFilePath, "/") // Remove leading slash from the path

	// Ensure that the parent directories for the output file exist
	outputDir := filepath.Dir(outputFilePath)
	err = os.MkdirAll("_gen/"+outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Create the output file in the _gen directory
	outputFile, err := os.Create("_gen/" + outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Execute the template (base template + content + partials)
	err = tmpl.Execute(outputFile, nil)
	if err != nil {
		return err
	}

	return nil
}
