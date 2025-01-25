package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king", "queen", "jack", "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	case "other":
		return 0
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	total := card1Value + card2Value

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"

	case total == 21: 
	if dealerValue == 10 || dealerCard == "ace" {
			return "S"
		}
		return "W"

	case total >= 17 && total <= 20:
		return "S"

	case total >= 12 && total <= 16:
		if dealerValue >= 7 {
			return "H"
		}
		return "S"

	case total <= 11:
		return "H"
	}
	return ""
}
