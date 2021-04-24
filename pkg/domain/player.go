package domain

type Player struct {
	Id string
	Role string
	AttackCoefficient int
	DefenceCoefficient int
	TwoPointCoefficient int
	ThreePointCoefficient int
}

func NewPlayer(role string, attackCoefficient int, defenceCoefficient int, twoPointCoefficient int, threePointCoefficient int) *Player {
	id, _ := generateId()
	return &Player{
		Id: id,
		Role: role,
		AttackCoefficient: attackCoefficient,
		DefenceCoefficient: defenceCoefficient,
		TwoPointCoefficient: twoPointCoefficient,
		ThreePointCoefficient: threePointCoefficient,
	}
}

func (p Player) Score() int {
	twoPointRate := 100 / (p.TwoPointCoefficient + p.ThreePointCoefficient) * p.TwoPointCoefficient
	if randomInt(0, 100) <= twoPointRate {
		return 2
	}
	return 3
}
