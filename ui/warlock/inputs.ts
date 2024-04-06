import * as InputHelpers from '../core/components/input_helpers.js';
import { Player } from '../core/player.js';
import { ItemSlot, Spec } from '../core/proto/common.js';
import {
	WarlockOptions_Armor as Armor,
	WarlockOptions_MaxFireboltRank as MaxFireboltRank,
	WarlockOptions_Summon as Summon,
	WarlockOptions_WeaponImbue as WeaponImbue,	WarlockRune} from '../core/proto/warlock.js';
import { ActionId } from '../core/proto_utils/action_id.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const ArmorInput = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecWarlock, Armor>({
	fieldName: 'armor',
	values: [
		{ value: Armor.NoArmor, tooltip: 'No Armor' },
		{ actionId: player => player.getMatchingSpellActionId([
			{ id: 706, minLevel: 20, maxLevel: 39 },
			{ id: 11733, minLevel: 40, maxLevel: 49 },
			{ id: 11734, minLevel: 50, maxLevel: 59 },
			{ id: 11735, minLevel: 60 },
		]), value: Armor.DemonArmor },
	],
});

export const WeaponImbueInput = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecWarlock, WeaponImbue>({
	fieldName: 'weaponImbue',
	values: [
		{ value: WeaponImbue.NoWeaponImbue, tooltip: 'No Weapon Stone' },
		{ actionId: player => player.getMatchingItemActionId([
			{ id: 1254, minLevel: 28, maxLevel: 35 },
			{ id: 13699, minLevel: 36, maxLevel: 45 },
			{ id: 13700, minLevel: 46, maxLevel: 55 },
			{ id: 13701, minLevel: 56 },
		]), value: WeaponImbue.Firestone },
		{ actionId: player => player.getMatchingItemActionId([
			{ id: 5522, minLevel: 36, maxLevel: 47 },
			{ id: 13602, minLevel: 48, maxLevel: 59 },
			{ id: 13603, minLevel: 60 },
		]), value: WeaponImbue.Spellstone },
	],
	showWhen: player => player.getEquippedItem(ItemSlot.ItemSlotOffHand) == null && player.getLevel() >= 28,
	changeEmitter: (player: Player<Spec.SpecWarlock>) => player.changeEmitter,
});

export const PetInput = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecWarlock, Summon>({
	fieldName: 'summon',
	values: [
		{ value: Summon.NoSummon, tooltip: 'No Pet' },
		{ actionId: () => ActionId.fromSpellId(688), value: Summon.Imp },
		{ actionId: () => ActionId.fromSpellId(697), value: Summon.Voidwalker },
		{ actionId: () => ActionId.fromSpellId(712), value: Summon.Succubus },
		{ actionId: () => ActionId.fromSpellId(691), value: Summon.Felhunter },
		{ actionId: () => ActionId.fromSpellId(427733), value: Summon.Felguard, showWhen: player => player.getEquippedItem(ItemSlot.ItemSlotWrist)?.rune?.id == WarlockRune.RuneBracerSummonFelguard },
	],
	changeEmitter: (player: Player<Spec.SpecWarlock>) => player.changeEmitter,
});

export const ImpFireboltRank = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecWarlock, MaxFireboltRank>({
	fieldName: 'maxFireboltRank',
	showWhen: player => player.getSpecOptions().summon == Summon.Imp,
	values: [
		{ value: MaxFireboltRank.NoMaximum, tooltip: 'Max' },
		{ actionId: () => ActionId.fromSpellId(3110), value: MaxFireboltRank.Rank1 },
		{ actionId: () => ActionId.fromSpellId(7799), value: MaxFireboltRank.Rank2 },
		{ actionId: () => ActionId.fromSpellId(7800), value: MaxFireboltRank.Rank3 },
		{ actionId: () => ActionId.fromSpellId(7801), value: MaxFireboltRank.Rank4 },
		{ actionId: () => ActionId.fromSpellId(7802), value: MaxFireboltRank.Rank5 },
		{ actionId: () => ActionId.fromSpellId(11762), value: MaxFireboltRank.Rank6 },
		{ actionId: () => ActionId.fromSpellId(11763), value: MaxFireboltRank.Rank7 },
	],
	changeEmitter: (player: Player<Spec.SpecWarlock>) => player.changeEmitter,
});

export const PetPoolManaInput = InputHelpers.makeSpecOptionsBooleanInput<Spec.SpecWarlock>({
	fieldName: 'petPoolMana',
	label: "No Pet Management",
	labelTooltip: 'Should Pet keep trying to cast on every mana regen instead of waiting for mana',
	changeEmitter: (player: Player<Spec.SpecWarlock>) => player.changeEmitter,
});
