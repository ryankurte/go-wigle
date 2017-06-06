package wigle

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWiggle(t *testing.T) {

	w := NewWiGLE()

	username, password := os.Getenv("WIGLE_USER"), os.Getenv("WIGLE_PASS")
	if username == "" || password == "" {
		t.Errorf("Could not find WiGLE username or password")
		t.FailNow()
	}

	t.Run("Fetches login cookie", func(t *testing.T) {
		err := w.Login(username, password)
		assert.Nil(t, err)
	})

}
