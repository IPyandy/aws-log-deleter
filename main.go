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
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func main() {

	region := flag.String("region", "", "--region=[AWS-REGION], defaults to aws config profile")
	flag.Parse()

	// Set background context, concurrency safety
	ctx := context.Background()

	// Load AWS client config
	// aws_key && aws_secret amongst others
	// set the region
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("Unable to load SDK config %v\n", err)
	}

	// If region flag is set
	if *region != "" {
		cfg.Region = *region
	}

	// Create cloudwatch client
	cw := cloudwatchlogs.New(cfg)

	lg := getLogGroups(ctx, cw)

	// check if there are log groups to delete
	if len(lg) > 0 {
		deleteLogGroups(ctx, lg, cw)

		return
	}

	fmt.Println("No log groups to delete")
}

// getLogGroups returns the logGroups if any acquired from the call to AWS CloudWatch-Logs
func getLogGroups(ctx context.Context, cw *cloudwatchlogs.CloudWatchLogs) []cloudwatchlogs.LogGroup {

	// will cancel request after 2 seconds if log groups are not received
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// parameters for filtering the log groups
	limit := aws.Int64(50)
	params := cloudwatchlogs.DescribeLogGroupsInput{
		Limit: limit,
	}

	// Create a request for the log groups within a region
	logReq := cw.DescribeLogGroupsRequest(&params)
	resp, err := logReq.Send(ctx)
	if err != nil {
		log.Fatalf("Unable to request logs %v\n", err)
	}

	return resp.LogGroups
}

// deleteLogGroups runs the delete operation on l[] for each group contianed
func deleteLogGroups(ctx context.Context, l []cloudwatchlogs.LogGroup, cw *cloudwatchlogs.CloudWatchLogs) {

	// will cancel deletion after 10 seconds of activity (in case of hung process)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// go through each log group
	for _, v := range l {

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
}
