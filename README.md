# HLTV
[![Go Reference](https://pkg.go.dev/badge/github.com/jmjp/hltv.svg)](https://pkg.go.dev/github.com/jmjp/hltv)
<p>Basic HLTV scraper, fetch data from html and transform into model


### usage
```
go get github.com/jmjp/hltv
```

```go
import "github.com/jmjp/hltv"
```

## Methods

```golang
FetchTeamById(10831)
```
<details><summary>Output</summary>
<p>

```json
{
	"team": {
		"id": 10831,
		"name": "Entropiq",
		"logo": "https://img-cdn.hltv.org/teamlogo/_x9xMubeeb-fqq_21BJ6Fd.png?ixlib=java-2.1.0&w=50&s=3ab3a2cbfb2ed2e30c1a71ca8113ae39",
		"country": "Europe",
		"world_ranking": 110,
		"weeks_in_top30": 0,
		"average_player_age": 22.8,
		"coach_name": "Niclas 'enkay J' Krumhorn",
		"lineup": [
			{
				"id": 19069,
				"name": "forsyy",
				"status": "STARTER",
				"time_on_team": 1,
				"maps_Played": 14,
				"rating": 1.08
			},
			{
				"id": 13026,
				"name": "c0llins",
				"status": "STARTER",
				"time_on_team": 19,
				"maps_Played": 14,
				"rating": 0.93
			},
			{
				"id": 20110,
				"name": "Marix",
				"status": "STARTER",
				"time_on_team": 19,
				"maps_Played": 14,
				"rating": 1.16
			},
			{
				"id": 18697,
				"name": "oxygeN",
				"status": "STARTER",
				"time_on_team": 19,
				"maps_Played": 14,
				"rating": 0.92
			},
			{
				"id": 5796,
				"name": "tiziaN",
				"status": "STARTER",
				"time_on_team": 19,
				"maps_Played": 14,
				"rating": 0.97
			}
		]
	}
}
```
</p>
</details>

```golang
FetchPlayerById(7998)
```

<details><summary>Output</summary>
<p>

```json
{
	"player": {
		"id": 7998,
		"ign": "s1mple",
		"name": " Oleksandr Kostyliev",
		"age": 25,
		"image": "https://img-cdn.hltv.org/playerbodyshot/4NNsZrSGWLr9mZNt0Pe3KS.png?ixlib=java-2.1.0&w=400&s=4484cc99087121a6f9877d3742717444",
		"team": {
			"id": 4608,
			"name": "natus-vincere"
		},
		"presence_in_top20": "#4('16), #8('17), #1('18), #2('19), #2('20), #1('21), #1('22)",
		"major_wins": 1,
		"stats": {
			"rating": 1.1,
			"kills_per_round": 0.71,
			"headshots": 42.7,
			"maps_played": 30,
			"deaths_per_round": 0.63,
			"rounds_contributed": 70.9
		}
	}
}
```

</p>
</details>


```golang
FetchMatches()
```
<details><summary>Output</summary>
<p>

```json
{
	"matches": [
		{
			"id": 2365857,
			"team_a": {
				"id": 10567,
				"name": "SAW"
			},
			"team_b": {
				"id": 4501,
				"name": "ALTERNATE aTTaX"
			},
			"format": "bo3",
			"event": "CCT Central Europe Series 7",
			"live": true,
			"date": "2023-08-15T12:42:27.5093321-03:00"
		},
		{
			"id": 2366043,
			"team_a": {
				"id": 11251,
				"name": "Eternal Fire"
			},
			"team_b": {
				"id": 12371,
				"name": "ORKS"
			},
			"format": "bo3",
			"event": "Pinnacle Cup V",
			"live": true,
			"date": "2023-08-15T12:42:27.5093321-03:00"
		},
		{
			"id": 2366064,
			"team_a": {
				"id": 11883,
				"name": "9 Pandas"
			},
			"team_b": {
				"id": 10831,
				"name": "Entropiq"
			},
			"format": "bo3",
			"event": "CCT East Europe Series 1",
			"live": true,
			"date": "2023-08-15T12:42:27.5093321-03:00"
		},
		{
			"id": 2366055,
			"team_a": {
				"id": 11737,
				"name": "EYEBALLERS"
			},
			"team_b": {
				"id": 10854,
				"name": "los kogutos"
			},
			"format": "bo3",
			"event": "CCT North Europe Series 7",
			"live": true,
			"date": "2023-08-15T12:42:27.5093321-03:00"
		},
		{
			"id": 2366065,
			"team_a": {
				"id": 10459,
				"name": "ARCRED"
			},
			"team_b": {
				"id": 11628,
				"name": "ThunderFlash"
			},
			"format": "bo3",
			"event": "CCT East Europe Series 1",
			"live": false,
			"date": "2023-08-15T13:00:00-03:00"
		},
		{
			"id": 2365858,
			"team_a": {
				"id": 11861,
				"name": "Aurora"
			},
			"team_b": {
				"id": 11883,
				"name": "9 Pandas"
			},
			"format": "bo3",
			"event": "CCT Central Europe Series 7",
			"live": false,
			"date": "2023-08-15T14:00:00-03:00"
		},
		{
			"id": 2366056,
			"team_a": {
				"id": 12334,
				"name": "Preasy"
			},
			"team_b": {
				"id": 12394,
				"name": "BetBoom"
			},
			"format": "bo3",
			"event": "CCT North Europe Series 7",
			"live": false,
			"date": "2023-08-15T14:30:00-03:00"
		},
		{
			"id": 2366044,
			"team_a": {
				"id": 10695,
				"name": "Sampi"
			},
			"team_b": {
				"id": 12399,
				"name": "QM"
			},
			"format": "bo3",
			"event": "Pinnacle Cup V",
			"live": false,
			"date": "2023-08-15T15:30:00-03:00"
		},
		{
			"id": 2365838,
			"team_a": {
				"id": 10894,
				"name": "Case"
			},
			"team_b": {
				"id": 11814,
				"name": "Corinthians"
			},
			"format": "bo3",
			"event": "RedZone PRO League 2023 Season 5",
			"live": false,
			"date": "2023-08-15T18:00:00-03:00"
		},
		{
			"id": 2366057,
			"team_a": {
				"id": 12376,
				"name": "M80"
			},
			"team_b": {
				"id": 11631,
				"name": "Unjustified"
			},
			"format": "bo3",
			"event": "ESL Challenger League Season 46 North America",
			"live": false,
			"date": "2023-08-15T21:00:00-03:00"
		},
		{
			"id": 2366058,
			"team_a": {
				"id": 9799,
				"name": "Bad News Bears"
			},
			"team_b": {
				"id": 5479,
				"name": "Mythic"
			},
			"format": "bo3",
			"event": "ESL Challenger League Season 46 North America",
			"live": false,
			"date": "2023-08-15T21:00:00-03:00"
		},
		{
			"id": 2365859,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-16T05:00:00-03:00"
		},
		{
			"id": 2366147,
			"team_a": {
				"id": 12255,
				"name": "KS"
			},
			"team_b": {
				"id": 10873,
				"name": "Enterprise"
			},
			"format": "bo3",
			"event": "European Pro League 2nd Division Season 5",
			"live": false,
			"date": "2023-08-16T05:00:00-03:00"
		},
		{
			"id": 2366152,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-16T05:30:00-03:00"
		},
		{
			"id": 2366146,
			"team_a": {
				"id": 11523,
				"name": "Zero Tenacity"
			},
			"team_b": {
				"id": 10854,
				"name": "los kogutos"
			},
			"format": "bo3",
			"event": "European Pro League 2nd Division Season 5",
			"live": false,
			"date": "2023-08-16T07:30:00-03:00"
		},
		{
			"id": 2365785,
			"team_a": {
				"id": 9806,
				"name": "Apeks"
			},
			"team_b": {
				"id": 9928,
				"name": "GamerLegion"
			},
			"format": "bo3",
			"event": "Gamers8 2023",
			"live": false,
			"date": "2023-08-16T08:00:00-03:00"
		},
		{
			"id": 2365786,
			"team_a": {
				"id": 5752,
				"name": "Cloud9"
			},
			"team_b": {
				"id": 4991,
				"name": "fnatic"
			},
			"format": "bo3",
			"event": "Gamers8 2023",
			"live": false,
			"date": "2023-08-16T08:00:00-03:00"
		},
		{
			"id": 2365860,
			"team_a": {
				"id": 10621,
				"name": "1WIN"
			},
			"team_b": {
				"id": 12260,
				"name": "Looking4Org"
			},
			"format": "bo3",
			"event": "CCT Central Europe Series 7",
			"live": false,
			"date": "2023-08-16T08:00:00-03:00"
		},
		{
			"id": 2366149,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-16T08:30:00-03:00"
		},
		{
			"id": 2366153,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-16T08:30:00-03:00"
		},
		{
			"id": 2365861,
			"team_a": {
				"id": 11811,
				"name": "Monte"
			},
			"team_b": {
				"id": 11251,
				"name": "Eternal Fire"
			},
			"format": "bo3",
			"event": "CCT Central Europe Series 7",
			"live": false,
			"date": "2023-08-16T11:00:00-03:00"
		},
		{
			"id": 2365787,
			"team_a": {
				"id": 7175,
				"name": "Heroic"
			},
			"team_b": {
				"id": 11283,
				"name": "Falcons"
			},
			"format": "bo3",
			"event": "Gamers8 2023",
			"live": false,
			"date": "2023-08-16T11:30:00-03:00"
		},
		{
			"id": 2365788,
			"team_a": {
				"id": 9215,
				"name": "MIBR"
			},
			"team_b": {
				"id": 9565,
				"name": "Vitality"
			},
			"format": "bo3",
			"event": "Gamers8 2023",
			"live": false,
			"date": "2023-08-16T11:30:00-03:00"
		},
		{
			"id": 2366154,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-16T11:30:00-03:00"
		},
		{
			"id": 2366150,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-16T12:00:00-03:00"
		},
		{
			"id": 2365862,
			"team_a": {
				"id": 11164,
				"name": "Into the Breach"
			},
			"team_b": {
				"id": 12289,
				"name": "Espionage"
			},
			"format": "bo3",
			"event": "CCT Central Europe Series 7",
			"live": false,
			"date": "2023-08-16T14:00:00-03:00"
		},
		{
			"id": 2366155,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-16T14:30:00-03:00"
		},
		{
			"id": 2365789,
			"team_a": {
				"id": 5973,
				"name": "Liquid"
			},
			"team_b": {
				"id": 4869,
				"name": "ENCE"
			},
			"format": "bo3",
			"event": "Gamers8 2023",
			"live": false,
			"date": "2023-08-16T15:00:00-03:00"
		},
		{
			"id": 2365790,
			"team_a": {
				"id": 4608,
				"name": "Natus Vincere"
			},
			"team_b": {
				"id": 8297,
				"name": "FURIA"
			},
			"format": "bo3",
			"event": "Gamers8 2023",
			"live": false,
			"date": "2023-08-16T15:00:00-03:00"
		},
		{
			"id": 2366151,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-16T15:30:00-03:00"
		},
		{
			"id": 2366059,
			"team_a": {
				"id": 12376,
				"name": "M80"
			},
			"team_b": {
				"id": 9799,
				"name": "Bad News Bears"
			},
			"format": "bo3",
			"event": "ESL Challenger League Season 46 North America",
			"live": false,
			"date": "2023-08-16T21:00:00-03:00"
		},
		{
			"id": 2366145,
			"team_a": {
				"id": 11712,
				"name": "Sashi"
			},
			"team_b": {
				"id": 10905,
				"name": "PGE Turow"
			},
			"format": "bo3",
			"event": "European Pro League 2nd Division Season 5",
			"live": false,
			"date": "2023-08-17T05:00:00-03:00"
		},
		{
			"id": 2366156,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T05:30:00-03:00"
		},
		{
			"id": 2366160,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T05:30:00-03:00"
		},
		{
			"id": 2365863,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T07:30:00-03:00"
		},
		{
			"id": 2366148,
			"team_a": {
				"id": 6978,
				"name": "Singularity"
			},
			"team_b": {
				"id": 8209,
				"name": "Pompa"
			},
			"format": "bo3",
			"event": "European Pro League 2nd Division Season 5",
			"live": false,
			"date": "2023-08-17T07:30:00-03:00"
		},
		{
			"id": 2365791,
			"team_a": {
				"id": 10278,
				"name": "9INE"
			},
			"team_b": {
				"id": 5995,
				"name": "G2"
			},
			"format": "bo3",
			"event": "Gamers8 2023",
			"live": false,
			"date": "2023-08-17T08:00:00-03:00"
		},
		{
			"id": 2366157,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T08:30:00-03:00"
		},
		{
			"id": 2366161,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T08:30:00-03:00"
		},
		{
			"id": 2365864,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T10:30:00-03:00"
		},
		{
			"id": 2365792,
			"team_a": {
				"id": 6667,
				"name": "FaZe"
			},
			"team_b": {
				"id": 5378,
				"name": "Virtus.pro"
			},
			"format": "bo3",
			"event": "Gamers8 2023",
			"live": false,
			"date": "2023-08-17T11:30:00-03:00"
		},
		{
			"id": 2366158,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T11:30:00-03:00"
		},
		{
			"id": 2366162,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T11:30:00-03:00"
		},
		{
			"id": 2365865,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T13:30:00-03:00"
		},
		{
			"id": 2366159,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T14:30:00-03:00"
		},
		{
			"id": 2366163,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T14:30:00-03:00"
		},
		{
			"id": 2365793,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-17T15:00:00-03:00"
		},
		{
			"id": 2365930,
			"team_a": {
				"id": 12144,
				"name": "BESTIA"
			},
			"team_b": {
				"id": 11249,
				"name": "Meta"
			},
			"format": "bo3",
			"event": "CBCS Masters 2023",
			"live": false,
			"date": "2023-08-17T16:10:00-03:00"
		},
		{
			"id": 2365931,
			"team_a": {
				"id": 10669,
				"name": "Paquetá"
			},
			"team_b": {
				"id": 4773,
				"name": "paiN"
			},
			"format": "bo3",
			"event": "CBCS Masters 2023",
			"live": false,
			"date": "2023-08-17T16:10:00-03:00"
		},
		{
			"id": 2365932,
			"team_a": {
				"id": 11768,
				"name": "ODDIK"
			},
			"team_b": {
				"id": 11837,
				"name": "Fluxo"
			},
			"format": "bo3",
			"event": "CBCS Masters 2023",
			"live": false,
			"date": "2023-08-17T19:40:00-03:00"
		},
		{
			"id": 2365933,
			"team_a": {
				"id": 8113,
				"name": "Sharks"
			},
			"team_b": {
				"id": 10994,
				"name": "O PLANO"
			},
			"format": "bo3",
			"event": "CBCS Masters 2023",
			"live": false,
			"date": "2023-08-17T19:40:00-03:00"
		},
		{
			"id": 2366164,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-18T05:30:00-03:00"
		},
		{
			"id": 2365794,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-18T08:00:00-03:00"
		},
		{
			"id": 2366165,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-18T08:30:00-03:00"
		},
		{
			"id": 2365795,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-18T11:30:00-03:00"
		},
		{
			"id": 2366166,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-18T11:30:00-03:00"
		},
		{
			"id": 2366167,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-18T14:30:00-03:00"
		},
		{
			"id": 2365796,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-18T15:00:00-03:00"
		},
		{
			"id": 2366168,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-19T05:30:00-03:00"
		},
		{
			"id": 2366169,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-19T08:30:00-03:00"
		},
		{
			"id": 2365797,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-19T11:30:00-03:00"
		},
		{
			"id": 2366170,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-19T11:30:00-03:00"
		},
		{
			"id": 2366171,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-19T14:30:00-03:00"
		},
		{
			"id": 2365798,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-19T15:00:00-03:00"
		},
		{
			"id": 2366172,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-20T07:00:00-03:00"
		},
		{
			"id": 2366173,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-20T10:00:00-03:00"
		},
		{
			"id": 2366174,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-20T13:00:00-03:00"
		},
		{
			"id": 2365799,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-20T14:00:00-03:00"
		},
		{
			"id": 2366175,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-21T12:00:00-03:00"
		},
		{
			"id": 2366176,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-21T15:30:00-03:00"
		},
		{
			"id": 2366177,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-22T12:00:00-03:00"
		},
		{
			"id": 2366178,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-22T15:30:00-03:00"
		},
		{
			"id": 2366179,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-23T12:00:00-03:00"
		},
		{
			"id": 2366180,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-23T15:30:00-03:00"
		},
		{
			"id": 2366181,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-24T12:00:00-03:00"
		},
		{
			"id": 2366182,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-24T15:30:00-03:00"
		},
		{
			"id": 2366183,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-25T12:00:00-03:00"
		},
		{
			"id": 2366184,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-25T15:30:00-03:00"
		},
		{
			"id": 2366185,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-26T14:00:00-03:00"
		},
		{
			"id": 2366066,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-30T11:00:00-03:00"
		},
		{
			"id": 2366067,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-30T11:00:00-03:00"
		},
		{
			"id": 2366068,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-30T14:30:00-03:00"
		},
		{
			"id": 2366069,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-30T14:30:00-03:00"
		},
		{
			"id": 2366070,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-31T07:30:00-03:00"
		},
		{
			"id": 2366071,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-31T07:30:00-03:00"
		},
		{
			"id": 2366072,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-31T11:00:00-03:00"
		},
		{
			"id": 2366073,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-08-31T14:30:00-03:00"
		},
		{
			"id": 2366074,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-01T11:00:00-03:00"
		},
		{
			"id": 2366075,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-01T14:30:00-03:00"
		},
		{
			"id": 2366076,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-02T07:30:00-03:00"
		},
		{
			"id": 2366077,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-02T11:00:00-03:00"
		},
		{
			"id": 2366078,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-02T14:30:00-03:00"
		},
		{
			"id": 2366079,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-03T07:30:00-03:00"
		},
		{
			"id": 2366080,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-03T11:00:00-03:00"
		},
		{
			"id": 2366081,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-03T14:30:00-03:00"
		},
		{
			"id": 2366082,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-06T11:00:00-03:00"
		},
		{
			"id": 2366083,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-06T11:00:00-03:00"
		},
		{
			"id": 2366084,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-06T14:30:00-03:00"
		},
		{
			"id": 2366085,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-06T14:30:00-03:00"
		},
		{
			"id": 2366086,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-07T07:30:00-03:00"
		},
		{
			"id": 2366087,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-07T11:00:00-03:00"
		},
		{
			"id": 2366088,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-07T14:30:00-03:00"
		},
		{
			"id": 2366089,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-07T14:30:00-03:00"
		},
		{
			"id": 2366090,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-08T11:00:00-03:00"
		},
		{
			"id": 2366091,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-08T14:30:00-03:00"
		},
		{
			"id": 2366092,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-09T07:30:00-03:00"
		},
		{
			"id": 2366093,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-09T11:00:00-03:00"
		},
		{
			"id": 2366094,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-09T14:30:00-03:00"
		},
		{
			"id": 2366095,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-10T07:30:00-03:00"
		},
		{
			"id": 2366096,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-10T11:00:00-03:00"
		},
		{
			"id": 2366097,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-10T14:30:00-03:00"
		},
		{
			"id": 2366098,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-13T11:00:00-03:00"
		},
		{
			"id": 2366099,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-13T11:00:00-03:00"
		},
		{
			"id": 2366100,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-13T14:30:00-03:00"
		},
		{
			"id": 2366101,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-13T14:30:00-03:00"
		},
		{
			"id": 2366102,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-14T07:30:00-03:00"
		},
		{
			"id": 2366103,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-14T11:00:00-03:00"
		},
		{
			"id": 2366104,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-14T14:30:00-03:00"
		},
		{
			"id": 2366105,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-14T14:30:00-03:00"
		},
		{
			"id": 2366106,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-15T11:00:00-03:00"
		},
		{
			"id": 2366107,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-15T14:30:00-03:00"
		},
		{
			"id": 2366108,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-16T07:30:00-03:00"
		},
		{
			"id": 2366109,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-16T11:00:00-03:00"
		},
		{
			"id": 2366110,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-16T14:30:00-03:00"
		},
		{
			"id": 2366111,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-17T07:30:00-03:00"
		},
		{
			"id": 2366112,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-17T11:00:00-03:00"
		},
		{
			"id": 2366113,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-17T14:30:00-03:00"
		},
		{
			"id": 2366114,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-20T11:00:00-03:00"
		},
		{
			"id": 2366115,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-20T11:00:00-03:00"
		},
		{
			"id": 2366116,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-20T14:30:00-03:00"
		},
		{
			"id": 2366117,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-20T14:30:00-03:00"
		},
		{
			"id": 2366118,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-21T07:30:00-03:00"
		},
		{
			"id": 2366119,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-21T11:00:00-03:00"
		},
		{
			"id": 2366120,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-21T14:30:00-03:00"
		},
		{
			"id": 2366121,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-21T14:30:00-03:00"
		},
		{
			"id": 2366122,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-22T11:00:00-03:00"
		},
		{
			"id": 2366123,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-22T14:30:00-03:00"
		},
		{
			"id": 2366124,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-23T07:30:00-03:00"
		},
		{
			"id": 2366125,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-23T11:00:00-03:00"
		},
		{
			"id": 2366126,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-23T14:30:00-03:00"
		},
		{
			"id": 2366127,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-24T07:30:00-03:00"
		},
		{
			"id": 2366128,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-24T11:00:00-03:00"
		},
		{
			"id": 2366129,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-24T14:30:00-03:00"
		},
		{
			"id": 2366130,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-26T10:30:00-03:00"
		},
		{
			"id": 2366131,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-26T10:30:00-03:00"
		},
		{
			"id": 2366132,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-26T14:00:00-03:00"
		},
		{
			"id": 2366133,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-26T14:00:00-03:00"
		},
		{
			"id": 2366134,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-27T10:30:00-03:00"
		},
		{
			"id": 2366135,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-27T10:30:00-03:00"
		},
		{
			"id": 2366136,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-27T14:00:00-03:00"
		},
		{
			"id": 2366137,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-27T14:00:00-03:00"
		},
		{
			"id": 2366138,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-28T10:30:00-03:00"
		},
		{
			"id": 2366139,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-28T14:00:00-03:00"
		},
		{
			"id": 2366140,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-29T10:30:00-03:00"
		},
		{
			"id": 2366141,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-29T14:00:00-03:00"
		},
		{
			"id": 2366142,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-30T10:30:00-03:00"
		},
		{
			"id": 2366143,
			"team_a": null,
			"team_b": null,
			"format": "bo3",
			"event": null,
			"live": false,
			"date": "2023-09-30T14:00:00-03:00"
		},
		{
			"id": 2366144,
			"team_a": null,
			"team_b": null,
			"format": "bo5",
			"event": null,
			"live": false,
			"date": "2023-10-01T11:00:00-03:00"
		}
	]
}
```

</p>
</details>

