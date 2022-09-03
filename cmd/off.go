package cmd

import (
	"github.com/Iot-hack-labs/gpio2pwr/internal/PowerStrip"
	"github.com/spf13/cobra"
)

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn off an outlet",
	Long: `Turn off an outlet

off Fan
`,
	Run: func(cmd *cobra.Command, args []string) {
		ps := PowerStrip.New()
		defer ps.Close()

		outlet, _ := cmd.Flags().GetString("outlet")
		ps.Off(outlet)

	},
}

func init() {
	rootCmd.AddCommand(offCmd)

	offCmd.Flags().StringP("outlet", "o", "", "Outlet to turn on")

}
