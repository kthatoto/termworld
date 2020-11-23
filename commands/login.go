package commands

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
	"bytes"
	"net/http"
	"errors"

	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(loginCommand)
}

var email string

var loginCommand = &cobra.Command{
	Use: "login",
	Short: "Login command",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(email) == 0 {
			fmt.Print("Please enter email: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			email = string(scanner.Text())
		}

		type RequestBody struct {
			Email string `json:"email"`
		}
		json, _ := json.Marshal(RequestBody{email})
		u := "http://localhost:8080/login/new"
		resp, err := http.Post(u, "application/json", bytes.NewBuffer(json))
		if err != nil {
			return err
		}
		if (resp.StatusCode != 201) {
			return errors.New("Request failed")
		}
		fmt.Printf("Sent mail to %s\nPlease check and click link on the mail to verify\n", email)

		u = "http://localhost:8080/login"
		var token string
		for {
			resp, err := http.Post(u, "application/json", bytes.NewBuffer(json))
			if err != nil {
				return err
			}
			if resp.StatusCode == 408 {
				continue
			} else if resp.StatusCode == 200 {
				break
			} else {
				return errors.New("Request failed")
			}
		}
		fmt.Println("Login finished!")

		return nil
	},
}

func init() {
	loginCommand.PersistentFlags().StringVar(&email, "email", "", "Email to login")
}
