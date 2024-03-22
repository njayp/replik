package pkg

import (
	"encoding/json"
	"testing"
)

type Item struct {
	Value int
	Next  *Item
}

func TestRecJSON(t *testing.T) {
	ll := Item{Value: 3, Next: &Item{Value: 4}}
	bytes, err := json.Marshal(ll)
	if err != nil {
		t.Fatal(err)
	}
	box := Item{}
	err = json.Unmarshal(bytes, &box)
	if err != nil {
		t.Fatal(err)
	}
	if box.Next.Value != 4 {
		t.Error("???")
	}
}

func TestClosedCh(t *testing.T) {
	ch := make(chan *struct{})
	close(ch)
	a := <-ch
	t.Log(a)
}
