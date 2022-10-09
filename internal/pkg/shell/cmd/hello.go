package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type helloCmdFlag struct {
	name string
}

var hcf helloCmdFlag

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Print hello",
	Long:  "Print hello message",
	Run:   func(cmd *cobra.Command, args []string) { runHelloCmd(cmd, args) },
}

func init() {
	rootCmd.AddCommand(helloCmd)

	helloCmd.Flags().StringVar(&hcf.name, "name", "", "random name")
	// helloCmd.Flags().StringVarP(&hcf.name, "name", "n", "", "random name") // If you want to use shorthand -n
	// helloCmd.MarkFlagRequired("name") // If you want to make the name flag required
}

func runHelloCmd(cmd *cobra.Command, args []string) {
	if hcf.name != "" {
		fmt.Printf("Hello %s!\n", hcf.name)
	} else {
		fmt.Println("Hello there!")
	}
}
