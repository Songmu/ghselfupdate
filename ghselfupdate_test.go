package ghselfupdate

import (
	"testing"
)

func TestDetectSlug(t *testing.T) {
	_, err := detectSlug()
	if err == nil {
		t.Errorf("error should be occured but nil")
	}
}
