package com.obarra.minesweeperclient.model;

public class GameResponse {
    private Integer gameId;
    private Integer rows;
    private Integer columns;
    private Integer mineAmount;

    public Integer getGameId() {
        return gameId;
    }

    public void setGameId(Integer gameId) {
        this.gameId = gameId;
    }
    public Integer getRows() {
        return rows;
    }

    public void setRows(Integer rows) {
        this.rows = rows;
    }

    public Integer getColumns() {
        return columns;
    }

    public void setColumns(Integer columns) {
        this.columns = columns;
    }

    public Integer getMineAmount() {
        return mineAmount;
    }

    public void setMineAmount(Integer mineAmount) {
        this.mineAmount = mineAmount;
    }

    @Override
    public String toString() {
        return "GameResponse{" +
                "gameId=" + gameId +
                ", rows=" + rows +
                ", columns=" + columns +
                ", mineAmount=" + mineAmount +
                '}';
    }
}
