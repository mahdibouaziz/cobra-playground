package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate-pass",
	Short: "Random password generator",
	Long: `Random password generator built with Go
in order to learn how to use cobra`,
	Run: generatePassword,
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charset := "abcdefghtriksjdfdqjlkqndfJLHQJFHQSJKDFHQSDKFSQDFqfdkjfhqhiuiahrf"

	if isDigits {
		charset += "0123456789"
	}

	if isSpecialChars {
		charset += "$?#@!(){}[]-_*€<>="
	}

	password := make([]byte, length)
	for i := range password {
		password[i] += charset[rand.Intn(len(charset))]
	}

	fmt.Println("Your password is:", string(password))

}
