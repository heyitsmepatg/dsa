package main

func main() {
	reverseBits(uint32(43261596))
}

func reverseBits(num uint32) uint32 {
	// fmt.Println(bits.Reverse32(num))
	// return bits.Reverse32(num)
	ret := uint32(0)
	power := uint32(31)
	for num != 0 {
		ret += (num & 1) << power
		num = num >> 1
		power -= 1
	}
	return ret
}
