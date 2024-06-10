package helloWorld

import (
	"github.com/stretchr/testify/assert"
	"tbTool/tests"
	"testing"
)

var (
	_ = tests.DioInit()
)

func TestGetInfo(t *testing.T) {
	result := NewHelloWorldServiceImpl().GetInfo()
	assert.NotEmpty(t, result)
	assert.IsType(t, "", result)
}
