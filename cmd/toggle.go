package cmd

import (
	"github.com/Iot-hack-labs/gpio2pwr/internal/PowerStrip"

	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle power to an outlet",
	Long:  `Toggle power to an outlet`,
	Run: func(cmd *cobra.Command, args []string) {
		ps := PowerStrip.New()
		defer ps.Close()
		outlet, _ := cmd.Flags().GetString("outlet")
		ps.Toggle(outlet)

	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)

	toggleCmd.Flags().StringP("outlet", "o", "", "Outlet to turn on")
}
