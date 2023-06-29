package main

import (
	"fmt"
	"github.com/openziti/zrok/tui"
	"github.com/openziti/zrok/zrokdir"
	"github.com/spf13/cobra"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(newConsoleCommand().cmd)
}

type consoleCommand struct {
	cmd *cobra.Command
}

func newConsoleCommand() *consoleCommand {
	cmd := &cobra.Command{
		Use:   "console",
		Short: "Open the web console",
		Args:  cobra.ExactArgs(0),
	}
	command := &consoleCommand{cmd}
	cmd.Run = command.run
	return command
}

func (cmd *consoleCommand) run(_ *cobra.Command, _ []string) {
	zrd, err := zrokdir.Load()
	if err != nil {
		tui.Error("unable to load zrokdir", err)
	}

	apiEndpoint, _ := zrd.ApiEndpoint()
	if err := exec.Command("open", apiEndpoint).Run(); err != nil {
		tui.Error(fmt.Sprintf("unable to open '%v'", apiEndpoint), err)
	}
}
