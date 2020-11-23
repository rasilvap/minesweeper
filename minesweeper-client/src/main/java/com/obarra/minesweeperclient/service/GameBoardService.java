package com.obarra.minesweeperclient.service;

import com.obarra.minesweeperclient.client.MinesweeperClient;
import com.obarra.minesweeperclient.model.*;
import com.obarra.minesweeperclient.utils.StateGameEnum;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class GameBoardService {
    private final static Logger LOGGER = LoggerFactory.getLogger(GameBoardService.class);

    private final MinesweeperClient minesweeperClient;
    private final GameBoardRender gameBoardRender;

    @Autowired
    public GameBoardService(final MinesweeperClient minesweeperClient, final GameBoardRender gameBoardRender) {
        this.minesweeperClient = minesweeperClient;
        this.gameBoardRender = gameBoardRender;
    }

    public GameBoard createGameBoard(final Integer rows, final Integer columns, final Integer mineAmount) {
        final var gameResponse = minesweeperClient.create(GameRequest.of(rows, columns, mineAmount));
        LOGGER.info(String.format("Created gameId: %s", gameResponse.getGameId()));

        final var gameBoard = new GameBoard();
        gameBoard.setState(StateGameEnum.NEW);
        gameBoard.setMineAmount(mineAmount);
        gameBoard.setGameId(gameResponse.getGameId());

        final var board = gameBoardRender.generateEmptyBoard(rows, columns);
        gameBoard.setBoard(board);

        return gameBoard;
    }

    public GameBoard playMovement(final GameBoard gameBoard, final Integer row, final Integer column) {
        if (!isValidMovement(gameBoard, row, column)) {
            LOGGER.info(String.format("Invalid  row: %s, column: %s", row, column));
            return gameBoard;
        }

        if (gameBoard.getBoard().get(row).get(column).equals("F")) {
            return markTile(gameBoard, row, column, "");
        }

        final var playResponse = minesweeperClient.play(gameBoard.getGameId(), PlayRequest.of(row, column));
        return gameBoardRender.updateGameBoard(gameBoard, playResponse);
    }

    public GameBoard markTile(final GameBoard gameBoard, final Integer row, final Integer column, final String mark) {
        minesweeperClient.mark(gameBoard.getGameId(), MarkRequest.flagBuilder(row, column));
        var board = gameBoard.getBoard();
        if ("".equals(mark)) {
            board.get(row).set(column, "");
            gameBoard.setMineAmount(gameBoard.getMineAmount() + 1);
        } else if ("FLAG".equals(mark)) {
            board.get(row).set(column, "F");
            gameBoard.setMineAmount(gameBoard.getMineAmount() - 1);
        }

        return gameBoard;
    }

    private boolean isValidMovement(final GameBoard gameBoard, final Integer row, final Integer column) {
        var tileValue = gameBoard.getBoard().get(row).get(column);
        return tileValue.isEmpty() || "F".equals(tileValue);
    }

    public GameResponse getGame(final Integer gameId) {
        return minesweeperClient.getGame(gameId);
    }

}
