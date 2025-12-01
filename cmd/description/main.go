package main

import (
	"log"

	"github.com/Andersson19/aoc-2025/internal/aoc"
)

func main() {
	year, day, cookie, err := aoc.ParseYearDayCookieFlags()
	if err != nil {
		log.Fatal(err)
	}

	err = aoc.DownloadDayDescription(year, day, cookie)
	if err != nil {
		log.Fatalf("Failed to download description for day %d (%d): %v", day, year, err)
	}
}
