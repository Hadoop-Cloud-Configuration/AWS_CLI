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
		
	}else if args[0]=="list"{
		
	}else if args[0]=="ip"{
		
	}else if args[0]=="stop"{
		
	}
	
}
