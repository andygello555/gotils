// Map functions such as deep copying and equality testing.
package maps

import (
	"fmt"
	"github.com/go-test/deep"
	"strings"
	"testing"
)

// Clones a map deeply using recursion.
func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}

// Used in tests to check equality between two interface{}s.
//
// This takes into account orderings of slices.
func JsonMapEqualTest(t *testing.T, actual, expected interface{}, forString string) {
	if diff := deep.Equal(actual, expected); diff != nil {
		var errB strings.Builder
		errB.WriteString(fmt.Sprintf("Difference between actual and expected for %s (Left = Actual, Right = Expected)\n", forString))
		for _, d := range diff {
			errB.WriteString(fmt.Sprintf("\t%s\n", d))
		}
		t.Error(errB.String())
	}
}
