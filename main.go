// Statistics of new registered cars (new ones and second handed ones) in Slovenia
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// Output value
type car struct {
	Name       string
	Count      uint64
	NewCount   uint64
	OldCount   uint64
	Percentage uint64
}

// Counters
type counter struct {
	Count      uint64
	NewCount   uint64
	OldCount   uint64
	Percentage uint64
}

// Funs starts here
func main() {
	// Parse flags and args
	percPtr := flag.Bool("percentage", false, "Ordered by percentage of new car registrations (default: by new car registrations)")
	allPtr := flag.Bool("all", false, "Ordered by all car registrations (default: by new car registrations)")
	brandPtr := flag.Bool("brand", false, "Grouped by brand")
	filterPtr := flag.String("period", "", "Filtered by year (e.g., 2023) or month + year (e.g., 082023 for August 2023)")
	petrolPtr := flag.Bool("petrol", false, "Filtered by petrol engine")
	dieselPtr := flag.Bool("diesel", false, "Filtered by diesel engine")
	electricPtr := flag.Bool("electric", false, "Filtered by non-fuel engine")
	personalPtr := flag.Bool("personal", false, "Filtered by personal owners")
	businessPtr := flag.Bool("business", false, "Filtered by business owners (personal flag disables that flag)")
	countPtr := flag.Int64("top", 99999, "Show first N values")

	flag.Parse()
	args := flag.Args()

	// Mode and search prefix word
	all := true
	searchFor := ""

	// Sum counters
	sum := 0
	newSum := 0
	oldSum := 0

	if len(args) == 1 {
		all = false
		searchFor = args[0]
	}

	searchFor = strings.ToUpper(searchFor)

	modelMap := map[string]*counter{}
	modelList := []car{}

	files, err := ioutil.ReadDir("stats/")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read each CSV file in stats subdirectory
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.Contains(file.Name(), *filterPtr) {
			continue
		}

		f, err := os.Open("stats/" + file.Name())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		csvReader := csv.NewReader(f)
		csvReader.Comma = ';'

		line := 0
		firstLine := true

		for {
			// Read and parse line
			record, err := csvReader.Read()

			if err != nil {
				break
			}

			if firstLine {
				firstLine = false
				continue
			}

			dateFirstReg := record[0]
			dateFirstRegInSlo := record[1]
			status := strings.Trim(record[4], " ")
			brand := strings.Trim(strings.ToUpper(record[20]), " ")
			category := strings.Trim(strings.ToUpper(record[33]), " ")
			model := strings.Trim(strings.ToUpper(record[100]), " ")
			engine := strings.Trim(strings.ToUpper(record[48]), " ")
			ownership := strings.Trim(strings.ToUpper(record[10]), " ")

			if strings.Contains(brand, "KODA") {
				brand = "SKODA"
			}

			name := brand

			if !*brandPtr {
				name += " " + model
			}

			newCar := dateFirstReg == dateFirstRegInSlo

			// Accept registrations only (excluding unregistrations)
			if status != "1" {
				continue
			}
			
			// Accept cars only
			if category != "OSEBNI AVTOMOBIL" {
				continue
			}

			if *petrolPtr && engine != "BENCIN" {
				continue
			}

			if *dieselPtr && engine != "DIZEL" {
				continue
			}

			if *electricPtr && engine != "NI GORIVA" {
				continue
			}

			if *personalPtr && ownership != "F" {
				continue
			}

			if !*personalPtr && *businessPtr && ownership != "P" {
				continue
			}

			if !all {
				if !strings.HasPrefix(name, searchFor) {
					continue
				}
			}

			line++

			value, ok := modelMap[name]

			if !ok {
				value = &counter{0, 0, 0, 0}
				modelMap[name] = value
			}

			value.Count = value.Count + 1
			sum += 1

			if newCar {
				value.NewCount = value.NewCount + 1
				newSum += 1
			} else {
				value.OldCount = value.OldCount + 1
				oldSum += 1
			}

			value.Percentage = 100 * value.NewCount / value.Count
		}
	}

	// Make output list from map
	for key, value := range modelMap {
		el := car{key, value.Count, value.NewCount, value.OldCount, value.Percentage}
		modelList = append(modelList, el)
	}

	// Sort
	if *percPtr {
		sort.Slice(modelList, func(i, j int) bool {
			return modelList[i].Percentage > modelList[j].Percentage
		})
	} else if *allPtr {
		sort.Slice(modelList, func(i, j int) bool {
			return modelList[i].Count > modelList[j].Count
		})
	} else {
		sort.Slice(modelList, func(i, j int) bool {
			return modelList[i].NewCount > modelList[j].NewCount
		})
	}

	// Print out
	fmt.Println("+------+----------------------------------------------------+--------+--------+--------+------+")
	fmt.Println("| #    | BRAND AND MODEL                                    | NEW    | OLD    | SUM    | PERC |")
	fmt.Println("+------+----------------------------------------------------+--------+--------+--------+------+")

	hasMore := false

	for i, value := range modelList {
		if i >= int(*countPtr) {
			hasMore = true
			break
		}

		if value.Count > 0 {
			fmt.Printf("| %4d | %-50s | %6d | %6d | %6d | %3d%% |\n", i+1, value.Name, value.NewCount, value.OldCount, value.Count, value.Percentage)
		}
	}

	if hasMore {
		fmt.Println("|  ... | ...                                                |    ... |    ... |    ... |  ... |")
	}

	sumPerc := 0

	if sum != 0 {
		sumPerc = 100 * newSum / sum
	}

	fmt.Println("+------+----------------------------------------------------+--------+--------+--------+------+")
	fmt.Printf("|      | SUM                                                | %6d | %6d | %6d | %3d%% |\n", newSum, oldSum, sum, sumPerc)
	fmt.Println("+------+----------------------------------------------------+--------+--------+--------+------+")
}
