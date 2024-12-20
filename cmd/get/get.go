package get

import (
	"github.com/mahdibouaziz/cobra-playground/cmd/get/users"
	"github.com/spf13/cobra"
)

var GetCommand = &cobra.Command{
	Use:        "get",
	Short:      "Get data based on an api",
	Long:       "CLI subcommand to fetch data based on an API",
	Aliases:    []string{"fetch"},
	SuggestFor: []string{"list"},
	Example: `  # Get the first 10 users
  cobra-playground get users
  # Get the first 20 users 
  cobra-playground get users -l 20
  # Get users between 30 and 45
  cobra-playground get users -b 30,45
  # Get users wide format
  cobra-playground get users -o wide
	`,
}

func init() {
	GetCommand.AddCommand(users.UsersCommand)
}
