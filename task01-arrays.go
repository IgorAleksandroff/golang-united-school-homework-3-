package homework

func average(input [15]float32) (result float32) {
	for _, num := range input {
		result += num
	}
	return result / float32(len(input))
}
