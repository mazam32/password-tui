package cmd

import (
	"fmt"
	"password-manager/internal"

	"github.com/spf13/cobra"
)

func NewAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add [service] [password]",
		Short: "Add a new password",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			service, password := args[0], args[1]

			store, err := internal.NewDBManager()
			if err != nil {
				fmt.Println("Error loading passwords:", err)
				return
			}

			err = store.AddPassword(service, password)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Password added successfully!")
			}
		},
	}
}
