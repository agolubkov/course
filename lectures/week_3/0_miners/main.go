package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := []string{"rock", "ore", "rock", "ore", "rock", "ore", "rock", "ore", "rock", "ore", "rock", "ore", "rock", "rock", "ore"}
	oreChannel := make(chan string, 2)
	minedOreChan := make(chan string, 2)
	// Разведчик
	go func(mine []string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item // передаем данные в oreChannel
			}
		}
	}(theMine)
	// Добытчик
	go func() {
		for {
			foundOre := <-oreChannel // чтение из канала oreChannel
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" // передаем данные в minedOreChan
		}
	}()
	// Переработчик
	go func() {
		for {
			minedOre := <-minedOreChan // чтение данных из minedOreChan
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}

	}()
	<-time.After(time.Second * 5) // Все еще можете игнорировать
}
