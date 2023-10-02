package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	sess, err := discordgo.New("Bot MTE1ODQ5NDA2MjQxMjgyODc5Mg.GuXNI_.3wbEMXmgeFxLLvoFraBLHVhWQUI22-ZkG7TgXo")
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		// if m.Content == "hellos" {
		// 	s.ChannelMessageSend(m.ChannelID, "world!")
		// }
		if strings.Contains(m.Content, "!hello") {
			s.ChannelMessageSend(m.ChannelID, "world!")
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