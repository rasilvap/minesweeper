package com.obarra.minesweeperclient.model;

import java.util.Arrays;

public class GameDTO {
    private TileDTO[][] board;
    private Integer rows;
    private Integer columns;
    private Integer flagAmount;

    public TileDTO[][] getBoard() {
        return board;
    }

    public void setBoard(TileDTO[][] board) {
        this.board = board;
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

    public Integer getFlagAmount() {
        return flagAmount;
    }

    public void setFlagAmount(Integer flagAmount) {
        this.flagAmount = flagAmount;
    }


    @Override
    public String toString() {
        return "GameDTO{" +
                "board=" + Arrays.toString(board) +
                ", rows=" + rows +
                ", columns=" + columns +
                ", flagAmount=" + flagAmount +
                '}';
    }
}

