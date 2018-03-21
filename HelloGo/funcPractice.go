package main

func main() {
	
}

//计算float64类型的slice的平均值
func average(slice []float64) (avg float64) {
	sum := 0.0
	switch len(slice) {
	case 0:
		avg = 0
	default:
		for _, v := range slice {
			sum += v
		}
		avg = sum /float64(len(slice))

	}
}