package get

import (
	"github.com/mahdibouaziz/cobra-playground/cmd/get/users"
	"github.com/spf13/cobra"
)

var GetCommand = &cobra.Command{
	Use:   "get",
	Short: "Get data based on an api",
	Long:  "CLI subcommand to fetch data based on an API",
}

func init() {
	GetCommand.AddCommand(users.UsersCommand)
}
