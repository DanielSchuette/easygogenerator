// gensite can be compiled to an executable and is the CLI of the easygogenerator package
//
// flags defined in this file:
/*
-header: file path to the header of your website

-footer: file path to the footer of your website

-body: file path to the footer of your website
*/
package main

import (
	"flag"
	"fmt"
	"log"

	egg "github.com/DanielSchuette/easygogenerator"
)

// define command line flags
var (
	header = flag.String("header", "", "file path to the header of your website")
	footer = flag.String("footer", "", "file path to the footer of your website")
	body   = flag.String("body", "", "file path to the footer of your website")
)

func main() {
	// defer done statement
	defer fmt.Println("done")

	// parse flags
	flag.Parse()

	// check if at least header and body are provided
	if *header == "" || *body == "" {
		log.Fatal("please provide a valid file path to at least a header file and a file with the content of your page\n(see --help)")
	}

	// get content of website from the respective file paths
	s := egg.GetSiteContent(*header, *body, *footer)

	// create a new website
	if err := egg.NewSite(s); err != nil {
		fmt.Printf("error while creating new website: %s\n", err)
	}

	// save website to a new file
	if err := s.Save("some_file_path"); err != nil {
		fmt.Printf("error while saving new website: %s\n", err)
	}
}
