package cmd

import (
	"fmt"
	"os"

	"github.com/mahdibouaziz/cobra-playground/cmd/get"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra-playground",
	Short: "Cobra CLI Playground ",
	Long: `A CLI practice cobra and know how to deal with it
Cobra is really an amazing CLI tool`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// Do stuff here
	// 	fmt.Println("ARGS", args)
	// },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("length", "l", 8, "length of the password")
	generateCmd.Flags().BoolP("digits", "d", false, "include digits in the generated password")
	generateCmd.Flags().BoolP("special-chars", "s", false, "include special characters in the generated password")

	rootCmd.AddCommand(get.GetCommand)
	get.GetCommand.PersistentFlags().IntP("limit", "l", 10, "Users limits")
	get.GetCommand.PersistentFlags().IntSliceP("between","b",nil,"Get users between 2 values")
	get.GetCommand.MarkFlagsMutuallyExclusive("limit","between")
}
