package com.obarra.minesweeperclient.utils;

import com.obarra.minesweeperclient.model.PlayResponse;
import com.obarra.minesweeperclient.model.TileDTO;

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

    public static List<List<String>> render(List<List<String>> board, PlayResponse playResponse) {
        var updatedBoard = playResponse.getGame().getBoard();

        if (StateGameEnum.LOST.name().equals(playResponse.getStateGame())
                || StateGameEnum.WON.name().equals(playResponse.getStateGame())) {
            updateBoard(board, playResponse.getGame().getBoard());
            System.out.println("Game Finish" + playResponse.getStateGame());
            return null;
        }

        System.out.println("The Game is Running" + playResponse.getStateGame());
        updateBoardRUN(board, playResponse.getGame().getBoard());

        return null;
    }

    private static void updateBoard(List<List<String>> currentBoard, TileDTO[][] resultBoard) {
        for (TileDTO[] rows : resultBoard) {
            for (TileDTO tileDTO : rows) {
                //TODO de debe ser null ver
                if (tileDTO.getIsMine() != null && tileDTO.getIsMine()) {
                    currentBoard.get(tileDTO.getRow()).set(tileDTO.getColumn(), "B");
                } else if (StateTileEnum.CLEAR.name().equals(tileDTO.getState()) &&
                        (tileDTO.getSurroundingMineCount() == null || tileDTO.getSurroundingMineCount() == 0)) {
                    System.out.println("CLEAR.......");
                    currentBoard.get(tileDTO.getRow()).set(tileDTO.getColumn(), "C");
                } else if (tileDTO.getSurroundingMineCount() != null && tileDTO.getSurroundingMineCount() > 0) {
                    System.out.println("NUMBERED.......");
                    currentBoard.get(tileDTO.getRow()).set(tileDTO.getColumn(), tileDTO.getSurroundingMineCount().toString());
                } else {
                    System.out.println("NOTHING.......");
                }
            }
        }
    }

    private static void updateBoardRUN(List<List<String>> currentBoard, TileDTO[][] resultBoard) {
        for (TileDTO[] rows : resultBoard) {
            for (TileDTO tileDTO : rows) {
                if (StateTileEnum.CLEAR.name().equals(tileDTO.getState()) &&
                        (tileDTO.getSurroundingMineCount() == null || tileDTO.getSurroundingMineCount() == 0)) {
                    System.out.println("CLEAR.......");
                    currentBoard.get(tileDTO.getRow()).set(tileDTO.getColumn(), "C");
                } else if (tileDTO.getSurroundingMineCount() != null && tileDTO.getSurroundingMineCount() > 0) {
                    System.out.println("NUMBERED.......");
                    currentBoard.get(tileDTO.getRow()).set(tileDTO.getColumn(), tileDTO.getSurroundingMineCount().toString());
                } else {
                    System.out.println("NOTHING.......");
                }
            }
        }
    }
}
