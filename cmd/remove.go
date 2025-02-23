package cmd

import (
	"fmt"
	"password-manager/internal"

	"github.com/spf13/cobra"
)

func NewRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "remove [service]",
		Short: "Delete a stored password",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			service := args[0]

			store, err := internal.NewDBManager()
			if err != nil {
				fmt.Println("Error loading passwords:", err)
				return
			}

			err = store.RemovePassword(service)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Password removed successfully!")
			}
		},
	}
}
