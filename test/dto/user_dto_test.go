package dto

import (
	"seriesmanager-services/dto"
	"testing"
)

func TestTrimSpace(t *testing.T) {
	got := dto.UserCreateDto{
		Email: "test@test.com	",
		Username: "    garrou		",
		Password: "	testtest		",
		Confirm: "testtest		",
	}
	want := dto.UserCreateDto{
		Email:    "test@test.com",
		Username: "garrou",
		Password: "testtest",
		Confirm:  "testtest",
	}
	got.TrimSpace()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestIsValidPasswordLenNok(t *testing.T) {
	user := dto.UserCreateDto{
		Email:    "test@test.com",
		Username: "garrou",
		Password: "test",
		Confirm:  "test",
	}
	got := user.IsValid()

	if got != false {
		t.Errorf("got %t, wanted true", got)
	}
}

func TestIsValidPasswordDiffNok(t *testing.T) {
	user := dto.UserCreateDto{
		Email:    "test@test.com",
		Username: "garrou",
		Password: "test",
		Confirm:  "testtest",
	}
	got := user.IsValid()

	if got != false {
		t.Errorf("got %t, wanted false", got)
	}
}

func TestIsValidUsernameNok(t *testing.T) {
	user := dto.UserCreateDto{
		Email:    "test@test.com",
		Username: "gu",
		Password: "testtest",
		Confirm:  "testtest",
	}
	got := user.IsValid()

	if got != false {
		t.Errorf("got %t, wanted true", got)
	}
}

func TestIsValidOk(t *testing.T) {
	user := dto.UserCreateDto{
		Email:    "test@test.com",
		Username: "garrou",
		Password: "testtest",
		Confirm:  "testtest",
	}
	got := user.IsValid()

	if got != true {
		t.Errorf("got %t, wanted true", got)
	}
}
