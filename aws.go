//_author Jesse XU
package main
import (
	"fmt"
	"os"
	"flag"
)
func main() {
	var id string
	var instancetype string
	var instance string
	var key string
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagSet.StringVar(&id, "id", "ami-1df0ac78", "Image ID")
	flagSet.StringVar(&instancetype, "type", "t2.micro", "instancetype")
	flagSet.StringVar(&key, "key", "", "key name")
	flagSet.StringVar(&instance, "instance", "", "Instance Id")
	flagSet.Usage = func() {
	    fmt.Printf("\nUsage: ./aws [options]\n\n")
		fmt.Printf("\t aws create\n")
		fmt.Printf("\t aws create -id <ami-1df0ac78> -type <t2.micro> -key <keyname>\n")
		fmt.Printf("\t aws stop -instance <instance_id>\n")
		// flag.PrintDefaults()
	}
	flagSet.Parse(os.Args[2:])
	args:=os.Args[1:]
	if args[0]=="create"{
		// id:="ami-1df0ac78"
		region:="us-east-1"
// 		instanceType:="t2.micro"
// 		keyName:="jesse"
		createAWSInstance(region,id,instancetype,key)
	}else if args[0]=="stop"{
		region:="us-east-1"
		// instanceId:="i-c93d361e"
		stopAWSInstance(region,instance)
	}
	
}
