package unique

import (
	"math/rand"
	"strings"
	"time"
)

// Options is the input argument for Generate function
type Options struct {
	dictionaries [][]string
	length       int
	separator    *string
	seed         int64
	style        string
}

var defaultSeparator = "_"

func (o *Options) fillDefaults() *Options {
	if o.dictionaries == nil {
		o.dictionaries = [][]string{Adjectives, Colors, Animals}
	}

	if o.length == 0 {
		defaultLength := 3

		if dictLength := len(o.dictionaries); dictLength < defaultLength {
			o.length = dictLength
		} else {
			o.length = 3
		}
	}

	if o.seed == 0 {
		o.seed = time.Now().Unix()
	}

	if o.separator == nil {
		o.separator = &defaultSeparator
	}

	if o.style == "" {
		o.style = "lowercase"
	}

	return o
}

// New creates random unique names
func New(o Options) string {
	uniqueNames := []string{}

	config := o.fillDefaults()

	for _, dict := range config.dictionaries[0:config.length] {
		s := rand.NewSource(config.seed)

		r := rand.New(s)
		n := r.Intn(len(dict))
		word := dict[n]

		switch config.style {
		case "uppercase":
			word = strings.ToUpper(word)
		case "lowercase":
			word = strings.ToLower(word)
		case "titlecase":
			word = strings.Title(word)
		}

		uniqueNames = append(uniqueNames, word)
	}

	return strings.Join(uniqueNames, *config.separator)
}
