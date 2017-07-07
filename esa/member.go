package esa

import (
	"net/url"
	"net/http"
	"io"
	"encoding/json"
)

type ScreenName struct {
	MaxPerPage int                `json:"max_per_page"`
	Members    []ScreenNameMember `json:"members"`
	NextPage   interface{}        `json:"next_page"`
	Page       int                `json:"page"`
	PerPage    int                `json:"per_page"`
	PrevPage   interface{}        `json:"prev_page"`
	TotalCount int                `json:"total_count"`
}

type ScreenNameMember struct {
	Email      string `json:"email"`
	Icon       string `json:"icon"`
	Name       string `json:"name"`
	PostsCount int    `json:"posts_count"`
	ScreenName string `json:"screen_name"`
}

type GetClient struct {
	Token    string
	TeamName string
	Endpoint url.URL
}

func Members(teamName string, token string) []string {
	client := GetClient{TeamName: teamName, Token: token}
	client.Endpoint = client.generateEndpoint("/v1/teams/"+client.TeamName+"/members")

	screenName := client.getApi()

	return screenName.membersArray()
}

func (c GetClient) generateEndpoint(path string) url.URL {
	endpoint := url.URL{}
	endpoint.Scheme = "https"
	endpoint.Host = Host
	endpoint.Path = path

	return endpoint
}

func (c GetClient) getApi() ScreenName {
	req, _ := http.NewRequest("GET", c.Endpoint.String(), nil)
	req.Header.Set("Authorization", "Bearer "+c.Token)

	httpClient := new(http.Client)
	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("Bad request at fetching esa members.")
	}

	screenName := jsonParse(resp.Body)

	return screenName
}

func jsonParse(jsonRaw io.ReadCloser) ScreenName {
	var screenName ScreenName
	json.NewDecoder(jsonRaw).Decode(&screenName)

	return screenName
}

func (s ScreenName) membersArray() []string {
	var screenNames []string

	for _, members := range s.Members {
		screenNames = append(screenNames, members.ScreenName)
	}

	return screenNames
}

