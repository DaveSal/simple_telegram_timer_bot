package main

import (
	"log"
	"os"
	"github.com/yanzay/tbot"
  "strconv"
  "time"
)

func main() {
	bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN_TIMER"))

	if err != nil {
		log.Fatal(err)
	}

  bot.HandleFunc("/timer {seconds}", timerHandler)
	bot.ListenAndServe()
}

func timerHandler(m *tbot.Message) {
  secondsStr := m.Vars["seconds"]
  seconds, err := strconv.Atoi(secondsStr)

  if err != nil {
          m.Reply("Вы ввели некорректные данные.")
          return
  }

  m.Replyf("Таймер запущен на %d секунд.", seconds)
  time.Sleep(time.Duration(seconds) * time.Second)
  m.Reply("Время вышло!")
}

