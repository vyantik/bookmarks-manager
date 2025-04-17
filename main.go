package main

import "fmt"

func main() {
	bookmarks := make(bookmarkMap)

Menu:
	for {
		userChoice := getUserInput()
		clearScreen()
		switch userChoice {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			fmt.Println("Exiting...")
			break Menu
		default:
			fmt.Println("Invalid choice")
		}
	}

}
