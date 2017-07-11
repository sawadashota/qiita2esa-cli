package esa

import (
	"github.com/sawadashota/qiita-posts-go"
	"net/url"
	"bytes"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strings"
)

const  Host = "api.esa.io"

type PostEnvelope struct {
	Post Post `json:"post"`
}

type Post struct {
	Name     string `json:"name"`
	BodyMd   string `json:"body_md"`
	Category string `json:"category"`
	Wip      bool `json:"wip"`
	Message  string `json:"message"`
	User     string `json:"user"`
}

type Client struct {
	Token    string
	TeamName string
	Endpoint url.URL
	Values   *bytes.Buffer
}

func Create(qiitaPost qiita.Post) Post {
	title := strings.Replace(qiitaPost.Title, "/", "-", -1)

	return Post{
		Name:     title,
		BodyMd:   qiitaPost.Body,
		Category: "Import/Qiita::Team",
		Wip:      false,
		Message:  "Import from Qiita::Team via qiita2esa-cli.",
		User:     qiitaPost.User.ID,
	}
}

func (p Post) PostTeam(teamName string, token string) (int, string) {
	client := Client{Token: token, TeamName: teamName}
	client.Endpoint = client.generateEndpoint("/v1/teams/"+client.TeamName+"/posts")
	client.Values = p.setJsonValues()

	return client.postApi()
}

func (c Client) generateEndpoint(path string) url.URL {
	endpoint := url.URL{}
	endpoint.Scheme = "https"
	endpoint.Host = Host
	endpoint.Path = path

	return endpoint
}

func (p Post)setJsonValues() *bytes.Buffer {
	postEnvelope := PostEnvelope{Post: p}
	requestBody, err := json.Marshal(postEnvelope)

	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(requestBody)
}


func (c Client) postApi() (int, string) {
	req, _ := http.NewRequest("POST", c.Endpoint.String(), c.Values)
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	httpClient := new(http.Client)
	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, "Fail to post Esa"
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, string(body)
}
