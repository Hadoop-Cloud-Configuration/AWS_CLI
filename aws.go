//_author Jesse XU
package main
import (
	"fmt"
	// "os"
	"flag"
)
func main() {
	flag.Usage = func() {
	    fmt.Printf("\nUsage: ./aws [options]\n\n")
		
		fmt.Printf("\t aws create\n")
		fmt.Printf("\t aws list\n")
		fmt.Printf("\t aws stop\n")
		// flag.PrintDefaults()
	}
	flag.Parse()
	args:=flag.Args();
	if args[0]=="create"{
		id:="ami-1624987f"
		region:="us-east-1"
		instanceType:="t1.micro"
		createAWSInstance(region,id,instanceType)
		
	}else if args[0]=="list"{
		// region:="us-east-1"
	}else if args[0]=="ip"{
		
	}else if args[0]=="stop"{
		region:="us-east-1"
		instanceId:="i-c93d361e"
		stopAWSInstance(region,instanceId)
	}
	
}
