package cwp

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	// "strings"

	"os"
)

type OpenWeatherMap struct {
	url   string
	place   string
	// group string
}

// type Group struct {
// 	Guid     string `json:"guid"`
// 	IsActive bool   `json:"is_active"`
// }

// func (openweathermap *OpenWeatherMap) Groups(config *Config) ([]*Group, error) {
// 	request, err := http.NewRequest("GET", openweathermap.url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	data, err := sendRequest(request, config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return parseGroups(data)
// }

// func parseGroups(data []byte) ([]*Group, error) {
// 	result := struct {
// 		Groups []*Group `json:"groups"`
// 	}{}
// 	err := json.Unmarshal(data, &result)
//
// 	if err != nil {
// 		return nil, err
// 	}
// 	returnValue := []*Group{}
// 	for _, g := range result.Groups {
// 		if g.IsActive {
// 			returnValue = append(returnValue, g)
// 		}
// 	}
// 	return returnValue, nil
// }

// func NewBitly(group string) *Bitly {
// 	return &Bitly{url: "https://api-ssl.bitly.com/v4/", group: group}
// }

// func NewOpenWeatherMap(group string) *OpenWeatherMap {
func NewOpenWeatherMap() *OpenWeatherMap {
	return &OpenWeatherMap{url: "https://api.openweathermap.org/data/2.5/weather"}
}

// func (openweathermap *OpenWeatherMap) buildUrl(endpoint string) string {
// 	return fmt.Sprintf("%s/%s", openweathermap.url, endpoint)
// }

// func handleGroup(config *Config, openweathermap *OpenWeatherMap) (string, error) {
// 	// if openweathermap.group != "" {
// 	// 	return openweathermap.group, nil
// 	// }
// 	gs, err := openweathermap.Groups(config)
// 	if err != nil {
// 		return "", err
// 	}
// 	if len(gs) == 0 {
// 		return "", fmt.Errorf("no active groups")
// 	}
// 	return gs[0].Guid, nil
// }

func (openweathermap *OpenWeatherMap) List(config *Config) ([]*Weather, error) {
	// group, err := handleGroup(config, openweathermap)
	// if err != nil {
	// 	return nil, err
	// }

        // request, err := http.NewRequest("GET", openweathermap.buildUrl(fmt.Sprintf("/groups/%s/bitlinks?size=20", group)), nil)
        request, err := http.NewRequest("GET", openweathermap.url, nil)
	if err != nil {
		return nil, err
	}
	// data, err := sendRequest(request, config)
        data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	return handleListResponse(data)
}

func handleListResponse(data []byte) ([]*Weather, error) {
	result := struct {
		Links []*Weather `json:"links"`
	}{}
	err := json.Unmarshal(data, &result)
	// return removeDeletedLinks(result.Links, group), err
	return nil, err
}

// func removeDeletedLinks(links []*ShortenUrl, group string) []*ShortenUrl {
// 	result := []*ShortenUrl{}
// 	for _, link := range links {
// 		if !link.IsDeleted {
// 			link.Group = group
// 			result = append(result, link)
// 		}
// 	}
// 	return result
// }

func (openweathermap *OpenWeatherMap) GetWeather(place string, config *Config) (*Weather, error) {
	// reader := strings.NewReader(fmt.Sprintf(`{"place": "%s", "group_guid": "%s"}`, place, openweathermap.group))

	// reader := strings.NewReader(fmt.Sprintf(`{
	// 	"q": "%s",
	// 	"appid": "%s",
	// 	"units": "metric",
	// 	"lang": "ja",
	// 	}`, place, token))
        // fmt.Println(reader)

	// request, err := http.NewRequest("POST", openweathermap.url, reader)

	// if err != nil {
	// 	return nil, err
	// }
	// data, err := sendRequest(request, config)
	// if err != nil {
	// 	return nil, err
	// }

	data, err := sendRequest(openweathermap, place, config)
        // data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	return handleApiResponse(data)
}

func handleApiResponse(data []byte) (*Weather, error) {
	result := &Weather{}
	fmt.Println("result:", string(data))
	err := json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	// result.Group, err = findGroup(data)
        os.Exit(0)
	return result, err
}

// func findGroup(data []byte) (string, error) {
// 	r := struct {
// 		References struct {
// 			Group string `json:"group"`
// 		} `json:"references"`
// 	}{}
// 	err := json.Unmarshal(data, &r)
// 	uri := r.References.Group
// 	index := strings.LastIndex(uri, "/")
// 	return uri[index+1:], err
// }

// func (openweathermap *OpenWeatherMap) Delete(config *Config, shortenURL string) error {
// 	request, err := http.NewRequest("DELETE", openweathermap.buildUrl("bitlinks/"+strings.TrimPrefix(shortenURL, "https://")), nil)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = sendRequest(request, config)
// 	return err
// }

// func (openweathermap *OpenWeatherMap) QRCode(config *Config, shortenURL string) ([]byte, error) {
// 	return []byte{}, fmt.Errorf("not implement yet")
// }

func sendRequest(openweathermap *OpenWeatherMap, place string, config *Config) ([]byte, error) {
	params := url.Values{}
	params.Set("q", place)
	params.Set("appid", config.Token)
	params.Set("units", "metric")
	params.Set("lang", "ja")

	response, err := http.Get(openweathermap.url + "?" + params.Encode())
	// response, err := sendRequestImpl(request, config)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	return handleResponse(response)
}

// func sendRequestImpl(request *http.Request, config *Config) (*http.Response, error) {
// 	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Token))
// 	request.Header.Add("Content-Type", "application/json")
// 	client := &http.Client{}
// 	return client.Do(request)
// }

func handleResponse(response *http.Response) ([]byte, error) {
	if response.StatusCode/100 != 2 {
		data, _ := io.ReadAll(response.Body)
		fmt.Println("response body:", string(data))
		return []byte{}, fmt.Errorf("response status code %d", response.StatusCode)
	}
	return io.ReadAll(response.Body)
}
