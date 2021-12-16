package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func AreAnagram(str1 string, str2 string) bool {
	// adapted from https://www.geeksforgeeks.org/check-whether-two-strings-are-anagram-of-each-other/
	// Get lengths of both strings
	n1 := len(str1)
	n2 := len(str2)

	// If length of both strings is not same, then
	// they cannot be anagrams
	if n1 != n2 {
		return false
	}

	// don't print match if it's the same exact word
	if str1 == str2 {
		return false
	}
	// Sort both strings
	str1 = SortString(str1)
	str2 = SortString(str2)

	// Compare sorted strings
	for i := 0; i < n1; i++ {
		if str1[i] != str2[i] {
			return false
		}
	}

	return true
}
func FindHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	filename := "word-list/en.txt"

	query := r.URL.Query()
	needles, present := query["word1"]
	if !present || len(needles) == 0 {
		io.WriteString(w, "no word given")
		return
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		io.WriteString(w, "File reading error")
		return
	}
	lines := strings.Split(string(data), "\n")
	anagrams := []string{}
	for i := 0; i < len(lines); i++ {
		if AreAnagram(lines[i], needles[0]) {
			anagrams = append(anagrams, lines[i])
		}
	}
	export := "[\"" + strings.Join(anagrams, "\", \"") + "\"]"
	io.WriteString(w, export)

}

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	message := "false"
	// get strings from query
	query := r.URL.Query()
	word1, present1 := query["word1"]
	word2, present2 := query["word2"]

	if !present1 || !present2 || len(word1[0]) == 0 || len(word2[0]) == 0 {
		io.WriteString(w, "Must give two words like '?word1=<first word>&word2=<second word>'")
		return
	}
	if AreAnagram(word1[0], word2[0]) {
		message = "true"
	}
	io.WriteString(w, "{\"anagrams\" : "+message+"}")

}

func main() {
	http.HandleFunc("/find", FindHandler)
	http.HandleFunc("/compare", CompareHandler)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
