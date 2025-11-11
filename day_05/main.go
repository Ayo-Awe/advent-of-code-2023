package main

import (
	"fmt"
	"log"
	"math"
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
			offset := mapInfo.dest - mapInfo.src
			return src + offset
		}

	}

	// any source number that isn't mapped corresponds to the same dest number
	return src
}

func (c *CategoryMap) RangeToDestination(rangeStart, rangeEnd int) [][]int {
	sourceRanges := [][]int{{rangeStart, rangeEnd}}
	var destRanges [][]int

outer:
	for i := 0; i < len(sourceRanges); i++ {
		rge := sourceRanges[i]
		start, end := rge[0], rge[1]

		for _, mapInfo := range c.Content {
			destRange, unmapped := mapInfo.ApplyRange(start, end)

			if len(destRange) != 0 {
				if len(unmapped) != 0 {
					sourceRanges = append(sourceRanges, unmapped...)
				}

				destRanges = append(destRanges, destRange)
				continue outer
			}
		}

		// any source range that isn't mapped corresponds to the same dest range
		destRanges = append(destRanges, []int{start, end})
	}

	return destRanges
}

func (c *CategoryMap) AddMapInfo(mapInfo *MapInfo) {
	c.Content = append(c.Content, *mapInfo)
}

func (mi *MapInfo) ApplyRange(srcStart, srcEnd int) (destRange []int, unmapped [][]int) {
	rangeStart, rangeEnd := mi.src, mi.src+mi.length
	// no overlap
	if srcEnd <= rangeStart || rangeEnd <= srcStart {
		return nil, [][]int{{srcStart, srcEnd}}
	}

	overlapStart := int(math.Max(float64(srcStart), float64(rangeStart)))
	overlapEnd := int(math.Min(float64(srcEnd), float64(rangeEnd)))

	offset := mi.dest - mi.src
	destRange = []int{overlapStart + offset, overlapEnd + offset}

	if overlapStart > srcStart {
		unmapped = append(unmapped, []int{srcStart, overlapStart})
	}

	if overlapEnd < srcEnd {
		unmapped = append(unmapped, []int{overlapEnd, srcEnd})
	}

	return destRange, unmapped
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

func (a Almanac) GetDestinationRanges(inputRanges [][]int, srcCategory string, destCategory string) [][]int {
	targetCategory, currentCategory := destCategory, srcCategory
	srcRanges := inputRanges

	for currentCategory != targetCategory {

		categoryMap, ok := a[currentCategory]
		if !ok {
			log.Fatalf("Category not found: %v\n", currentCategory)
		}

		// Parse to destination ranges
		var destRanges [][]int
		for _, srcRange := range srcRanges {
			d := categoryMap.RangeToDestination(srcRange[0], srcRange[1])
			destRanges = append(destRanges, d...)
		}

		srcRanges = destRanges
		currentCategory = categoryMap.Destination
	}

	return srcRanges
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
	puzzleTwo(seeds, almanac)
}

func puzzleTwo(seeds []int, almanac *Almanac) {
	var seedRanges [][]int

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := seeds[i] + seeds[i+1]
		seedRanges = append(seedRanges, []int{start, end})
	}

	locationRanges := almanac.GetDestinationRanges(seedRanges, "seed", "location")
	minLocation := 1000000000000
	for _, loc := range locationRanges {
		if loc[0] < minLocation {
			minLocation = loc[0]
		}
	}

	fmt.Printf("Puzzle two solution: %v\n", minLocation)
}
