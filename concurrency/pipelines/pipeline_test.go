package pipelines

import (
	"testing"
)

func TestLaunchPipeline(t *testing.T) {
	data := [][]int{
		{3, 14},
		{5, 55},
	}

	var res int
	for _, tt := range data {
		res = LaunchPipeline(tt[0])
		if res != tt[1] {
			t.Fatal()
		}

		t.Logf("%d == %d\n", res, tt[1])
	}
}
