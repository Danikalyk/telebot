package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	// Создаем файл для логов
	logFile, err := os.OpenFile("bot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Устанавливаем вывод логов в файл
	log.SetOutput(logFile)

	// Получаем токен бота из переменной окружения
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("Необходимо установить переменную окружения TELEGRAM_BOT_TOKEN")
	}

	// Настраиваем бота
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// Обработчик команды /start
	b.Handle("/start", func(m *tb.Message) { handleStart(b, m) })

	// Обработчик нажатия кнопки "Меню"
	b.Handle(&tb.ReplyButton{Text: "Меню"}, func(m *tb.Message) { handleMenu(b, m) })

	// Обработчик команды /help
	b.Handle("/help", func(m *tb.Message) { handleHelp(b, m) })

	// Обработчик команды /reverse <PAYLOAD>
	b.Handle("/reverse", func(m *tb.Message) { handleReverse(b, m) })

	// Обработчик текстовых сообщений
	// b.Handle(tb.OnText, handleText)

	// Запускаем бота
	fmt.Println("Бот запущен!")
	b.Start()
}
