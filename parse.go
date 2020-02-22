package main

import (
	"strconv"
	"strings"
)

func getName(body string) string {
	return getValueBySubstring(body, "\"name\":\"", 1)
}

func getSource(body string) string {
	return getValueBySubstring(body, "\"n\":\"", 1)
}

func getRarity(body string) string {
	start, end := getTypeIndexRange(body)
	rarities := []string{"Poor", "Common", "Uncommon", "Rare", "Epic", "Legendary"}

	for i := 0; i < len(rarities); i++ {
		if strings.Contains(body[start:end], rarities[i]) {
			return rarities[i]
		}
	}
	return "Error"
}

func getItemType(body string) string {
	start, end := getTypeIndexRange(body)
	itemtype := []string{"Weapons", "Armor"}

	for i := 0; i < len(itemtype); i++ {
		if strings.Contains(body[start:end], itemtype[i]) {
			return itemtype[i]
		}
	}
	return "Error"
}

func getSubType(body string) string {
	start, end := getTypeIndexRange(body)
	subtype := []string{
		"Cloth Armor", "Leather Armor", "Mail Armor", "Plate Armor",
		"Trinkets", "Amulets", "Daggers", "Shields", "Fist Weapons",
		"One-Handed Axes", "One-Handed Swords", "One-Handed Maces",
		"Two-Handed Axes", "Two-Handed Swords", "Two-Handed Maces",
		"Rings", "Staves", "Cloaks", "Guns", "Bows", "Crossbows",
		"Wands", "Off-hand Frills", "Polearms",
	}

	for i := 0; i < len(subtype); i++ {
		if strings.Contains(body[start:end], subtype[i]) {
			return subtype[i]
		}
	}
	return "Error"
}

func getTypeIndexRange(body string) (int, int) {
	return strings.IndexAny(body, "0123456789"), strings.Index(body, "<table><tr><td><!--nstart-->")
}

func getItemLevel(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"level\":", 0))
	return value
}

func getSlot(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"slot\":", 0))
	return value
}

func getTopEnd(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"dmgmax1\":", 0))
	return value
}

func getDPS(body string) float32 {
	value, _ := strconv.ParseFloat(getValueBySubstring(body, "\"dps\":", 0), 32)
	return float32(value)
}

func getStamina(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"sta\":", 0))
	return value
}

func getStrength(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"str\":", 0))
	return value
}

func getAgility(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"agi\":", 0))
	return value
}

func getIntellect(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"int\":", 0))
	return value
}

func getSpirit(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"spi\":", 0))
	return value
}

func getMp5(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"manargn\":", 0))
	return value
}

func getSpellDamage(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"splpwr\":", 0))
	return value
}

func getSpellHealing(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"splheal\":", 0))
	return value
}

func getMeleeAttackPower(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"mleatkpwr\":", 0))
	return value
}

func getRangedAttackPower(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"rgdatkpwr\":", 0))
	return value
}

func getMeleeHit(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"mlehitpct\":", 0))
	return value
}

func getRangedHit(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"rgdhitpct\":", 0))
	return value
}

func getSpellHit(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"splhitpct\":", 0))
	return value
}

func getMeleeCrit(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"mlecritstrkpct\":", 0))
	return value
}

func getRangedCrit(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"rgdcritstrkpct\":", 0))
	return value
}

func getSpellCrit(body string) int {
	value, _ := strconv.Atoi(getValueBySubstring(body, "\"splcritstrkpct\":", 0))
	return value
}

func getValueBySubstring(body string, sub string, valueType int) string {
	length := len(sub)
	index := strings.Index(body, sub)
	value := ""

	for i := index + length; ; i++ {
		if !strings.Contains(body, sub) {
			break
		}
		char := body[i : i+1]
		if valueType == 0 {
			_, err := strconv.ParseFloat(char, 64)
			if err != nil && char != "." {
				break
			}
		}
		if valueType == 1 {
			if strings.Contains(char, "\"") {
				break
			}
		}
		value += char
	}
	return value
}
