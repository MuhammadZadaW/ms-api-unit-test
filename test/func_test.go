package test

import (
	"ms-api-unit-test/utility"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCheckAlphaSpace(t *testing.T) {
	result := utility.CheckAlphaSpace("Hello World")
	assert.Equal(t, true, result)
}

func TestCheckAlphaSpaceMultiple(t *testing.T) {

	dataList := []struct{
		id int
		data string
		expected bool
	}{
		{1, "Hello World", true},
		{2, "HelloWorld", true},
		{3, "Hello World!", false},
		{4, "123", false},
		{5, "123a", false},
	}

	for _, data := range dataList {
		t.Run(data.data, func(t *testing.T) {
			result := utility.CheckAlphaSpace(data.data)
			assert.Equal(t, data.expected, result)
		})
	}

}
