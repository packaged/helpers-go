package helpers

import (
	"net/url"

	"github.com/gorilla/schema"
)

func FormURLDecode(r interface{}, b string) error {
	q, e := url.ParseQuery(b)
	if e != nil {
		return e
	}

	s := schema.NewDecoder()
	s.IgnoreUnknownKeys(true)
	return s.Decode(r, q)
}
