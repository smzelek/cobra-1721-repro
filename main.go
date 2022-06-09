package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := NewRootCmd()
	rootCmd.AddCommand(Level2Cmd())
	rootCmd.SilenceErrors = true
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// Disable Cobra's default behavior to print command help text on run-time errors.
		cmd.SilenceUsage = true
	}

	err := rootCmd.Execute()
	fmt.Println(err)
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "top-level",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Args: ExactValidArgs(1),
	}
	return cmd
}

func Level2Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "level-2",
		ValidArgs: []string{
			"check-key",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Args: ExactValidArgs(1),
	}
	cmd.AddCommand(Level3Cmd())
	return cmd
}

func Level3Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "check-key",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	return cmd
}

// *****************

func ExactValidArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if err := ExactArgs(n)(cmd, args); err != nil {
			return err
		}
		return OnlyValidArgs(cmd, args)
	}
}

func ExactArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
		}
		return nil
	}
}

func OnlyValidArgs(cmd *cobra.Command, args []string) error {
	if len(cmd.ValidArgs) > 0 {
		var validArgs []string
		for _, v := range cmd.ValidArgs {
			validArgs = append(validArgs, strings.Split(v, "\t")[0])
		}

		for _, v := range args {
			if !stringInSlice(v, validArgs) {
				fmt.Println("HERE:", v, args[0])
				return fmt.Errorf("invalid argument %q for %q%s", v, cmd.CommandPath(), findSuggestions(cmd, args[0]))
			}
		}
	}
	return nil
}

func findSuggestions(c *cobra.Command, arg string) string {
	if c.DisableSuggestions {
		return ""
	}
	if c.SuggestionsMinimumDistance <= 0 {
		c.SuggestionsMinimumDistance = 2
	}
	suggestionsString := ""
	if suggestions := c.SuggestionsFor(arg); len(suggestions) > 0 {
		suggestionsString += "\n\nDid you mean this?\n"
		for _, s := range suggestions {
			suggestionsString += fmt.Sprintf("\t%v\n", s)
		}
	}
	return suggestionsString
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
