/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chatgptdash",
	Short: "Will Call Chat GPT and get the information from the command line it self",
	Long:  `chatgptdash -t`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var openaiToken string
		flagval, err := cmd.Flags().GetBool("token")
		if err != nil {
			return
		}
		if flagval {
			openaiToken = args[0]
		}

		if openaiToken == "" {
			openaiToken = os.Getenv("OPENAI_TOKEN")
		}
		value := openaiToken

		// if value == "" {
		// 	reader2 := bufio.NewReader(os.Stdin)
		// 	fmt.Print("No ENV Variable as OPENAI_TOKEN found. Please enter the Open AI Token from  https://platform.openai.com/account/api-keys: ")

		// 	value, err := reader2.ReadString('\n')

		// 	if err != nil {
		// 		panic(err)
		// 	}

		// 	// fmt.Println(value)

		// }

		// print(args[0])
		// text := args[0]

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ask to chat gpt: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		c := openai.NewClient(value)

		resp, err := c.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    "user",
						Content: text,
					},
				},
			},
		)

		if err != nil {
			panic(err)
		}

		fmt.Println(resp.Choices[0].Message.Content)
		fmt.Println()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("token", "t", false, "Open AI Token")
	// rootCmd.Flags()
}
