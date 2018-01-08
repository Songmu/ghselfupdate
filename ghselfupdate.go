package ghselfupdate

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func init() {
	if _, ok := os.LookupEnv("GHSELFUPDATE_DEBUG"); ok {
		selfupdate.EnableLog()
	}
}

var pkgSlug string
var slugReg = regexp.MustCompile(`/(?:src|vendor)/github.com/([^/]+/[^/]+)`)

func getPkgSlug() string {
	if pkgSlug == "" {
		_, fname, _, _ := runtime.Caller(1)
		slug, err := fname2Slug(fname)
		if err != nil {
			log.Fatal(err)
		}
		pkgSlug = slug
	}
	return pkgSlug
}

func fname2Slug(s string) (string, error) {
	s = filepath.ToSlash(s)
	matches := slugReg.FindStringSubmatch(s)
	if len(matches) < 2 {
		return "", fmt.Errorf("can't detect github slug (owner/repo) from binary")
	}
	return matches[1], nil
}

// Do binary self update
func Do(ver string) error {
	semv, err := semver.Parse(ver)
	if err != nil {
		return err
	}
	slug, err := detectSlug()
	if err != nil {
		return err
	}

	latest, err := selfupdate.UpdateSelf(semv, slug)
	if err != nil {
		return err
	}
	if latest.Version.Equals(semv) {
		log.Println("Current binary is the latest version", ver)
	} else {
		log.Println("Successfully updated to version", latest.Version)
	}
	return nil
}

func detectSlug() (slug string, err error) {
	// 1 : this private function `detectSlug`
	// 2 : caller (function in this package. e.g. `Do`)
	// 3-: extern
	for i := 3; true; i++ {
		_, fname, _, ok := runtime.Caller(i)
		if !ok {
			return "", fmt.Errorf("no slug deteced from caller")
		}
		fname = filepath.ToSlash(fname)
		if strings.Contains(fname, getPkgSlug()) {
			continue
		}
		slug, err = fname2Slug(fname)
		break
	}
	return
}
