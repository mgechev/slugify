package slugify_test

import (
	"testing"

	"github.com/mgechev/slugify"
)

var slugified = map[string]string{
	"foobar":         "foobar",
	"foo bar":        "foo-bar",
	"foo bar ж":      "foo-bar-zh",
	"foo <3 ж":       "foo-love-zh",
	"😊foo <3 ж":      "foo-love-zh",
	"foo 😊 qux 😊":    "foo-qux",
	"% qux":          "qux",
	"%20barqux":      "20barqux",
	"Случаен низ":    "sluchaen-niz",
	"I don't like 🕺": "i-dont-like",
}

func TestSlufigy(t *testing.T) {
	for original, transformed := range slugified {
		if result := slugify.Transform(original); result != transformed {
			t.Errorf("%s should equal %s but equals %s", original, transformed, result)
		}
	}
}
