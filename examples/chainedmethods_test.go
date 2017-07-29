package examples

import (
	"fmt"
	c "jecolasurdo/go-chaining"
	"jecolasurdo/go-chaining/behavior"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func addOne(value interface{}) (interface{}, error) {
	return value.(int) + 1, nil
}

func multiplyBySix(value interface{}) (interface{}, error) {
	return value.(int) * 6, nil
}

func convertToString(value interface{}) (interface{}, error) {
	return strconv.Itoa(value.(int)), nil
}

func sendToThePrinter(value interface{}) error {
	_, err := fmt.Print(value)
	return err
}

func Test_ChainItAllTogether(t *testing.T) {
	chain := c.New()
	var initialValue interface{} = 1
	initialBehavior := c.ActionArg{
		Value:    initialValue,
		Behavior: behavior.InjectSuppliedValue,
	}
	chain.ApplyUnaryIface(addOne, initialBehavior)
	// chain.ApplyUnaryIface(multiplyBySix, c.ActionArg{})
	// chain.ApplyUnaryIface(convertToString, c.ActionArg{})
	// chain.ApplyUnary(sendToThePrinter, c.ActionArg{})

	result, err := chain.Flush()
	//_, ok := result.(interface{}).
	//assert.True(t, ok)
	var flar interface{} = 2
	assert.Equal(t, flar, result)
	assert.Nil(t, err)
}