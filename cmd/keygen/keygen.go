package main

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func main() {

	fmt.Print("Write password for hashing:")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nОшибка чтения пароля:", err)
		return
	}
	fmt.Println()

	password := string(bytePassword)

	// Генерируем хэш. bcrypt.DefaultCost (значение 10) — оптимальный баланс скорости и защиты
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error: Ошибка при генерации хэша: %v", err)
	}

	// Выводим результат в консоль
	fmt.Printf("Пароль: %s\n", password)
	fmt.Printf("Хэш: %s\n", string(hash))
}
