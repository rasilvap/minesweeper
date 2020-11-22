package com.obarra.minesweeperclient.utils;

import java.util.ArrayList;
import java.util.List;

public final class GameBoardRenderUtil {
    private GameBoardRenderUtil() {
    }

    public static List<List<String>> generateEmptyBoard(Integer rows, Integer columns) {
        var board = new ArrayList<List<String>>();
        for (int i = 0; i < rows; i++) {
            var row = new ArrayList<String>();
            board.add(row);
            for (int j = 0; j < columns; j++) {
                row.add("");
            }
        }
        return board;
    }
}
