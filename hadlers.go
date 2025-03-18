package main

import (
	"fmt"
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Функция для переворачивания строки
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Обработчик команды /start
func handleStart(b *tb.Bot, m *tb.Message) {
	log.Println("Обработчик команды /start вызван")

	// Создаем reply-клавиатуру
	menuBtn := tb.ReplyButton{
		Text: "Меню",
	}
	replyKeys := [][]tb.ReplyButton{
		{menuBtn},
	}

	// Отправляем картинку
	photo := &tb.Photo{File: tb.FromDisk("img/tevirp_cat.jpg")}
	b.Send(m.Sender, photo)

	// Отправляем приветственное сообщение с клавиатурой
	b.Send(m.Sender, "Тевирп! Отправь мне любое слово с помощью команды /reverse, и я переверну его. Пример: /reverse Привет", &tb.ReplyMarkup{
		ReplyKeyboard:       replyKeys,
		ResizeReplyKeyboard: true,
	})
}

// Обработчик нажатия кнопки "Меню"
func handleMenu(b *tb.Bot, m *tb.Message) {
	log.Println("Нажата кнопка Меню")
	b.Send(m.Sender, "Доступные команды:\n/start - Начать работу с ботом\n/help - Показать инструкцию\n/reverse <слово> - Перевернуть слово")
}

// Обработчик команды /help
func handleHelp(b *tb.Bot, m *tb.Message) {
	log.Println("Обработчик команды /help вызван")

	// Создаем inline кнопку
	inlineBtn := tb.InlineButton{
		Unique: "help_link",
		Text:   "Инструкция на нашем YouTube канале!",
		URL:    "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	}

	// Создаем клавиатуру с кнопкой
	inlineKeys := [][]tb.InlineButton{
		{inlineBtn},
	}

	// Отправляем сообщение с кнопкой
	b.Send(m.Sender, "Инструкция по использованию бота", &tb.ReplyMarkup{
		InlineKeyboard: inlineKeys,
	})
}

// Обработчик команды /reverse <PAYLOAD>
func handleReverse(b *tb.Bot, m *tb.Message) {
	log.Println("Обработчик команды /reverse вызван")
	if !m.Private() {
		return
	}
	if len(m.Payload) != 0 {
		reversed := reverseString(m.Payload)
		b.Send(m.Sender, fmt.Sprintf("Перевернутое слово: %s", reversed))

		fmt.Println(m.Payload)
	} else {
		b.Send(m.Sender, "Вы не ввели слово. Пример использования: /reverse Привет")
	}
}

// Обработчик текстовых сообщений
// func handleText(m *tb.Message) {
// 	// Переворачиваем текст
// 	reversed := reverseString(m.Text)
// 	b.Send(m.Sender, fmt.Sprintf("Перевернутое слово: %s", reversed))
// }
