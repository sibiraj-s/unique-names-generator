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
	if got := New(Options{dictionaries: [][]string{{"a"}, {"b"}}}); got != "a_b" {
		t.Errorf(`Expected to generate "a-b" but got: "%s"`, got)
	}

	// should generate random names given only one dictonary
	if got := New(Options{dictionaries: [][]string{{"a"}}}); got != "a" {
		t.Errorf(`Expected to generate "a" but got: "%s"`, got)
	}

	// should generate string with 3 random words, without any config
	if got := New(Options{}); len(strings.Split(got, "_")) != 3 {
		t.Errorf("Random string should be of default length 3")
	}

	// should generate string with given length
	if got := New(Options{length: 2}); len(strings.Split(got, "_")) != 2 {
		t.Errorf("Generated string is expected to be of length 2")
	}

	// should generate string with given separator
	hypen := "-"
	if got := New(Options{separator: &hypen}); got != "" {
		match, _ := regexp.MatchString(`^\w+-\w+-\w+$`, got)

		if match == false {
			t.Errorf("String should be generated with underscore as separator")
		}
	}

	// should be able to concat strings without a seperator
	emptySeparator := ""
	if got := New(Options{dictionaries: [][]string{{"a"}, {"b"}}, separator: &emptySeparator}); got != "ab" {
		t.Errorf(`Expected ab but got %s`, got)
	}

	seededUnique1 := New(Options{seed: 10})
	seededUnique2 := New(Options{seed: 10})

	// should generate random string with given separator `underscrore`
	if seededUnique1 != seededUnique2 {
		t.Errorf("Should generate same unique with given seed.")
	}

	// should format the string to uppercase
	if got := New(
		Options{
			dictionaries: [][]string{{"test"}, {"uppercase"}, {"style"}},
			style:        "uppercase",
		},
	); got != "TEST_UPPERCASE_STYLE" {
		t.Errorf(`Expect to format in uppercase style but got "%s"`, got)
	}

	// should format the string to lowercase
	if got := New(
		Options{
			dictionaries: [][]string{{"TEST"}, {"LOWERCASE"}, {"STYLE"}},
			style:        "lowercase",
		},
	); got != "test_lowercase_style" {
		t.Errorf(`Expect to format in lowercase style but got "%s"`, got)
	}

	if got := New(
		Options{
			dictionaries: [][]string{{"test"}, {"titlecase"}, {"style"}},
			separator:    &emptySeparator,
			style:        "titlecase",
		},
	); got != "TestTitlecaseStyle" {
		t.Errorf(`Expect to format in titlecase style but got "%s"`, got)
	}
}
