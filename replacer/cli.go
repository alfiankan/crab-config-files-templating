package replacer

import (
	"fmt"

	"github.com/spf13/cobra"
)

// color TUI
const (
	greenTemplate = "\033[1;32m%s\033[0m"
	redTemplate   = "\033[1;31m%s\033[0m"
)

// RootCLI main cli command
func RootCLI() *cobra.Command {
	var replaceKV []string
	var replaceWithStringKV []string
	var replacableKV []ReplacableKV

	rootCmd := &cobra.Command{
		Use:   "crab",
		Short: "Crab CLI config files templating",
		Long:  `Crab CLI Dynamic configuration file templating tool for kubernetes manifest or general configuration files`,
		Run: func(cmd *cobra.Command, args []string) {

			var inputFilePath string
			var outputFilePath string
			var verbose bool
			var errFlag error

			inputFilePath, errFlag = cmd.Flags().GetString("file")
			outputFilePath, errFlag = cmd.Flags().GetString("output")
			verbose, errFlag = cmd.Flags().GetBool("verbose")

			if errFlag != nil {
				fmt.Printf(redTemplate+" Something went wrong %s \n", "[ERROR]", errFlag.Error())
			}

			if outputFilePath == "" {
				outputFilePath = inputFilePath
			}
			for _, v := range replaceKV {
				replacableKV = append(replacableKV, ReplacableKV{
					KV:     v,
					Quotes: false,
				})
			}
			for _, v := range replaceWithStringKV {
				replacableKV = append(replacableKV, ReplacableKV{
					KV:     v,
					Quotes: true,
				})
			}

			replacer := NewReplacer(inputFilePath, outputFilePath, false)
			if err := replacer.Run(replacableKV, verbose); err != nil {
				fmt.Printf(redTemplate+" Something went wrong %s \n", "[ERROR]", err.Error())
			}
			if verbose {
				fmt.Printf(greenTemplate+" Crab output result at %s \n", "[DONE]", outputFilePath)
			}

		},
	}

	// set flag input
	// grep flag input
	rootCmd.Flags().StringP("file", "f", "", "Input File Path")
	rootCmd.Flags().StringP("output", "o", "", "Output File Path")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose process")
	rootCmd.Flags().StringSliceVarP(&replaceKV, "replace", "r", []string{}, "key value replacable eg. -r namespace=default")
	rootCmd.Flags().StringSliceVarP(&replaceWithStringKV, "quotes", "q", []string{}, "replace including quotes eg. -q url=http://domain.com result 'http://domain.com'")

	return rootCmd
}
