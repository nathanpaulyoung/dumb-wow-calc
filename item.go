package main

import "fmt"

type item struct {
	ID                int
	Name              string
	Source            string
	Rarity            string
	ItemType          string
	SubType           string
	ItemLevel         int
	TopEnd            float32
	DPS               float32
	Stamina           int
	Strength          int
	Agility           int
	Intellect         int
	Spirit            int
	Mp5               int
	SpellDamage       int
	SpellHealing      int
	MeleeAttackPower  int
	RangedAttackPower int
	MeleeHit          int
	RangedHit         int
	SpellHit          int
	MeleeCrit         int
	RangedCrit        int
	SpellCrit         int
	Weights           weights
}

func newItem() *item {
	item := new(item)
	item.ID = 0
	item.Name = ""
	item.Source = ""
	item.Rarity = ""
	item.ItemType = ""
	item.SubType = ""
	item.ItemLevel = 0
	item.TopEnd = 0.0
	item.DPS = 0.0
	item.Stamina = 0
	item.Strength = 0
	item.Agility = 0
	item.Intellect = 0
	item.Spirit = 0
	item.Mp5 = 0
	item.SpellDamage = 0
	item.SpellHealing = 0
	item.MeleeAttackPower = 0
	item.RangedAttackPower = 0
	item.MeleeHit = 0
	item.RangedHit = 0
	item.SpellHit = 0
	item.MeleeCrit = 0
	item.RangedCrit = 0
	item.SpellCrit = 0
	item.Weights = []*weight{}
	return item
}

var (
	cloth    = "Cloth Armor"
	leather  = "Leather Armor"
	mail     = "Mail Armor"
	plate    = "Plate Armor"
	shield   = "Shields"
	trinket  = "Trinkets"
	neck     = "Amulets"
	ring     = "Rings"
	cloak    = "Cloaks"
	dagger   = "Daggers"
	fist     = "Fist Weapons"
	mace1h   = "One-Handed Maces"
	mace2h   = "Two-Handed Maces"
	axe1h    = "One-Handed Axes"
	axe2h    = "Two-Handed Axes"
	sword1h  = "One-Handed Swords"
	sword2h  = "Two-Handed Swords"
	staff    = "Staves"
	polearm  = "Polearms"
	bow      = "Bows"
	gun      = "Guns"
	crossbow = "Crossbows"
	wand     = "Wands"
	offhand  = "Off-hand Frills"
)

func (i *item) isUsableByClass(class string) (bool, error) {
	var usable []string

	switch class {
	case "Druid":
		usable = []string{cloth, leather, trinket, neck, ring, cloak, dagger, fist, mace1h, mace2h, staff, offhand}
	case "Hunter":
		usable = []string{leather, mail, dagger, fist, axe1h, axe2h, sword1h, sword2h, staff, polearm, bow, gun, crossbow, trinket, neck, ring, cloak, offhand}
	case "Mage":
		usable = []string{cloth, trinket, neck, ring, cloak, dagger, sword1h, staff, offhand, wand}
	case "Priest":
		usable = []string{cloth, trinket, neck, ring, cloak, dagger, mace1h, staff, offhand, wand}
	case "Rogue":
		usable = []string{leather, trinket, neck, ring, cloak, dagger, fist, mace1h, sword1h, bow, gun, crossbow}
	case "Shaman":
		usable = []string{cloth, leather, mail, shield, trinket, neck, ring, cloak, dagger, fist, mace1h, mace2h, axe1h, axe2h, staff, offhand}
	case "Warlock":
		usable = []string{cloth, trinket, neck, ring, cloak, dagger, sword1h, staff, offhand, wand}
	case "Warrior":
		usable = []string{leather, mail, plate, shield, trinket, neck, ring, cloak, dagger, fist, mace1h, mace2h, axe1h, axe2h, sword1h, sword2h, staff, polearm, bow, gun, crossbow}
	default:
		return false, fmt.Errorf("%s is not a supported class", class)
	}

	for _, subtype := range usable {
		if subtype == i.SubType {
			return true, nil
		}
	}
	return false, nil
}

func (i *item) isUsableByRole(role string) (bool, error) {
	switch role {
	case "Melee":
		if (i.MeleeAttackPower > 0 || i.MeleeHit > 0 || i.MeleeCrit > 0 ||
			i.Strength > 0 || i.Agility > 0) && i.SpellDamage == 0 &&
			i.SpellHealing == 0 && i.SpellHit == 0 && i.SpellCrit == 0 {
			return true, nil
		}
		return false, nil
	case "Magic":
		if (i.SpellDamage > 0 || i.SpellHit > 0 || i.SpellCrit > 0 ||
			i.Intellect > 0 || i.Spirit > 0) && i.SpellHealing == 0 &&
			i.MeleeAttackPower == 0 && i.MeleeHit == 0 && i.MeleeCrit == 0 &&
			i.Strength == 0 && i.Agility == 0 {
			return true, nil
		}
		return false, nil
	case "Healer":
		if (i.SpellHealing > 0 || i.SpellDamage > 0 || i.SpellCrit > 0 ||
			i.Intellect > 0 || i.Spirit > 0) && i.SpellHit == 0 &&
			i.MeleeAttackPower == 0 && i.MeleeHit == 0 && i.MeleeCrit == 0 &&
			i.Strength == 0 && i.Agility == 0 {
			return true, nil
		}
		return false, nil
	case "Ranged":
		if (i.RangedAttackPower > 0 || i.RangedHit > 0 || i.RangedCrit > 0 ||
			i.Agility > 0) && i.Strength == 0 && i.SpellDamage == 0 &&
			i.SpellHealing == 0 && i.SpellHit == 0 && i.SpellCrit == 0 {
			return true, nil
		}
		return false, nil
	default:
		return false, fmt.Errorf("%s is not a supported role", role)
	}
}

func (i *item) setItemWeights(ws weights) error {
	if melee, err := i.isUsableByRole("Melee"); melee && err == nil {
		if warrior, err := i.isUsableByClass("Warrior"); warrior && err == nil {
			i.Weights = append(i.Weights, ws.getWeightsByClass("Warrior").getWeightsByRole("Melee")...)
		} else if err != nil {
			return err
		}
		if shaman, err := i.isUsableByClass("Shaman"); shaman && err == nil {
			i.Weights = append(i.Weights, ws.getWeightsByClass("Shaman").getWeightsByRole("Melee")...)
		} else if err != nil {
			return err
		}
		if rogue, err := i.isUsableByClass("Rogue"); rogue && err == nil {
			i.Weights = append(i.Weights, ws.getWeightsByClass("Rogue").getWeightsByRole("Melee")...)
		} else if err != nil {
			return err
		}
		if druid, err := i.isUsableByClass("Druid"); druid && err == nil {
			i.Weights = append(i.Weights, ws.getWeightsByClass("Druid").getWeightsByRole("Melee")...)
		} else if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if ranged, err := i.isUsableByRole("Ranged"); ranged && err == nil {
		i.Weights = append(i.Weights, ws.getWeightsByClass("Hunter").getWeightsByRole("Ranged")...)
	} else if err != nil {
		return err
	}

	if magic, err := i.isUsableByRole("Magic"); magic && err == nil {
		if mage, err := i.isUsableByClass("Mage"); mage && err == nil {
			i.Weights = append(i.Weights, ws.getWeightsByClass("Mage").getWeightsByRole("Magic")...)
		} else if err != nil {
			return err
		}
		if warlock, err := i.isUsableByClass("Warlock"); warlock && err == nil {
			i.Weights = append(i.Weights, ws.getWeightsByClass("Warlock").getWeightsByRole("Magic")...)
		} else if err != nil {
			return err
		}
		if shaman, err := i.isUsableByClass("Shaman"); shaman && err == nil {
			i.Weights = append(i.Weights, ws.getWeightsByClass("Shaman").getWeightsByRole("Magic")...)
		} else if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if healer, err := i.isUsableByRole("Healer"); healer && err == nil {
		if druid, err := i.isUsableByClass("Druid"); druid && err == nil {
			i.Weights = append(i.Weights, ws.getWeightsByClass("Druid").getWeightsByRole("Healer")...)
		} else if err != nil {
			return err
		}
		if priest, err := i.isUsableByClass("Priest"); priest && err == nil {
			i.Weights = append(i.Weights, ws.getWeightsByClass("Priest").getWeightsByRole("Healer")...)
		} else if err != nil {
			return err
		}
		if shaman, err := i.isUsableByClass("Shaman"); shaman && err == nil {
			i.Weights = append(i.Weights, ws.getWeightsByClass("Shaman").getWeightsByRole("Healer")...)
		} else if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
