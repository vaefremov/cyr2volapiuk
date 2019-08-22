// Package cyr2volapiuk implements utility methods to convert
// strings containing Cyrillic (currently, Russian) letters into
// their Latin representation using something like a phonetic
// representation.
package cyr2volapiuk

import (
	"strings"
)

// Basic sets of chars that are related to operating this package
//
//
const (
	CyrCapitals        = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЬЪЫЭЮЯ"
	CyrSmall           = "абвгдеёжзийклмнопрстуфхцчшщьъыэюя"
	ASCIILetters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits             = "0123456789"
	AllowedInFilenames = "_-+=.@:,"
)

// String accepts a string of chars and convert every Cyrillic char
// (strictly speaking, only the Russian alphabet if supported so far)
// into its ASCII representation.
// All chars that do not fall into Cyrillic/ASCII Letter/Digit category are
// replaced by _ (underscore).
func String(s string) string {
	var b strings.Builder
	for _, r := range s {
		rep, ok := convDict[r]
		if !ok {
			switch {
			case (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r <= '9' && r >= '0'):
				b.WriteRune(r)
			case r == ' ':
				b.WriteRune(' ')
			default:
				b.WriteRune('_')
			}
		} else {
			b.WriteString(rep)
		}
	}
	return b.String()
}

// StringPermissive is similar to String, but it does not replace unhandled
// characters with underscore, this chars are passed through as they are.
func StringPermissive(s string) string {
	var b strings.Builder
	for _, r := range s {
		rep, ok := convDict[r]
		if ok {
			b.WriteString(rep)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// FileName is aimed at being used to generate file names from plain text strings.
// These file names can be safely used in the cases where the underlying file system either does not
// support Cyrillic (or other non-ASCII Unicode characters) encoding, or when using Cyrillic letters in file names
// may cause problems. Also, only "safe" non-letter characters are left,
// list of these safe characters is available through the AllowedInFilenames constant.
func FileName(s string) string {
	s = StringPermissive(s)
	var b strings.Builder
	for _, r := range s {
		if _, ok := isFilenameSafe[r]; ok {
			b.WriteRune(r)
		} else {
			b.WriteRune('_')
		}
	}
	return b.String()
}

func init() {
	isFilenameSafe = make(map[rune]bool)
	for _, c := range ASCIILetters + AllowedInFilenames + Digits {
		isFilenameSafe[c] = true
	}
}

var isFilenameSafe map[rune]bool

var convDict = map[rune]string{
	'А': "A",
	'Б': "B",
	'В': "W",
	'Г': "G",
	'Д': "D",
	'Е': "E",
	'Ё': "Yo",
	'Ж': "Zh",
	'З': "Z",
	'И': "I",
	'Й': "J",
	'К': "K",
	'Л': "L",
	'М': "M",
	'Н': "N",
	'О': "O",
	'П': "P",
	'Р': "R",
	'С': "S",
	'Т': "T",
	'У': "U",
	'Ф': "F",
	'Х': "H",
	'Ц': "C",
	'Ч': "Ch",
	'Ш': "Sh",
	'Щ': "Sch",
	'Ь': "",
	'Ъ': "",
	'Ы': "Y",
	'Э': "E",
	'Ю': "Yu",
	'Я': "Ya",
	'а': "a",
	'б': "b",
	'в': "w",
	'г': "g",
	'д': "d",
	'е': "e",
	'ё': "yo",
	'ж': "zh",
	'з': "z",
	'и': "i",
	'й': "j",
	'к': "k",
	'л': "l",
	'м': "m",
	'н': "n",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'у': "u",
	'ф': "f",
	'х': "h",
	'ц': "c",
	'ч': "ch",
	'ш': "sh",
	'щ': "sch",
	'ь': "",
	'ъ': "",
	'ы': "y",
	'э': "e",
	'ю': "yu",
	'я': "ya",
}
