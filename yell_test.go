package yell

import (
	"regexp"
	"testing"
)

func TestNationalityString(t *testing.T) {
	nationality := "Iranian"
	want := regexp.MustCompile(`\b` + nationality + `\b`)
	msg, err := FreeSoftware(nationality)
	if msg == "" || err != nil {
		t.Fatalf("FreeSoftware(%v) = %q, %#v, want match for %#q", nationality, msg, err, want)
	}
}

func TestNationalityEmptySting(t *testing.T) {
	nationality := ""
	msg, err := FreeSoftware(nationality)
	if msg != "" || err == nil {
		t.Fatalf(`FreeSoftware(%v) = %q, %s, want "", error`, nationality, msg, err)
	}
}
