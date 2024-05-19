package llm

import (
	"fmt"

	"github.com/fermyon/spin-go-sdk/internal/fermyon/spin/v2.0.0/llm"
	"github.com/ydnar/wasm-tools-go/cm"
)

// The model use for inferencing
const (
	Llama2Chat        InferencingModel = "llama2-chat"
	CodellamaInstruct InferencingModel = "codellama-instruct"
)

type InferencingParams llm.InferencingParams
type InferencingResult llm.InferencingResult
type InferencingModel llm.InferencingModel
type EmbeddingsResult llm.EmbeddingsResult

// Infer performs inferencing using the provided model and prompt with the
// given optional parameters.
func Infer(model string, prompt string, params *InferencingParams) (InferencingResult, error) {
	var iparams = cm.None[llm.InferencingParams]()
	if params != nil {
		iparams = cm.Some(llm.InferencingParams(*params))
	}

	result := llm.Infer(llm.InferencingModel(model), prompt, iparams)
	if result.IsErr() {
		return InferencingResult{}, errorVariantToError(*result.Err())
	}

	return InferencingResult(*result.OK()), nil
}

// GenerateEmbeddings generates the embeddings for the supplied list of text.
func GenerateEmbeddings(model InferencingModel, text []string) (*EmbeddingsResult, error) {
	result := llm.GenerateEmbeddings(llm.EmbeddingModel(model), cm.ToList[[]string](text))
	if result.IsErr() {
		return &EmbeddingsResult{}, errorVariantToError(*result.Err())
	}

	embeddingResult := EmbeddingsResult(*result.OK())
	return &embeddingResult, nil
}

func errorVariantToError(err llm.Error) error {
	switch {
	case llm.ErrorModelNotSupported() == err:
		return fmt.Errorf("model not supported")
	case err.RuntimeError() != nil:
		return fmt.Errorf("runtime error %s", *err.RuntimeError())
	case err.InvalidInput() != nil:
		return fmt.Errorf("invalid input %s", *err.InvalidInput())
	default:
		return fmt.Errorf("no error provided by host implementation")
	}
}
