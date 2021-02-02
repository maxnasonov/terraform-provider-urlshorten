package urlshorten

import (
	"fmt"

	"github.com/Jeffail/gabs"
	"github.com/go-resty/resty/v2"
)

func supportedServices() map[string]string {
	return map[string]string{"cuttly": "yes", "tinyurl": "yes", "bitly": "yes", "v.gd": "yes"}
}

type Service struct {
	Name     string
	Password string
}

func NewClient(service, password string) (*Service, error) {
	if _, ok := supportedServices()[service]; ok {
		s := &Service{Name: service, Password: password}
		return s, nil
	}
	return nil, fmt.Errorf("%s is not a supported client", service)
}

func (s *Service) extractURL(resp *resty.Response, location string) (*string, error) {
	var newURL string
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("Could not get short URL")
	}
	strResponse := resp.String()
	if location == "" {
		return &strResponse, nil
	}
	jsonParsed, err := gabs.ParseJSON([]byte(strResponse))
	if err != nil {
		return nil, err
	}
	shortURL, ok := jsonParsed.Path(location).Data().(string)
	if !ok || shortURL == "" {
		err = fmt.Errorf("Response did not have short url")
		return nil, err
	}
	newURL = shortURL
	return &newURL, nil
}

func (s *Service) logResponse(resp *resty.Response) {
	if resp != nil {
		fmt.Printf("Response Status Code: %v", resp.StatusCode())
		fmt.Printf("Response Time: %v", resp.Time())
		fmt.Printf("Header: %v", resp.Header())
		fmt.Printf("Response Received At: %v", resp.ReceivedAt())
		fmt.Printf("Response Body: %v", resp.String())
	}
}

func (s *Service) ShortenURL(URL string) (*string, error) {
	var err error
	var resp *resty.Response

	if s == nil {
		return nil, fmt.Errorf("Not found service")
	}

	if s.Name == "cuttly" {
		request := resty.New().R().
			SetQueryParams(map[string]string{
				"key":   s.Password,
				"short": URL,
			})
		resp, err = request.Get("https://cutt.ly/api/api.php")
		s.logResponse(resp)
		if err != nil {
			return nil, err
		}
		return s.extractURL(resp, "url.shortLink")
	}
	if s.Name == "tinyurl" {
		request := resty.New().R().
			SetQueryParams(map[string]string{
				"url": URL,
			})
		resp, err = request.Get("https://tinyurl.com/api-create.php")
		s.logResponse(resp)
		if err != nil {
			return nil, err
		}
		return s.extractURL(resp, "")
	}
	if s.Name == "bitly" {
		request := resty.New().R().
			SetHeaders(map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", s.Password),
				"Content-Type":  "application/json",
			}).
			SetBody(map[string]string{
				"domain":   "bit.ly",
				"long_url": URL,
			})
		resp, err = request.Post("https://api-ssl.bitly.com/v4/shorten")
		s.logResponse(resp)
		if err != nil {
			return nil, err
		}
		return s.extractURL(resp, "link")
	}
	if s.Name == "v.gd" {
		request := resty.New().R().
			SetQueryParams(map[string]string{
				"format": "json",
				"url":    URL,
			})
		resp, err = request.Get("https://v.gd/create.php")
		s.logResponse(resp)
		if err != nil {
			return nil, err
		}
		return s.extractURL(resp, "shorturl")
	}
	return nil, fmt.Errorf("%s is not a supported client", s.Name)
}
