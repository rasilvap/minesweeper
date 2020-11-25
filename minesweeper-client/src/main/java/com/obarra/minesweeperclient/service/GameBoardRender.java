package com.obarra.minesweeperclient.service;

import com.obarra.minesweeperclient.enums.StateGameEnum;
import com.obarra.minesweeperclient.enums.StateTileEnum;
import com.obarra.minesweeperclient.model.GameBoard;
import com.obarra.minesweeperclient.model.PlayResponse;
import com.obarra.minesweeperclient.model.TileDTO;
import com.obarra.minesweeperclient.utils.GameBoardConst;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

@Service
public class GameBoardRender {
    private final static Logger LOGGER = LoggerFactory.getLogger(GameBoardRender.class);

    public List<List<String>> generateEmptyBoard(Integer rows, Integer columns) {
        return IntStream.range(0, rows)
                .mapToObj(x -> IntStream.range(0, columns)
                        .mapToObj(y -> GameBoardConst.COVERED_NOT_PLAYED)
                        .collect(Collectors.toList()))
                .collect(Collectors.toList());
    }

    public GameBoard updateGameBoard(final GameBoard gameBoard, final PlayResponse playResponse) {
        final var updatedGameBoard = new GameBoard();
        updatedGameBoard.setGameId(gameBoard.getGameId());
        updatedGameBoard.setMineAmount(gameBoard.getMineAmount());

        var gameState = getGameState(playResponse.getStateGame());
        updatedGameBoard.setState(gameState);

        var updatedBoard = updateBoard(gameBoard.getBoard(), playResponse.getGame().getBoard());

        updatedGameBoard.setBoard(updatedBoard);

        return updatedGameBoard;
    }

    private StateGameEnum getGameState(final String stateGameResponse) {
        final var gameState = StateGameEnum.valueOf(stateGameResponse);
        switch (gameState) {
            //TODO show only a point
            case WON:
                LOGGER.info("Flawless Victory");
                break;
            //TODO show only mines
            case LOST:
                LOGGER.info("Game Over");
                break;
            //TODO show one o any points
            case RUNNING:
                LOGGER.info("The Game is Running");
                break;
            default:
                throw new IllegalStateException("Invalid Game State: " + stateGameResponse);
        }
        return gameState;
    }

    private static List<List<String>> updateBoard(List<List<String>> currentBoard, TileDTO[][] resultBoard) {
        List<List<String>> board = copyBoard(currentBoard);
        for (TileDTO[] rows : resultBoard) {
            for (TileDTO tileDTO : rows) {
                if (tileDTO.getMine()) {
                    board.get(tileDTO.getRow()).set(tileDTO.getColumn(), GameBoardConst.MINE_TILE);
                } else if (tileDTO.getSurroundingMineCount() == 0 &&
                        (StateTileEnum.CLEAR.name().equals(tileDTO.getState()) ||
                                StateTileEnum.COVERED.name().equals(tileDTO.getState()))) {
                    board.get(tileDTO.getRow()).set(tileDTO.getColumn(), GameBoardConst.COVERED_TILE);
                } else if (tileDTO.getSurroundingMineCount() > 0) {
                    board.get(tileDTO.getRow()).set(tileDTO.getColumn(), tileDTO.getSurroundingMineCount().toString());
                }
            }
        }

        return board;
    }

    private static List<List<String>> copyBoard(final List<List<String>> board) {
        var copyBoard = new ArrayList<List<String>>();
        for (var row : board) {
            var copyRow = new ArrayList<String>();
            copyBoard.add(copyRow);
            copyRow.addAll(row);
        }
        return copyBoard;
    }
}
