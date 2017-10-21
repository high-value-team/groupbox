package interactions

import (
	"encoding/hex"
	crand "crypto/rand"
	"fmt"
	mrand "math/rand"
	"time"
)


// source: http://www.ashishbanerjee.com/home/go/go-generate-uuid
// API doc: https://golang.org/pkg/math/rand/#Read
func GenerateKey() string {
	uuid := make([]byte, 16)
	crand.Read(uuid)

	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid)
}


type NicknameGenerator struct {
	adjectives     []string
	nouns          []string
	generated 	   map[string]int
}

func NewNicknameGenerator() *NicknameGenerator {
	return newNicknameGeneratorExplicit(
		// not yet assigned: AEHNOQRTUWXY
		[]string{"Mighty", "Golden", "Flying", "Black", "Valiant", "Smart", "Indigo", "Charming", "Dreamy", "Luminous", "Pink", "Kinky", "Jumping"},

		// source for nouns: https://en.wikipedia.org/wiki/List_of_animal_names
		// not yet assigned: GILMNOQUVXY
		[]string{"Eagle", "Panda", "Raven", "Wolf", "Jaguar", "Dolphin", "Tiger", "Shark", "Koala", "Zebra", "Albatross", "Fox", "Barracuda", "Cheetah", "Hornet"},
	)
}

func newNicknameGeneratorExplicit(adjectives []string, nouns []string) *NicknameGenerator {
	o := NicknameGenerator{
		adjectives:     adjectives,
		nouns:          nouns,
		generated:      map[string]int{},
	}
	return &o
}

func (g *NicknameGenerator) Next() string {
	const MAX_RETRIES int = 10
	var nickname string

	mrand.Seed(time.Now().Unix())
	var retries int
	for retries = 0; retries < MAX_RETRIES; retries++ {
		adjectiveIndex := mrand.Intn(len(g.adjectives))
		nounIndex := mrand.Intn(len(g.nouns))

		nickname = fmt.Sprintf("%s %s", g.adjectives[adjectiveIndex], g.nouns[nounIndex])

		if _, ok := g.generated[nickname]; !ok {
			g.generated[nickname] = 0
			break;
		}
	}
	if retries == MAX_RETRIES {
		g.generated[nickname]++
		nickname = fmt.Sprintf("%s(%d)", nickname, g.generated[nickname])
	}
	return nickname
}