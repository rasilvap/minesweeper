package com.obarra.minesweeperclient.service;

import com.obarra.minesweeperclient.client.MinesweeperClient;
import com.obarra.minesweeperclient.enums.StateGameEnum;
import com.obarra.minesweeperclient.model.GameBoard;
import com.obarra.minesweeperclient.model.GameRequest;
import com.obarra.minesweeperclient.model.GameResponse;
import com.obarra.minesweeperclient.model.MarkRequest;
import com.obarra.minesweeperclient.model.PlayRequest;
import com.obarra.minesweeperclient.utils.GameBoardConst;
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

        if (gameBoard.getBoard().get(row).get(column).equals(GameBoardConst.FLAG_TILE)) {
            return markTile(gameBoard, row, column, GameBoardConst.PLAY_MOVEMENT);
        }

        final var playResponse = minesweeperClient.play(gameBoard.getGameId(), PlayRequest.of(row, column));
        return gameBoardRender.updateGameBoard(gameBoard, playResponse);
    }

    public GameBoard markTile(final GameBoard gameBoard, final Integer row, final Integer column, final String mark) {
        minesweeperClient.mark(gameBoard.getGameId(), MarkRequest.flagBuilder(row, column));
        var board = gameBoard.getBoard();
        if (GameBoardConst.PLAY_MOVEMENT.equals(mark)) {
            board.get(row).set(column, GameBoardConst.COVERED_NOT_PLAYED);
            gameBoard.setMineAmount(gameBoard.getMineAmount() + 1);
        } else if (GameBoardConst.FLAG_MOVEMENT.equals(mark)) {
            board.get(row).set(column, GameBoardConst.FLAG_TILE);
            gameBoard.setMineAmount(gameBoard.getMineAmount() - 1);
        }

        return gameBoard;
    }

    private boolean isValidMovement(final GameBoard gameBoard, final Integer row, final Integer column) {
        var tileValue = gameBoard.getBoard().get(row).get(column);
        return tileValue.isEmpty() || GameBoardConst.FLAG_TILE.equals(tileValue);
    }

    public GameResponse getGame(final Integer gameId) {
        return minesweeperClient.getGame(gameId);
    }

}
