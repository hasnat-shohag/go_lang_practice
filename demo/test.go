package test

import (
	"fmt"
	"strings"
)

// Formatter interface
type Formatter interface {
	Format([]string) string
}

// CSVFormatter struct implementing the Formatter interface
type CSVFormatter struct{}

func (c CSVFormatter) Format(data []string) string {
	return strings.Join(data, ", ")
}

// HTMLFormatter struct implementing the Formatter interface
type HTMLFormatter struct{}

func (h HTMLFormatter) Format(data []string) string {
	return "<ul><li>" + strings.Join(data, "</li><li>") + "</li></ul>"
}

// FormatData function to format data using a given Formatter
func FormatData(formatter Formatter, data []string) {
	switch f := formatter.(type) {
	case CSVFormatter:
		fmt.Println("Using CSVFormatter:")
		fmt.Println(f.Format(data))
	case HTMLFormatter:
		fmt.Println("Using HTMLFormatter:")
		fmt.Println(f.Format(data))
	default:
		fmt.Println("Unsupported formatter type")
	}
}

func main() {
	// Create instances of CSVFormatter and HTMLFormatter
	csvFormatter := CSVFormatter{}
	htmlFormatter := HTMLFormatter{}

	// Test the FormatData function with CSVFormatter
	dataCSV := []string{"John Doe", "30", "Software Engineer"}
	FormatData(csvFormatter, dataCSV)

	// Test the FormatData function with HTMLFormatter
	dataHTML := []string{"Apples", "Bananas", "Oranges"}
	FormatData(htmlFormatter, dataHTML)
}
