package domain

type MatchPlayer struct {
	Id string
	PlayerId string
	TeamId string
	MatchId string
	Assist bool
	Score int
}

func NewMatchPlayer(teamId string, matchId string, playerId string, assist bool, score int) *MatchPlayer {
	id, _ := generateId()
	return &MatchPlayer{
		Id: id,
		PlayerId: playerId,
		TeamId: teamId,
		MatchId: matchId,
		Assist: assist,
		Score: score,
	}
}