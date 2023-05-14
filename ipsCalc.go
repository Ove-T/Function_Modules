package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func ipsCalc(input string) (ip [4]uint8, subnetmask1 [4]uint8) {
	pattern := "\\d+"
	matcher := regexp.MustCompile(pattern)
	matched := matcher.FindAllString(input, -1)
	ip = [4]uint8{}
	if len(matched) == 5 {
		for i, _ := range ip {
			temp, _ := strconv.Atoi(matched[i])
			ip[i] = uint8(temp)
		}
	}
	subnetMask, _ := strconv.Atoi(matched[4])

	subnetmask1 = [4]uint8{}
	for i := 0; i < 4; i++ {
		if subnetMask >= 8 {
			subnetmask1[i] = 2<<7 - 1
			subnetMask = subnetMask - 8
		} else {
			if subnetMask == 0 {
				subnetmask1[i] = 0
			} else {
				out := 0
				for ; subnetMask >= 1; subnetMask-- {
					out += 2 << (7 - subnetMask)
				}
				subnetmask1[i] = uint8(out)
			}
		}

	}

	IpRange(ip, subnetmask1)
	return
}

func IpRange(ip [4]uint8, subnetmask [4]uint8) {
	// [255 255 255 0] --> [0 0 0 255]
	subnetmask1 := [4]uint8{}
	for i := 0; i < 4; i++ {
		subnetmask1[i] = ^subnetmask[i]
	}

	// [127.0.0.1] & [255 255 255 0] --> [127.0.0.0]
	ip1 := [4]uint8{}
	for i := 0; i < 4; i++ {
		ip1[i] = ip[i] & subnetmask[i]
	}

	// [127.0.0.0] | [0 0 0 255] --> [127 0 0 255]
	ipMax := ip1
	for i := 0; i < 4; i++ {
		ipMax[i] = ip1[i] | subnetmask1[i]
	}
	ipMin := ip1
	// [127 0 0 0] --> [127 0 0 1]
	ipMin[3] = ip1[3] + 1
	//ip2 := ipMax
	// [127 0 0 255] --> [127 0 0 254]
	ipMax[3] = ipMax[3] - 1
	fmt.Printf("%d,%d", ipMin, ipMax)
}

func main() {
	var a = "192.168.0.10/24"
	ipsCalc(a)
}
