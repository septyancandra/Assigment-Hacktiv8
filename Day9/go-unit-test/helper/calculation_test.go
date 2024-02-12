package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(10, 10)
	fmt.Printf("Hasil %d \n", result)

	assert.Equal(t, 20, result, " Result has to be 20")
	fmt.Println("TestSum Eksekusi terhenti")

}

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)
	require.Equal(t, 40, result, "Result has to be 40")
	fmt.Println("TestFailSum Eksekusi terhenti")

}
