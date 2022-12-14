package instance

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func splitInstanceType(s string) (string, string) {
	slice := strings.Split(s, ".")
	return slice[0], slice[1]
}

func Offerings(locationType string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("unable to load AWS config, %v", err)
	}

	ec2Client := ec2.NewFromConfig(cfg)
	awsLocationType := types.LocationTypeRegion
	switch locationType {
	case "availability-zone":
		awsLocationType = types.LocationTypeAvailabilityZone
	case "availability-zone-id":
		awsLocationType = types.LocationTypeAvailabilityZoneId
	case "region":
		awsLocationType = types.LocationTypeRegion
	default:
		return fmt.Errorf("Unknown AWS location type, '%v'", locationType)
	}

	do := true
	instanceTypesByLocation := make(map[string]map[string][]string)
	input := ec2.DescribeInstanceTypeOfferingsInput{
		LocationType: awsLocationType,
	}

	for do {
		output, err := ec2Client.DescribeInstanceTypeOfferings(context.TODO(), &input)
		if err != nil {
			return fmt.Errorf("unable to DescribeInstanceTypeOfferings, %v", err)
		}

		for _, o := range output.InstanceTypeOfferings {
			instanceClass, instanceSize := splitInstanceType(string(o.InstanceType))
			location := *o.Location
			instanceTypes := instanceTypesByLocation[location]
			if instanceTypes == nil {
				instanceTypes = make(map[string][]string)
				instanceTypesByLocation[location] = instanceTypes
			}

			instanceTypes[instanceClass] = append(instanceTypes[instanceClass], instanceSize)
		}

		if output.NextToken != nil {
			input.NextToken = output.NextToken
		} else {
			do = false
		}
	}

	orderedSize := map[string]int{
		"nano":      -4,
		"micro":     -3,
		"small":     -2,
		"medium":    -1,
		"large":     0,
		"xlarge":    1,
		"2xlarge":   2,
		"3xlarge":   3,
		"4xlarge":   4,
		"6xlarge":   6,
		"8xlarge":   8,
		"9xlarge":   9,
		"10xlarge":  10,
		"12xlarge":  12,
		"16xlarge":  16,
		"18xlarge":  18,
		"24xlarge":  24,
		"32xlarge":  32,
		"48xlarge":  48,
		"56xlarge":  56,
		"112xlarge": 112,
		"metal":     999,
	}
	for _, l := range instanceTypesByLocation {
		for _, ss := range l {
			sort.Slice(ss, func(i, j int) bool {
				return orderedSize[ss[i]] < orderedSize[ss[j]]
			})
		}
	}

	b, err := json.MarshalIndent(instanceTypesByLocation, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshall instance types by location, '%v'", err)
	}

	fmt.Printf("%s\n", string(b))
	return nil
}
