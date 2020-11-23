package com.obarra.minesweeperclient.service;

import com.obarra.minesweeperclient.model.GameBoard;
import com.obarra.minesweeperclient.model.PlayResponse;
import com.obarra.minesweeperclient.model.TileDTO;
import com.obarra.minesweeperclient.utils.StateGameEnum;
import com.obarra.minesweeperclient.utils.StateTileEnum;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public final class GameBoardRender {

    public List<List<String>> generateEmptyBoard(Integer rows, Integer columns) {
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

    public GameBoard updateGameBoard(final GameBoard gameBoard, final PlayResponse playResponse) {
        final var updatedGameBoard = new GameBoard();
        updatedGameBoard.setGameId(gameBoard.getGameId());

        var gameState = getGameState(playResponse.getStateGame());
        updatedGameBoard.setState(gameState);

        updateBoard(gameBoard.getBoard(), playResponse.getGame().getBoard());
        updatedGameBoard.setBoard(gameBoard.getBoard());

        return updatedGameBoard;
    }

    private StateGameEnum getGameState(final String stateGameResponse) {
        final var gameState = StateGameEnum.valueOf(stateGameResponse);
        switch (gameState) {
            case WON:
                System.out.println("Flawless Victory");
                break;
            case LOST:
                System.out.println("Game Over");
                break;
            case RUNNING:
                System.out.println("The Game is Running");
                break;
            default:
                throw new IllegalStateException("Invalid Game State: " + stateGameResponse);
        }
        return gameState;
    }


    private static void updateBoard(List<List<String>> currentBoard, TileDTO[][] resultBoard) {
        for (TileDTO[] rows : resultBoard) {
            for (TileDTO tileDTO : rows) {
                if (tileDTO.getMine() != null && tileDTO.getMine()) {
                    currentBoard.get(tileDTO.getRow()).set(tileDTO.getColumn(), "B");
                    System.out.println("BOOM.......");
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
}
