package easygogenerator

import (
	"testing"
)

type testContent struct {
	in   []string
	want string
}

func TestGetSiteContent(t *testing.T) {
	cases := []testContent{
		{
			// test if header, body, and footer are correctly included in new website
			in:   []string{"header", "body", "footer"},
			want: "header\nbody\nfooter",
		},
		{
			// test if header, body, and footer replacement (i.e. comment) are correctly included in new website
			in:   []string{"header", "body", ""},
			want: "header\nbody\n<!-- no footer -->",
		},
	}
	for idx, c := range cases {
		s := GetSiteContent(c.in[0], c.in[1], c.in[2])
		if err := NewSite(s); err != nil {
			t.Errorf("error creating new test site %d: %s\n", idx, err)
		}
		got := s.String()
		if got != c.want {
			t.Errorf("GetSiteContent(%s, %s, %s) == %s, want %s", c.in[0], c.in[1], c.in[2], got, c.want)
		}
	}
}
