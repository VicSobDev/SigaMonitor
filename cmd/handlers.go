package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/VicSobDev/sigamonitor/bot" // Import external packages
	"github.com/VicSobDev/sigamonitor/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

// struct named Message to hold message data.
type Message struct {
	id       int64
	district bot.District
}

// sendBotStartMessage sends a start message to the Telegram channel.
func (b *BotApp) sendBotStartMessage() {
	msg := tgbotapi.NewMessageToChannel(b.config.ChannelID, "Bot Started Successfully")
	if _, err := b.tgBot.Send(msg); err != nil {
		b.logger.Info("Error sending start message:", zap.Error(err))
	}
}

// processCheck is the main function that performs the checking process.
func (b *BotApp) processCheck() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	checkService, err := service.NewService(ctx)
	if err != nil {
		b.logger.Error("Error creating check service", zap.Error(err))
		return
	}

	districts, err := checkService.Check()
	if err != nil {
		b.logger.Error("Error checking for available hours", zap.Error(err))
		return
	}

	b.updateDistrictAvailability(districts)
}

// updateDistrictAvailability updates the availability of districts.
func (b *BotApp) updateDistrictAvailability(districts []bot.District) {
	var notAvailable = []Message{}
	var active = []bot.District{}

	for key, district := range b.messageMap {
		if found := findDistrict(districts, district.Name); !found {
			notAvailable = append(notAvailable, Message{district: district, id: key})
			b.logger.Info(district.Name + " is not available anymore")
		} else {
			active = append(active, district)
			b.logger.Info(district.Name + " is still available")
		}
	}

	b.handleNotAvailableDistricts(notAvailable)
	b.handleNewAvailableDistricts(districts, active)
}

// findDistrict checks if a district exists in the provided list.
func findDistrict(districts []bot.District, name string) bool {
	for _, d := range districts {
		if d.Name == name {
			return true
		}
	}
	return false
}

// handleNotAvailableDistricts handles districts that are no longer available.
func (b *BotApp) handleNotAvailableDistricts(notAvailable []Message) {
	for _, message := range notAvailable {
		b.logger.Info(fmt.Sprintf("%s is not available anymore", message.district.Name))
		messages := formatDistrictsMessage([]bot.District{message.district}, false)
		for _, msgText := range messages {
			msg := tgbotapi.NewEditMessageText(int64(b.config.ChannelIDAsInt), int(message.id), msgText)
			if _, err := b.tgBot.Send(msg); err != nil {
				b.logger.Error("Error updating district availability", zap.Error(err))
			}
		}
		delete(b.messageMap, message.id)
	}
}

// handleNewAvailableDistricts handles newly available districts.
func (b *BotApp) handleNewAvailableDistricts(districts, active []bot.District) {
	for _, district := range districts {
		if !findDistrict(active, district.Name) {
			b.logger.Info(fmt.Sprintf("%s is available", district.Name))
			b.sendDistrictsInfo([]bot.District{district})
		}
	}
}

// formatDistrictsMessage formats messages for districts.
func formatDistrictsMessage(districts []bot.District, available bool) []string {
	var messages []string
	status := "Indisponível ❌"
	if available {
		status = "Disponível ✅"
	}

	for _, district := range districts {
		for _, locality := range district.Locality {
			for _, attendancePlace := range locality.AttendancePlaces {
				message := fmt.Sprintf(
					"Novo Horário Disponível! 🎉🎉🎉\n\n----------------------------------------\nDistrito: %s\nLocalidade: %s\nLocal: %s\nHorários: %v\nStatus: %s\n----------------------------------------\nPara agendar acesse: https://siga.marcacaodeatendimento.pt/Marcacao/MarcacaoInicio\n",
					district.Name, locality.Name, attendancePlace.Name, attendancePlace.AvailableHours, status)
				messages = append(messages, message)
			}
		}
	}

	return messages
}

// sendDistrictsInfo sends information about available districts.
func (b *BotApp) sendDistrictsInfo(districts []bot.District) {
	messages := formatDistrictsMessage(districts, true)
	for i, messageText := range messages {
		log.Println(messageText)
		msg := tgbotapi.NewMessageToChannel(b.config.ChannelID, messageText)
		message, err := b.tgBot.Send(msg)
		if err != nil {
			log.Printf("Error sending districts info: %v\n", err)
		}
		b.messageMap[int64(message.MessageID)] = districts[i]
	}
}
