package load

import (
	"game/ebitd"
	"game/td"
	"log"
)

func createShootingArcherAnim() *td.Animation {
	var runningArcherPaths = getShootingArcherPaths()
	scale := td.Scale{WFactor: 3, HFactor: 3}
	duration := 7
	animation, err := ebitd.CreateAnim(runningArcherPaths, scale, duration)
	if err != nil {
		log.Fatalf("Error loading archers: %v", err)
	}
	return animation
}

func createRunningArcherAnim() *td.Animation {
	var runningArcherPaths = getRunningArcherPaths()
	scale := td.Scale{WFactor: 3, HFactor: 3}
	duration := 7
	animation, err := ebitd.CreateAnim(runningArcherPaths, scale, duration)
	if err != nil {
		log.Fatalf("Error loading archers: %v", err)
	}
	return animation
}

func createSpearAnim() *td.Animation {
	scale := td.Scale{WFactor: 1, HFactor: 1}
	duration := 7
	paths := []string{lancePath}
	spearAnim, err := ebitd.CreateAnim(paths, scale, duration)
	if err != nil {
		log.Fatalf("Error loading spear: %v", err)
	}
	return spearAnim
}

func Animations() map[string]map[string]*td.Animation {
	var spearAnim *td.Animation = createSpearAnim()
	var runningArcherAnim *td.Animation = createRunningArcherAnim()
	var shootingArcherAnim *td.Animation = createShootingArcherAnim()

	archerAnim := make(map[string]*td.Animation, 2)
	archerAnim["move"] = runningArcherAnim
	archerAnim["shoot"] = shootingArcherAnim

	regularSpear := make(map[string]*td.Animation)
	regularSpear[""] = spearAnim

	animations := make(map[string]map[string]*td.Animation, 1)
	animations["archer"] = archerAnim
	animations["spear"] = regularSpear
	return animations
}
