package ghselfupdate

import "testing"

func TestGetPkgSlug(t *testing.T) {
	if getPkgSlug() != "Songmu/ghselfupdate" {
		t.Errorf("something went wrong")
	}
}
