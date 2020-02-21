package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type weight struct {
	FileName          string
	Class             string  `yaml:"class"`
	Role              string  `yaml:"role"`
	TopEnd            float32 `yaml:"top-end-dmg"`
	DPS               float32 `yaml:"dps"`
	Strength          float32 `yaml:"strength"`
	Agility           float32 `yaml:"agility"`
	Intellect         float32 `yaml:"intellect"`
	Spirit            float32 `yaml:"spirit"`
	Mp5               float32 `yaml:"mp5"`
	SpellDamage       float32 `yaml:"spell-damage"`
	SpellHealing      float32 `yaml:"spell-healing"`
	MeleeAttackPower  float32 `yaml:"melee-attack-power"`
	RangedAttackPower float32 `yaml:"ranged-attack-power"`
	MeleeHit          float32 `yaml:"melee-hit"`
	RangedHit         float32 `yaml:"ranged-hit"`
	SpellHit          float32 `yaml:"spell-hit"`
	MeleeCrit         float32 `yaml:"melee-crit"`
	RangedCrit        float32 `yaml:"ranged-crit"`
	SpellCrit         float32 `yaml:"spell-crit"`
}

type weights []*weight

func newWeight(file string) *weight {
	weight := new(weight)
	weight.FileName = file
	return weight
}

func (w *weight) getWeightsFromFile() {
	yamlFile, err := ioutil.ReadFile("weightfiles/" + w.FileName + ".yaml")
	if err != nil {
		log.Printf("Weight file not found: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &w)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func (ws weights) getWeightsByClass(class string) weights {
	var out weights

	for i := 0; i < len(ws); i++ {
		if ws[i].Class == class {
			out = append(out, ws[i])
		}
	}

	return out
}

func (ws weights) getWeightsByRole(role string) weights {
	var out weights

	for i := 0; i < len(ws); i++ {
		if ws[i].Role == role {
			out = append(out, ws[i])
		}
	}

	return out
}

func (ws weights) averageWeights() *weight {
	w := newWeight("")

	for i := 0; i < len(ws); i++ {
		w.TopEnd += ws[i].TopEnd
		w.DPS += ws[i].DPS
		w.Strength += ws[i].Strength
		w.Agility += ws[i].Agility
		w.Intellect += ws[i].Intellect
		w.Spirit += ws[i].Spirit
		w.Mp5 += ws[i].Mp5
		w.SpellDamage += ws[i].SpellDamage
		w.SpellHealing += ws[i].SpellHealing
		w.MeleeAttackPower += ws[i].MeleeAttackPower
		w.RangedAttackPower += ws[i].RangedAttackPower
		w.MeleeHit += ws[i].MeleeHit
		w.RangedHit += ws[i].RangedHit
		w.SpellHit += ws[i].SpellHit
		w.MeleeCrit += ws[i].MeleeCrit
		w.RangedCrit += ws[i].RangedCrit
		w.SpellCrit += ws[i].SpellCrit
	}

	w.TopEnd /= float32(len(ws) + 1)
	w.DPS /= float32(len(ws) + 1)
	w.Strength /= float32(len(ws) + 1)
	w.Agility /= float32(len(ws) + 1)
	w.Intellect /= float32(len(ws) + 1)
	w.Spirit /= float32(len(ws) + 1)
	w.Mp5 /= float32(len(ws) + 1)
	w.SpellDamage /= float32(len(ws) + 1)
	w.SpellHealing /= float32(len(ws) + 1)
	w.MeleeAttackPower /= float32(len(ws) + 1)
	w.RangedAttackPower /= float32(len(ws) + 1)
	w.MeleeHit /= float32(len(ws) + 1)
	w.RangedHit /= float32(len(ws) + 1)
	w.SpellHit /= float32(len(ws) + 1)
	w.MeleeCrit /= float32(len(ws) + 1)
	w.RangedCrit /= float32(len(ws) + 1)
	w.SpellCrit /= float32(len(ws) + 1)

	return w
}
