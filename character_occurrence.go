package main

import "fmt"

//Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
//Each letter in magazine can only be used once in ransomNote.

func canConstruct(ransomNote string, magazine string) bool {
	char := make(map[string]int)

	for i := 0; i < len(magazine); i++ {
		char[string(magazine[i])]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if _, ok := char[string(ransomNote[i])]; !ok {
			fmt.Println("no matching characters")
			return false
		}
		char[string(ransomNote[i])]--
		if char[string(ransomNote[i])] < 0 {
			fmt.Println("not enough matching characters")
			return false
		}
	}
	fmt.Println("verification was successful")
	return true
}
