package lordeckoder

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"testing"
)

type decodeTest struct {
	deckCode string
	expected []string
}

type DecodeSuite struct {
	suite.Suite
	TestCases []decodeTest
}

func (s *DecodeSuite) SetupSuite() {
	file, err := os.Open("./test/DeckCodesTestData.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		test := decodeTest{deckCode: scanner.Text(), expected: []string{}}
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			test.expected = append(test.expected, line)
		}
		s.TestCases = append(s.TestCases, test)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (s *DecodeSuite) TestDecode() {
	for _, test := range s.TestCases {
		var cards []string
		deck, err := Decode(test.deckCode)
		for _, card := range deck.Cards {
			cards = append(cards, fmt.Sprintf("%d:%s", card.Count, card.Card))
		}
		s.NoError(err, "error decoding a deck code")
		s.ElementsMatch(cards, test.expected, "decks are not matching")
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(DecodeSuite))
}
