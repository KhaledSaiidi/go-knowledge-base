package lab

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
