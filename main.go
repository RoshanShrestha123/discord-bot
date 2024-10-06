package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TOKEN")

	if token == "" {
		log.Fatalln("No Token provided, please login to the discord app to generate the new token")
	}

	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(messageHandler)

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer discord.Close()

	fmt.Println("The bot is online!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(s.State.User.ID)
	fmt.Println(m.Author.ID)
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Println(m.Content)
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "!pong [Roshan is making his own bot! please ignore this message]")
	}

}
