package com.obarra.minesweeperclient.controller;

import java.util.List;

public class BoardGame {
    private List<List<String>> board;

    public List<List<String>> getBoard() {
        return board;
    }

    public void setBoard(List<List<String>> board) {
        this.board = board;
    }
}
