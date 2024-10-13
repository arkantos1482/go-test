package main

import (
	"fmt"
	"strings"
)

func main() {

	oldString := "0x38756901815c531d730f9793690656bf01B732D7 0x141b1Be5a4D73ED486e82e1930Df3Acb700d3b35 0x8d6d6d99D1a586619e822cdB6447d7136951D9a7 0x584067948BF575C06025A17574B60F76B3de6484 0x2Ed7D951b3bA8EA1FF717b922f0A315685901D44 0x6a827f4cb63a037241f406A2F6FF3dD32b80EA93 0x442F07C79E4b83CE9ed566CaC1b319088f588316 0x60B767Bc4786B5585bABA61e30c9cF7cc099976D"
	newString := "0x38756901815c531d730f9793690656bf01B732D7 0x141b1Be5a4D73ED486e82e1930Df3Acb700d3b35 0x8d6d6d99D1a586619e822cdB6447d7136951D9a7 0x584067948BF575C06025A17574B60F76B3de6484 0x2Ed7D951b3bA8EA1FF717b922f0A315685901D44 0x6a827f4cb63a037241f406A2F6FF3dD32b80EA93 0x442F07C79E4b83CE9ed566CaC1b319088f588316 0x60B767Bc4786B5585bABA61e30c9cF7cc099976D"

	oldArray := strings.Split(oldString, " ")
	newArray := strings.Split(newString, " ")

	// compare new array with old array and give a list of add, remove, update and unchanged
	added := []string{}
	removed := []string{}
	updated := []string{}
	unchanged := []string{}

	for _, newItem := range newArray {
		found := false
		for _, old := range oldArray {
			if newItem == old {
				found = true
				break
			}
		}
		if !found {
			added = append(added, newItem)
		}
	}

	for _, old := range oldArray {
		found := false
		for _, newItem := range newArray {
			if old == newItem {
				found = true
				break
			}
		}
		if !found {
			removed = append(removed, old)
		}
	}

	for _, newItem := range newArray {
		found := false
		for _, old := range oldArray {
			if newItem == old {
				found = true
				break
			}
		}
		if found {
			updated = append(updated, newItem)
		}
	}

	for _, old := range oldArray {
		found := false
		for _, newItem := range newArray {
			if old == newItem {
				found = true
				break
			}
		}
		if found {
			unchanged = append(unchanged, old)
		}
	}

	// print the result
	fmt.Printf("Added: %v\n", added)
	fmt.Printf("Removed: %v\n", removed)
	fmt.Printf("Updated: %v\n", updated)
	fmt.Printf("Unchanged: %v\n", unchanged)

}
