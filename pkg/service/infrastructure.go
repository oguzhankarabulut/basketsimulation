package service

import "basketsimulation/pkg/domain"

type MatchRepository interface {
	Save (m *domain.Match) error
	Latest() (*domain.Match, error)
}

type MatchPlayerRepository interface {
	Save(mp *domain.MatchPlayer) error
	ScoreCount(matchId string, playerId string, score int) (int64, error)
	AssistCount(matchId string, playerId string) (int64, error)
}
