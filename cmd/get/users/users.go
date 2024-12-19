package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var UsersCommand = &cobra.Command{
	Use:   "users",
	Short: "Get users data",
	Long: `Get users data based on
https://dummyjson.com/users API`,
	Run: getUsers,
}

type user struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Gender    string
	Email     string
	Phone     string
}

type apiResponse struct {
	Users []user
	Total int
	Skip  int
	Limit int
}

const API = "https://dummyjson.com/users"

func getUsers(cmd *cobra.Command, args []string) {
	// Parsing flags
	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		fmt.Printf("Error parsing limit flag: %v\n", err)
		return
	}

	// Make the GET request
	resp, err := http.Get(API)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check for HTTP Error
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Http error %v\n", err)
		return
	}

	// Read the response body
	var apiResponse apiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		fmt.Printf("Error Parsing data %v\n", err)
		return
	}

	// fmt.Println("Total", apiResponse.Total)
	printTable(apiResponse.Users[:limit])
}

func printTable(users []user) {
	// Create a new tabwriter for formatted output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	// Print the header
	fmt.Fprintln(w, "ID\tFirst Name\tLast Name\tAge\tGender\tEmail")
	// Print user data
	for _, u := range users {
		fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%s\t%s\n", u.ID, u.FirstName, u.LastName, u.Age, u.Gender, u.Email)
	}
	// Flush the writer to output the table
	w.Flush()
}
