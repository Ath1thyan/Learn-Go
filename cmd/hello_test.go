package main

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"Athi"},
			result: "Hello, Athi!",
		},
		{
			items: []string{"Athi", "Matt"},
			result: "Hello, Athi, Matt!",
		},
	}

	// want := "Hello, test!"
	// got := Say([]string{"test"})

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}

	// if want != got {
	// 	t.Errorf("wanted %s, got %s", want, got)
	// }
}