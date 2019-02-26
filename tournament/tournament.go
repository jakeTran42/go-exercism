package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type standings struct {
	name                                    string
	mapsPlayed, wins, draws, losses, points int
}

// Future refactor?
// var teamStandings map[string]standings
//
// func (standings *teamStandings) countResult (team1, team2, result string) (*teamStandings, error) {
//
// }

// Tally is a function that inputs a reader with match results and
// outputs a formatted table to writer.
func Tally(reader io.Reader, writer io.Writer) error {
	teamStandings := make(map[string]standings)

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reader)

	for {
		byteLine, err := buffer.ReadBytes('\n')

		stringLine := string(byteLine)

		if stringLine == "" {
			break
		}

		if stringLine == "\n" || string(stringLine[0]) == "#" {
			if err != nil {
				break
			}
			continue
		}
		split := strings.Split(stringLine, ";")

		if len(split) != 3 {
			return errors.New("Results incorrectly formatted")
		}

		team1 := split[0]
		team2 := split[1]
		result := strings.Trim(split[2], "\n")

		switch {
		case result == "win":
			if standing, hasStanding := teamStandings[team1]; hasStanding {
				standing.mapsPlayed++
				standing.wins++
				standing.points += 3
				teamStandings[team1] = standing
			} else {
				teamStandings[team1] = standings{team1, 1, 1, 0, 0, 3}
			}

			if standing, hasStanding := teamStandings[team2]; hasStanding {
				standing.mapsPlayed++
				standing.losses++
				teamStandings[team2] = standing
			} else {
				teamStandings[team2] = standings{team2, 1, 0, 0, 1, 0}
			}
		case result == "loss":
			if standing, hasStanding := teamStandings[team1]; hasStanding {
				standing.mapsPlayed++
				standing.losses++
				teamStandings[team1] = standing
			} else {
				teamStandings[team1] = standings{team1, 1, 0, 0, 1, 0}
			}

			if standing, hasStanding := teamStandings[team2]; hasStanding {
				standing.mapsPlayed++
				standing.wins++
				standing.points += 3
				teamStandings[team2] = standing
			} else {
				teamStandings[team2] = standings{team2, 1, 1, 0, 0, 3}
			}
		case result == "draw":
			if standing, hasStanding := teamStandings[team1]; hasStanding {
				standing.mapsPlayed++
				standing.draws++
				standing.points++
				teamStandings[team1] = standing
			} else {
				teamStandings[team1] = standings{team1, 1, 0, 1, 0, 1}
			}

			if standing, hasStanding := teamStandings[team2]; hasStanding {
				standing.mapsPlayed++
				standing.draws++
				standing.points++
				teamStandings[team2] = standing
			} else {
				teamStandings[team2] = standings{team2, 1, 0, 1, 0, 1}
			}
		default:
			return errors.New("Invalid match result")
		}

		if err != nil {
			break
		}
	}

	var teams []standings
	for _, standing := range teamStandings {
		teams = append(teams, standing)
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points > teams[j].points {
			return true
		}
		if teams[i].points < teams[j].points {
			return false
		}
		return teams[i].name < teams[j].name
	})

	writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, standing := range teams {
		standingString := fmt.Sprintf("% -30s", standing.name)
		standingString += fmt.Sprintf(" |% 3d", standing.mapsPlayed)
		standingString += fmt.Sprintf(" |% 3d", standing.wins)
		standingString += fmt.Sprintf(" |% 3d", standing.draws)
		standingString += fmt.Sprintf(" |% 3d", standing.losses)
		standingString += fmt.Sprintf(" |% 3d\n", standing.points)

		writer.Write([]byte(standingString))
	}

	return nil
}
