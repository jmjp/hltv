package hltv

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmjp/hltv/models"

	"github.com/PuerkitoBio/goquery"
)

func FetchPlayerById(id int) (*models.PlayerInfo, error) {
	uri := fmt.Sprint("/player/", id, "/xpskd291")
	res, err := Fetch(uri)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res.Body))
	if err != nil {
		log.Fatal(err)
	}
	var player models.PlayerInfo
	player.Id = id
	player.Ign = doc.Find(".playerNickname").Text()
	player.Name = doc.Find(".playerRealname").Text()
	age, _ := strconv.Atoi(strings.ReplaceAll(doc.Find(".playerAge span[itemprop='text']").Text(), " years", ""))
	player.Age = &age
	top20 := doc.Find(".playerTop20 .top20ListRight").Text()
	majorWinds, err := strconv.ParseInt(doc.Find(".majorWinner b").Text(), 10, 16)
	if err != nil {
		majorWinds = 0
	}
	majorInt := int16(majorWinds)
	player.MajorWins = &majorInt
	image, _ := doc.Find(".bodyshot-img").Attr("src")
	player.Stats = stats(doc)
	if top20 != "" {
		player.Top20 = &top20
	}
	if image != "" {
		player.Image = &image
	}

	team_link, _ := doc.Find(".playerTeam a").Attr("href")
	if team_link != "" {
		team := strings.Split(team_link, "/")
		id, _ := strconv.Atoi(team[2])
		name := team[3]
		player.CurrentTeam = &models.Team{
			Id:   id,
			Name: name,
		}
	}
	return &player, nil
}

func FetchPlayerStatsById(id int) (*models.PlayerSummary, error) {
	uri := fmt.Sprint("/stats/players/", id, "/svgpading")
	res, err := Fetch(uri)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res.Body))
	if err != nil {
		log.Fatal(err)
	}
	var player models.PlayerSummary
	player.Id = id
	player.Ign = doc.Find(".summaryNickname").Text()
	player.Name = doc.Find(".summaryRealname div.text-ellipsis").Text()
	age, _ := strconv.Atoi(strings.ReplaceAll(doc.Find(".summaryPlayerAge").Text(), " years", ""))
	player.Age = &age
	doc.Find(".summaryStatBreakdownRow").Each(func(i int, s *goquery.Selection) {
		stat := models.StatEntry{
			Name:        strings.TrimSpace(s.Find(".summaryStatBreakdownSubHeader").Text()),
			Value:       strings.TrimSpace(s.Find(".summaryStatBreakdownDataValue").Text()),
			Description: strings.TrimSpace(s.Find(".summaryStatBreakdownDescription").Text()),
		}
		player.Stats = append(player.Stats, stat)
	})
	team_link, _ := doc.Find(".SummaryTeamname a").Attr("href")
	if team_link != "" {
		team := strings.Split(team_link, "/")
		id, _ := strconv.Atoi(team[3])
		name := team[4]
		player.Team = &models.Team{
			Id:   id,
			Name: name,
		}
	}
	image, _ := doc.Find(".summaryBodyshot").Attr("src")
	if image != "" {
		player.Image = &image
	}
	player.Metrics = metrics(doc)
	return &player, nil
}

func FetchPlayerStatsByIdWithContext(id int, ctx context.Context, ch chan *models.PlayerSummary) {
	p, err := FetchPlayerStatsById(id)
	if err != nil {
		ctx.Err()
	} else {
		ctx.Done()
		ch <- p
	}
}
func FetchPlayerByIdWithContext(id int, ctx context.Context, ch chan *models.PlayerInfo) {
	p, err := FetchPlayerById(id)
	if err != nil {
		ctx.Err()
	} else {
		ctx.Done()
		ch <- p
	}
}

func stats(doc *goquery.Document) *models.PlayerInfoStats {
	var playerStats models.PlayerInfoStats
	doc.Find(".player-stat").Each(func(i int, s *goquery.Selection) {
		statName := s.Find("b").Text()
		statValue := s.Find(".statsVal p").Text()

		float_stat, err := strconv.ParseFloat(strings.ReplaceAll(statValue, "%", ""), 64)
		if err != nil {
			float_stat = 0.0
		}

		int_stat, err := strconv.Atoi(statValue)
		if err != nil {
			int_stat = 0
		}

		switch statName {
		case "Rating 2.0":
			playerStats.Rating = float32(float_stat)
		case "Kills per round":
			playerStats.KillsPerRound = float32(float_stat)
		case "Headshots":
			playerStats.Headshots = float32(float_stat)
		case "Maps played":
			playerStats.MapsPlayed = int_stat
		case "Deaths per round":
			playerStats.DeathsPerRound = float32(float_stat)
		case "Rounds contributed":
			playerStats.RoundsContrib = float32(float_stat)
		}
	})
	return &playerStats
}

func metrics(doc *goquery.Document) *models.PlayerMetrics {
	var playerStats models.PlayerMetrics
	doc.Find(".stats-rows").Each(func(i int, col *goquery.Selection) {
		col.Find(".stats-row").Each(func(j int, s *goquery.Selection) {
			statName := s.Find("span:nth-child(1)").Text()
			statValue := s.Find("span:nth-child(2)").Text()

			float_stat, err := strconv.ParseFloat(strings.ReplaceAll(statValue, "%", ""), 64)
			if err != nil {
				float_stat = 0.0
			}

			int_stat, err := strconv.Atoi(statValue)
			if err != nil {
				int_stat = 0
			}

			switch statName {
			case "Total kills":
				playerStats.TotalKills = int_stat
			case "Headshot %":
				playerStats.HeadshotPercent = float32(float_stat)
			case "Total deaths":
				playerStats.TotalDeaths = int_stat
			case "K/D Ratio":
				playerStats.KDRatio = float32(float_stat)
			case "Damage / Round":
				playerStats.DamagePerRound = float32(float_stat)
			case "Grenade dmg / Round":
				playerStats.GrenadeDmgRound = float32(float_stat)
			case "Maps played":
				playerStats.MapsPlayed = int_stat
			case "Rounds played":
				playerStats.RoundsPlayed = int_stat
			case "Kills / round":
				playerStats.KillsPerRound = float32(float_stat)
			case "Assists / round":
				playerStats.AssistsPerRound = float32(float_stat)
			case "Deaths / round":
				playerStats.DeathsPerRound = float32(float_stat)
			case "Saved by teammate / round":
				playerStats.SavedByTeammate = float32(float_stat)
			case "Saved teammates / round":
				playerStats.SavedTeammates = float32(float_stat)
			case "Rating 1.0":
				playerStats.Rating1_0 = float32(float_stat)
			}
		})
	})
	return &playerStats
}
