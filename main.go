// Copyright 2019 Yandy Ramirez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func main() {

	// Set background context, concurrency safety
	ctx := context.Background()

	// Load AWS client config
	// aws_key && aws_secret amongst others
	// set the region
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("Unable to load SDK config %v\n", err)
	}

	// TODO: add cli flag to specify region or get from profile
	cfg.Region = endpoints.UsEast2RegionID

	// Create cloudwatch client
	cw := cloudwatchlogs.New(cfg)

	// parameters for filtering the log groups
	limit := aws.Int64(50)
	params := cloudwatchlogs.DescribeLogGroupsInput{
		Limit: limit,
	}

	// Create a request for the log groups within a region
	a := cw.DescribeLogGroupsRequest(&params)
	resp, err := a.Send(ctx)
	if err != nil {
		log.Fatalf("Unable to request logs %v\n", err)
	}

	// check if there are log groups to delete
	if len(resp.LogGroups) > 0 {
		// go through each log group
		for _, v := range resp.LogGroups {

			name := v.LogGroupName
			params := &cloudwatchlogs.DeleteLogGroupInput{
				LogGroupName: name,
			}

			// Create delete request
			del := cw.DeleteLogGroupRequest(params)
			// Send the request to delete the log group
			_, err := del.Send(ctx)
			if err != nil {
				log.Fatalf("Could not delete log group %v\n", err)
			}
			fmt.Printf("Deleted log group %v\n", *name)
		}

		return
	}

	fmt.Println("No log groups to delete")
}
