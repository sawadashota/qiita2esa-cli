package main

import (
	"qiita2esa-cli/qiita"
	"strconv"
	"qiita2esa-cli/esa"
	"net/http"
	"time"
	"flag"
	"os"
)

const RequestInterval = 10

func main() {
	var (
		statusCode int
		body string
		qiitaPost qiita.Post
	)

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	qiitaTeamName := fs.String("q", "", "Qiita::Team Name")
	qiitaToken := fs.String("qToken", "", "Qiita::Team Access Token")
	esaTeamName := fs.String("e", "", "esa team name if not same Qiita::Team Name")
	esaToken := fs.String("eToken", "", "esa Access Token")
	fs.Parse(os.Args[1:])

	if *qiitaTeamName == "" {
		panic("Please type -q (Qiita::Team Name)")
	}

	if *qiitaToken == "" {
		panic("Please type -qToken (Qiita::Team Access Token)")
	}

	if *esaToken == "" {
		panic("Please type -eToken (esa Access Token)")
	}

	if *esaTeamName == "" {
		esaTeamName = qiitaTeamName
	}

	esaMembers := esa.Members(*esaTeamName, *esaToken)

	i := 1
	for {
		statusCode, qiitaPost = qiita.GetPost(i, *qiitaTeamName, *qiitaToken)

		if statusCode != 200 {
			println(strconv.Itoa(i-1) + " posts processed")
			break
		}

		if !esa.ExistMember(esaMembers, qiitaPost.User.ID) {
			qiitaPost.User.ID = "esa_bot"
		}

		println("Processing : " + qiitaPost.ID)
		statusCode, body = esa.Create(qiitaPost).PostTeam(*esaTeamName, *esaToken)

		if statusCode != http.StatusCreated {
			println(statusCode)
			println(qiitaPost.ID + " : " + body)
		}

		time.Sleep(RequestInterval * time.Second)
		i++
	}

}
