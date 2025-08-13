package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to name", func(t *testing.T) {
		got := Hello("Luiz", "en")
		want := "Hello, Luiz"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' if no name is supplied", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola, nombre', if spanish is selected", func(t *testing.T) {
		got := Hello("Luiz", "es")
		want := "Hola, Luiz"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola, mundo', if spanish name is not supplied", func(t *testing.T) {
		got := Hello("", "es")
		want := "Hola, mundo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("if lang is missing, should return in english", func(t *testing.T) {
		got := Hello("Luiz", "")
		want := "Hello, Luiz"

		assertCorrectMessage(t, got, want)
	})
	t.Run("if lang and name is missing, should return Hello, world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
