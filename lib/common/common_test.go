package common_test

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/assert"
)

func Test_GetLines(t *testing.T) {
	str := `line1
line2
line3`
	assert.Equal(t, []string{"line1", "line2", "line3"}, common.GetLines(str))
}
