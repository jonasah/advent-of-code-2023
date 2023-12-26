package day25

import (
	"testing"

	"github.com/jonasah/advent-of-code-2023/lib/common"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 54, Part1(testInput, []string{"jqt nvd", "bvb cmg", "hfx pzl"}))
	require.Equal(t, 598120, Part1(realInput, []string{"zmq pgh", "ldk bkm", "rsm bvc"}))
}

var realInput = common.GetInput(25)

const testInput = `jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr`
