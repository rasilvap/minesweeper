package engine

import (
	"minesweeper-API/models/dto"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/obarra-dev/minesweeper"
)

func Test_buildPlayResponse(t *testing.T) {
	t.Run("in board 3x8 when running", func(t *testing.T) {
		mines := []minesweeper.Mine{{Row: 1, Column: 1}}
		mwNew := minesweeper.New(3, 8, mines)
		mw := mwNew.Play(0, 5, minesweeper.TypeMoveClean)

		//ect
		got := buildPlayResponse(mw)

		assert.Equal(t, dto.GameStateRunning, got.Game.State)
		assert.Equal(t, 8, got.Game.Columns)
		assert.Equal(t, 3, got.Game.Rows)
		assert.Equal(t, got.Game.Rows*got.Game.Columns, len(got.Game.Board))
	})

	t.Run("in board 3x3", func(t *testing.T) {
		mines := []minesweeper.Mine{{Row: 1, Column: 1}}

		t.Run("when lost", func(t *testing.T) {
			mwNew := minesweeper.New(3, 3, mines)
			mw := mwNew.Play(1, 1, minesweeper.TypeMoveClean)

			//ect
			got := buildPlayResponse(mw)

			assert.Equal(t, dto.GameStateLost, got.Game.State)
			assert.Equal(t, 3, got.Game.Columns)
			assert.Equal(t, 3, got.Game.Rows)
			assert.Equal(t, got.Game.Rows*got.Game.Columns, len(got.Game.Board))
		})

		t.Run("when lost and play again", func(t *testing.T) {
			mwNew := minesweeper.New(3, 3, mines)

			mw := mwNew.Play(0, 0, minesweeper.TypeMoveClean)
			got := buildPlayResponse(mw)
			assert.Equal(t, got.Game.State, dto.GameStateRunning)

			mw = mwNew.Play(1, 1, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, dto.GameStateLost, got.Game.State)

			mw = mwNew.Play(0, 2, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, dto.GameStateLost, got.Game.State)
		})

		t.Run("when won", func(t *testing.T) {
			mwNew := minesweeper.New(3, 3, mines)

			mw := mwNew.Play(0, 0, minesweeper.TypeMoveClean)
			got := buildPlayResponse(mw)
			assert.Equal(t, got.Game.State, dto.GameStateRunning)

			mw = mwNew.Play(0, 1, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, dto.GameStateRunning, got.Game.State)

			mw = mwNew.Play(0, 2, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, dto.GameStateRunning, got.Game.State)

			mw = mwNew.Play(1, 0, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, dto.GameStateRunning, got.Game.State)

			mw = mwNew.Play(1, 2, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, dto.GameStateRunning, got.Game.State)

			mw = mwNew.Play(2, 0, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, dto.GameStateRunning, got.Game.State)

			mw = mwNew.Play(2, 1, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, dto.GameStateRunning, got.Game.State)

			mw = mwNew.Play(2, 2, minesweeper.TypeMoveClean)
			got = buildPlayResponse(mw)
			assert.Equal(t, dto.GameStateWon, got.Game.State)
		})
	})
}
