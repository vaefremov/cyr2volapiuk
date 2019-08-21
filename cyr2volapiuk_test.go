package cyr2volapiuk_test

import (
	"testing"

	"github.com/vaefremov/cyr2volapiuk"
)

func TestString(t *testing.T) {
	cases := []struct{ in, wanted string }{
		{"АБВ", "ABW"},
		{cyr2volapiuk.CyrCapitals, "ABWGDEYoZhZIJKLMNOPRSTUFHCChShSchYEYuYa"},
		{cyr2volapiuk.CyrSmall, "abwgdeyozhzijklmnoprstufhcchshschyeyuya"},
		{cyr2volapiuk.ASCIILetters, cyr2volapiuk.ASCIILetters},
		{cyr2volapiuk.Digits, cyr2volapiuk.Digits},
		{"AB CD", "AB CD"},
		{cyr2volapiuk.AllowedInFilenames, "________"},
	}
	for _, c := range cases {
		if got := cyr2volapiuk.String(c.in); got != c.wanted {
			t.Errorf("From: %s, expected: %s, got: %s", c.in, c.wanted, got)
		}
	}
}
func TestStringPermissive(t *testing.T) {
	cases := []struct{ in, wanted string }{
		{"АБВ", "ABW"},
		{cyr2volapiuk.CyrCapitals, "ABWGDEYoZhZIJKLMNOPRSTUFHCChShSchYEYuYa"},
		{cyr2volapiuk.CyrSmall, "abwgdeyozhzijklmnoprstufhcchshschyeyuya"},
		{cyr2volapiuk.ASCIILetters, cyr2volapiuk.ASCIILetters},
		{cyr2volapiuk.Digits, cyr2volapiuk.Digits},
		{"AB CD", "AB CD"},
		{cyr2volapiuk.AllowedInFilenames, cyr2volapiuk.AllowedInFilenames},
	}
	for _, c := range cases {
		if got := cyr2volapiuk.StringPermissive(c.in); got != c.wanted {
			t.Errorf("From: %s, expected: %s, got: %s", c.in, c.wanted, got)
		}
	}
}

func TestFileName(t *testing.T) {
	cases := []struct{ in, wanted string }{
		{cyr2volapiuk.CyrCapitals, "ABWGDEYoZhZIJKLMNOPRSTUFHCChShSchYEYuYa"},
		{cyr2volapiuk.CyrSmall, "abwgdeyozhzijklmnoprstufhcchshschyeyuya"},
		{cyr2volapiuk.ASCIILetters, cyr2volapiuk.ASCIILetters},
		{cyr2volapiuk.Digits, cyr2volapiuk.Digits},
		{"AB CD", "AB_CD"},
		{cyr2volapiuk.AllowedInFilenames, cyr2volapiuk.AllowedInFilenames},
	}
	for _, c := range cases {
		if got := cyr2volapiuk.FileName(c.in); got != c.wanted {
			t.Errorf("From: %s, expected: %s, got: %s", c.in, c.wanted, got)
		}
	}
}

func BenchmarkFileName(b *testing.B) {
	testString := cyr2volapiuk.CyrCapitals + cyr2volapiuk.CyrSmall + cyr2volapiuk.AllowedInFilenames
	for n := 0; n < b.N; n++ {
		cyr2volapiuk.FileName(testString)
	}
}
