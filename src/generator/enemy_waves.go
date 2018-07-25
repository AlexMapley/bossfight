package generator

import (
	"../creatures"
)

type Waves struct {
	enemy_waves []*EnemyWave
}

type EnemyWave struct {
	enemies []*creatures.Creature
}
