package qiita

import (
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	Token    string
	TeamName string
	Endpoint url.URL
	Params   url.Values
}

func GetPost(page int, teamName string, token string) (int, Post) {
	client := Client{TeamName: teamName, Token: token}
	client.Endpoint = client.generateEndpoint("/api/v2/items")
	client.Params = setValues(strconv.Itoa(page))

	return client.request()
}

func (c Client) generateEndpoint(path string) url.URL {
	endpoint := url.URL{}
	endpoint.Scheme = "https"
	endpoint.Host = c.TeamName + ".qiita.com"
	endpoint.Path = path

	return endpoint
}

func setValues(page string) url.Values {
	values := url.Values{}
	values.Add("per_page", "1")
	values.Add("page", page)

	return values
}

func (c Client) request() (int, Post) {
	req, _ := http.NewRequest("GET", c.Endpoint.String(), nil)
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.URL.RawQuery = c.Params.Encode()

	httpClient := new(http.Client)
	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, Post{}
	}

	post := JsonParse(resp.Body)

	return resp.StatusCode, post
}