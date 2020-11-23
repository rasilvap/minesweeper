package com.obarra.minesweeperclient.model;

import com.obarra.minesweeperclient.utils.StateGameEnum;

import java.util.List;

public class GameBoard {
    private StateGameEnum state;
    private Integer gameId;
    private List<List<String>> board;
    private Integer mineAmount;

    public StateGameEnum getState() {
        return state;
    }

    public void setState(StateGameEnum state) {
        this.state = state;
    }

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

    public Integer getMineAmount() {
        return mineAmount;
    }

    public void setMineAmount(Integer mineAmount) {
        this.mineAmount = mineAmount;
    }

    @Override
    public String toString() {
        return "GameBoard{" +
                "state=" + state +
                ", gameId=" + gameId +
                ", board=" + board +
                ", mineAmount=" + mineAmount +
                '}';
    }
}
