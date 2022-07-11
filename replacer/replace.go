package replacer

import (
	"fmt"
	"os"
	"strings"
)

// ReplacableKV hold key and value
// key represent target template to be replaced
// value represent next value
type ReplacableKV struct {
	KV     string
	Quotes bool
}

// Replacer struct object for replace
// inputPath => input file path
// dryRun run without saving result
type Replacer struct {
	inputPath  string
	outputPath string
	dryRun     bool
}

// NewReplacer create new Replacer
func NewReplacer(inputPath, outputPath string, dryRun bool) *Replacer {
	return &Replacer{
		inputPath,
		outputPath,
		dryRun,
	}
}

func (rep *Replacer) readFile() (string, error) {
	content, err := os.ReadFile(rep.inputPath)
	return string(content), err
}

// Run running replace process
// will replace {{templateName}} format
func (rep *Replacer) Run(kv []ReplacableKV, verbose bool) error {

	content, err := rep.readFile()
	if err != nil {
		return err
	}

	// subtituting Key Value replacing template {{templateName}}
	for _, rk := range kv {
		v := strings.Split(rk.KV, "=")
		newValue := v[1]
		if rk.Quotes {
			newValue = fmt.Sprintf(`"%s"`, v[1])
		}
		content = strings.ReplaceAll(content, fmt.Sprintf("{{%s}}", v[0]), newValue)
		if verbose {
			fmt.Printf(greenTemplate+" from %s to %s \n", "[REPLACED]", v[0], newValue)
		}
	}

	if rep.dryRun {
		os.Stdout.Write([]byte(content))
	} else {
		if err := os.WriteFile(rep.outputPath, []byte(content), 0644); err != nil {
			return err
		}
	}
	return nil
}
