package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync/atomic"
	"time"
	"unicode"
)

const (
	maxPlayers       = 13
	checkLetterLimit = 5
	csvFile          = "cities15000.csv"
	csvColumn        = 1 // [utf8]name=0, asciiname=1
)

var (
	letters  []rune
	byletter = make(map[rune][]string, 26)

	canBegin  = make(chan bool)
	endOfGame = make(chan bool)
	turn      uint64

	allPlayers = [maxPlayers]string{
		"James", "Mary", "Rob", "Patti", "John",
		"Linda", "Mike", "Jenny", "David", "Beth",
		"Bill", "Barbara", "Rich",
	}
)

func cleanName(name string) string {
	name, _, _ = strings.Cut(name, " (")
	name, _, _ = strings.Cut(name, ",")
	return strings.Trim(name, "'`‘’0123456789 ")
}

func firstLetter(name string) rune {
	return unicode.ToLower([]rune(name)[0])
}

func findCity(letter rune) string {
	cities := byletter[letter]
	if len(cities) == 0 {
		return ""
	}
	choice := rand.Intn(len(cities))
	city := cities[choice]

	// "Hide" already played city
	cities[choice], cities[len(cities)-1] = cities[len(cities)-1], cities[choice]
	byletter[letter] = cities[:len(cities)-1]

	return city
}

func askNumPlayers() (numPlayers int) {
	var input string
	fmt.Print("Enter number of players (from 2 to 13): ")
	fmt.Scanln(&input)
	if len(input) < 1 {
		log.Fatal("Exiting after empty input")
	}
	_, err := fmt.Sscanf(input, "%d", &numPlayers)
	if err != nil {
		log.Fatal(err)
	}
	if numPlayers < 2 || numPlayers > maxPlayers {
		log.Fatal("The number of players should be from 2 to 13")
	}

	return numPlayers
}

func readCities() {
	var unique = make(map[string]struct{}, 24737)

	f, err := os.Open(csvFile)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer f.Close()

	csvData := csv.NewReader(f)

	// skip the header
	if _, err = csvData.Read(); err != nil {
		log.Fatal("Error reading the CSV file header:", err)
	}

	for {
		// CSV fields: name,asciiname,country,geonameid
		rec, err := csvData.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading the CSV record:", err)
		}

		// Get city's name (asciiname)
		cityname := cleanName(rec[csvColumn])

		if _, ok := unique[cityname]; !ok {
			firstLetter := firstLetter(cityname)
			byletter[firstLetter] = append(byletter[firstLetter], cityname)
			unique[cityname] = struct{}{}
			letters = append(letters, firstLetter)
		}

	}
	canBegin <- true
}

func main() {
	go readCities()
	numPlayers := askNumPlayers()

	<-canBegin

	rand.Seed(time.Now().UTC().UnixNano())
	randomLetter := letters[rand.Intn(len(letters))]
	cities := byletter[randomLetter]
	startingCity := cities[rand.Intn(len(cities))]

	fmt.Println("Start of game =", startingCity)

	rand.Shuffle(len(allPlayers), func(i, j int) {
		allPlayers[i], allPlayers[j] = allPlayers[j], allPlayers[i]
	})

	ch1 := make(chan string)
	ch := ch1
	var ch2 chan string
	for id := 1; id <= numPlayers; id++ {
		if id == numPlayers {
			ch2 = ch1
		} else {
			ch2 = make(chan string)
		}
		go play(id, allPlayers[id-1], ch, ch2)
		ch = ch2
	}

	ch1 <- startingCity

	go func() {
		fmt.Scanln()
		endOfGame <- true
	}()

	<-endOfGame

	fmt.Println("End of game on", turn, "turn")
}

func play(id int, name string, in <-chan string, out chan<- string) {
	for {
		inCity := <-in

		filteredCity := []rune(strings.Map(func(r rune) rune {
			if !unicode.IsLetter(r) {
				return -1
			}
			return r
		}, inCity))

		// Find outCity by trying last letter, then previous, up to len-5 (or filteredCity is out of letters)
		outCity := ""
		for try := 0; try < checkLetterLimit && try < len(filteredCity); try++ {
			letter := filteredCity[len(filteredCity)-try-1]
			outCity = findCity(unicode.ToLower(letter))
			if outCity != "" {
				break
			}
		}
		if outCity == "" {
			endOfGame <- true
			return
		}

		log.Printf("[%v] %v says\t%v\n", id, name, outCity)
		atomic.AddUint64(&turn, 1)
		// logrus.WithFields(logrus.Fields{
		// 	"id":       id,
		// 	"name":     name,
		// 	"quest":    inCity,
		// 	"response": outCity,
		// }).Info("New Round")
		out <- outCity
	}
}
