package main

func main() {
}

type Member struct {
	HitPoint   int
	MagicPoint int
}

/*
 * damageFlagで切り替えてしまっている！！
 * damageFlagという命名も微妙。isMagicAttackとかならまだなし。インタフェースに切り出すのがベスト。
 */

type Damage struct {
	Member *Member
}

func NewDamage(member Member) Damage {
	return Damage{Member: &member}
}

func (m *Damage) Damage(damageFlag bool, damageAmount int) {
	if damageFlag {
		m.Member.HitPoint -= damageAmount
	} else {
		m.Member.MagicPoint -= damageAmount
	}
}
