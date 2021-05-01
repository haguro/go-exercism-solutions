package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const rowTemplate = "%-31s|%3v |%3v |%3v |%3v |%3v\n"

type standing struct {
	team   string
	played int
	won    int
	drawn  int
	lost   int
	points int
}

// Tally generates a team standings table from a list of match outcomes
func Tally(input io.Reader, output io.Writer) error {
	standings, err := parseMatchOutcomes(input)
	if err != nil {
		return err
	}

	sort.SliceStable(standings, func(i, j int) bool {
		if standings[i].points == standings[j].points {
			return standings[i].team < standings[j].team
		}
		return standings[i].points > standings[j].points
	})

	output.Write([]byte(fmt.Sprintf(rowTemplate, "Team", "MP", "W", "D", "L", "P")))
	for _, standing := range standings {
		output.Write([]byte(fmt.Sprintf(rowTemplate, standing.team, standing.played, standing.won, standing.drawn, standing.lost, standing.points)))
	}
	return nil
}

func parseMatchOutcomes(input io.Reader) ([]*standing, error) {
	var tokens []string
	var standings []*standing

	standingsMap := make(map[string]*standing)
	s := bufio.NewScanner(input)
	for s.Scan() {
		line := s.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		tokens = strings.Split(line, ";")
		if len(tokens) < 3 {
			return nil, errors.New("malformed input")
		}

		team1, team2, outcome := tokens[0], tokens[1], tokens[2]
		if outcome != "win" && outcome != "loss" && outcome != "draw" {
			return nil, errors.New("invalid match outcome")
		}

		if _, ok := standingsMap[team1]; !ok {
			r := standing{team: team1}
			standings = append(standings, &r)
			standingsMap[team1] = &r
		}
		if _, ok := standingsMap[team2]; !ok {
			r := standing{team: team2}
			standings = append(standings, &r)
			standingsMap[team2] = &r
		}

		standingsMap[team1].played++
		standingsMap[team2].played++
		switch outcome {
		case "win":
			standingsMap[team1].won++
			standingsMap[team1].points += 3
			standingsMap[team2].lost++
		case "loss":
			standingsMap[team1].lost++
			standingsMap[team2].won++
			standingsMap[team2].points += 3
		case "draw":
			standingsMap[team1].drawn++
			standingsMap[team1].points++
			standingsMap[team2].drawn++
			standingsMap[team2].points++
		}
	}
	return standings, nil
}
