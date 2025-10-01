package main

import "fmt"

func main() {
	bookmarks := make(map[string]string)

	for {
		menu := menu()

		switch menu {
		case 1:
			showBookmarksList(bookmarks)
		case 2:
			bookmarks = addBookmark(bookmarks)
		case 3:
			bookmarks = deleteBookmark(bookmarks)
		default:
			return
		}
	}
}

func menu() int {
	var userChoise int
	fmt.Printf(`
1. Посмотреть закладки
2. Добавить закладку
3. Удалить закладку
4. Выйти
`)
	fmt.Scanln(&userChoise)

	return userChoise
}

func showBookmarksList(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("У вас ещё нет закладок")
		return
	}
	for key, value := range bookmarks {
		fmt.Printf(`
Имя: %s
Адрес: %s
		`, key, value)
	}
}

func addBookmark(bookmarks map[string]string) map[string]string {
	var bookmarkName, bookmarkAdress string
	fmt.Print("Введите имя закладки: ")
	fmt.Scan(&bookmarkName)
	fmt.Print("Введите адрес закладки: ")
	fmt.Scan(&bookmarkAdress)
	bookmarks[bookmarkName] = bookmarkAdress
	return bookmarks
}

func deleteBookmark(bookmarks map[string]string) map[string]string {
	var bookmarkName string
	fmt.Println("Введите имя закладки которую хотите удалить")
	fmt.Scanln(&bookmarkName)

	_, exists := bookmarks[bookmarkName]
	if exists {
		delete(bookmarks, bookmarkName)
		fmt.Printf("Закладка %s удалена", bookmarkName)
	} else {
		fmt.Printf("Нет такой закладки: %s", bookmarkName)
	}
	return bookmarks
}
