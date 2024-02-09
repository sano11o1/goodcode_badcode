package main

func main() {
	m := Member{Level: 10}
	magicMap := map[MagicType]Magic{
		MagicTypeFire:     Fire{Member: m},
		MagicTypeHellFire: HellFire{Member: m},
	}
	game := NewGame(m, magicMap)

	game.magicAttack(MagicTypeFire)

	game.magicAttack(MagicTypeHellFire)
}

type Game struct {
	Member Member
	Magics map[MagicType]Magic
}

func NewGame(member Member, magics map[MagicType]Magic) Game {
	return Game{
		Member: member,
		Magics: magics,
	}
}

type Member struct {
	Level int
}

type Magic interface {
	/*
		TODO:戻り値をオブジェクトに書き換える
	*/
	Name() string
	ConstMagicPoint() int
	AttackPower() int
	ConstTechnicalPoint() int
}

type Fire struct {
	Member Member
}

func NewMember(member Member) Fire {
	return Fire{Member: member}
}

func (f Fire) Name() string {
	return "ファイア"
}

func (f Fire) ConstMagicPoint() int {
	return 2
}

func (f Fire) AttackPower() int {
	return 20 + int(f.Member.Level*3)
}

func (f Fire) ConstTechnicalPoint() int {
	return 0
}

type HellFire struct {
	Member Member
}

func NewHellFire(member Member) HellFire {
	return HellFire{Member: member}
}

func (f HellFire) Name() string {
	return "地獄の業火"
}

func (f HellFire) ConstMagicPoint() int {
	return 16
}

func (f HellFire) AttackPower() int {
	return 160 + int(f.Member.Level*5)
}

func (f HellFire) ConstTechnicalPoint() int {
	return 16000000
}

type MagicType string

const MagicTypeFire MagicType = "fire"
const MagicTypeHellFire MagicType = "hellfire"

var magicMap = map[MagicType]Magic{
	MagicTypeFire:     Fire{},
	MagicTypeHellFire: HellFire{},
}

// 魔法攻撃を実行する
func (g Game) magicAttack(mType MagicType) {
	magic := g.Magics[mType]
	g.showMagicName(magic)
	g.consumeMagicPoint(magic)
	g.consumeTechnicalPoint(magic)
}

// 魔法名を表示する
func (g Game) showMagicName(magic Magic) {
	println(magic.Name())
}

// 魔法力を消費する
func (g Game) consumeMagicPoint(magic Magic) int {
	return magic.ConstMagicPoint()
}

// テクニカルポイントを消費する
func (g Game) consumeTechnicalPoint(magic Magic) int {
	return magic.ConstTechnicalPoint()
}
