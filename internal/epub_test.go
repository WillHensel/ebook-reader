package internal

import (
	"testing"
)

func TestReadEpub(t *testing.T) {
	path := "../test-data/epub/Dracula - Bram Stoker.epub"

	epub, err := ReadEpub(path)

	if err != nil {
		t.Errorf("Got error %v", err)
	}

	if epub.Title != "Dracula" {
		t.Errorf("epub.Title - wanted: 'Dracula', Got: '%s'", epub.Title)
	}
}
