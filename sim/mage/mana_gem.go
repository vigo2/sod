package mage

import (
	"time"

	"github.com/wowsims/sod/sim/core"
)

// Technically 4 different items
const ManaGemRanks = 4

var ManaGemItemId = [ManaGemRanks + 1]int32{0, 5514, 5513, 8007, 8008}
var ManaGemManaRestored = [ManaGemRanks + 1][]float64{{0}, {375, 425}, {550, 650}, {775, 925}, {1000, 1200}}
var ManaGemLevel = [ManaGemRanks + 1]int{0, 23, 38, 48, 58}

func (mage *Mage) registerManaGemCD() {
	mage.ManaGem = make([]*core.Spell, ManaGemRanks+1)
	cdTimer := mage.NewTimer()

	for rank := 1; rank <= ManaGemRanks; rank++ {
		config := mage.newManaGemCooldown(rank, cdTimer)

		if config.RequiredLevel <= int(mage.Level) {
			mage.ManaGem[rank] = mage.RegisterSpell(config)

			mage.AddMajorCooldown(core.MajorCooldown{
				Spell:    mage.ManaGem[rank],
				Priority: int32(ManaGemManaRestored[rank][0]),
				Type:     core.CooldownTypeMana,
				ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
					return character.CurrentManaPercent() < .6
				},
			})
		}
	}
}

func (mage *Mage) newManaGemCooldown(rank int, cdTimer *core.Timer) core.SpellConfig {
	itemID := ManaGemItemId[rank]
	manaRestoredLow := ManaGemManaRestored[rank][0]
	manaRestoredHigh := ManaGemManaRestored[rank][1]

	actionID := core.ActionID{ItemID: itemID}
	manaMetrics := mage.NewManaMetrics(actionID)

	var remainingManaGems int
	mage.RegisterResetEffect(func(sim *core.Simulation) {
		remainingManaGems = 1
	})

	return core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagAPL,

		RequiredLevel: ManaGemLevel[rank],

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: time.Minute * 2,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			mage.AddMana(sim, sim.Roll(manaRestoredLow, manaRestoredHigh), manaMetrics)

			remainingManaGems--
			if remainingManaGems == 0 {
				// Disable this cooldown since we're out of emeralds.
				mage.GetMajorCooldown(actionID).Disable()
			}
		},

		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return remainingManaGems != 0
		},
	}
}
