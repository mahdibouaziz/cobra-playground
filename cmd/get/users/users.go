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
	BirthDate string
}

type apiResponse struct {
	Users []user
	Total int
	Skip  int
	Limit int
}

const API = "https://dummyjson.com/users"

func getUsers(cmd *cobra.Command, args []string) {
	// Limit flag
	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		fmt.Printf("Error parsing limit flag: %v\n", err)
		return
	}

	// Between flag
	between, err := cmd.Flags().GetIntSlice("between")
	if err != nil {
		fmt.Printf("Error parsing between flag: %v\n", err)
		return
	}

	if len(between) != 0 && len(between) != 2 {
		fmt.Printf("between flag accepts exactly 2 parameters")
		return
	}

	// Output flag
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		fmt.Printf("Error parsing output flag: %v\n", err)
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

	if len(between) != 0 {
		printTable(apiResponse.Users[between[0]:between[1]], output)
	} else {
		printTable(apiResponse.Users[:limit], output)
	}
}

func printTable(users []user, outputFormat string) {
	// Create a new tabwriter for formatted output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	// Print the header
	if outputFormat == "wide" {
		fmt.Fprintln(w, "ID\tFirst Name\tLast Name\tAge\tGender\tEmail\tPhone\tBirthDate")
		// Print user data
		for _, u := range users {
			fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%s\t%s\t%s\t%s\n", u.ID, u.FirstName, u.LastName, u.Age, u.Gender, u.Email, u.Phone, u.BirthDate)
		}
	} else {
		fmt.Fprintln(w, "ID\tFirst Name\tLast Name\tAge\tGender\tEmail")
		// Print user data
		for _, u := range users {
			fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%s\t%s\n", u.ID, u.FirstName, u.LastName, u.Age, u.Gender, u.Email)
		}
	}

	// Flush the writer to output the table
	w.Flush()
}
