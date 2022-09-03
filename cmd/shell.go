package cmd

import (
	"github.com/Iot-hack-labs/gpio2pwr/internal/PowerStrip"
	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
	"io"
	"log"
	"strconv"
	"strings"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "An interactive shell for controlling power outlets",
	Long:  `An interactive shell for controlling power outlets`,
	Run: func(cmd *cobra.Command, args []string) {
		ps := PowerStrip.New()
		defer ps.Close()

		l, err := readline.NewEx(&readline.Config{
			Prompt:          "\033[31m»\033[0m ",
			HistoryFile:     "/tmp/readline.tmp",
			AutoComplete:    completer,
			InterruptPrompt: "^C",
			EOFPrompt:       "exit",

			HistorySearchFold:   true,
			FuncFilterInputRune: filterInput,
		})
		if err != nil {
			panic(err)
		}
		defer l.Close()
		//l.CaptureExitSignal()

		log.SetOutput(l.Stderr())
		for {
			line, err := l.Readline()
			if err == readline.ErrInterrupt {
				if len(line) == 0 {
					break
				} else {
					continue
				}
			} else if err == io.EOF {
				break
			}

			line = strings.TrimSpace(line)
			switch {

			case line == "all on":
				ps.AllOn()

			case line == "all off":
				ps.AllOff()

			case line == "all toggle":
				ps.ToggleAll()

			case strings.HasPrefix(line, "on "):
				outlet := line[3:]
				ps.On(outlet)
			case strings.HasPrefix(line, "off "):
				outlet := line[4:]
				ps.Off(outlet)
			case strings.HasPrefix(line, "toggle "):

				outlet := line[7:]
				ps.Toggle(outlet)

			case line == "help":
				usage(l.Stderr())
			case line == "?":
				usage(l.Stderr())

			case line == "exit":
				goto exit
			case line == "q":
				goto exit
			default:
				log.Println("unknown command:", strconv.Quote(line))
			}
		}
	exit:
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func usage(w io.Writer) {
	io.WriteString(w, "commands:\n")
	io.WriteString(w, completer.Tree("    "))
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem("all",
		readline.PcItem("on"),
		readline.PcItem("off"),
		readline.PcItem("toggle"),
	),

	readline.PcItem("on",
		readline.PcItem(PowerStrip.Fan),
		readline.PcItem(PowerStrip.Lamp),
		readline.PcItem(PowerStrip.RotatingLight),
		readline.PcItem(PowerStrip.Speakers),
	),
	readline.PcItem("off",
		readline.PcItem(PowerStrip.Fan),
		readline.PcItem(PowerStrip.Lamp),
		readline.PcItem(PowerStrip.RotatingLight),
		readline.PcItem(PowerStrip.Speakers),
	),
	readline.PcItem("toggle",
		readline.PcItem(PowerStrip.Fan),
		readline.PcItem(PowerStrip.Lamp),
		readline.PcItem(PowerStrip.RotatingLight),
		readline.PcItem(PowerStrip.Speakers),
	),

	readline.PcItem("quit"),
	readline.PcItem("q"),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}
