package cmd

import (
	"fmt"

	"github.com/go-hasaki/hasaki/cmd/command"
	"github.com/go-hasaki/hasaki/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "hasaki",
	Example: "hasaki new demo-api",
	Short:   "\n _   _    __    ___    __    _  _  ____ \n( )_( )  /__\\  / __)  /__\\  ( )/ )(_  _)\n ) _ (  /(__)\\ \\__ \\ /(__)\\  )  (  _)(_ \n(_) (_)(__)(__)(___/(__)(__)(_)\\_)(____) ",
	Version: fmt.Sprintf("Hasaki %s - Copyright (c) 2023 Hasaki\nReleased under the Apache-2.0 License.\n\n", version.Version),
}

func init() {
	rootCmd.AddCommand(command.NewCmd)
	rootCmd.AddCommand(command.WireCmd)
	rootCmd.AddCommand(command.UpgradeCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
