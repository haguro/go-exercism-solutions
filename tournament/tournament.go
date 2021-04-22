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

type byPointsAndTeams []*standing

// Len returns the number of elements in the collection.
func (s byPointsAndTeams) Len() int { return len(s) }

// Swap swaps the elements with indexes i and j.
func (s byPointsAndTeams) Swap(i int, j int) { s[i], s[j] = s[j], s[i] }

// Less reports whether the element with index i must sort before the element with index j.
func (s byPointsAndTeams) Less(i int, j int) bool {
	if s[i].points == s[j].points {
		return s[i].team < s[j].team
	}
	return s[i].points > s[j].points
}

// Tally generates a team standings table from a list of match outcomes
func Tally(input io.Reader, output io.Writer) error {
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
			return errors.New("malformed input")
		}

		team1, team2, outcome := tokens[0], tokens[1], tokens[2]
		if outcome != "win" && outcome != "loss" && outcome != "draw" {
			return errors.New("invalid match outcome")
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
	sort.Stable(byPointsAndTeams(standings))
	output.Write([]byte(fmt.Sprintf(rowTemplate, "Team", "MP", "W", "D", "L", "P")))
	for _, standing := range standings {
		output.Write([]byte(fmt.Sprintf(rowTemplate, standing.team, standing.played, standing.won, standing.drawn, standing.lost, standing.points)))
	}
	return nil
}
