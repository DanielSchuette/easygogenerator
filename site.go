package easygogenerator

import "fmt"

// Site holds all the information that is required to build the website, i.e. the header, an optional footer, and site body
type Site struct {
	Header  []byte // content of site header
	Body    []byte // content of site body
	Footer  []byte // content of site footer
	Website []byte // content of the complete website
}

// GetSiteContent takes the file paths to the header, body, and an optional footer for the new site and returns a populated *Site struct that can be passed to NewSite() to generate the website
func GetSiteContent(header, body, footer string) *Site {
	// create a new pointer to a site
	s := new(Site)

	// TODO: open files and copy content to *Site
	s.Header = []byte(header)
	s.Body = []byte(body)

	// if footer is empty, replace it with a comment
	if footer == "" {
		s.Footer = []byte("<!-- no footer -->")
	} else {
		s.Footer = []byte(footer)
	}
	return s
}

// Save saves a Site to a html file as specified by path
func (s *Site) Save(path string) error {
	// print done statement
	fmt.Printf("saved website to file: %s\n", path)
	// TODO: print website to new output file
	// return no error
	return nil
}

// The String method is useful because now Site implements the Stringer interface and the content of *Site gets printed in a custom format
func (s *Site) String() string {
	return string(s.Website)
}
