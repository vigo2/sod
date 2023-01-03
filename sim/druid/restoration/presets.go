package restoration

import (
	"github.com/wowsims/wotlk/sim/core"
	"github.com/wowsims/wotlk/sim/core/proto"
)

var CelestialFocusTalents = &proto.DruidTalents{
	Genesis:         5,
	Moonglow:        3,
	NaturesMajesty:  2,
	NaturesGrace:    3,
	NaturesSplendor: true,
	NaturesReach:    1,
	CelestialFocus:  3,

	ImprovedMarkOfTheWild: 2,
	NaturesFocus:          3,
	Subtlety:              2,
	NaturalShapeshifter:   3,
	Intensity:             3,
	OmenOfClarity:         true,
	MasterShapeshifter:    2,
	TranquilSpirit:        1,
	ImprovedRejuvenation:  3,
	NaturesSwiftness:      true,
	GiftOfNature:          5,
	EmpoweredTouch:        2,
	NaturesBounty:         3,
	LivingSpirit:          3,
	Swiftmend:             true,
	EmpoweredRejuvenation: 5,
	Revitalize:            3,
	TreeOfLife:            true,
	ImprovedTreeOfLife:    3,
	GiftOfTheEarthmother:  5,
	WildGrowth:            true,
}

var StandardGlyphs = &proto.Glyphs{
	Major1: int32(proto.DruidMajorGlyph_GlyphOfWildGrowth),
	Major2: int32(proto.DruidMajorGlyph_GlyphOfSwiftmend),
	Major3: int32(proto.DruidMajorGlyph_GlyphOfNourish),
}

var FullConsumes = &proto.Consumes{
	Flask:           proto.Flask_FlaskOfBlindingLight,
	Food:            proto.Food_FoodBlackenedBasilisk,
	DefaultPotion:   proto.Potions_SuperManaPotion,
	PrepopPotion:    proto.Potions_DestructionPotion,
	DefaultConjured: proto.Conjured_ConjuredDarkRune,
}

var PlayerOptionsStandard = &proto.Player_RestorationDruid{
	RestorationDruid: &proto.RestorationDruid{
		Talents: CelestialFocusTalents,
		Options: &proto.RestorationDruid_Options{
			InnervateTarget: &proto.RaidTarget{TargetIndex: 0}, // self innervate
		},
		Rotation: &proto.RestorationDruid_Rotation{},
	},
}

var P1Gear = core.EquipmentSpecFromJsonString(`{"items": [
	{
		"id": 44007,
		"enchant": 3819,
		"gems": [
			41401,
			40017
		]
	},
	{
		"id": 40071
	},
	{
		"id": 39719,
		"enchant": 3809,
		"gems": [
			39998
		]
	},
	{
		"id": 40723,
		"enchant": 3859
	},
	{
		"id": 44002,
		"enchant": 3832,
		"gems": [
			39998,
			40026
		]
	},
	{
		"id": 44008,
		"enchant": 2332,
		"gems": [
			39998,
			0
		]
	},
	{
		"id": 40460,
		"enchant": 3246,
		"gems": [
			40017,
			0
		]
	},
	{
		"id": 40561,
		"enchant": 3601,
		"gems": [
			39998
		]
	},
	{
		"id": 40379,
		"enchant": 3719,
		"gems": [
			39998,
			40017
		]
	},
	{
		"id": 40558,
		"enchant": 3606
	},
	{
		"id": 40719
	},
	{
		"id": 40375
	},
	{
		"id": 37111
	},
	{
		"id": 40432
	},
	{
		"id": 40395,
		"enchant": 3834
	},
	{
		"id": 39766
	},
	{
		"id": 40342
	}
]}`)