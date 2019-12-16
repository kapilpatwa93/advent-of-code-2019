package main

import (
	"fmt"
	"strings"
)

func main() {
	input := [] string {
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	orbitMap, parentMap := getOrbitMap(&input)
	getHopCount(orbitMap, parentMap)

}
func isPresent(obj string, orbitMap *map[string][]string, queue []string) (func(string) bool) {
	return func(parent string) bool{
		isPresent := false

		for _, orbiter := range (*orbitMap)[parent] {
			if orbiter == obj {
				isPresent = true
				break
			}
			queue = append(queue, orbiter)
		}
		return isPresent
	}
}
func getHopCount(orbitMap *map[string][]string, parentMap *map[string]string) {

	youParent := (*parentMap)["YOU"]
	objQueue := []string{}
	objQueue = append(objQueue, youParent)
	isSantaPresent := isPresent("SAN", orbitMap, objQueue)
	found := false
	parent := youParent
	var obj string
	for i:=0; i<len(*parentMap); i++ {

		obj = objQueue[0]
		fmt.Println("obj", obj)
		if len(objQueue) > 1 {
			objQueue = objQueue[1:]
		} else {
			objQueue = []string{}
		}
		found = isSantaPresent(parent)
		if found {
			break
		}
	}
	fmt.Print("found", found, obj)

}

func getOrbitMap(input *[]string) (*map[string][]string,*map[string]string) {
	orbitMap := map[string][]string{}
	parentMap := map[string]string{}
	for i := range *input {
		var str = strings.Split((*input)[i],")")
		 val := orbitMap[str[0]]
		parentMap[str[1]] = str[0]
		 	orbitMap[str[0]] = append(val, str[1])
	}
	return &orbitMap, &parentMap
}
