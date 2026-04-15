package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card{
        case "ace":
        	return 11
        case "ten", "jack", "queen", "king":
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
        default:
        	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.

// FirstTurn determines the decision for the first turn based on the cards.
func FirstTurn(card1, card2, dealerCard string) string {
	myScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
	// 1. If you have a pair of aces, always split (P).
	case card1 == "ace" && card2 == "ace":
		return "P"

	// 2. Blackjack logic (Score of 21)
	case myScore == 21:
		// If dealer has Ace, 10, or face card (10+ points), Stand (S).
		if dealerScore >= 10 {
			return "S"
		}
		// Otherwise, Automatically win (W).
		return "W"

	// 3. Score between 17 and 20: Always Stand (S).
	case myScore >= 17 && myScore <= 20:
		return "S"

	// 4. Score between 12 and 16: Stand (S) unless dealer has 7 or higher.
	case myScore >= 12 && myScore <= 16:
		if dealerScore >= 7 {
			return "H" // Hit
		}
		return "S" // Stand

	// 5. Score 11 or lower: Always Hit (H).
	default:
		return "H"
	}
}

