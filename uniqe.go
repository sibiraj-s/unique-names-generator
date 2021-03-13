package unique

import (
	"math/rand"
	"strings"
	"time"
)

// GenerateParams is the input argument for Generate function
type GenerateParams struct {
	dictionaries [][]string
	length       int
	separator    *string
	seed         int64
	style        string
}

var defaultSeparator = "_"

func (p *GenerateParams) fillDefaults() *GenerateParams {
	if p.dictionaries == nil {
		p.dictionaries = [][]string{Adjectives, Colors, Animals}
	}

	if p.length == 0 {
		defaultLength := 3

		if dictLength := len(p.dictionaries); dictLength < defaultLength {
			p.length = dictLength
		} else {
			p.length = 3
		}
	}

	if p.seed == 0 {
		p.seed = time.Now().Unix()
	}

	if p.separator == nil {
		p.separator = &defaultSeparator
	}

	if p.style == "" {
		p.style = "lowercase"
	}

	return p
}

// New creates random unique names
func New(p GenerateParams) string {
	uniqueNames := []string{}

	config := p.fillDefaults()

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
