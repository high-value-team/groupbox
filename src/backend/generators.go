package backend

import (
	"encoding/hex"
	"crypto/rand"
	"fmt"
	"math"
)


// source: http://www.ashishbanerjee.com/home/go/go-generate-uuid
// API doc: https://golang.org/pkg/math/rand/#Read
func generateKey() string {
	uuid := make([]byte, 16)
	rand.Read(uuid)

	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid)
}


/* Nicknames are generated as pairs of words: first an adjective, then a noun, e.g. "Golden Tiger".
   Both words are taken from respective lists and will lead to nicknames generated for all boxes in the same order.
   Within a box all nicknames are different, across boxes they are the same.

   Lists of a fixed length only support a fixed number of nicknames. Should the number of box members exceed
   this max. number of different nicknames a suffix is appended to ensure uniqueness.
   This should not happen, though, since such large numbers of members are very unlikely.
 */
type NicknameGenerator struct {
	adjectives     []string
	nouns          []string
	adjectiveIndex int
	nounIndex      int
}

func NewNicknameGenerator() *NicknameGenerator {
	return NewNicknameGenerator_explicit(
		// AEHNOQRTUWXY
		[]string{"Mighty", "Golden", "Flying", "Black", "Valiant", "Smart", "Indigo", "Charming", "Dreamy", "Luminous", "Pink", "Kinky", "Jumping"},

		// source for nouns: https://en.wikipedia.org/wiki/List_of_animal_names
		// GILMNOQUVXY
		[]string{"Eagle", "Panda", "Raven", "Wolf", "Jaguar", "Dolphin", "Tiger", "Shark", "Koala", "Zebra", "Albatross", "Fox", "Barracuda", "Cheetah", "Hornet"},
	)
}

func NewNicknameGenerator_explicit(adjectives []string, nouns []string) *NicknameGenerator {
	o := NicknameGenerator{
		adjectives:     adjectives,
		nouns:          nouns,
		adjectiveIndex: 0,
		nounIndex:      0,
	}
	return &o
}

func (g *NicknameGenerator) next() string {
	ai := int(math.Mod(float64(g.adjectiveIndex),float64(len(g.adjectives))))
	ni := int(math.Mod(float64(g.nounIndex),float64(len(g.nouns))))
	fmt.Printf("  next before: %v, %v -> %v/%v\n", g.adjectiveIndex, g.nounIndex, ai, ni)

	if ni == len(g.nouns)-1 { g.adjectiveIndex += 1 }
	g.nounIndex += 1

	nickname := fmt.Sprintf("%s %s", g.adjectives[ai], g.nouns[ni])
	if g.adjectiveIndex >= len(g.adjectives) {
		suffix := fmt.Sprintf("(%v.%v)", g.adjectiveIndex, g.nounIndex)
		nickname += suffix
	}

	fmt.Printf("  next: %s\n", nickname)
	return nickname
}