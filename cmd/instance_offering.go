package cmd

import (
	"fmt"

	"github.com/bharath23/awstool/pkg/instance"
	"github.com/spf13/cobra"
)

var offeringsCmd = &cobra.Command{
	Use:   "offerings",
	Short: "list the different classes of EC2 instances available",
	Long:  "list the different classes of EC2 instances available for the availability zone",
	RunE:  run,
}

func init() {
	offeringsCmd.Flags().String("location-type", "region", "location type of offerings, possible values [availability-zone, availability-zone-id, region]")
	instanceCmd.AddCommand(offeringsCmd)
}

func run(cmd *cobra.Command, args []string) error {
	locationType, err := cmd.Flags().GetString("location-type")
	if err != nil {
		return err
	}

	switch locationType {
	case "availability-zone", "availability-zone-id", "region":
		return instance.Offerings(locationType)
	default:
		return fmt.Errorf("bad value for location-type '%s'", locationType)
	}

	return fmt.Errorf("UNIMPLEMENTED!")
}
