package qiita

import (
	"io"
	"encoding/json"
)

type Post struct {
	Body         string      `json:"body"`
	Coediting    bool        `json:"coediting"`
	CreatedAt    string      `json:"created_at"`
	ID           string      `json:"id"`
	Private      bool        `json:"private"`
	RenderedBody string      `json:"rendered_body"`
	Title        string      `json:"title"`
	UpdatedAt    string      `json:"updated_at"`
	URL          string      `json:"url"`
	User         QiitaUser  `json:"user"`
}

type QiitaUser struct {
	Description       interface{} `json:"description"`
	FacebookID        interface{} `json:"facebook_id"`
	FolloweesCount    int         `json:"followees_count"`
	FollowersCount    int         `json:"followers_count"`
	GithubLoginName   interface{} `json:"github_login_name"`
	ID                string      `json:"id"`
	ItemsCount        int         `json:"items_count"`
	LinkedinID        interface{} `json:"linkedin_id"`
	Location          interface{} `json:"location"`
	Name              string      `json:"name"`
	Organization      interface{} `json:"organization"`
	PermanentID       int         `json:"permanent_id"`
	ProfileImageURL   string      `json:"profile_image_url"`
	TwitterScreenName interface{} `json:"twitter_screen_name"`
	WebsiteURL        interface{} `json:"website_url"`
}

func JsonParse(jsonRaw io.ReadCloser) []Post {
	var posts []Post
	json.NewDecoder(jsonRaw).Decode(&posts)

	if len(posts) < 1 {
		panic("Post should be array.")
	}

	return posts
}
