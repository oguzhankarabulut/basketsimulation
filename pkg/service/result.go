package service

import (
	"basketsimulation/pkg/domain"
	"log"
)

type ResultService interface {
	Result() *MatchResponse
}

type resultService struct {
	matchRepository       MatchRepository
	matchPlayerRepository MatchPlayerRepository
}

func NewResultService(m MatchRepository, mp MatchPlayerRepository) *resultService {
	return &resultService{matchRepository: m, matchPlayerRepository: mp}
}

type SinglePlayerResponse struct {
	Id                   string
	TwoPointScore        int64
	TwoPointScoreCount   int64
	ThreePointScore      int64
	ThreePointScoreCount int64
	TotalScore           int64
	AssistCount          int64
}

func NewSinglePlayerResponse(
	id string,
	twoPointScore int64,
	twoPointScoreCount int64,
	threePointScore int64,
	threePointScoreCount int64,
	TotalScore int64,
	AssistCount int64,
) *SinglePlayerResponse {
	return &SinglePlayerResponse{
		Id:                   id,
		TwoPointScore:        twoPointScore,
		TwoPointScoreCount:   twoPointScoreCount,
		ThreePointScore:      threePointScore,
		ThreePointScoreCount: threePointScoreCount,
		TotalScore:           TotalScore,
		AssistCount:          AssistCount,
	}
}

type SingleTeamResponse struct {
	Id          string
	Score       int
	PlayerStats []*SinglePlayerResponse
}

func NewSingleTeamResponse(id string, score int, playerStats []*SinglePlayerResponse) *SingleTeamResponse {
	return &SingleTeamResponse{
		Id:          id,
		Score:       score,
		PlayerStats: playerStats,
	}
}

type MatchResponse struct {
	Id          string
	AttackCount int
	TeamOne     *SingleTeamResponse
	TeamTwo     *SingleTeamResponse
}

func NewMatchResponse(id string, attackCount int, teamOne *SingleTeamResponse, teamTwo *SingleTeamResponse) *MatchResponse {
	return &MatchResponse{Id: id, AttackCount: attackCount, TeamOne: teamOne, TeamTwo: teamTwo}
}

func (s resultService) Result() *MatchResponse {
	match, err := s.matchRepository.Latest()
	if err != nil {
		log.Println(err)
		return nil
	}

	teamOneResponse := NewSingleTeamResponse(
		match.TeamOne.Id,
		match.TeamOneScore,
		s.teamPlayersStats(match.Id, match.TeamOne),
	)
	teamTwoResponse := NewSingleTeamResponse(
		match.TeamTwo.Id,
		match.TeamTwoScore,
		s.teamPlayersStats(match.Id, match.TeamTwo),
	)

	return NewMatchResponse(match.Id, match.AttackCount, teamOneResponse, teamTwoResponse)

}

func (s resultService) teamPlayersStats(matchId string, t *domain.Team) []*SinglePlayerResponse {
	teamPlayers := make([]*SinglePlayerResponse, 5)
	for i, p := range t.Players {
		twoPointScoreCount, err := s.matchPlayerRepository.ScoreCount(matchId, p.Id, 2)
		if err != nil {
			log.Println(err)
		}
		twoPointScore := twoPointScoreCount * 2

		threePointScoreCount, err := s.matchPlayerRepository.ScoreCount(matchId, p.Id, 3)
		if err != nil {
			log.Println(err)
		}
		threePointScore := threePointScoreCount * 3

		assistCount, err := s.matchPlayerRepository.AssistCount(matchId, p.Id)
		if err != nil {
			log.Println(err)
		}
		teamPlayers[i] = NewSinglePlayerResponse(
			p.Id,
			twoPointScore,
			twoPointScoreCount,
			threePointScore,
			threePointScoreCount,
			twoPointScore+threePointScore,
			assistCount,
		)
	}

	return teamPlayers
}
