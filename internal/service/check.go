package service

import (
	"errors"

	"github.com/VicSobDev/sigamonitor/bot"
	"go.uber.org/zap"
)

// Check checks for available hours.
func (s *Service) Check() ([]bot.District, error) {
	s.logger.Info("Checking for available hours")

	client, err := bot.NewBot(s.ctx)
	if err != nil {
		return nil, errors.New("Error creating bot client: " + err.Error())
	}

	if err := s.loadCookies(client); err != nil {
		return nil, errors.New("Error loading cookies: " + err.Error())
	}

	districts, err := s.getDistricts(client)
	if err != nil {
		return nil, errors.New("Error getting districts: " + err.Error())
	}

	if err := s.getLocalitiesAndAttendancePlaces(client, districts); err != nil {
		return nil, errors.New("Error getting localities and attendance places: " + err.Error())
	}

	availableDistricts := s.getAvailableDistricts(client, districts)
	if len(availableDistricts) != 0 {
		s.logger.Info("Got available hours", zap.Any("availableHours", availableDistricts))
	}

	return availableDistricts, nil
}

// loadCookies loads cookies for the bot client.
func (s *Service) loadCookies(client *bot.Bot) error {
	s.logger.Info("Loading cookies")
	return client.LoadCookies()
}

// getDistricts retrieves districts.
func (s *Service) getDistricts(client *bot.Bot) ([]bot.District, error) {
	s.logger.Info("Getting districts")
	return client.GetDistricts()
}

// getLocalitiesAndAttendancePlaces retrieves localities and attendance places.
func (s *Service) getLocalitiesAndAttendancePlaces(client *bot.Bot, districts []bot.District) error {
	for i, district := range districts {
		s.logger.Info("Getting localities for: " + district.Name)
		localities, err := client.GetLocalities(district)
		if err != nil {
			s.logger.Error("Error getting localities for: "+district.Name, zap.Error(err))
			continue
		}
		districts[i].Locality = localities

		for j, locality := range localities {
			if locality.Name == "ALL PLACES" {
				continue
			}
			s.logger.Info("Getting attendance places for: " + district.Name + " " + locality.Name)
			attendancePlaces, err := client.GetAttendancePlaces(district, locality)
			if err != nil {
				s.logger.Error("Error getting attendance places for: "+district.Name+" "+locality.Name, zap.Error(err))
				continue
			}
			districts[i].Locality[j].AttendancePlaces = attendancePlaces
		}
	}
	return nil
}

// getAvailableDistricts retrieves available districts.
func (s *Service) getAvailableDistricts(client *bot.Bot, districts []bot.District) []bot.District {
	var result []bot.District
	for _, district := range districts {
		for _, locality := range district.Locality {
			for _, attendancePlace := range locality.AttendancePlaces {

				availableHours, err := client.GetAvailableHours(district, locality, attendancePlace)
				if err != nil {
					s.logger.Error("Error getting available hours for: "+district.Name+" "+locality.Name+" "+attendancePlace.Name, zap.Error(err))
					continue
				}

				attendancePlace.AvailableHours = availableHours

				if len(attendancePlace.AvailableHours) != 0 {
					result = append(result, bot.District{
						Name: district.Name,
						ID:   district.ID,
						Locality: []bot.Locality{
							{
								ID:   locality.ID,
								Name: locality.Name,
								AttendancePlaces: []bot.AttendancePlace{
									{
										ID:             attendancePlace.ID,
										Name:           attendancePlace.Name,
										AvailableHours: attendancePlace.AvailableHours,
									},
								},
							},
						},
					})
				}
			}
		}
	}
	return result
}
