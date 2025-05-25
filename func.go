package main

import (
	"sort"
)

//-------------------------------------------------------------------------------------------------------------------
// 374. Guess Number Higher or Lower

// We are playing the Guess Game. The game is as follows:
// I pick a number from 1 to n. You have to guess which number I picked.
// Every time you guess wrong, I will tell you whether the number I picked is
// higher or lower than your guess.
// You call a pre-defined API int guess(int num), which returns three possible
// results:

// -1: Your guess is higher than the number I picked (i.e. num > pick).
// 1: Your guess is lower than the number I picked (i.e. num < pick).
// 0: your guess is equal to the number I picked (i.e. num == pick).
// Return the number that I picked.

// Example 1:
// 	Input: n = 10, pick = 6
// 	Output: 6

func guess(num int) int {
	pick := 6
	switch {
	case pick > num:
		return 1
	case pick < num:
		return -1
	}
	return 0
}

func guessNumber(n int) int {
	min := 1
	max := n
	for {
		mid := min + (max-min)/2
		res := guess(mid)
		switch res {
		case 1:
			min = mid + 1
		case -1:
			max = mid - 1
		case 0:
			return mid
		}
	}
}

//-------------------------------------------------------------------------------------------------------------------
// 2300. Successful Pairs of Spells and Potions

// You are given two positive integer arrays spells and potions, of length
// n and m respectively, where spells[i] represents the strength of the ith
// spell and potions[j] represents the strength of the jth potion.
// You are also given an integer success. A spell and potion pair is considered
// successful if the product of their strengths is at least success.
// Return an integer array pairs of length n where pairs[i] is the number of
// potions that will form a successful pair with the ith spell.
// Example 1:
// Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7
// Output: [4,0,3]
// Explanation:
// - 0th spell: 5 * [1,2,3,4,5] = [5,10,15,20,25]. 4 pairs are successful.
// - 1st spell: 1 * [1,2,3,4,5] = [1,2,3,4,5]. 0 pairs are successful.
// - 2nd spell: 3 * [1,2,3,4,5] = [3,6,9,12,15]. 3 pairs are successful.
// Thus, [4,0,3] is returned.
// Example 2:
// Input: spells = [3,1,2], potions = [8,5,8], success = 16
// Output: [2,0,2]
// Explanation:
// - 0th spell: 3 * [8,5,8] = [24,15,24]. 2 pairs are successful.
// - 1st spell: 1 * [8,5,8] = [8,5,8]. 0 pairs are successful.
// - 2nd spell: 2 * [8,5,8] = [16,10,16]. 2 pairs are successful.
// Thus, [2,0,2] is returned.

func tf2300(s []int, aim int) int {
	min := 1
	max := len(s) - 1

	if s[0] >= aim {
		return len(s)
	}
	if s[len(s)-1] < aim {
		return 0
	}

	for max >= min {
		mid := min + (max-min)/2
		if s[mid] >= aim && s[mid-1] < aim {
			return len(s) - mid
		}
		if s[mid] < aim {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, len(spells))
	sort.Ints(potions)

	for i, v := range spells {
		threshold := (success + int64(v) - 1) / int64(v)
		tmp := tf2300(potions, int(threshold))
		res[i] = tmp
	}
	return res
}

//-------------------------------------------------------------------------------------------------------------------
