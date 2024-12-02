package day5

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/kevinlowe0x3F7/aoc-go/shared"
)

// MapEntry holds the start and end names, along with the associated numbers
type MapEntry struct {
	StartName   string
	EndName     string
	Destination int
	Source      int
	Range       int
}

// ParsedData holds the seeds and the map entries
type ParsedData struct {
	Seeds   []int
	Entries []MapEntry
}

func Day5() {
	data, err := parseFile("./day5/day5pt1.txt")
	if err != nil {
		log.Fatal(err)
	}

	typeToRange := make(shared.Multimap[string, MapEntry])
	for _, entry := range data.Entries {
		typeToRange.Put(entry.StartName, entry)
	}

	minLocation := math.MaxInt
	for _, seed := range data.Seeds {
		location := getLocation(seed, typeToRange)
		minLocation = min(minLocation, location)
	}

	fmt.Println("Finished calculation part 1")
	fmt.Println(minLocation)

	for i := 0; i < len(data.Seeds); i += 2 {
		seed := data.Seeds[i]
		seedRange := data.Seeds[i+1]
		fmt.Printf("Starting seed range, starting seed: %d, range: %d\n", seed, seedRange)
		getMinLocationFromRange(seed, seedRange, typeToRange)
	}
	// locationCache := make(map[int]int)
	/*
		for i := 0; i < len(data.Seeds); i += 2 {
			seed := data.Seeds[i]
			seedRange := data.Seeds[i+1]
			fmt.Printf("Starting seed range, starting seed: %d, range: %d\n", seed, seedRange)
				for s := seed; s < seed+seedRange; s++ {
					location, ok := locationCache[s]
					if !ok {
						location = getLocation(s, typeToRange)
						locationCache[s] = location
					}
					minLocation = min(minLocation, location)
				}
			lowerLocation := getLocation(seed, typeToRange)
			upperLocation := getLocation(seed+seedRange-1, typeToRange)
			fmt.Printf("Seed range, seed: %d, location: %d\n", seed, getLocation(seed, typeToRange))
			fmt.Printf("Seed range, seed + seedRange - 1: %d, location: %d\n", seed+seedRange-1, getLocation(seed+seedRange-1, typeToRange))
			minLocation = min(minLocation, lowerLocation)
			minLocation = min(minLocation, upperLocation)
		}
	*/

	fmt.Println("Finished calculation part 2")
	fmt.Println(minLocation)
}

func getLocation(seed int, typeToRange shared.Multimap[string, MapEntry]) int {
	var soil int
	var fertilizer int
	var water int
	var light int
	var temperature int
	var humidity int
	var location int
	// seed to soil
	seedRanges := typeToRange.Get("seed")
	soil = seed
	for _, seedRange := range seedRanges {
		if seed >= seedRange.Source && seed < seedRange.Source+seedRange.Range {
			soil = seed - seedRange.Source + seedRange.Destination
			break
		}
	}
	// soil to fertilizer
	soilRanges := typeToRange.Get("soil")
	fertilizer = soil
	for _, soilRange := range soilRanges {
		if soil >= soilRange.Source && soil < soilRange.Source+soilRange.Range {
			fertilizer = soil - soilRange.Source + soilRange.Destination
			break
		}
	}

	// fertilizer to water
	fertilizerRanges := typeToRange.Get("fertilizer")
	water = fertilizer
	for _, fertilizerRange := range fertilizerRanges {
		if fertilizer >= fertilizerRange.Source && fertilizer < fertilizerRange.Source+fertilizerRange.Range {
			water = fertilizer - fertilizerRange.Source + fertilizerRange.Destination
			break
		}
	}

	// water to light
	waterRanges := typeToRange.Get("water")
	light = water
	for _, waterRange := range waterRanges {
		if water >= waterRange.Source && water < waterRange.Source+waterRange.Range {
			light = water - waterRange.Source + waterRange.Destination
			break
		}
	}

	// light to temperature
	lightRanges := typeToRange.Get("light")
	temperature = light
	for _, lightRange := range lightRanges {
		if light >= lightRange.Source && light < lightRange.Source+lightRange.Range {
			temperature = light - lightRange.Source + lightRange.Destination
			break
		}
	}

	// temperature to humidity
	temperatureRanges := typeToRange.Get("temperature")
	humidity = temperature
	for _, temperatureRange := range temperatureRanges {
		if temperature >= temperatureRange.Source && temperature < temperatureRange.Source+temperatureRange.Range {
			humidity = temperature - temperatureRange.Source + temperatureRange.Destination
			break
		}
	}

	// humidity to location
	humidityRanges := typeToRange.Get("humidity")
	location = humidity
	for _, humidityRange := range humidityRanges {
		if humidity >= humidityRange.Source && humidity < humidityRange.Source+humidityRange.Range {
			location = humidity - humidityRange.Source + humidityRange.Destination
			break
		}
	}

	return location
}

type Range struct {
	start int
	end   int
}

func getMinLocationFromRange(seed int, range_ int, typeToRange shared.Multimap[string, MapEntry]) int {
	endSeed := seed + range_ - 1
	ranges := make([]Range, 0)
	seedRanges := typeToRange.Get("seed")
	for _, seedRange := range seedRanges {
		fmt.Println("seedRange")
		fmt.Println(seedRange)
		if endSeed < seedRange.Source || seedRange.Source+seedRange.Range-1 < seed {
			fmt.Println("hit no connecting case")
			continue
		} else if seed <= seedRange.Source {
			fmt.Println("hit seed range to the right case")
			end := min(endSeed, seedRange.Source+seedRange.Range-1)
			ranges = append(ranges, Range{seedRange.Destination, end - seed + seedRange.Destination})
		} else {
			fmt.Println("hit seed to the left case")
			start := max(seed, seedRange.Source)
			end := min(seed, seedRange.Source+seedRange.Range-1)
			ranges = append(ranges, Range{convert(start, seedRange), convert(end, seedRange)})
		}
	}

	fmt.Println("ranges")
	fmt.Println(ranges)
	return 0
}

func convert(sourceValue int, range_ MapEntry) int {
	if range_.Source < range_.Destination {
		return sourceValue + range_.Destination - range_.Source
	} else {
		return sourceValue - range_.Destination - range_.Source
	}
}

// parseFile parses the input file and returns a ParsedData struct
func parseFile(filename string) (ParsedData, error) {
	var data ParsedData

	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse the first line containing the seeds
	if scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "seeds:") {
			seedsStr := strings.TrimPrefix(line, "seeds:")
			seeds := strings.Fields(seedsStr)
			for _, seed := range seeds {
				seedInt, err := strconv.Atoi(seed)
				if err != nil {
					return data, fmt.Errorf("invalid seed number: %v", err)
				}
				data.Seeds = append(data.Seeds, seedInt)
			}
		} else {
			return data, fmt.Errorf("expected seeds line but got: %s", line)
		}
	}

	var currentStartName, currentEndName string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}

		if strings.HasSuffix(line, "map:") {
			// This is an X-to-Y name
			namePart := strings.TrimSuffix(line, " map:")
			names := strings.Split(namePart, "-to-")
			if len(names) != 2 {
				return data, fmt.Errorf("invalid format for name: %s", namePart)
			}
			currentStartName = names[0]
			currentEndName = names[1]
		} else {
			// This is a line of numbers, parse them
			numbers := strings.Fields(line)
			if len(numbers) != 3 {
				return data, fmt.Errorf("invalid number of entries in line: %s", line)
			}

			dest, err := strconv.Atoi(numbers[0])
			if err != nil {
				return data, fmt.Errorf("invalid destination number: %v", err)
			}
			src, err := strconv.Atoi(numbers[1])
			if err != nil {
				return data, fmt.Errorf("invalid source number: %v", err)
			}
			rng, err := strconv.Atoi(numbers[2])
			if err != nil {
				return data, fmt.Errorf("invalid range number: %v", err)
			}

			entry := MapEntry{
				StartName:   currentStartName,
				EndName:     currentEndName,
				Destination: dest,
				Source:      src,
				Range:       rng,
			}
			data.Entries = append(data.Entries, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return data, err
	}

	return data, nil
}
