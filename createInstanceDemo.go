package main
import (
	"log"
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

	params := &ec2.RunInstancesInput{
    	ImageId:      aws.String("ami-1624987f"),
    	InstanceType: aws.String("t1.micro"),
    	MinCount:     aws.Int64(1),
    	MaxCount:     aws.Int64(1),
	}

	runResult, err := svc.RunInstances(params)
	if err != nil {
    		log.Println("Could not create instance", err)
    		return
	}

	fmt.Println(runResult)
}
