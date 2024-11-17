package stats

import (
	"math"
	"sort"
)

func GetMean(array []int) float64{
	sum := 0;
	for _, value := range array{
		sum += value
	}
	if len(array) == 0{
		return 0
	}
	return float64(sum)/float64(len(array))
}

func GetMedian(array []int) float64{
	if len(array) == 0{
		return 0
	}
	if len(array) % 2 == 0{
		return GetMean(array)
	}else{
		return float64(array[len(array)/2])
	}
}

func GetMode(array []int) (mode int){ 

	if len(array) == 0{
		return 0
	}
	sort.Ints(array)
	countMap := make(map[int] int)
	for _, value := range array{
		countMap[value] += 1
	}
	maxСount := 0
	for key, count := range countMap{
		if  count > maxСount{
			mode = key
			maxСount = count
		}
	}
	return
}

func GetStandardDeviation(array []int) float64{
	if len(array) == 0{
		return 0
	}
	average := GetMean(array)
	var sum float64
	for _, value := range array{
		sum += math.Pow((float64(value) - average),2)
	}
	return math.Sqrt(sum/float64(len(array)))
}