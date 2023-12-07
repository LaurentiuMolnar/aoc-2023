package solutions

import "testing"

func TestIsFive(t *testing.T) {
	hand := "AAAAA"
	ok := IsFive(hand)

	if !ok {
		t.Fatal("IsFive failed for", hand)
	}
}

func TestIsFour(t *testing.T) {
	hands := []string{"1111A", "12222"}

	for _, hand := range hands {
		if !IsFour(hand) {
			t.Fatal("IsFour failed for", hand)
		}
	}
}

func TestIsFull(t *testing.T) {
	goodHands := []string{"11122", "AA222"}
	badHands := []string{"11112", "AA122", "AAAAA"}

	for _, h := range goodHands {
		if !IsFull(h) {
			t.Fatal("IsFull failed for", h)
		}
	}

	for _, h := range badHands {
		if IsFull(h) {
			t.Fatal("IsFull flagged non-full as full house:", h)
		}
	}
}

func TestIsPair(t *testing.T) {
	goodHands := []string{"1122A", "A2233"}
	badHands := []string{"222A3", "11222"}

	for _, h := range goodHands {
		ok, _ := IsPair(h)
		if !ok {
			t.Fatal("IsPair failed for", h)
		}
	}

	for _, h := range badHands {
		ok, _ := IsPair(h)
		if ok {
			t.Fatal("IsPair flagged non-pair as pair:", h)
		}
	}
}

func TestIsThree(t *testing.T) {
	goodHands := []string{"1112A", "A2223", "12333"}
	badHands := []string{"11222", "11112", "11111"}

	for _, h := range goodHands {
		ok := IsThree(h)
		if !ok {
			t.Fatal("IsThree failed for", h)
		}
	}

	for _, h := range badHands {
		ok := IsThree(h)
		if ok {
			t.Fatal("IsThree flagged non-three as three:", h)
		}
	}
}

func TestIsTwoPair(t *testing.T) {
	goodHands := []string{"1122A", "A2233"}
	badHands := []string{"222A3", "11222", "11A23", "A2234"}

	for _, h := range goodHands {
		ok := IsTwoPair(h)
		if !ok {
			t.Fatal("IsTwoPair failed for", h)
		}
	}

	for _, h := range badHands {
		ok := IsTwoPair(h)
		if ok {
			t.Fatal("IsTwoPair flagged non-two-pair as two-pair:", h)
		}
	}
}
