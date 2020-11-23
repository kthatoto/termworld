package commands

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
	"bytes"
	"net/http"
	"errors"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		param, _ := json.Marshal(RequestBody{email})
		u := "http://localhost:8080/login/new"
		resp, err := http.Post(u, "application/json", bytes.NewBuffer(param))
		if err != nil {
			return err
		}
		if (resp.StatusCode != 201) {
			return errors.New("Request failed")
		}
		fmt.Printf("Sent mail to %s\nPlease check and click link on the mail to verify\n", email)

		u = "http://localhost:8080/login"
		for {
			resp, err := http.Post(u, "application/json", bytes.NewBuffer(param))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode == 408 {
				continue
			} else if resp.StatusCode == 200 {
				bytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				type ResponseBody struct {
					Token string `json:"token"`
				}
				var responseBody ResponseBody
				if err := json.Unmarshal(bytes, &responseBody); err != nil {
					return err
				}
				viper.Set("token", responseBody.Token)
				viper.WriteConfig()
				break
			} else {
				return errors.New("Request failed")
			}
		}
		fmt.Println("Login success!")

		return nil
	},
}

func init() {
	loginCommand.PersistentFlags().StringVar(&email, "email", "", "Email to login")
}
