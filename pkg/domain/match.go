package domain

type Match struct {
	Id string
	TeamOne *Team
	TeamTwo *Team
	TeamOneScore int
	TeamTwoScore int
	AttackCount int
}

func NewMatch(t1 *Team, t2 *Team) *Match {
	id, _ := generateId()
	return &Match{
		Id: id,
		TeamOne: t1,
		TeamTwo: t2,
	}
}

func (m *Match) UpdateTeamOneScore(s int) {
	m.TeamOneScore = m.TeamOneScore + s
}

func (m *Match) UpdateTeamTwoScore(s int) {
	m.TeamTwoScore = m.TeamTwoScore + s
}

func (m *Match) UpdateAttackCount() {
	m.AttackCount = m.AttackCount + 1
}