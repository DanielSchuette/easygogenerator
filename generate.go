package easygogenerator

import "fmt"

// NewSite generates a new website from a header.html and footer.html (optional) and saves it to an output .html file with '_complete.html' extension
func NewSite(s *Site) error {
	// defer done statement
	defer fmt.Println("site generated")

	// paste header, body, and footer (if it exists) together
	ws := fmt.Sprintf("%s\n%s\n%s", string(s.Header), string(s.Body), string(s.Footer))
	s.Website = ([]byte(ws))

	// print the site content and return no error
	fmt.Println(s)
	return nil
}
