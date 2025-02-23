package cmd

import (
	"fmt"
	"password-manager/internal"

	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all stored services",
		Run: func(cmd *cobra.Command, args []string) {
			store, err := internal.NewDBManager()
			if err != nil {
				fmt.Println("Error loading passwords:", err)
				return
			}

			services, err := store.ListServices()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if len(services) == 0 {
				fmt.Println("No passwords stored.")
			} else {
				fmt.Println("Stored services:")
				for _, service := range services {
					fmt.Println("-", service)
				}
			}
		},
	}
}
