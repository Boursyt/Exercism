package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten":
		return 10
	case card == "jack":
		return 10
	case card == "queen":
		return 10
	case card == "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Calculate the sum of the player's cards and the dealer's card value.
	playerSum := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {
	// Rule 1: Always split a pair of aces.
	case card1 == "ace" && card2 == "ace":
		return "P"

	// Rule 2: Handle Blackjack (sum 21) BEFORE other score ranges.
	case playerSum == 21:
		// Rule 2a: Dealer does NOT have Ace or 10-value card -> Win
		if dealerCard != "ace" && dealerValue != 10 {
			return "W"
		}
		// Rule 2b: Dealer HAS an Ace or 10-value card -> Stand
		return "S" // This covers the (dealerCard == "ace" || dealerValue == 10) case

	// Rule 3: Score is [17, 20].
	case playerSum >= 17 && playerSum <= 20:
		return "S"

	// Rule 4: Score is [12, 16].
	case playerSum >= 12 && playerSum <= 16:
		// Rule 4b: Dealer has 7 or higher -> Hit
		if dealerValue >= 7 {
			return "H"
		}
		// Rule 4a: Dealer has less than 7 (2-6) -> Stand
		return "S"

	// Rule 5: Score is 11 or lower.
	case playerSum <= 11:
		return "H"

	// Default case (should not be reached if all rules are covered)
	default:
		return "H"
	}

}
