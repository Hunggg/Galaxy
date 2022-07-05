package cmd

import (
	"log"

	publishmetadata "github.com/metrogalaxy-io/metrogalaxy-api/cmd/publish_metadata"
	"github.com/spf13/cobra"
)

var (
	metronionWearableMetadataCsvFile string
)

var publishWearableMetadata = &cobra.Command{
	Use:   "publish-wearable-metadata",
	Short: "Publish Metronion Wearable metadata to persistent storage",
	Long:  `Publish Metronion Wearable metadata to persistent storage`,
	Run: func(cmd *cobra.Command, args []string) {
		if metronionWearableMetadataCsvFile != "" {
			s := publishmetadata.NewPublishMetadataService()
			if err := s.RunPublishWearableMetadata(metronionWearableMetadataCsvFile); err != nil {
				log.Panic(err)
			}
		} else {
			log.Panicf("missing input wearable metadata CSV file")
		}
	},
}

func init() {
	RootCmd.AddCommand(publishWearableMetadata)

	publishWearableMetadata.Flags().StringVar(&metronionWearableMetadataCsvFile, "file", "", "Input wearable metadata CSV file")
}
