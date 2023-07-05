package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-hasaki/hasaki/config"
	"github.com/spf13/cobra"
)

var UpgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade the hasaki command.",
	Long:    "Upgrade the hasaki command.",
	Example: "hasaki upgrade",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("go install %s\n", config.HasakiCmd)
		cmd := exec.Command("go", "install", config.HasakiCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("go install %s error\n", err)
		}
		fmt.Printf("\n🎉 Hasaki upgrade successfully!\n\n")
	},
}
