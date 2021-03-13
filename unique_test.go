package unique

import (
	"regexp"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {

	// should generate random string correctly
	if got := New(GenerateParams{}); got != "" {
		match, _ := regexp.MatchString(`^\w+_\w+_\w+$`, got)

		if match == false {
			t.Errorf("Invalid string generated")
		}
	}

	// should generate random names given only two dictionaries
	if got := New(GenerateParams{dictionaries: [][]string{{"a"}, {"b"}}}); got != "a_b" {
		t.Errorf(`Expected to generate "a-b" but got: "%s"`, got)
	}

	// should generate random names given only one dictonary
	if got := New(GenerateParams{dictionaries: [][]string{{"a"}}}); got != "a" {
		t.Errorf(`Expected to generate "a" but got: "%s"`, got)
	}

	// should generate string with 3 random words, without any config
	if got := New(GenerateParams{}); len(strings.Split(got, "_")) != 3 {
		t.Errorf("Random string should be of default length 3")
	}

	// should generate string with given length
	if got := New(GenerateParams{length: 2}); len(strings.Split(got, "_")) != 2 {
		t.Errorf("Generated string is expected to be of length 2")
	}

	// should generate string with given separator
	hypen := "-"
	if got := New(GenerateParams{separator: &hypen}); got != "" {
		match, _ := regexp.MatchString(`^\w+-\w+-\w+$`, got)

		if match == false {
			t.Errorf("String should be generated with underscore as separator")
		}
	}

	// should be able to concat strings without a seperator
	emptySeparator := ""
	if got := New(GenerateParams{dictionaries: [][]string{{"a"}, {"b"}}, separator: &emptySeparator}); got != "ab" {
		t.Errorf(`Expected ab but got %s`, got)
	}

	seededUnique1 := New(GenerateParams{seed: 10})
	seededUnique2 := New(GenerateParams{seed: 10})

	// should generate random string with given separator `underscrore`
	if seededUnique1 != seededUnique2 {
		t.Errorf("Should generate same unique with given seed.")
	}
}
