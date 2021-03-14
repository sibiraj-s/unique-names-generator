package unique

import (
	"math/rand"
	"strings"
	"time"
)

// Options is the input argument for Generate function
type Options struct {
	Dictionaries [][]string
	Length       int
	Separator    *string
	Seed         int64
	Style        string
}

var defaultSeparator = "_"

func (o *Options) fillDefaults() *Options {
	if o.Dictionaries == nil {
		o.Dictionaries = [][]string{Adjectives, Colors, Animals}
	}

	if o.Length == 0 {
		defaultLength := 3

		if dictLength := len(o.Dictionaries); dictLength < defaultLength {
			o.Length = dictLength
		} else {
			o.Length = 3
		}
	}

	if o.Seed == 0 {
		o.Seed = time.Now().Unix()
	}

	if o.Separator == nil {
		o.Separator = &defaultSeparator
	}

	if o.Style == "" {
		o.Style = "lowercase"
	}

	return o
}

// New creates random unique names
func New(o Options) string {
	uniqueNames := []string{}

	config := o.fillDefaults()

	for _, dict := range config.Dictionaries[0:config.Length] {
		s := rand.NewSource(config.Seed)

		r := rand.New(s)
		n := r.Intn(len(dict))
		word := dict[n]

		switch config.Style {
		case "uppercase":
			word = strings.ToUpper(word)
		case "lowercase":
			word = strings.ToLower(word)
		case "titlecase":
			word = strings.Title(word)
		}

		uniqueNames = append(uniqueNames, word)
	}

	return strings.Join(uniqueNames, *config.Separator)
}
