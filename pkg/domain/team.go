package domain

type Team struct {
	Id string
	Players []*Player
}

func NewTeam (players []*Player) *Team	{
	id, _ := generateId()
	return &Team{
		Id: id,
		Players: players,
	}
}

func (t Team) AttackRate() int {
	pp := t.Players
	ar := 0
	for _, p := range pp {
		ar += p.AttackCoefficient
	}
	return ar
}

func (t Team) DefenceRate() int {
	pp := t.Players
	dr := 0
	for _, p := range pp {
		dr += p.DefenceCoefficient
	}
	return dr
}

func (t Team) Scorer () *Player {
	pp := t.Players

	var chance []int
	for i, p := range pp {
		for j := 0; j <= p.AttackCoefficient; j++ {
			chance = append(chance, i)
		}
	}
	pi := chance[randomInt(0, len(chance)-1)]
	return t.Players[pi]
}

func (t Team) Assister (scorer *Player) *Player {
	pp := t.Players

	var chance []int
	for i, p := range pp {
		if p != scorer {
			for j := 0; j <= p.AttackCoefficient; j++ {
				chance = append(chance, i)
			}
		}
	}
	pi := chance[randomInt(0, len(chance)-1)]
	return t.Players[pi]
}