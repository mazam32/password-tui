package cmd

import (
	"fmt"
	"password-manager/internal"

	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "get [service]",
		Short: "Retrieve a stored password",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			service := args[0]

			store, err := internal.NewDBManager()
			if err != nil {
				fmt.Println("Error loading passwords:", err)
				return
			}

			password, err := store.GetPassword(service)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Password:", password)
			}
		},
	}
}
