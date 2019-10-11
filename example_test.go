package semver

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/version"
)

func ExampleSemVer() {
	fmt.Printf("%v\n", version.MustParseSemantic("3.1.0").LessThan(version.MustParseSemantic("3.1.0-rev4")))
	fmt.Printf("%v\n", version.MustParseSemantic("3.1.0-rev4").LessThan(version.MustParseSemantic("3.1.0")))
	fmt.Printf("%v\n", version.MustParseSemantic("3.1.0-rev4").LessThan(version.MustParseSemantic("3.1.0-rev5")))
	fmt.Printf("%v\n", version.MustParseSemantic("3.1.0-rev4").LessThan(version.MustParseSemantic("3.2.0")))

	// Output:
	// false
	// true
	// true
	// true
	// false
}
