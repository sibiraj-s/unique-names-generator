package unique

import (
	"regexp"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {

	// should generate random string correctly
	if got := New(Options{}); got != "" {
		match, _ := regexp.MatchString(`^\w+_\w+_\w+$`, got)

		if match == false {
			t.Errorf("Invalid string generated")
		}
	}

	// should generate random names given only two dictionaries
	if got := New(Options{Dictionaries: [][]string{{"a"}, {"b"}}}); got != "a_b" {
		t.Errorf(`Expected to generate "a_b" but got: "%s"`, got)
	}

	// should generate random names given only one dictionary
	if got := New(Options{Dictionaries: [][]string{{"a"}}}); got != "a" {
		t.Errorf(`Expected to generate "a" but got: "%s"`, got)
	}

	// should generate string with 3 random words, without any config
	if got := New(Options{}); len(strings.Split(got, "_")) != 3 {
		t.Errorf("Random string should be of default length 3")
	}

	// should generate string with given length
	if got := New(Options{Length: 2}); len(strings.Split(got, "_")) != 2 {
		t.Errorf("Generated string is expected to be of length 2")
	}

	// should generate string with given separator
	hyphen := "-"
	if got := New(Options{Separator: &hyphen}); got != "" {
		match, _ := regexp.MatchString(`^\w+-\w+-\w+$`, got)

		if match == false {
			t.Errorf("String should be generated with hyphen as separator")
		}
	}

	// should be able to concat strings without a separator
	emptySeparator := ""
	if got := New(Options{Dictionaries: [][]string{{"a"}, {"b"}}, Separator: &emptySeparator}); got != "ab" {
		t.Errorf(`Expected ab but got %s`, got)
	}

	seededUnique1 := New(Options{Seed: 10})
	seededUnique2 := New(Options{Seed: 10})

	// should generate same random string with given seed
	if seededUnique1 != seededUnique2 {
		t.Errorf("Should generate same unique with given seed.")
	}

	// should format the string to uppercase
	if got := New(
		Options{
			Dictionaries: [][]string{{"test"}, {"uppercase"}, {"style"}},
			Style:        "uppercase",
		},
	); got != "TEST_UPPERCASE_STYLE" {
		t.Errorf(`Expect to format in uppercase style but got "%s"`, got)
	}

	// should format the string to lowercase
	if got := New(
		Options{
			Dictionaries: [][]string{{"TEST"}, {"LOWERCASE"}, {"STYLE"}},
			Style:        "lowercase",
		},
	); got != "test_lowercase_style" {
		t.Errorf(`Expect to format in lowercase style but got "%s"`, got)
	}

	if got := New(
		Options{
			Dictionaries: [][]string{{"test"}, {"titlecase"}, {"style"}},
			Separator:    &emptySeparator,
			Style:        "titlecase",
		},
	); got != "TestTitlecaseStyle" {
		t.Errorf(`Expect to format in titlecase style but got "%s"`, got)
	}
}
