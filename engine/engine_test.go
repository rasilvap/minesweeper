package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/obarra-dev/minesweeper"
)


func Test_buildPlayResponse(t *testing.T) {
	t.Run("in board 3x8 when running", func(t *testing.T) {
		mines := [][2]int{{1, 1}}
		mwNew := minesweeper.NewMinesweeper(3, 8, mines)
		mw := mwNew.Play(0, 5, minesweeper.TypeMoveClean)

		//ect
		got := buildPlayResponse(mw)

		assert.Equal(t, "RUNNING", got.StateGame)
		assert.Equal(t, 8, got.Game.Columns)
		assert.Equal(t, 3, got.Game.Rows)
		assert.Equal(t, got.Game.Rows, len(got.Game.Board))
		assert.Equal(t, 6, len(got.Game.Board[0]))
	})

	t.Run("in board 3x3", func(t *testing.T) {
		mines := [][2]int{{1, 1}}

		t.Run("when lost", func(t *testing.T) {
			mwNew := minesweeper.NewMinesweeper(3, 3, mines)
			mw := mwNew.Play(1, 1, minesweeper.TypeMoveClean)

			//ect
			got := buildPlayResponse(mw)

			assert.Equal(t, "LOST", got.StateGame)
			assert.Equal(t, 3, got.Game.Columns)
			assert.Equal(t, 3, got.Game.Rows)
			assert.Equal(t, got.Game.Rows, len(got.Game.Board))
			assert.Equal(t, got.Game.Columns, len(got.Game.Board[0]))
		})

		t.Run("when lost and play again", func(t *testing.T) {
			mwNew := minesweeper.NewMinesweeper(3, 3, mines)

			mw := mwNew.Play(0, 0, minesweeper.TypeMoveClean)
			got := buildPlayResponse(mw)
			assert.Equal(t, got.StateGame, "RUNNING")

			mw = mwNew.Play(1, 1, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, "LOST", got.StateGame)

			//TODO Bug
			mw = mwNew.Play(0, 2, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, "LOST", got.StateGame)
		})

		t.Run("when won", func(t *testing.T) {
			mwNew := minesweeper.NewMinesweeper(3, 3, mines)

			mw := mwNew.Play(0, 0, minesweeper.TypeMoveClean)
			got := buildPlayResponse(mw)
			assert.Equal(t, got.StateGame, "RUNNING")

			mw = mwNew.Play(0, 1, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, "RUNNING", got.StateGame)

			mw = mwNew.Play(0, 2, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, "RUNNING", got.StateGame)

			mw = mwNew.Play(1, 0, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, "RUNNING", got.StateGame)

			mw = mwNew.Play(1, 2, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, "RUNNING", got.StateGame)

			mw = mwNew.Play(2, 0, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, "RUNNING", got.StateGame)

			mw = mwNew.Play(2, 1, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, "RUNNING", got.StateGame)

			mw = mwNew.Play(2, 2, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, "WON", got.StateGame)
		})
	})
}
