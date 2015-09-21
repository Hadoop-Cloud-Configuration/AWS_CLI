package main
import (
//        "log"
        "fmt"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/defaults"
        "github.com/aws/aws-sdk-go/service/ec2"
)
/**
 * Don't hard-code your credentials!
 * Export the following environment variables instead:
 *
 * export AWS_ACCESS_KEY_ID='AKID'
 * export AWS_SECRET_ACCESS_KEY='SECRET'
 */

func main(){
defaults.DefaultConfig.Region = aws.String("us-east-1")
svc := ec2.New(nil)

params := &ec2.StopInstancesInput{
InstanceIds: []*string{ // Required
		aws.String("i-b5fd8c60"), // Required
		// More values...
	},
	DryRun: aws.Bool(false),
	Force:  aws.Bool(true),
	
}
resp, err := svc.StopInstances(params)

if err != nil {
	// Print the error, cast err to awserr.Error to get the Code and
	// Message from an error.
	fmt.Println(err.Error())
	return
}

// Pretty-print the response data.
fmt.Println(resp)
}
