package main

import (
	"fmt"
	"time"
)

type Item interface {
	Borrow(user *User)
	Return(user *User)
	Describe() string
	IsAvailable() bool
}

type BasicInfo struct {
	Title     string
	Author    string
	Available bool
}

type Book struct {
	BasicInfo
	ISBN string
}

func (u *User) BorrowItem(item Item) {
	if item.IsAvailable() {
		item.Borrow(u)
		u.Borrowed = append(u.Borrowed, item)
	}
}

func (u *User) ReturnItem(item Item) {
	item.Return(u)
	for i, b := range u.Borrowed {
		if b == item {
			u.Borrowed = append(u.Borrowed[:i], u.Borrowed[i+1:]...)
			break
		}
	}
}

func (b *Book) Borrow(user *User) {
	if b.Available {
		b.Available = false
		fmt.Printf("%s взял книгу '%s'.\n", user.Name, b.Title)
	} else {
		fmt.Printf("Книга '%s' недоступна.\n", b.Title)
	}
}

func (b *Book) Return(user *User) {
	b.Available = true
	fmt.Printf("%s вернул книгу '%s'.\n", user.Name, b.Title)
}

func (b *Book) Describe() string {
	return fmt.Sprintf("Книга '%s' автора '%s' с ISBN '%s'", b.Title, b.Author, b.ISBN)
}

func (b *Book) IsAvailable() bool {
	return b.Available
}

type User struct {
	Name     string
	Borrowed []Item
}

type Library struct {
	Items        []Item
	Transactions []*Transaction
	Registered   []*User
}

func (l *Library) AddItem(item Item) {
	l.Items = append(l.Items, item)
}

type Transaction struct {
	User     *User
	Item     Item
	Type     string
	DateTime time.Time
}

func main() {
	library := &Library{}

	// Добавление книг в библиотеку
	library.AddItem(&Book{BasicInfo: BasicInfo{Title: "Гарри Поттер", Author: "Дж.К. Роулинг", Available: true}, ISBN: "1234567890"})

	fmt.Println("Добро пожаловать в библиотеку!")
	for {
		fmt.Println("Введите ваше имя:")
		var name string = "Nurzhan"
		// fmt.Scanln(&name)

		var currentUser *User
		for _, u := range library.Registered {
			if u.Name == name {
				currentUser = u
			}
		}

		if currentUser == nil {
			currentUser = &User{Name: name}
			library.Registered = append(library.Registered, currentUser)
		}

		fmt.Println("Выберите действие:")
		fmt.Println("1. Просмотреть список книг")
		fmt.Println("2. Взять книгу")
		fmt.Println("3. Вернуть книгу")
		fmt.Println("4. Выход")

		var choice int = 4

		switch choice {
		case 1:
			for _, item := range library.Items {
				fmt.Println(item.Describe())
			}
		case 2:
			fmt.Println("Введите название книги, которую хотите взять:")
			var bookTitle string
			fmt.Scanln(&bookTitle)

			var bookToBorrow Item
			for _, item := range library.Items {
				if b, ok := item.(*Book); ok && b.Title == bookTitle {
					bookToBorrow = item
					break
				}
			}

			if bookToBorrow != nil {
				currentUser.BorrowItem(bookToBorrow)
			} else {
				fmt.Println("Книга не найдена.")
			}
		case 3:
			fmt.Println("Введите название книги, которую хотите вернуть:")
			var bookTitle string
			fmt.Scanln(&bookTitle)

			var bookToReturn Item
			for _, item := range currentUser.Borrowed {
				if b, ok := item.(*Book); ok && b.Title == bookTitle {
					bookToReturn = item
					break
				}
			}

			if bookToReturn != nil {
				currentUser.ReturnItem(bookToReturn)
			} else {
				fmt.Println("Книга не найдена.")
			}
		case 4:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный выбор.")
		}
	}
}
