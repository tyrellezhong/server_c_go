package generics

import "fmt"

// SumIntsOrFloats 计算map中所有值的总和
// 使用泛型实现，支持int64和float64类型的值
// 参数:
//   - K: 可比较的key类型
//   - V: int64或float64类型的value
//   - m: 输入的map
//
// 返回值:
//   - V类型的求和结果
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func TestSumIntsOrFloats() {
	intMap := map[string]int64{"a": 1, "b": 2, "c": 3}
	floatMap := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	fmt.Printf("Sums %d, %f", SumIntsOrFloats[string, int64](intMap), SumIntsOrFloats(floatMap))
}
