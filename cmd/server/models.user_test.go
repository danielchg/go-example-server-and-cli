package main

import "testing"

func TestUsernameAvailability(t *testing.T) {
	saveList()

	if !isUserNameAvailable("newuser") {
		t.Fail()
	}

	if isUserNameAvailable("user1") {
		t.Fail()
	}

	registerNewUser("newuser", "newpass")

	if isUserNameAvailable("newuser") {
		t.Fail()
	}

	restoreLists()
}

func TestVAlidUserRegistration(t *testing.T) {
	saveList()

	u, err := registerNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()
}

func TestInvalidUserRegistration(t *testing.T) {
	saveList()

	u, err := registerNewUser("user1", "pass1")

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = registerNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}

func TestUserValidity(t *testing.T) {
	if !isUserValid("user1", "pass1") {
		t.Fail()
	}

	if isUserValid("user2", "pass1") {
		t.Fail()
	}

	if isUserValid("user1", "") {
		t.Fail()
	}

	if isUserValid("", "pass1") {
		t.Fail()
	}

	if isUserValid("User1", "pass1") {
		t.Fail()
	}
}
