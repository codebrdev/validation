package validation

import (
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	if got, err := Name("  Ana  "); err != nil || got != "Ana" {
		t.Errorf("esperava Ana, obteve %q (%v)", got, err)
	}
	for _, in := range []string{"", "   ", strings.Repeat("a", MaxTextLen+1)} {
		if _, err := Name(in); err != ErrInvalidName {
			t.Errorf("%q: esperava ErrInvalidName, obteve %v", in, err)
		}
	}
	if _, err := Name(strings.Repeat("é", MaxTextLen)); err != nil {
		t.Errorf("o limite deve contar caracteres, não bytes: %v", err)
	}
}

func TestEmail(t *testing.T) {
	if got, err := Email(" ANA@Example.com "); err != nil || got != "ana@example.com" {
		t.Errorf("esperava ana@example.com, obteve %q (%v)", got, err)
	}
	for _, in := range []string{"", "nao-email", "Ana <a@b.com>", "a@b.com, c@d.com", strings.Repeat("a", 250) + "@b.com"} {
		if _, err := Email(in); err != ErrInvalidEmail {
			t.Errorf("%q: esperava ErrInvalidEmail, obteve %v", in, err)
		}
	}
}

func TestPassword(t *testing.T) {
	if err := Password("12345678"); err != nil {
		t.Errorf("senha válida rejeitada: %v", err)
	}
	for _, in := range []string{"", "1234567", strings.Repeat("a", MaxPasswordLen+1)} {
		if err := Password(in); err != ErrInvalidPassword {
			t.Errorf("%d bytes: esperava ErrInvalidPassword, obteve %v", len(in), err)
		}
	}
}
