package newsfeed

import (
	"testing"
)

func TestAdd(t *testing.T)  {
	feed := New()
	feed.Add(Item{Title: "some title", Post: "Some Post"})
	if len(feed.Items) == 0 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T)  {
	feed := New()
	feed.Add(Item{})
	feed.Add(Item{})
	feed.Add(Item{})
	feed.Add(Item{})
	if len(feed.GetAll()) != 4 {
		t.Errorf("Unable to retrieve all items")
	} 
}
