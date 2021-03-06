package replacer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendReplacableKV(t *testing.T) {
	replacerKV := []ReplacableKV{}
	addReplacableKv(&replacerKV, "hello=world", true)
	assert.Equal(t, replacerKV[0].KV, "hello=world")
}

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
	replacer := NewReplacer(
		"nginx-deployment.yaml",
		"out.yaml",
		true,
	)

	t.Run("Read input file path to strin", func(t *testing.T) {
		err := replacer.Run(replacerKV, true)
		assert.Nil(t, err)
	})
}

func TestCobraCLI(t *testing.T) {
	t.Run("create CLI must be no error", func(t *testing.T) {
		cli := RootCLI()
		err := cli.Execute()
		assert.Nil(t, err)
	})
}
