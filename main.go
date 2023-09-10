// Statistics of new registered cars (new ones and second handed ones) in Slovenia
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"encoding/csv"
	"io/ioutil"
)

// Output value
type car struct {
	Name string
	Count uint64
	NewCount uint64
	OldCount uint64
	Percentage uint64
}

// Counters
type counter struct {
	Count uint64
	NewCount uint64
	OldCount uint64
	Percentage uint64
}

// Funs starts here
func main() {
	// Parse flags and args
	percPtr := flag.Bool("p", false, "Ordered by percentage")
	allPtr := flag.Bool("a", false, "Ordered by all registrations")
	brandPtr := flag.Bool("b", false, "Grouped by brand")

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

		f, err := os.Open("stats/" + file.Name())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	
		csvReader := csv.NewReader(f)
		csvReader.Comma = ';';
	
		line := 0;
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
			brand := strings.Trim(strings.ToUpper(record[20]), " ")
			category := strings.Trim(strings.ToUpper(record[33]), " ")
			model := strings.Trim(strings.ToUpper(record[100]), " ")
	
			if strings.Contains(brand, "KODA") {
				brand = "SKODA"
			}
	
			name := brand

			if !*brandPtr {
				name += " " + model
			}
	
			newCar := dateFirstReg == dateFirstRegInSlo
	
			// Accept cars only
			if category != "OSEBNI AVTOMOBIL" {
				continue
			}

			if !all {
				if !strings.HasPrefix(name, searchFor) {
					continue
				}
			}
	
			line++;
	
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
	fmt.Println("+------+----------------------------------------------------+-------+-------+-------+------+")
	fmt.Println("| #    | BRAND AND MODEL                                    | NEW   | OLD   | SUM   | PERC |")
	fmt.Println("+------+----------------------------------------------------+-------+-------+-------+------+")

	for i, value := range modelList {
		if value.Count > 0 {
			fmt.Printf("| %4d | %-50s | %5d | %5d | %5d | %3d%% |\n", i + 1, value.Name, value.NewCount, value.OldCount, value.Count, value.Percentage)
		}
	}

	sumPerc := 0

	if sum != 0 {
		sumPerc = 100 * newSum / sum		
	}

	fmt.Println("+------+----------------------------------------------------+-------+-------+-------+------+")
	fmt.Printf("|      | SUM                                                | %5d | %5d | %5d | %3d%% |\n", newSum, oldSum, sum, sumPerc)
	fmt.Println("+------+----------------------------------------------------+-------+-------+-------+------+")
}
