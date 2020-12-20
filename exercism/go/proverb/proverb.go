package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var proverb []string

	if len(rhyme) == 0 {
		return proverb
	}

	for i, word := range rhyme[:len(rhyme)-1] {
		p := fmt.Sprintf("For want of a %s the %s was lost.", word, rhyme[i+1])
		proverb = append(proverb, p)
	}

	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return proverb
}
