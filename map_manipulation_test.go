package main

import (
	"testing"
)

var (
	data = []string{"1", "23", "23", "4", "23", "5", "6", "1", "4", "6", "5", "6", "6", "6", "1", "4", "23", "1", "4", "5", "6"}
)

func BenchmarkSetMapWithCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetMapWithCheck(data)
	}
}

func BenchmarkSetMapWithoutCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetMapWithoutCheck(data)
	}
}

func SetMapWithoutCheck(rankingSlots []string) map[string]float32 {
	var sellerScoreMap = make(map[string]float32, len(rankingSlots))
	for _, slot := range rankingSlots {
		sellerScoreMap[slot] = -1
	}
	return sellerScoreMap
}

func SetMapWithCheck(rankingSlots []string) map[string]float32 {
	var sellerScoreMap = make(map[string]float32, len(rankingSlots))
	for _, slot := range rankingSlots {
		if _, exist := sellerScoreMap[slot]; !exist {
			sellerScoreMap[slot] = -1
		}

	}
	return sellerScoreMap
}
