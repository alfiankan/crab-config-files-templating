package replacer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {

	replacerKV := []ReplacableKV{}
	replacerKV = append(replacerKV, ReplacableKV{
		KV:     "namespace=development",
		Quotes: false,
	})
	replacerKV = append(replacerKV, ReplacableKV{
		KV:     "traceUrl=http://trace.productio.com",
		Quotes: true,
	})

	t.Run("Read input file path to strin", func(t *testing.T) {
		replacer := NewReplacer(
			"nginx-deployment.yaml",
			"out.yaml",
			true,
		)
		err := replacer.Run(replacerKV, true)
		assert.Nil(t, err)
	})
}
