package variables

import (
	"fmt"

	"github.com/fermyon/spin-go-sdk/internal/fermyon/spin/v2.0.0/variables"
	"github.com/fermyon/spin-go-sdk/internal/wasi/cli/v0.2.0/stdout"
	"github.com/ydnar/wasm-tools-go/cm"
)

// Get an application variable value for the current component.
//
// The name must match one defined in in the component manifest.
func Get(key string) (string, error) {
	result := variables.Get(key)
	if result.IsErr() {
		return "", errorVariantToError(*result.Err())
	}

	stdout.GetStdout().Write(cm.ToList([]byte(fmt.Sprintf("key: %s, val is %q \n", key, *result.OK()))))
	return *result.OK(), nil
}

func errorVariantToError(err variables.Error) error {
	switch {
	case err.InvalidName() != nil:
		return fmt.Errorf(*err.InvalidName())
	case err.Provider() != nil:
		return fmt.Errorf(*err.Provider())
	case err.Undefined() != nil:
		return fmt.Errorf(*err.Undefined())
	default:
		if err.Other() != nil {
			return fmt.Errorf(*err.Other())
		}

		return fmt.Errorf("no error provided by host implementation")
	}
}
