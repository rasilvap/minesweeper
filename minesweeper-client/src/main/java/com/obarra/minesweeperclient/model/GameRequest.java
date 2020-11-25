package com.obarra.minesweeperclient.model;

import java.util.Objects;

public class GameRequest {
    private Integer rows;
    private Integer columns;
    private Integer mineAmount;

    public static GameRequest of(Integer rows, Integer columns, Integer mineAmount) {
        final var gameRequest = new GameRequest();
        gameRequest.setRows(rows);
        gameRequest.setColumns(columns);
        gameRequest.setMineAmount(mineAmount);
        return gameRequest;
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
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        GameRequest that = (GameRequest) o;
        return Objects.equals(rows, that.rows) &&
                Objects.equals(columns, that.columns) &&
                Objects.equals(mineAmount, that.mineAmount);
    }

    @Override
    public int hashCode() {
        return Objects.hash(rows, columns, mineAmount);
    }
}
