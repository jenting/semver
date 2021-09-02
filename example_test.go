package semver

import (
	"fmt"

	"github.com/blang/semver"
)

func ExampleSemVer() {
	_, err := semver.ParseRange(">=v1.16.0-0")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	_, err = semver.ParseRange(">=1.16.0-0")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// Output:
	// Could not parse Range ">=v1.16.0-0": Could not parse comparator ">=v" in ">=v1.16.0-0"
}
