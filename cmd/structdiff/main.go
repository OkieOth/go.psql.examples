package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "structdiff",
	Short: "Tool to compare the structure of two Postgresql servers",
	Long: `This tool will parse the data dictionaries of two postgresql servers and compare
	the structure of the tables, views, indexes, sequences and functions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World ... from inside the root command!")
	},
}

func main() {
	rootCmd.Execute()
}
