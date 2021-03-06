package cmd

import (
	"fmt"
	"github.com/jamesroutley/awscm/core"
	"github.com/spf13/cobra"
	"strings"
)

// ls prints the available AWS profiles to stdout.
func ls(cmd *cobra.Command, args []string) {
	fmt.Println(strings.Join(core.Profiles(), "\n"))
}
