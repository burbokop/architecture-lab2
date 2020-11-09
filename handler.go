package lab2

import (
	"io"
	"io/ioutil"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buf, err := ioutil.ReadAll(ch.Input)

	if err != nil && err != io.EOF {
		return err
	}
	res, err := PostfixToPrefix(string(buf))
	if err != nil {
		return err
	}
	_, err = ch.Output.Write([]byte(res))
	if err != nil {
		return err
	}
	return nil
}
