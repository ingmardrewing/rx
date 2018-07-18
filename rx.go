package rx

import "regexp"

// create a convenience struct for dealing with the
// rather unwieldy (compared to Perl) regexes in golang
func NewRx(pattern string) (*rx, error) {
	r := new(rx)
	r.pattern = pattern
	err := r.compile()
	if err != nil {
		return nil, err
	}
	return r, nil
}

type rx struct {
	text    string
	pattern string
	regex   *regexp.Regexp
}

func (r *rx) compile() error {
	regex, err := regexp.Compile(r.pattern)
	if err != nil {
		return err
	}
	r.regex = regex
	return nil
}

// use the regex to substitute all matching
// parts of the param text with the param replacement
func (r *rx) SubstituteAllOccurences(text, replacement string) string {
	return r.regex.ReplaceAllString(text, replacement)
}
