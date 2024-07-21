package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

const numHashes = 100000000

func main() {
	// Initialize counters for each byte value (0-255)
	byteCount := make([]int, 256)

	// Loop to generate [32]byte from Unix time in nanoseconds and calculate hashes
	for i := 0; i < numHashes; i++ {
		var randomBytes [32]byte
		unixTimeNano0 := time.Now().UnixNano()
		unixTimeNano1 := time.Now().UnixNano()
		unixTimeNano2 := time.Now().UnixNano()
		unixTimeNano3 := time.Now().UnixNano()
		binary.LittleEndian.PutUint64(randomBytes[0:], uint64(unixTimeNano0))
		binary.LittleEndian.PutUint64(randomBytes[8:], uint64(unixTimeNano1))
		binary.LittleEndian.PutUint64(randomBytes[16:], uint64(unixTimeNano2))
		binary.LittleEndian.PutUint64(randomBytes[24:], uint64(unixTimeNano3))
		copy(randomBytes[8:], randomBytes[0:8])
		copy(randomBytes[16:], randomBytes[0:8])
		copy(randomBytes[24:], randomBytes[0:8])

		hash := sha256.Sum256(randomBytes[:])

		// Count occurrences of each byte in the hash
		for _, b := range hash {
			byteCount[b]++
		}
	}

	// Calculate expected count per byte (assuming uniform distribution)
	expectedCount := float64(numHashes) * 32 / 256 // 32 bytes per hash

	minDiff := math.MaxFloat64
	maxDiff := -math.MaxFloat64
	totalAbsoluteDifference := 0.0
	totalCount := 0
	hitCount := 0
	hitIndexes := []int{}

	// Track zero-hits
	zeroHits := 0
	zeroHitIndexes := []int{}

	// Analyze and print byte distribution
	fmt.Println("Byte Distribution Analysis:")
	fmt.Println("Byte | Count  | Expected       | Difference")
	for b, count := range byteCount {
		difference := math.Abs(float64(count) - expectedCount)
		totalCount += count
		AbsoluteDifference := difference
		totalAbsoluteDifference += AbsoluteDifference
		if AbsoluteDifference < minDiff {
			minDiff = difference
		}
		if AbsoluteDifference > maxDiff {
			maxDiff = difference
		}
		if count == 0 {
			zeroHits++
			zeroHitIndexes = append(zeroHitIndexes, b)
		}
		if AbsoluteDifference == 0 {
			hitCount++
			hitIndexes = append(hitIndexes, b)
		}
		fmt.Printf("%02x   | %6d | %15.4f | %15.2f\n", b, count, expectedCount, difference)
	}

	averageAbsoluteDifference := totalAbsoluteDifference / 256
	fmt.Printf("Average Absolute Difference: %.6f\n", averageAbsoluteDifference)
	fmt.Printf("Minimum Difference: %.6f\n", minDiff)
	fmt.Printf("Maximum Difference: %.6f\n", maxDiff)
	fmt.Printf("Total count: %d\n", totalCount)
	fmt.Printf("Zero hits: %d\n", zeroHits)
	fmt.Printf("Zero hit indexes: %v\n", zeroHitIndexes)
	fmt.Printf("     hits: %d\n", hitCount)
	fmt.Printf("     hit indexes: %v\n", hitIndexes)
	fmt.Println("timestamp bytes")
}
