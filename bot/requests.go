package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"go.uber.org/zap"
	"golang.org/x/net/html"
)

func (c *Bot) LoadCookies() error {

	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, pesquisaAssuntoEndpoint, nil)
	if err != nil {
		c.logger.Error("error while creating request for cookies", zap.Error(err))
		return err
	}

	header := c.GetHeaders()
	header.Del("Origin")
	req.Header = header

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error while making request for cookies", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	bodyRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error while reading response body for cookies", zap.Error(err))
		return err
	}

	if resp.StatusCode != http.StatusOK {
		c.logger.Error("error while getting cookies", zap.String("status", resp.Status), zap.String("body", string(bodyRaw)))
		return err
	}

	reqBody := url.Values{}
	reqBody.Add("txtPesquisa", "renewal+")

	req, err = http.NewRequestWithContext(c.ctx, http.MethodPost, pesquisaAssuntosEndpoint, strings.NewReader(reqBody.Encode()))
	if err != nil {
		c.logger.Error("error while creating request for cookies", zap.Error(err))
		return err
	}

	header = c.GetHeaders()
	header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err = c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error while making request for cookies", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	bodyRaw, err = io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error while reading response body for cookies", zap.Error(err))
		return err
	}

	if resp.StatusCode != http.StatusOK {
		c.logger.Error("error while getting cookies", zap.String("status", resp.Status), zap.String("body", string(bodyRaw)))
		return fmt.Errorf("error while getting cookies: %s", resp.Status)
	}

	reqBody = url.Values{}
	reqBody.Add("pagina", "0")
	reqBody.Add("numPaginas", "0")
	reqBody.Add("IdEntidade", EntityID)
	reqBody.Add("DescricaoEntidade", EnitityDescription)
	reqBody.Add("RequerAutenticacao", AuthRequired)
	reqBody.Add("NivelCategorias", "3")
	reqBody.Add("AtendimentoRemoto", "false")
	reqBody.Add("IdCategoria", CategoryID)
	reqBody.Add("IdSubcategoria", SubCategoryID)
	reqBody.Add("IdMotivo", ReasonID)
	reqBody.Add("NumeroEntidades", "0")
	reqBody.Add("TextoPesquisaAnterior", "renewal+")
	reqBody.Add("HtmlAutenticacaoUtilizador", "False")
	reqBody.Add("textoPesquisa", "renewal+")

	req, err = http.NewRequestWithContext(c.ctx, http.MethodPost, pesquisaAssuntoEndpoint, strings.NewReader(reqBody.Encode()))
	if err != nil {
		c.logger.Error("error while creating request for cookies", zap.Error(err))
		return err
	}

	header = c.GetHeaders()
	header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header = header

	resp, err = c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error while making request for cookies", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	bodyRaw, err = io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error while reading response body for cookies", zap.Error(err))
		return err
	}

	if resp.StatusCode != http.StatusFound {
		c.logger.Error("error while getting cookies", zap.String("status", resp.Status), zap.String("body", string(bodyRaw)))
		return fmt.Errorf("error while getting cookies: %s", resp.Status)
	}

	req, err = http.NewRequestWithContext(c.ctx, http.MethodGet, assuntoEndpoint, nil)
	if err != nil {
		c.logger.Error("error while creating request for cookies", zap.Error(err))
		return err
	}

	req.Header = c.GetHeaders()

	resp, err = c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error while making request for cookies", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	bodyRaw, err = io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error while reading response body for cookies", zap.Error(err))
		return err
	}

	if resp.StatusCode != http.StatusOK {
		c.logger.Error("error while getting cookies", zap.String("status", resp.Status), zap.String("body", string(bodyRaw)))
		return fmt.Errorf("error while getting cookies: %s", resp.Status)
	}

	reqBody = url.Values{}
	reqBody.Add("IdEntidade", EntityID)
	reqBody.Add("DescricaoEntidade", EnitityDescription)
	reqBody.Add("RequerAutenticacao", AuthRequired)
	reqBody.Add("AtendimentoRemoto", "False")
	reqBody.Add("TiposAtendimentoList", "")
	reqBody.Add("DescricaoCategoria", CategoryDescription)
	reqBody.Add("ListCategoriasAtivasPorEntidade", "System.Web.Mvc.SelectList")
	reqBody.Add("DescricaoSubcategoria", SubCategoryDescription)
	reqBody.Add("DescricaoMotivo", ReasonDescription)
	reqBody.Add("NivelCategorias", "3")
	reqBody.Add("regressoPagina", "False")
	reqBody.Add("PaginaAnterior", "PesquisaAssunto")
	reqBody.Add("TextoPesquisaAnterior", "renewal+")
	reqBody.Add("HtmlAutenticacaoUtilizador", "False")
	reqBody.Add("IdCategoria", CategoryID)
	reqBody.Add("IdSubcategoria", SubCategoryID)
	reqBody.Add("IdMotivo", ReasonID)
	reqBody.Add("NumCasos", NumCases)
	reqBody.Add("Email", "")
	reqBody.Add("proximoButton", "Próximo")

	req, err = http.NewRequestWithContext(c.ctx, http.MethodPost, assuntoEndpoint, strings.NewReader(reqBody.Encode()))
	if err != nil {
		c.logger.Error("error while creating request for cookies", zap.Error(err))
		return err
	}
	header = c.GetHeaders()
	header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header = header

	resp, err = c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error while making request for cookies", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	bodyRaw, err = io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error while reading response body for cookies", zap.Error(err))
		return err
	}

	if resp.StatusCode != http.StatusFound {
		c.logger.Error("error while getting cookies", zap.String("status", resp.Status), zap.String("body", string(bodyRaw)))
		return fmt.Errorf("error while getting cookies: %s", resp.Status)
	}

	return nil
}

func (c *Bot) GetDistricts() ([]District, error) {

	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, localEndpoint, nil)
	if err != nil {
		c.logger.Error("error while creating request for districts", zap.Error(err))
		return []District{}, err
	}

	req.Header = c.GetHeaders()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error while making request for districts", zap.Error(err))
		return []District{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error while reading response body for districts", zap.Error(err))
		return []District{}, err
	}

	doc, err := html.Parse(strings.NewReader(string(bodyBytes)))
	if err != nil {
		c.logger.Error("error while parsing response body for districts", zap.Error(err))
		return []District{}, err
	}

	var districts []District
	var parseOptions bool
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "select" {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "IdDistrito" {
					parseOptions = true
				}
			}
		}
		if parseOptions && n.Type == html.ElementNode && n.Data == "option" {
			var dist District
			for _, a := range n.Attr {
				if a.Key == "value" {
					if a.Val == "" {
						continue
					}
					fmt.Sscanf(a.Val, "%d", &dist.ID)

				}
			}
			if n.FirstChild != nil {
				if dist.ID != 0 {
					dist.Name = n.FirstChild.Data
				}
			}
			if dist.ID != 0 {
				districts = append(districts, dist)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
		if n.Type == html.ElementNode && n.Data == "select" {
			parseOptions = false
		}
	}

	f(doc)

	return districts, nil

}

type GetLocalityIDResponse []LocalityIdResponse

type LocalityIdResponse struct {
	IDConcelho int    `json:"IdConcelho"`
	Descricao  string `json:"Descricao"`
}

func (c *Bot) GetLocalities(District District) ([]Locality, error) {

	data := url.Values{}

	data.Add("IdEntidade", EntityID)
	data.Add("DescricaoEntidade", EnitityDescription)
	data.Add("RequerAutenticacao", AuthRequired)
	data.Add("IdCategoria", CategoryID)
	data.Add("DescricaoCategoria", CategoryDescription)
	data.Add("IdSubcategoria", SubCategoryID)
	data.Add("DescricaoSubcategoria", SubCategoryDescription)
	data.Add("IdMotivo", ReasonID)
	data.Add("DescricaoMotivo", ReasonDescription)
	data.Add("NumCasos", NumCases)
	data.Add("DescricaoDistrito", "")
	data.Add("DescricaoLocalidade", "")
	data.Add("DescricaoLocalAtendimento", "")
	data.Add("HtmlAutenticacaoUtilizador", HtmlAuthUser)
	data.Add("IdDistrito", fmt.Sprint(District.ID))
	requestBody := data.Encode()

	req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, pesquisaLocalidadeEndpoint, strings.NewReader(requestBody))
	if err != nil {
		c.logger.Error("error while creating request for district %s: %s", zap.String("district", District.Name), zap.Error(err))
		return []Locality{}, err
	}

	header := c.GetHeaders()
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header = header

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error while making request for district %s: %s", zap.String("district", District.Name), zap.Error(err))
		return []Locality{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error while reading response body for district %s: %s", zap.String("district", District.Name), zap.Error(err))
		return []Locality{}, err
	}

	var localityIDResponse GetLocalityIDResponse
	err = json.Unmarshal(bodyBytes, &localityIDResponse)
	if err != nil {
		c.logger.Error("error while unmarshalling response body for district %s: %s\n%s", zap.String("district", District.Name), zap.Error(err), zap.String("body", string(bodyBytes)))
		return []Locality{}, err
	}

	var localities []Locality
	for _, locality := range localityIDResponse {
		if locality.IDConcelho == -1 {
			continue
		}
		localities = append(localities, Locality{
			ID:   locality.IDConcelho,
			Name: locality.Descricao,
		})
	}

	if len(localities) == 0 {
		c.logger.Error("no localities found for district %s", zap.String("district", District.Name))
		return []Locality{}, fmt.Errorf("no localities found for district %s", District.Name)
	}

	return localities, nil
}

type GetAttendancePlaceResponse []AttendancePlaceResponse

type AttendancePlaceResponse struct {
	IDLocalAtendimento int    `json:"IdLocalAtendimento"`
	Descricao          string `json:"Descricao"`
}

func (c *Bot) GetAttendancePlaces(District District, Locality Locality) ([]AttendancePlace, error) {

	data := url.Values{}

	data.Add("IdEntidade", EntityID)
	data.Add("DescricaoEntidade", EnitityDescription)
	data.Add("RequerAutenticacao", AuthRequired)
	data.Add("IdCategoria", CategoryID)
	data.Add("DescricaoCategoria", CategoryDescription)
	data.Add("IdSubcategoria", SubCategoryID)
	data.Add("DescricaoSubcategoria", SubCategoryDescription)
	data.Add("IdMotivo", ReasonID)
	data.Add("DescricaoMotivo", ReasonDescription)
	data.Add("NumCasos", NumCases)
	data.Add("DescricaoDistrito", "")
	data.Add("DescricaoLocalidade", "")
	data.Add("DescricaoLocalAtendimento", "")
	data.Add("HtmlAutenticacaoUtilizador", HtmlAuthUser)
	data.Add("IdDistrito", fmt.Sprint(District.ID))
	data.Add("IdLocalidade", fmt.Sprint(Locality.ID))
	requestBody := data.Encode()

	req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, pesquisaLocalAtendimentoEndpoint, strings.NewReader(requestBody))
	if err != nil {
		c.logger.Error("error while creating request for locality %s: %s", zap.String("locality", Locality.Name), zap.Error(err))
		return []AttendancePlace{}, err
	}

	header := c.GetHeaders()

	header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header = header

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error while making request for locality %s: %s", zap.String("locality", Locality.Name), zap.Error(err))
		return []AttendancePlace{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error while reading response body for locality %s: %s", zap.String("locality", Locality.Name), zap.Error(err))
		return []AttendancePlace{}, err
	}

	var attendancePlaceResponse GetAttendancePlaceResponse

	err = json.Unmarshal(bodyBytes, &attendancePlaceResponse)
	if err != nil {
		c.logger.Error("error while unmarshalling response body for locality %s: %s\n%s", zap.String("locality", Locality.Name), zap.Error(err), zap.String("body", string(bodyBytes)))
		return []AttendancePlace{}, err
	}

	var attendancePlaces []AttendancePlace
	for _, attendancePlace := range attendancePlaceResponse {
		attendancePlaces = append(attendancePlaces, AttendancePlace{
			ID:   attendancePlace.IDLocalAtendimento,
			Name: attendancePlace.Descricao,
		})
	}

	if len(attendancePlaces) == 0 {
		c.logger.Error("no attendance places found for locality %s", zap.String("locality", Locality.Name))
		return []AttendancePlace{}, fmt.Errorf("no attendance places found for locality %s", Locality.Name)
	}

	return attendancePlaces, nil
}

func (c *Bot) GetAvailableHours(district District, locality Locality, attendancePlace AttendancePlace) ([]string, error) {

	var availableHours []string

	data := url.Values{}
	data.Add("IdEntidade", EntityID)
	data.Add("DescricaoEntidade", EnitityDescription)
	data.Add("RequerAutenticacao", AuthRequired)
	data.Add("IdCategoria", CategoryID)
	data.Add("DescricaoCategoria", CategoryDescription)
	data.Add("IdSubcategoria", SubCategoryID)
	data.Add("DescricaoSubcategoria", SubCategoryDescription)
	data.Add("IdMotivo", ReasonID)
	data.Add("DescricaoMotivo", ReasonDescription)
	data.Add("NumCasos", NumCases)
	data.Add("DescricaoDistrito", district.Name)
	data.Add("DescricaoLocalidade", locality.Name)
	data.Add("DescricaoLocalAtendimento", attendancePlace.Name)
	data.Add("HtmlAutenticacaoUtilizador", HtmlAuthUser)
	data.Add("IdDistrito", fmt.Sprint(district.ID))
	data.Add("IdLocalidade", fmt.Sprint(locality.ID))
	data.Add("IdLocalAtendimento", fmt.Sprint(attendancePlace.ID))
	data.Add("proximoButton", "Próximo")
	requestBody := data.Encode()

	req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, localEndpoint, strings.NewReader(requestBody))
	if err != nil {
		c.logger.Error("error while creating request for available hours: %s", zap.Error(err))
		return []string{}, err
	}

	header := c.GetHeaders()
	header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header = header

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error while making request for available hours: %s", zap.Error(err))
		return []string{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error while reading response body for available hours: %s", zap.Error(err))
		return []string{}, err
	}

	if resp.StatusCode != http.StatusFound {
		c.logger.Error("error while getting available hours: %s, %s", zap.String("status", resp.Status), zap.String("body", string(bodyBytes)))
		return []string{}, fmt.Errorf("error while getting available hours: %s", resp.Status)
	}

	req, err = http.NewRequestWithContext(c.ctx, http.MethodGet, horarioLocalEndpoint, nil)
	if err != nil {
		c.logger.Error("error while creating request for available hours: %s", zap.Error(err))
		return []string{}, err
	}

	req.Header = c.GetHeaders()

	resp, err = c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("error while making request for available hours: %s", zap.Error(err))
		return []string{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error while reading response body for available hours: %s", zap.Error(err))
		return []string{}, err
	}

	if resp.StatusCode != http.StatusOK {
		c.logger.Error("error while getting available hours: %s, %s", zap.String("status", resp.Status), zap.String("body", string(bodyBytes)))
		return []string{}, fmt.Errorf("error while getting available hours: %s", resp.Status)
	}
	if strings.Contains(string(bodyBytes), "ere are no appointment shedules available for the selec") {
		return []string{}, nil
	}
	pattern := `([01]?[0-9]|2[0-3]):([0-5][0-9]) - (0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-(\d{4})`

	r, err := regexp.Compile(pattern)
	if err != nil {
		c.logger.Error("error while compiling regex for available hours: %s", zap.Error(err))
		return []string{}, err
	}

	normalizedInput := strings.ReplaceAll(string(bodyBytes), "\n", " ")
	normalizedInput = strings.ReplaceAll(normalizedInput, ".", " ")
	normalizedInput = strings.TrimSpace(normalizedInput)

	splitInput := regexp.MustCompile(`\s{2,}`).Split(normalizedInput, -1)

	if strings.Contains(string(bodyBytes), "Suggestion of upcoming dates by time") {

		for _, str := range splitInput {
			if str != "" && r.MatchString(str) {
				str = strings.ReplaceAll(str, " ", "")
				str = strings.ReplaceAll(str, "<span>", "")
				str = strings.ReplaceAll(str, "</span>", "")
				availableHours = append(availableHours, str)
			}
		}
	} else {
		c.logger.Warn("unknown response for available hours: %s", zap.String("body", string(bodyBytes)))
		return []string{}, fmt.Errorf("unknown response for available hours: %s", string(bodyBytes))
	}

	return availableHours, nil
}
