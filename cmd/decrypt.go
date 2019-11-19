package cmd

import (
	"fmt"
	"log"

	"github.com/yaien/mask/crypto"

	"github.com/spf13/cobra"
)

// Decrypt get decrypt command
func Decrypt() *cobra.Command {
	var offset int

	command := &cobra.Command{
		Use:   "decrypt [string]",
		Short: "print decrypted string",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			message := args[0]
			result, err := crypto.Decrypt(message, offset)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
		},
	}

	command.Flags().IntVar(&offset, "offset", 1, "offset steps to rotate message")

	return command
}
