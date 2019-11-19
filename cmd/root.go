package cmd

import "github.com/spf13/cobra"

func Root() *cobra.Command {

	command := &cobra.Command{
		Use:   "mask",
		Short: "a mask tool for encrypt / decrypt strings",
	}

	return command
}
