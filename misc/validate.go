package misc

import "regexp"

func ValidateEmail(email string) bool {
	r, _ := regexp.Compile(EmailRegex)
	return r.MatchString(email)
}

func ValidateURL(url string) bool {
	r, _ := regexp.Compile(URLRegex)
	return r.MatchString(url)
}
