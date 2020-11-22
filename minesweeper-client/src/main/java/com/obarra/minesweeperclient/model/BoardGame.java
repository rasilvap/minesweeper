package com.obarra.minesweeperclient.model;

import java.util.List;

public class BoardGame {
    private Integer gameId;
    private List<List<String>> board;

    public Integer getGameId() {
        return gameId;
    }

    public void setGameId(Integer gameId) {
        this.gameId = gameId;
    }

    public List<List<String>> getBoard() {
        return board;
    }

    public void setBoard(List<List<String>> board) {
        this.board = board;
    }

    @Override
    public String toString() {
        return "BoardGame{" +
                "gameId=" + gameId +
                ", board=" + board +
                '}';
    }
}
