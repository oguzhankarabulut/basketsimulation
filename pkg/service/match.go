package service

import (
	"basketsimulation/pkg/domain"
	"log"
	_time "time"
)

type matchService struct {
	matchRepository MatchRepository
	matchPlayerRepository MatchPlayerRepository
}

func NewMatchService(m MatchRepository, mp MatchPlayerRepository) *matchService {
	return &matchService{matchRepository: m, matchPlayerRepository: mp}
}


func setTeams() (*domain.Team, *domain.Team) {
	//TODO: Get players from mongo
	var teamOnePlayers []*domain.Player
	teamOneGuard := domain.NewPlayer("pointGuard", 4, 3, 13, 5)
	teamOnePointGuard := domain.NewPlayer("shootingGuard", 7, 3, 8, 5)
	teamOneSmallForward := domain.NewPlayer("smallForward", 6, 5, 8, 2)
	teamOnePowerForward := domain.NewPlayer("powerForward", 3, 8, 13, 1)
	teamOneCenter := domain.NewPlayer("center", 3, 5, 8, 3)
	teamOnePlayers = append(teamOnePlayers, teamOneGuard, teamOnePointGuard, teamOneSmallForward, teamOnePowerForward, teamOneCenter)
	teamOne := domain.NewTeam(teamOnePlayers)

	var teamTwoPlayers []*domain.Player
	teamTwoGuard := domain.NewPlayer("pointGuard", 6, 4, 8, 3)
	teamTwoPointGuard := domain.NewPlayer("shootingGuard", 9, 5, 8, 5)
	teamTwoSmallForward := domain.NewPlayer("smallForward", 7, 6, 8, 2)
	teamTwoPowerForward := domain.NewPlayer("powerForward", 4, 9, 13, 1)
	teamTwoCenter := domain.NewPlayer("center", 5, 5, 9, 3)
	teamTwoPlayers = append(teamTwoPlayers, teamTwoGuard, teamTwoPointGuard, teamTwoSmallForward, teamTwoPowerForward, teamTwoCenter)
	teamTwo := domain.NewTeam(teamTwoPlayers)

	return teamOne, teamTwo
}

func (s matchService) Start() {
	teamOne, teamTwo := setTeams()
	match := domain.NewMatch(teamOne, teamTwo)
	s.Match(match, teamOne, teamTwo, 240)
}

func (s matchService) Match(m *domain.Match, t1 *domain.Team, t2 *domain.Team, time float64) *domain.Match {
	m.UpdateAttackCount()
	t1AttackPoint := t1.AttackRate()
	t2DefencePoint := t2.DefenceRate()
	t2DefenceRate := 100 / (t1AttackPoint + t2DefencePoint) * t2DefencePoint

	if t2DefenceRate <= randomInt(0, 100) {
		if m.TeamOne == t1 {
			scorer := m.TeamOne.Scorer()

			assister := m.TeamOne.Assister(scorer)
			if err := s.matchPlayerRepository.Save(domain.NewMatchPlayer(t1.Id, m.Id, assister.Id, true, 0)); err != nil {
				log.Println(err)
			}

			score := scorer.Score()
			m.UpdateTeamOneScore(score)
			if err := s.matchPlayerRepository.Save(domain.NewMatchPlayer(t1.Id, m.Id, scorer.Id, false, score)); err != nil {
				log.Println(err)
			}
		}
		if m.TeamTwo == t1 {
			scorer := m.TeamTwo.Scorer()

			assister := m.TeamTwo.Assister(scorer)
			if err := s.matchPlayerRepository.Save(domain.NewMatchPlayer(t1.Id, m.Id, assister.Id, true, 0)); err != nil {
				log.Println(err)
			}

			score := scorer.Score()
			m.UpdateTeamTwoScore(score)
			if err := s.matchPlayerRepository.Save(domain.NewMatchPlayer(t1.Id, m.Id, scorer.Id, false, score)); err != nil {
				log.Println(err)
			}
		}
	}

	if err := s.matchRepository.Save(m); err != nil {
		log.Println(err)
	}

	randomFloat :=  randomFloat()
	sleep := int(randomFloat * 100)
	_time.Sleep(_time.Duration(sleep) * _time.Millisecond * 10)
	time = time - randomFloat
	if time <= 0 {
		return m
	}

	return s.Match(m, t2, t1, time)
}