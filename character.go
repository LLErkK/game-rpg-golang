package main

type CharacterType int

const (
	KnightType CharacterType = iota
	RogueType
	BerserkerType
)

type Character struct {
	Name           string
	Type           CharacterType
	Level          int
	EXP            int
	EXPToNextLevel int
	HP             int
	MaxHP          int
	ATK            int
	DEF            int
	SPD            int
	CRIT           float64
}

type CharacterScaling struct {
	HPGrowth   int
	ATKGrowth  int
	DEFGrowth  int
	SPDGrowth  int
	CRITGrowth float64
	EXPScaling float64
}

var characterScalings = map[CharacterType]CharacterScaling{
	KnightType: {
		HPGrowth:   15, // Knight fokus ke HP dan DEF
		ATKGrowth:  2,
		DEFGrowth:  3,
		SPDGrowth:  1,
		CRITGrowth: 0.01,
		EXPScaling: 1.4, // EXP requirement scaling lebih rendah
	},
	RogueType: {
		HPGrowth:   8, // Rogue fokus ke SPD dan CRIT
		ATKGrowth:  3,
		DEFGrowth:  1,
		SPDGrowth:  4,
		CRITGrowth: 0.03,
		EXPScaling: 1.5, // Standard scaling
	},
	BerserkerType: {
		HPGrowth:   10, // Berserker fokus ke ATK
		ATKGrowth:  4,
		DEFGrowth:  1,
		SPDGrowth:  2,
		CRITGrowth: 0.01,
		EXPScaling: 1.6, // EXP requirement scaling lebih tinggi
	},
}

func GetStarterCharacters() []Character {
	return []Character{
		{
			Name: "Knight", Type: KnightType, Level: 1,
			HP: 120, MaxHP: 120, ATK: 25, DEF: 20, SPD: 10, CRIT: 0.1,
			EXP: 0, EXPToNextLevel: 100,
		},
		{
			Name: "Rogue", Type: RogueType, Level: 1,
			HP: 90, MaxHP: 90, ATK: 15, DEF: 10, SPD: 30, CRIT: 0.2,
			EXP: 0, EXPToNextLevel: 100,
		},
		{
			Name: "Berserker", Type: BerserkerType, Level: 1,
			HP: 80, MaxHP: 80, ATK: 35, DEF: 5, SPD: 20, CRIT: 0.05,
			EXP: 0, EXPToNextLevel: 100,
		},
	}
}

func (c *Character) GainExp(amount int) {
	c.EXP += amount
	for c.EXP >= c.EXPToNextLevel {
		c.EXP -= c.EXPToNextLevel
		c.LevelUp()
	}
}
func (c *Character) LevelUp() {
	c.Level++
	scaling := characterScalings[c.Type]

	// Update EXP requirement dengan scaling berbeda per karakter
	c.EXPToNextLevel = int(float64(c.EXPToNextLevel) * scaling.EXPScaling)

	// Apply stat growth berdasarkan tipe karakter
	c.MaxHP += scaling.HPGrowth
	c.ATK += scaling.ATKGrowth
	c.DEF += scaling.DEFGrowth
	c.SPD += scaling.SPDGrowth
	c.CRIT += scaling.CRITGrowth

	// Heal to full HP on level up
	c.HP = c.MaxHP

	println(c.Name + " leveled up to level " + itoa(c.Level) + "!")
}
