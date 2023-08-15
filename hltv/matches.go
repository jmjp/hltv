package hltv

import (
	"strconv"
	"strings"
	"time"

	"github.com/jmjp/hltv/models"

	"github.com/PuerkitoBio/goquery"
)

func FetchMatches() ([]models.UpcomingMatche, error) {
	res, err := Fetch("/matches")
	if err != nil {
		println(err.Error())
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res.Body))
	if err != nil {
		return nil, err
	}
	var upcomming_matches []models.UpcomingMatche

	doc.Find(".liveMatch-container").Add(".upcomingMatch").Each(func(i int, selection *goquery.Selection) {
		matches, _ := selection.Find(".a-reset").Attr("href")
		matchID, _ := strconv.Atoi(strings.Split(matches, "/")[2])
		matchTimestamp, _ := selection.Find(".matchTime").First().Attr("data-unix")
		date := UnixTimeStringToTime(matchTimestamp)
		eventName := selection.Find(".matchEventName").First().Text()
		team1Name := selection.Find(".team1 .matchTeamName").First().Text()
		team1IDStr, _ := selection.Attr("team1")
		team1ID, _ := strconv.Atoi(team1IDStr)

		team2Name := selection.Find(".team2 .matchTeamName").Last().Text()
		team2IDStr, _ := selection.Attr("team2")
		team2ID, _ := strconv.Atoi(team2IDStr)
		format := selection.Find(".matchMeta").Last().Text()
		live := selection.Find(".matchTime.matchLive").Text() == "LIVE"

		match := models.UpcomingMatche{
			Id:     matchID,
			Format: format,
			Live:   live,
		}
		if date.Year() > 2015 {
			match.Date = date
		} else {
			match.Date = time.Now()
		}
		if team1ID != 0 {
			match.TeamA = &models.Team{
				Id:   team1ID,
				Name: team1Name,
			}
		}
		if team2ID != 0 {
			match.TeamB = &models.Team{
				Id:   team2ID,
				Name: team2Name,
			}
		}
		if eventName != "" {
			match.Event = &eventName
		}
		upcomming_matches = append(upcomming_matches, match)
	})

	return upcomming_matches, nil
}

func UnixTimeStringToTime(unixTime string) time.Time {
	if len(unixTime) == 0 {
		return time.Unix(0, 0)
	}
	dateStartInt64, _ := strconv.ParseInt(unixTime[:len(unixTime)-3], 10, 64)
	return time.Unix(dateStartInt64, 0)
}
