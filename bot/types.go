package bot

import (
	"context"
	"log"
	"net/http"
	"net/http/cookiejar"

	"go.uber.org/zap"
)

type Bot struct {
	httpClient *http.Client
	Districts  []District
	logger     *zap.Logger
	ctx        context.Context
	Headers    http.Header
}

type District struct {
	Name     string
	ID       int
	Locality []Locality
}

type Locality struct {
	ID               int
	Name             string
	AttendancePlaces []AttendancePlace
}

type AttendancePlace struct {
	ID             int
	Name           string
	AvailableHours []string
}

func NewBot(ctx context.Context) (*Bot, error) {

	/*proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://localhost:8888")
	}*/

	headers := http.Header{
		"User-Agent":                {`Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:122.0) Gecko/20100101 Firefox/122.0`},
		"Accept":                    {`text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8`},
		"Accept-Language":           {`en-US,en;q=0.5`},
		"DNT":                       {`1`},
		"Sec-GPC":                   {`1`},
		"Connection":                {`keep-alive`},
		"Upgrade-Insecure-Requests": {`1`},
		"Sec-Fetch-Dest":            {`document`},
		"Sec-Fetch-Mode":            {`navigate`},
		"Sec-Fetch-Site":            {`none`},
		"Sec-Fetch-User":            {`?1`},
		"Origin":                    {`https://siga.marcacaodeatendimento.pt`},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Printf("Error creating cookie jar: %v\n", err)
		return nil, err
	}
	httpClient := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
			//Proxy:             proxy,
			//TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		},
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Printf("Error creating logger: %v\n", err)
		return nil, err
	}

	client := &Bot{
		httpClient: httpClient,
		logger:     logger,
		ctx:        ctx,
		Headers:    headers,
	}

	return client, nil
}
