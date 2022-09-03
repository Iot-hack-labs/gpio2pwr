package cmd

import (
	"github.com/Iot-hack-labs/gpio2pwr/internal/PowerStrip"
	"github.com/spf13/cobra"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn an outlet on",
	Long:  `Turn an outlet on`,
	Run: func(cmd *cobra.Command, args []string) {
		ps := PowerStrip.New()
		defer ps.Close()
		outlet, _ := cmd.Flags().GetString("outlet")
		ps.On(outlet)

	},
}

func init() {
	rootCmd.AddCommand(onCmd)

	onCmd.Flags().StringP("outlet", "o", "", "Outlet to turn on")
}
