package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	var count int

	// Obter o valor da variável de ambiente TOKEN
	godotenv.Load()
	token := os.Getenv("TOKEN")
	if token == "" {
		fmt.Println("A variável de ambiente TOKEN não foi configurada.")
		return
	}
	fmt.Println("Valor da variável de ambiente TOKEN:", token)

	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if strings.Contains(m.Content, "!hellos") {
			s.ChannelMessageSend(m.ChannelID, "world!")
			count = count + 1
		}
		if strings.Contains(m.Content, "!result") {
			message := fmt.Sprintf("valor de count: %d", count)
			count = 0
			s.ChannelMessageSend(m.ChannelID, message)
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)

	}
	defer sess.Close()

	fmt.Println("Bot ta comecando!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
