package version

import (
	bver "github.com/carzil/go-test/b/version"
)

func Version() string {
	return "b-" + bver.Version() + ", c-v1.1"
}
