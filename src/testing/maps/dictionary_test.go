package dictionary

import "testing"

func TestUpdate(t *testing.T) {
	dict := Dictionary{"test": "just a test"}

	dict.Update("test", "just another test")

	got, err := dict.Search("test")

	if err != nil {
		t.Fatal("error must not be nil")
	}

	assertStrings(t, got, "just another test")
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	dict.Add("test", "just a test")

	want := "just a test"
	got, err := dict.Search("test")

	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, want)
}

func TestSearch(t *testing.T) {

	dict := Dictionary{"test": "just a test"}
	t.Run("known word", func(t *testing.T) {
		got, err := dict.Search("test")
		want := "just a test"

		if err != nil {
			t.Fatal("error should not be nil")
		}

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected error")
		}

		assertErrors(t, err, want)
	})

}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}

