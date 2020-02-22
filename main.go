package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	IDs := getIDs("./ids.txt")

	var items []*item

	for _, itemID := range IDs {
		// Make HTTP GET request
		response, err := http.Get(fmt.Sprintf("https://classic.wowhead.com/item=%s?xml", itemID))
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)

		item := newItem()
		item.ID, _ = strconv.Atoi(itemID)
		item.Name = getName(string(body))
		item.Source = getSource(string(body))
		item.Rarity = getRarity(string(body))
		item.ItemType = getItemType(string(body))
		item.SubType = getSubType(string(body))
		item.ItemLevel = getItemLevel(string(body))
		item.TopEnd = float32(getTopEnd(string(body)))
		item.DPS = float32(getDPS(string(body)))
		item.Stamina = getStamina(string(body))
		item.Strength = getStrength(string(body))
		item.Agility = getAgility(string(body))
		item.Intellect = getIntellect(string(body))
		item.Spirit = getSpirit(string(body))
		item.Mp5 = getMp5(string(body))
		item.SpellDamage = getSpellDamage(string(body))
		item.SpellHealing = getSpellHealing(string(body))
		item.MeleeAttackPower = getMeleeAttackPower(string(body))
		item.RangedAttackPower = getRangedAttackPower(string(body))
		item.MeleeHit = getMeleeHit(string(body))
		item.RangedHit = getRangedHit(string(body))
		item.SpellHit = getSpellHit(string(body))
		item.MeleeCrit = getMeleeCrit(string(body))
		item.RangedCrit = getRangedCrit(string(body))
		item.SpellCrit = getSpellCrit(string(body))

		items = append(items, item)
	}

	var csv string
	weights := generateWeights()

	for i := 0; i < len(items); i++ {
		csv = buildCSV(items[i], weights, csv)
	}

	ioutil.WriteFile("gpvalues.csv", []byte(csv), 0644)
}

func getIDs(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic("File read error.")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func buildCSV(item *item, weights []*weight, csv string) string {
	if len(csv) == 0 {
		csv = "Item Name,Source,GP\n"
	}
	csv += fmt.Sprintf("\"%s\",%s,%f\n", item.Name, item.Source, calculateGP(item, weights))
	return csv
}

func calculateGP(item *item, ws weights) float32 {
	item.setItemWeights(ws)

	w := item.Weights.averageWeights()

	if item.MeleeAttackPower == item.RangedAttackPower {
		item.RangedAttackPower = 0
	}

	if item.MeleeHit == item.RangedHit {
		item.RangedHit = 0
	}

	if item.MeleeCrit == item.RangedCrit {
		item.RangedCrit = 0
	}

	itemdps := (item.DPS * w.DPS)
	itemstr := (float32(item.Strength) * w.Strength)
	itemagi := (float32(item.Agility) * w.Agility)
	itemint := (float32(item.Intellect) * w.Intellect)
	itemspi := (float32(item.Spirit) * w.Spirit)
	itemmp5 := (float32(item.Mp5) * w.Mp5)
	itemsd := (float32(item.SpellDamage) * w.SpellDamage)
	itemheal := (float32(item.SpellHealing) * w.SpellHealing)
	itemmap := (float32(item.MeleeAttackPower) * w.MeleeAttackPower)
	itemrap := (float32(item.RangedAttackPower) * w.RangedAttackPower)
	itemmh := (float32(item.MeleeHit) * w.MeleeHit)
	itemrh := (float32(item.RangedHit) * w.RangedHit)
	itemsh := (float32(item.SpellHit) * w.SpellHit)
	itemmc := (float32(item.MeleeCrit) * w.MeleeCrit)
	itemrc := (float32(item.RangedCrit) * w.RangedCrit)
	itemsp := (float32(item.SpellCrit) * w.SpellCrit)

	return itemdps + itemstr + itemagi + itemint + itemspi + itemmp5 + itemsd + itemheal + itemmap + itemrap + itemmh + itemrh + itemsh + itemmc + itemrc + itemsp
}

func generateWeights() []*weight {
	files := []string{
		"druid-kitty",
		"druid-restoration",
		"hunter",
		"mage",
		"priest-heal",
		"rogue",
		"shaman-elemental",
		"shaman-enhancement",
		"shaman-restoration",
		"warlock",
		"warrior-fury",
	}
	var weights []*weight

	for i := 0; i < len(files); i++ {
		weight := newWeight(files[i])
		weight.getWeightsFromFile()
		weights = append(weights, weight)
	}

	return weights
}
