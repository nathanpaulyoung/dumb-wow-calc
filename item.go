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
	return item
}

func (i *item) print() {
	fmt.Println("Item ID:", i.ID)
	fmt.Println("Name:", i.Name)
	fmt.Println("Source:", i.Source)
	fmt.Println("iLvl:", i.ItemLevel)
	fmt.Println("Top End Dmg:", i.TopEnd)
	fmt.Println("DPS:", i.DPS)
	fmt.Println("Stamina:", i.Stamina)
	fmt.Println("Strength:", i.Strength)
	fmt.Println("Agility:", i.Agility)
	fmt.Println("Intellect:", i.Intellect)
	fmt.Println("Spirit:", i.Spirit)
	fmt.Println("Mp5:", i.Mp5)
	fmt.Println("Spell Damage:", i.SpellDamage)
	fmt.Println("+Healing:", i.SpellHealing)
	fmt.Println("Melee AP:", i.MeleeAttackPower)
	fmt.Println("Ranged AP:", i.RangedAttackPower)
	fmt.Println("Melee Hit:", i.MeleeHit)
	fmt.Println("Ranged Hit:", i.RangedHit)
	fmt.Println("Spell Hit:", i.SpellHit)
	fmt.Println("Melee Crit:", i.MeleeCrit)
	fmt.Println("Ranged Crit:", i.RangedCrit)
	fmt.Println("Spell Crit:", i.SpellCrit)
}
