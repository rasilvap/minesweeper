package com.obarra.minesweeperclient.service;

import com.obarra.minesweeperclient.client.MinesweeperClient;
import com.obarra.minesweeperclient.model.*;
import com.obarra.minesweeperclient.utils.StateGameEnum;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class GameBoardService {

    private final MinesweeperClient minesweeperClient;
    private final GameBoardRender gameBoardRender;

    @Autowired
    public GameBoardService(final MinesweeperClient minesweeperClient, final GameBoardRender gameBoardRender) {
        this.minesweeperClient = minesweeperClient;
        this.gameBoardRender = gameBoardRender;
    }

    public GameBoard createGameBoard(final Integer rows, final Integer columns, final Integer mineAmount) {
        final var gameResponse = minesweeperClient.create(GameRequest.of(rows, columns, mineAmount));
        final var boardGame = new GameBoard();
        boardGame.setState(StateGameEnum.NEW);
        boardGame.setMineAmount(mineAmount);
        boardGame.setGameId(gameResponse.getGameId());
        System.out.println(gameResponse.getGameId());

        final var board = gameBoardRender.generateEmptyBoard(rows, columns);
        boardGame.setBoard(board);

        return boardGame;
    }

    public GameBoard playMovement(final GameBoard gameBoard, final Integer row, final Integer column) {
        final var playResponse = minesweeperClient.play(gameBoard.getGameId(), PlayRequest.of(row, column));
        System.out.println(playResponse);

        return gameBoardRender.updateGameBoard(gameBoard, playResponse);
    }

    public GameBoard markTile(final GameBoard gameBoard, final Integer row, final Integer column) {
        minesweeperClient.mark(gameBoard.getGameId(), MarkRequest.flagBuilder(row, column));
        System.out.println("Marking.....");
        var board = gameBoard.getBoard();
        board.get(row).set(column, "F");
        gameBoard.setMineAmount(gameBoard.getMineAmount() -1);
        return gameBoard;
    }

    public GameResponse getGame(final Integer gameId) {
        return minesweeperClient.getGame(gameId);
    }

}
