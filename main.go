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

const EsaRequestInterval = 10
const QiitaRequestInterval = 3.6

func main() {
	var (
		statusCode int
		body       string
		qiitaPost  qiita.Post
	)

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	qiitaTeamName := fs.String("q", "", "Qiita::Team Name")
	qiitaToken := fs.String("qToken", "", "Qiita::Team Access Token")
	esaTeamName := fs.String("e", "", "esa team name if not same Qiita::Team Name")
	esaToken := fs.String("eToken", "", "esa Access Token")
	startFrom := fs.String("start-from", "", "Skip until this Qiita::Team post ID")
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

	passedStartFrom := *startFrom == ""

	esaMembers := esa.Members(*esaTeamName, *esaToken)

	for i := 1; ; i++ {
		statusCode, qiitaPost = qiita.GetPost(i, *qiitaTeamName, *qiitaToken)

		if statusCode != 200 {
			println(strconv.Itoa(i-1) + " posts processed!")
			break
		}

		if !esa.ExistMember(esaMembers, qiitaPost.User.ID) {
			qiitaPost.User.ID = "esa_bot"
		}

		print("Processing : " + qiitaPost.ID + " ...")

		passedStartFrom = shouldSkip(passedStartFrom, *startFrom, qiitaPost.ID)
		if !passedStartFrom {
			println(" Skipped")
			time.Sleep(QiitaRequestInterval * 1000 * time.Millisecond)
			continue
		}

		statusCode, body = esa.Create(qiitaPost).PostTeam(*esaTeamName, *esaToken)

		if statusCode != http.StatusCreated {
			println("")
			println("-----------------------------------------------")
			println(statusCode)
			println(qiitaPost.ID + " : " + body)
			println("-----------------------------------------------")
		} else {
			println(" Complete!")
		}

		time.Sleep(EsaRequestInterval * 1000 * time.Millisecond)
	}

}

// esaへの投稿をスキップするかの判定
func shouldSkip(passedStartFrom bool, startFrom string, qiitaPostId string) bool {
	if passedStartFrom {
		return passedStartFrom
	}

	return startFrom == qiitaPostId
}
