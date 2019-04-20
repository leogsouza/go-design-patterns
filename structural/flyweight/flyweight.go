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

func (t *teamFlyweightFactory) GetTeam(team int) *Team {
	return nil
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return 0
}
