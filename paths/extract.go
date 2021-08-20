package paths

import "fmt"

func GetRunningArcher() (paths[]string) {
	for i := 0; i < 6; i++ {
		path := fmt.Sprintf("images/archer/adventurer-run-%02d.png", i)
		paths = append(paths, path)
	}
	return paths
}

func GetShootingArcher() (paths[]string) {
	for i := 0; i < 9; i++ {
		path := fmt.Sprintf("images/archer/adventurer-bow-%02d.png", i)
		paths = append(paths, path)
	}
	return paths
}
