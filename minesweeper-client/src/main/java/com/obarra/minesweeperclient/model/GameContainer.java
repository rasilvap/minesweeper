package com.obarra.minesweeperclient.model;

import java.util.Objects;

public class GameContainer {
    private GameBoard gameBoard;

    public GameBoard getGameBoard() {
        return gameBoard;
    }

    public void setGameBoard(GameBoard gameBoard) {
        this.gameBoard = gameBoard;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        GameContainer that = (GameContainer) o;
        return Objects.equals(gameBoard, that.gameBoard);
    }

    @Override
    public int hashCode() {
        return Objects.hash(gameBoard);
    }
}
