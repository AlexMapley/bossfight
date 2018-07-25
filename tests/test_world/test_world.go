package testworld

import (
	"testing"

	"$GOPATH/src/bossfight/src/generator"

	"github.com/faiface/pixel"
)

// Global environment variables
var (
	trees         []*pixel.Sprite
	enemies       []*pixel.Sprite
	matrices      []pixel.Matrix
	enemyMatrices []pixel.Matrix
)

func testGenerateEnvironment(*testing.T) {
	_ = generator.LoadTrees(*win)
}

func testGenerateCreatures(*testing.T) {
	_ = generator.EnemyGenerator(*win)
}

func testGenerateAll(*testing.T) {
	// Background elements
	_ = generator.LoadTrees(*win)
	_ = generator.EnemyGenerator(*win)

}
