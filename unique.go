package unique

import (
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Options is the input argument for Generate function
type Options struct {
	Dictionaries [][]string
	Length       int
	Separator    *string
	Seed         int64
	Style        string
}

var (
	defaultSeparator = "_"
	defaultStyle     = "lowercase"
	defaultLength    = 3
)

func fillDefaults(options Options) Options {
	if options.Dictionaries == nil {
		options.Dictionaries = [][]string{Adjectives, Colors, Animals}
	}

	if options.Length == 0 {
		if dictLength := len(options.Dictionaries); dictLength < defaultLength {
			options.Length = dictLength
		} else {
			options.Length = defaultLength
		}
	}

	if options.Seed == 0 {
		options.Seed = time.Now().UnixNano()
	}

	if options.Separator == nil {
		options.Separator = String(defaultSeparator)
	}

	if options.Style == "" {
		options.Style = defaultStyle
	}

	return options
}

// New creates random unique names
func New(options ...Options) string {
	uniqueNames := []string{}

	if len(options) == 0 {
		options = append(options, Options{})
	}

	o := fillDefaults(options[0])
	s := rand.NewSource(o.Seed)
	r := rand.New(s)

	for _, dict := range o.Dictionaries[0:o.Length] {
		n := r.Intn(len(dict))
		word := dict[n]

		switch o.Style {
		case "uppercase":
			word = strings.ToUpper(word)
		case "lowercase":
			word = strings.ToLower(word)
		case "titlecase":
			word = cases.Title(language.English).String(word)
		}

		uniqueNames = append(uniqueNames, word)
	}

	return strings.Join(uniqueNames, *o.Separator)
}
