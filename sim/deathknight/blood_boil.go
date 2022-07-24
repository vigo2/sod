package deathknight

import (
	"github.com/wowsims/wotlk/sim/core"
)

func (deathKnight *Deathknight) registerBloodBoilSpell() {
	deathKnight.BloodBoil = deathKnight.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 49941},
		SpellSchool: core.SpellSchoolShadow,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				cast.GCD = deathKnight.getModifiedGCD()
			},
		},

		ApplyEffects: core.ApplyEffectFuncAOEDamage(deathKnight.Env, core.SpellEffect{
			ProcMask:             core.ProcMaskSpellDamage,
			BonusSpellCritRating: 0.0,
			DamageMultiplier:     1.0,
			ThreatMultiplier:     1.0,

			BaseDamage: core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
					roll := (220.0-180.0)*sim.RandomFloat("Blood Boil") + 180.0
					return (roll + deathKnight.applyImpurity(hitEffect, spell.Unit)*0.06) *
						deathKnight.rageOfRivendareBonus(hitEffect.Target) *
						deathKnight.tundraStalkerBonus(hitEffect.Target)
				},
				TargetSpellCoefficient: 1,
			},
			OutcomeApplier: deathKnight.OutcomeFuncMagicHitAndCrit(deathKnight.spellCritMultiplierGoGandMoM()),

			OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Target == deathKnight.CurrentTarget {
					deathKnight.LastCastOutcome = spellEffect.Outcome
				}
				if spellEffect.Landed() && spellEffect.Target == deathKnight.CurrentTarget {
					dkSpellCost := deathKnight.DetermineOptimalCost(sim, 1, 0, 0)
					deathKnight.Spend(sim, spell, dkSpellCost)

					amountOfRunicPower := 10.0 + 2.5*float64(deathKnight.Talents.ChillOfTheGrave)
					deathKnight.AddRunicPower(sim, amountOfRunicPower, spell.RunicPowerMetrics())
				}
			},
		}),
	})
}

func (deathKnight *Deathknight) CanBloodBoil(sim *core.Simulation) bool {
	return deathKnight.CastCostPossible(sim, 0.0, 1, 0, 0) && deathKnight.BloodBoil.IsReady(sim)
}

func (deathKnight *Deathknight) CastBloodBoil(sim *core.Simulation, target *core.Unit) bool {
	if deathKnight.CanBloodBoil(sim) {
		deathKnight.BloodBoil.Cast(sim, target)
		return true
	}
	return false
}
