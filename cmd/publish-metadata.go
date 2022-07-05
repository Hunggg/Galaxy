package cmd

import (
	"log"

	publishmetadata "github.com/metrogalaxy-io/metrogalaxy-api/cmd/publish_metadata"
	"github.com/spf13/cobra"
)

var (
	metronionMetadataCsvFile string
)

var publishMetadata = &cobra.Command{
	Use:   "publish-metadata",
	Short: "Publish Metronion NFT metadata to Mongodb storage",
	Long:  `Publish Metronion NFT metadata to Mongodb storage`,
	Run: func(cmd *cobra.Command, args []string) {
		if metronionMetadataCsvFile != "" {
			s := publishmetadata.NewPublishMetadataService()
			if err := s.Run(metronionMetadataCsvFile); err != nil {
				log.Panic(err)
			}
		} else {
			log.Panicf("missing input metadata CSV file")
		}
	},
}

func init() {
	RootCmd.AddCommand(publishMetadata)

	publishMetadata.Flags().StringVar(&metronionMetadataCsvFile, "file", "", "Input metadata CSV file")
}
