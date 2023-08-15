package hltv

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jmjp/hltv/models"

	"github.com/PuerkitoBio/goquery"
)

func FetchTeamById(id int) (*models.TeamInfo, error) {
	uri := fmt.Sprint("/team/", id, "/xpskd291")
	res, err := Fetch(uri)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res.Body))
	if err != nil {
		return nil, err
	}
	var teamInfo models.TeamInfo
	// Use seletores goquery para extrair informações
	teamInfo.Id = id
	teamInfo.LogoURL = doc.Find(".profile-team-logo-container img").AttrOr("src", "")
	teamInfo.Country = doc.Find(".team-country img").AttrOr("title", "")
	teamInfo.Name = doc.Find(".profile-team-name").Text()
	world, err := strconv.Atoi(strings.Split(doc.Find(".profile-team-stat:nth-child(1) a").Text(), "#")[1])
	teamInfo.WorldRanking = uint8(world)
	weeks, err := strconv.Atoi(doc.Find(".profile-team-stat:nth-child(2) span.right").Text())
	if err != nil {
		weeks = 0
	}
	teamInfo.WeeksInTop30 = weeks
	avg, err := strconv.ParseFloat(doc.Find(".profile-team-stat:nth-child(3) span.right").Text(), 32)
	teamInfo.AveragePlayerAge = float32(avg)
	teamInfo.CoachName = strings.TrimSpace(doc.Find(".profile-team-stat:nth-child(4) a").Text())

	var playerInfoList []models.TeamPlayers

	doc.Find(".table-container.players-table tbody tr").Each(func(i int, s *goquery.Selection) {
		var playerInfo models.TeamPlayers
		phref, _ := s.Find(".playersBox-playernick-image").Attr("href")
		playerId, _ := strconv.Atoi(strings.Split(phref, "/")[2])
		playerInfo.Id = playerId
		playerInfo.Name = s.Find(".playersBox-playernick .text-ellipsis").Text()
		playerInfo.Status = s.Find(".player-status").Text()

		pattern := `[0-9]+`
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(s.Find(".players-cell.opacity-cell").Eq(0).Text(), -1)

		matchesLen := len(matches)

		if matchesLen == 1 {
			months, _ := strconv.ParseFloat(matches[0], 32)
			playerInfo.TimeOnTeam = float32(months)
		} else {
			years, _ := strconv.ParseFloat(matches[0], 32)
			months, _ := strconv.ParseFloat(matches[1], 32)
			playerInfo.TimeOnTeam = float32(years) + float32(months)/100
		}

		maps, _ := strconv.Atoi(s.Find(".players-cell.opacity-cell").Eq(1).Text())
		playerInfo.MapsPlayed = maps
		rating, _ := strconv.ParseFloat(s.Find(".rating-cell").Text(), 32)
		playerInfo.Rating = float32(rating)

		playerInfoList = append(playerInfoList, playerInfo)
	})
	teamInfo.Lineup = &playerInfoList

	return &teamInfo, nil
}
