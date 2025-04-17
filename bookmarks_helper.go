package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type bookmarkMap = map[string]string

func addBookmark(bookmarks bookmarkMap) {
	var bookmarkName string
	var bookmarkUrl string
	fmt.Print("Enter bookmark name: ")
	fmt.Scanln(&bookmarkName)
	fmt.Print("Enter bookmark url: ")
	fmt.Scanln(&bookmarkUrl)
	fmt.Println("Adding bookmark...")
	bookmarks[bookmarkName] = bookmarkUrl
	fmt.Println("Bookmark added")
}

func deleteBookmark(bookmarks bookmarkMap) {
	var bookmarkName string
	fmt.Print("Enter bookmark name: ")
	fmt.Scanln(&bookmarkName)
	fmt.Println("Deleting bookmark...")
	delete(bookmarks, bookmarkName)
	fmt.Println("Bookmark deleted")
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("No bookmarks found")
		return
	}
	for key, value := range bookmarks {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func getUserInput() int {
	var userChoice int

	fmt.Println("1. List bookmarks")
	fmt.Println("2. Add bookmark")
	fmt.Println("3. Delete bookmark")
	fmt.Println("4. Exit")

	fmt.Print("Enter your choice: ")

	fmt.Scanln(&userChoice)
	return userChoice
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
