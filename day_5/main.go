package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/ayo-awe/advent-of-code-2023/util"
)

type CategoryMap struct {
	Source      string
	Destination string
	Content     []MapInfo
}

type MapInfo struct {
	src    int
	dest   int
	length int
}

func NewCategoryMap(src string, dest string, content []MapInfo) *CategoryMap {
	return &CategoryMap{Source: src, Destination: dest, Content: content}
}

func NewMapInfo(dest, src, length int) *MapInfo {
	return &MapInfo{src, dest, length}
}

func (c *CategoryMap) ToDestination(src int) int {
	for _, mapInfo := range c.Content {
		rangeStart, rangeEnd := mapInfo.src, mapInfo.src+mapInfo.length
		withinRange := src >= rangeStart && src < rangeEnd

		// calc dest value
		if withinRange {
			offset := src - mapInfo.src
			return mapInfo.dest + offset
		}

	}

	// any source number that isn't mapped corresponds to the same dest number
	return src
}

func (c *CategoryMap) AddMapInfo(mapInfo *MapInfo) {
	c.Content = append(c.Content, *mapInfo)
}

type Almanac map[string]*CategoryMap

func NewAlmanac() *Almanac {
	return &Almanac{}
}

func (a Almanac) AddEntry(key string, c *CategoryMap) {
	a[key] = c
}

func (a Almanac) GetDestination(src int, srcCategory string, destCategory string) int {
	targetCategory, currentCategory := destCategory, srcCategory
	srcValue := src

	for currentCategory != targetCategory {
		categoryMap, ok := a[currentCategory]
		if !ok {
			log.Fatalf("Category not found: %v\n", currentCategory)
		}

		// Parse to destination value
		srcValue = categoryMap.ToDestination(srcValue)
		currentCategory = categoryMap.Destination
	}

	return srcValue
}

func Initialize() ([]int, *Almanac) {
	lines, err := util.GetInput()
	if err != nil {
		log.Fatalf("Failed: %v\n", err)
	}

	seedStrings := strings.Split(strings.TrimPrefix(lines[0], "seeds: "), " ")
	seeds := util.ToIntArray(seedStrings)

	almanac := NewAlmanac()
	var currentMap *CategoryMap

	for _, line := range lines[1:] {
		if line == "" && currentMap != nil {
			almanac.AddEntry(currentMap.Source, currentMap)
			currentMap = nil
		} else if strings.Contains(line, "map") {

			mapDetails := strings.Split(line, " ")
			srcDest := strings.Split(mapDetails[0], "-to-")
			currentMap = NewCategoryMap(srcDest[0], srcDest[1], nil)
		} else {
			if line == "" {
				continue
			}

			data := util.ToIntArray(strings.Split(line, " "))
			mapInfo := NewMapInfo(data[0], data[1], data[2])

			if currentMap == nil {
				log.Fatal("Unexpected nil current map")
			}

			currentMap.AddMapInfo(mapInfo)
		}

	}

	if currentMap != nil {
		almanac.AddEntry(currentMap.Source, currentMap)
	}

	return seeds, almanac
}

func puzzleOne(seeds []int, almanac *Almanac) {
	var locations []int

	for _, seed := range seeds {
		location := almanac.GetDestination(seed, "seed", "location")
		locations = append(locations, location)
	}

	fmt.Printf("Puzzle one solution: %v\n", slices.Min(locations))
}

func main() {
	seeds, almanac := Initialize()
	puzzleOne(seeds, almanac)
}
