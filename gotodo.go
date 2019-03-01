package gotodo

import (
	"errors"
	"fmt"
	"strings"
)

var notImplementedError = errors.New("An operation is not implemented ")

func TODO(messages ...string) {
	if len(messages) > 0 {
		panic(fmt.Sprintf("TODO: An operation is not implemented :%v", strings.Join(messages,"\n")))
	}
	panic("TODO: An operation is not implemented.")
}

func TODO_ERR(messages ...string) error {
	if len(messages) > 0 {
		return errors.New(fmt.Sprintf("An operation is not implemented %s", strings.Join(messages,"\n")))
	}
	return notImplementedError
}
