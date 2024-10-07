package data

import (
	"testing"

	"memcards.ristomcintosh.com/internal/validator"
)

func TestValidateDeck(t *testing.T) {
	tests := []struct {
		name        string
		deck        *Deck
		expectedMsg string
	}{
		{"deck name is required", &Deck{}, "name is required"},
		{"deck name should not be empty", &Deck{Name: ""}, "name is required"},
		{"deck name should be at least 3 chars long", &Deck{Name: "ab"}, "name should be at least 3 characters long"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.New()
			ValidateDeck(v, tt.deck)

			got := v.Errors["name"]
			want := tt.expectedMsg

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}

		})
	}
}
