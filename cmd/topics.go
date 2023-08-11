package cmd

import (
	"log"

	"github.com/rafalskolasinski/kafctl/pkg/topics"
	"github.com/spf13/cobra"
)

var getTopics = &cobra.Command{
	Use:   "topics",
	Short: "List topics",
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig(cfgFile)
		logConfig(config)
		err := topics.Get(config)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var createTopics = &cobra.Command{
	Use:   "topics <topic-1> [<topic-2> ...]",
	Short: "Create topic(s)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig(cfgFile)
		logConfig(config)

		numPartitions, _ := cmd.Flags().GetInt("num-partitions")
		replicationFactor, _ := cmd.Flags().GetInt("replication-factor")
		topics.Create(config, args, numPartitions, replicationFactor)
	},
}

var removeTopics = &cobra.Command{
	Use:   "topics <topic-1> [<topic-2> ...]",
	Short: "Remove topic(s)",
	Run: func(cmd *cobra.Command, args []string) {
		config := readConfig(cfgFile)
		logConfig(config)
		topics.Remove(config, args)
	},
}

func init() {
	getCmd.AddCommand(getTopics)
	createCmd.AddCommand(createTopics)
	removeCmd.AddCommand(removeTopics)

	createTopics.Flags().Int("num-partitions", 2, "number of partitions for new topic(s)")
	createTopics.Flags().Int("replication-factor", 1, "replication factor for new topic(s)")
}
