package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/framadhita4/quranlet-api/app/models"
	"github.com/framadhita4/quranlet-api/platform/database"
)

func readSurah(surahNumber int, wg *sync.WaitGroup, surahChan chan models.Surah) {
	defer wg.Done()

	file, err := os.Open(fmt.Sprintf("./pkg/reader/quran/surah/surah_%d/surah_info.json", surahNumber))
	if err != nil {
		fmt.Println("Error reading JSON:", err)
		return
	}
	defer file.Close()

	var surah models.Surah
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&surah)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	surahChan <- surah
}

func main() {
	var wg sync.WaitGroup
	var surahs []models.Surah
	surahChan := make(chan models.Surah, 114)

	for i := 1; i <= 114; i++ {
		wg.Add(1)
		readSurah(i, &wg, surahChan)
	}

	go func() {
		wg.Wait()
		close(surahChan)
	}()

	for surah := range surahChan {
		surahs = append(surahs, surah)
	}

	database.ConnectDB()

	database.DB.Create(&surahs)
}
