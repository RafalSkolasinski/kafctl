package cmd

import (
	"log"

	"github.com/rafalskolasinski/kafctl/pkg/messages"

	"github.com/spf13/cobra"
)

var getMessages = &cobra.Command{
	Use:   "messages <topic1> [<topic2> ... ]",
	Short: "Consume messages",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig(cfgFile)
		logConfig(config)
		groupID, _ := cmd.Flags().GetString("group-id")
		delay, _ := cmd.Flags().GetString("delay")
		err := messages.Consume(config, args, groupID, delay)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var createMessages = &cobra.Command{
	Use:   "messages <topic>",
	Short: "Produce messages",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig(cfgFile)
		logConfig(config)

		numMessages, _ := cmd.Flags().GetInt("num-messages")
		delay, _ := cmd.Flags().GetString("delay")
		key, _ := cmd.Flags().GetString("key")
		messages.Produce(config, args[0], numMessages, key, delay)
	},
}

func init() {
	getCmd.AddCommand(getMessages)
	createCmd.AddCommand(createMessages)

	getMessages.Flags().String("group-id", "default", "consumer group id")
	getMessages.Flags().StringP("delay", "d", "0s", "delay between consuming messages")

	createMessages.Flags().IntP("num-messages", "n", 10, "number of messages to produce")
	createMessages.Flags().StringP("key", "k", "", "key for messages")
	createMessages.Flags().StringP("delay", "d", "0s", "delay between producing messages")

}
