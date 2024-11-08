package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

func countAlphanum(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			count++
		}
	}
	return count
}

func calculatePoints(receipt Receipt) int {
	points := 0

	// Rule 1.
	points += countAlphanum(receipt.Retailer)

	// Rule 2,3
	total, _ := strconv.ParseFloat(receipt.Total, 64) // can simply compare with last to char == 00 but too hard coded and leaves little room for later changes in req
	if total == float64(int64(total)) {
		points += 50
	}
	//
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4,5
	points += (len(receipt.Items) / 2) * 5
	//
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6,7
	layout := "2006-01-02"
	date, _ := time.Parse(layout, receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}
	//
	layoutTime := "15:04"
	t, _ := time.Parse(layoutTime, receipt.PurchaseTime)
	if t.Hour() >= 14 && t.Hour() < 16 {
		points += 10
	}

	return points
}
