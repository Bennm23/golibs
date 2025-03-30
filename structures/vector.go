package structures

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Vector[T Number] []T

func (vec Vector[T]) Times(multiplier T) Vector[T] {
	res := make([]T, len(vec))
	for i, val := range vec{
		res[i] = val * multiplier
	}
	return res
}
func (vec Vector[T]) Divide(divisor T) Vector[T] {
	res := make([]T, len(vec))
	for i, val := range vec{
		res[i] = val / divisor
	}
	return res
}
func (vec Vector[T]) Plus(other Vector[T]) Vector[T] {
	res := make([]T, len(vec))
	for i, val := range vec{
		res[i] = val + other[i]
	}
	return res
}
func (vec Vector[T]) Minus(other Vector[T]) Vector[T] {
	res := make([]T, len(vec))
	for i, val := range vec{
		res[i] = val - other[i]
	}
	return res
}

func (vec Vector[T]) SimpleCross(other Vector[T]) T {
	if len(vec) != 2 || len(other) != 2 {
		panic("Can't Simple Cross These")
	}
	return (vec[0]*other[1]) - (vec[1]*other[0])
}