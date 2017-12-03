package main

import (
	"fmt"
	"log"
	"math"
)

var rows = [][]int{
	[]int{4347, 3350, 196, 162, 233, 4932, 4419, 3485, 4509, 4287, 4433, 4033, 207, 3682, 2193, 4223},
	[]int{648, 94, 778, 957, 1634, 2885, 1964, 2929, 2754, 89, 972, 112, 80, 2819, 543, 2820},
	[]int{400, 133, 1010, 918, 1154, 1008, 126, 150, 1118, 117, 148, 463, 141, 940, 1101, 89},
	[]int{596, 527, 224, 382, 511, 565, 284, 121, 643, 139, 625, 335, 657, 134, 125, 152},
	[]int{2069, 1183, 233, 213, 2192, 193, 2222, 2130, 2073, 2262, 1969, 2159, 2149, 410, 181, 1924},
	[]int{1610, 128, 1021, 511, 740, 1384, 459, 224, 183, 266, 152, 1845, 1423, 230, 1500, 1381},
	[]int{5454, 3936, 250, 5125, 244, 720, 5059, 202, 4877, 5186, 313, 6125, 172, 727, 1982, 748},
	[]int{3390, 3440, 220, 228, 195, 4525, 1759, 1865, 1483, 5174, 4897, 4511, 5663, 4976, 3850, 199},
	[]int{130, 1733, 238, 1123, 231, 1347, 241, 291, 1389, 1392, 269, 1687, 1359, 1694, 1629, 1750},
	[]int{1590, 1394, 101, 434, 1196, 623, 1033, 78, 890, 1413, 74, 1274, 1512, 1043, 1103, 84},
	[]int{203, 236, 3001, 1664, 195, 4616, 2466, 4875, 986, 1681, 152, 3788, 541, 4447, 4063, 5366},
	[]int{216, 4134, 255, 235, 1894, 5454, 1529, 4735, 4962, 220, 2011, 2475, 185, 5060, 4676, 4089},
	[]int{224, 253, 19, 546, 1134, 3666, 3532, 207, 210, 3947, 2290, 3573, 3808, 1494, 4308, 4372},
	[]int{134, 130, 2236, 118, 142, 2350, 3007, 2495, 2813, 2833, 2576, 2704, 169, 2666, 2267, 850},
	[]int{401, 151, 309, 961, 124, 1027, 1084, 389, 1150, 166, 1057, 137, 932, 669, 590, 188},
	[]int{784, 232, 363, 316, 336, 666, 711, 430, 192, 867, 628, 57, 222, 575, 622, 234},
}

func main() {
	part1 := 0
	part2 := 0

	for _, row := range rows {
		part1 += Checksum(row)
		part2 += Divsum(row)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func Checksum(input []int) int {
	var min = math.MaxInt32
	var max = math.MinInt32
	for _, v := range input {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	if min == math.MaxInt32 || max == math.MinInt32 {
		log.Fatalf("ERROR: min or max doesn't exist")
	}
	return max - min
}

func Divsum(input []int) int {
	var result = 0
	for i, a := range input {
		for j, b := range input {
			if i != j && a%b == 0 {
				result += a / b
			}
		}
	}
	return result
}