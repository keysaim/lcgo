package sort

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuick(t *testing.T) {

	require := require.New(t)

	a := []int{0, 0, 5, 9, 5, 1, 3, 8, 6, 6, 0, 6}
	log.Println(a)
	Quick(a)
	log.Println(a)
	require.True(IsSorted(a))

	a = RandList(60, 10)
	log.Println(a)
	Quick(a)
	log.Println(a)
	require.True(IsSorted(a))

	a = RandList(60, 100)
	log.Println(a)
	Quick(a)
	log.Println(a)
	require.True(IsSorted(a))

	a = RandList(60, 100)
	log.Println(a)
	Quick(a, true)
	log.Println(a)
	require.True(IsSorted(a))
}
