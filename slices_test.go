package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type product struct {
	ProductFid string
	Name       string
	Price      int
}

func TestPPull(t *testing.T) {
	products := []product{
		{ProductFid: "abc-123", Name: "Widget", Price: 100},
		{ProductFid: "def-456", Name: "Gadget", Price: 200},
		{ProductFid: "ghi-789", Name: "Doohickey", Price: 300},
	}

	result, err := PPull(products, "ProductFid")

	assert.NoError(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, "Widget", result["abc-123"].Name)
	assert.Equal(t, "Gadget", result["def-456"].Name)
	assert.Equal(t, "Doohickey", result["ghi-789"].Name)
}

func TestPPull_IntKey(t *testing.T) {
	products := []product{
		{ProductFid: "abc", Name: "Widget", Price: 100},
		{ProductFid: "def", Name: "Gadget", Price: 200},
	}

	result, err := PPull(products, "Price")

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Widget", result["100"].Name)
	assert.Equal(t, "Gadget", result["200"].Name)
}

func TestPPull_Empty(t *testing.T) {
	result, err := PPull([]product{}, "ProductFid")

	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestPPull_DuplicateKeys(t *testing.T) {
	products := []product{
		{ProductFid: "abc", Name: "First"},
		{ProductFid: "abc", Name: "Second"},
	}

	result, err := PPull(products, "ProductFid")

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Second", result["abc"].Name) // last wins
}

func TestPPull_InvalidField(t *testing.T) {
	products := []product{{ProductFid: "abc", Name: "Widget"}}

	_, err := PPull(products, "Missing")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestPPull_Pointer(t *testing.T) {
	products := []*product{
		{ProductFid: "abc-123", Name: "Widget"},
		{ProductFid: "def-456", Name: "Gadget"},
	}

	result, err := PPull(products, "ProductFid")

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Widget", result["abc-123"].Name)
	assert.Equal(t, "Gadget", result["def-456"].Name)
}
