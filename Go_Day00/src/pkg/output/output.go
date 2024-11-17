package output

import (
	"GoDay0/pkg/flagParser"
	"GoDay0/pkg/stats"
	"fmt"
)

func Output(numbers []int, userFlags flagParser.Flags) {
	if userFlags.Mean{
		fmt.Printf("Mean: %.2f\n", stats.GetMean(numbers))
	}
	if userFlags.Median{
		fmt.Printf("Median: %.2f\n", stats.GetMedian(numbers))
	}
	if userFlags.Mode{
		fmt.Printf("Mode: %d\n", stats.GetMode(numbers))
	}
	if userFlags.SD{
		fmt.Printf("SD: %.2f\n", stats.GetStandardDeviation(numbers))
	}

}