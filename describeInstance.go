package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Create an EC2 service object in the "us-west-2" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := ec2.New(&aws.Config{Region: aws.String("us-east-1")})

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	// resp has all of the response data, pull out instance IDs:
	fmt.Println("> Number of reservation sets: ", len(resp.Reservations))
	for idx, _ := range resp.Reservations {
		// fmt.Println("  > Number of instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {

			// dnsName = append(dnsName, *inst.NetworkInterfaces[0].PrivateDnsName)
			// ips = append(ips, *inst.NetworkInterfaces[0].Association.PublicIp)
			// fmt.Println("    - Instance ID: ", *inst)
			// if len(*inst.Tags) == 0 {
			// 	continue
			// }
			fmt.Printf("    - Instance ID: %s DNS: %s IP: %s \n", *inst.Tags[0].Value, *inst.NetworkInterfaces[0].PrivateDnsName, *inst.NetworkInterfaces[0].Association.PublicIp)
		}
	}
	// for _, x := range dnsName {
	// 	fmt.Println(x)
	// }
	// for _, x := range ips {
	// 	fmt.Println(x)
	// }

}
