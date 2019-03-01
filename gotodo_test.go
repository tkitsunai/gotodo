package gotodo_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tkitsunai/gotodo"
	"testing"
)

func TestTODO(t *testing.T) {
	assert.Panics(t, func() { gotodo.TODO("message") })
}

func TestTODOERR(t *testing.T) {
	err := gotodo.TODO_ERR("err")
	assert.Error(t, err, "err")
}
