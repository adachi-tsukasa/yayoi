package main

import (
	"fmt"
	"os"

	"github.com/adachi-tsukasa/yayoi"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	app := cli.NewApp()

	app.Name = "yayoi"
	app.Usage = "YAml convert json for You with Output / Input"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "STDOUT",
			Usage: "File output. Path specification required.",
		},
		// cli.BoolFlag{
		// 	Name: "eval, e",
		// },
		// cli.BoolFlag{
		// 	Name: "reverse, r",
		// },
		cli.BoolFlag{
			Name:  "minify, m",
			Usage: "Output File minifiy.",
		},
	}
	app.Action = func(command *cli.Context) error {
		var source string
		var outputBytes []byte

		if terminal.IsTerminal(0) {
			// if command.Bool("eval") {
			// 	source = "ARGV"
			// } else {
			if command.Args().First() == "" {
				fmt.Println("plz type the help command. `yayoi --help`")
				os.Exit(1)
			}
			source = "FILE"
			// }
		} else {
			source = "STDIN"
		}

		inputByte := yayoi.GetInputBytes(source, command.Args().First())
		outputBytes = yayoi.YamlToJason(inputByte, command.Bool("minify"))
		// outputBytes = yayoi.JsonToYaml(inputByte)
		yayoi.Output(outputBytes, command.String("output"))

		return nil
	}
	app.Run(os.Args)
}
