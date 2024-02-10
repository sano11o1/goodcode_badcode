package main

func main() {
	mem := Member{HitPoint: 100, MagicPoint: 100}
	magicPointDamage := NewMagicPointDamage(mem)
	hitPointDamage := NewHitPointDamage(mem)
	var damageMap = map[string]Damage{
		"magicPointDamage": magicPointDamage,
		"hitPointDamage":   hitPointDamage,
	}
	game := Game{damageMap: damageMap}
	game.ApplyDamage("magicPointDamage", 10)
	game.ApplyDamage("hitPointDamage", 20)
}

type Game struct {
	damageMap map[string]Damage
}

func (g Game) ApplyDamage(damageType string, damageAmount int) {
	g.damageMap[damageType].execute(damageAmount)
}

type Member struct {
	HitPoint   int
	MagicPoint int
}

type Damage interface {
	execute(damageAmount int)
}

type MagicPointDamage struct {
	Member *Member
}

func NewMagicPointDamage(member Member) MagicPointDamage {
	return MagicPointDamage{Member: &member}
}

type HitPointDamage struct {
	Member *Member
}

func (m MagicPointDamage) execute(damageAmount int) {
	m.Member.MagicPoint -= damageAmount
}

func NewHitPointDamage(member Member) HitPointDamage {
	return HitPointDamage{Member: &member}
}

func (h HitPointDamage) execute(damageAmount int) {
	h.Member.HitPoint -= damageAmount
}
