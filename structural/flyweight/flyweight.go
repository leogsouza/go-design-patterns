package flyweight

import (
	"time"
)

// Team types
const (
	TeamA = iota
	TeamB
)

// Team struct holds all informations about a team
type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

// Player represents a player of a team
type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

// HistoricalData represents all data related to the championship
type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

// Match contains informations about the match played between two teams
type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

// NewTeamFactory return a new instance of teamFlyweightFactory
func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}

func getTeamFactory(team int) Team {
	switch team {
	case TeamB:
		return Team{
			ID:   2,
			Name: "Team B",
		}
	default:
		return Team{
			ID:   1,
			Name: "Team A",
		}
	}
}

func (t *teamFlyweightFactory) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}

	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
