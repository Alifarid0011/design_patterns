package main

import "time"

const (
	TEAM_A = iota
	TEAM_B
	TEAM_C
)

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}
type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}
type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}
type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}
type TeamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func (t *TeamFlyweightFactory) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}
func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}
func (t *TeamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
func NewTeamFactory() TeamFlyweightFactory {
	return TeamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}
