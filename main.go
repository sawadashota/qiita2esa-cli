package main

import (
	"flag"
	"github.com/sawadashota/qiita-posts-go"
	"math"
	"net/http"
	"os"
	"qiita2esa-cli/esa"
	"strconv"
	"time"
)

const EsaRequestInterval = 12

func main() {
	var (
		qiitaStatusCode int
		esaStatusCode   int
		body            string
		qiitaPosts      []qiita.Post
		qiitaPost       qiita.Post
		key             int
		processID       int
	)

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	qiitaTeamName := fs.String("q", "", "Qiita::Team Name")
	qiitaToken := fs.String("qToken", "", "Qiita::Team Access Token")
	esaTeamName := fs.String("e", "", "esa team name if not same Qiita::Team Name")
	esaToken := fs.String("eToken", "", "esa Access Token")
	restartFromStr := fs.String("restart-from", "1", "Restart from process ID")
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

	restartFrom, err := strconv.Atoi(*restartFromStr)
	if err != nil {
		panic("-restart-from should be integer.")
	}

	startingQiitaPage := postPage(restartFrom, qiita.PagePerPost)

	esaMembers := esa.Members(*esaTeamName, *esaToken)

	for i := startingQiitaPage; ; i++ {

		qiitaStatusCode, qiitaPosts = qiita.Posts(i, *qiitaTeamName, *qiitaToken).Get()

		if qiitaStatusCode != 200 {
			println("-----------------------------------------------")
			println("Qiita Status Code: " + strconv.Itoa(qiitaStatusCode))
			println("-----------------------------------------------")
			break
		}

		for key, qiitaPost = range qiitaPosts {
			processID = (qiita.PagePerPost * (i - 1)) + (key + 1)

			if processID < restartFrom {
				continue
			}

			if !esa.ExistMember(esaMembers, qiitaPost.User.ID) {
				qiitaPost.User.ID = "esa_bot"
			}

			print(strconv.Itoa(processID) + ". Processing : " + qiitaPost.ID + " ...")

			esaStatusCode, body = esa.Create(qiitaPost).PostTeam(*esaTeamName, *esaToken)

			if esaStatusCode != http.StatusCreated {
				println("")
				println("-----------------------------------------------")
				println("esa Status Code: " + strconv.Itoa(esaStatusCode))
				println(qiitaPost.ID + " : " + body)
				println("See esa document: https://docs.esa.io/posts/102")
				println("-----------------------------------------------")
			} else {
				println(" Complete!")
			}

			time.Sleep(EsaRequestInterval * 1000 * time.Millisecond)
		}

		if len(qiitaPosts) < qiita.PagePerPost {
			println("-----------------------------------------------")
			println(strconv.Itoa(processID-restartFrom) + " posts processed!")
			println("-----------------------------------------------")
			break
		}
	}
}

// Qiitaは何ページ目から読み込むか
func postPage(processId int, pagePerPost int) int {
	return int(math.Floor(float64((processId-1)/pagePerPost))) + 1
}
