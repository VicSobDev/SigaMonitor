package bot

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"
)

const baseURL = "https://siga.marcacaodeatendimento.pt/Marcacao"

func TestRequests(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bot, err := NewBot(ctx)
	if err != nil {
		t.Fatalf("Failed to create bot: %v", err)
	}

	if err := bot.LoadCookies(); err != nil {
		t.Fatalf("Failed to load cookies: %v", err)
	}

	jar := bot.httpClient.Jar
	if jar == nil {
		t.Fatal("Expected http.Client.Jar to not be nil")
	}

	uri, err := url.Parse(baseURL)
	if err != nil {
		t.Fatalf("Failed to parse URL '%s': %v", baseURL, err)
	}

	cookies := jar.Cookies(uri)
	if len(cookies) == 0 {
		t.Fatal("Expected to find cookies, but found none")
	}

	if !cookieExists(cookies, "SIGA") {
		t.Fatal("Expected to find 'SIGA' cookie, but it was not found")
	}

	if !cookieExists(cookies, "SIGA_ME") {
		t.Fatal("Expected to find 'SIGA_ME' cookie, but it was not found")
	}

	districts, err := bot.GetDistricts()
	if err != nil {
		t.Fatalf("Failed to get districts: %v", err)
	}

	if len(districts) == 0 {
		t.Fatal("Expected to find districts, but found none")
	}

	locality, err := bot.GetLocalities(districts[0])
	if err != nil {
		t.Fatalf("Failed to get localities: %v", err)
	}

	if len(locality) == 0 {
		t.Fatal("Expected to find localities, but found none")
	}

	attendancePlaces, err := bot.GetAttendancePlaces(districts[0], locality[0])
	if err != nil {
		t.Fatalf("Failed to get attendance places: %v", err)
	}

	if len(attendancePlaces) == 0 {
		t.Fatal("Expected to find attendance places, but found none")
	}

	t.Logf("Found attendance places: %v", attendancePlaces)

	availableHours, err := bot.GetAvailableHours(districts[0], locality[0], attendancePlaces[0])
	if err != nil {
		t.Fatalf("Failed to get available hours: %v", err)
	}

	if len(availableHours) == 0 {
		t.Log("No available hours found")
		return
	}

	t.Logf("Found available hours: %v", availableHours)
}

func cookieExists(cookies []*http.Cookie, name string) bool {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return true
		}
	}
	return false
}
