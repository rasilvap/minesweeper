package com.obarra.minesweeperclient.service;

import com.obarra.minesweeperclient.client.MinesweeperClient;
import com.obarra.minesweeperclient.doubletest.GameBoardDT;
import com.obarra.minesweeperclient.model.*;
import com.obarra.minesweeperclient.utils.StateGameEnum;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.mockito.Mockito.*;

class GameBoardServiceTest extends GameBoardDT {
    private GameBoardRender gameBoardRender;
    private MinesweeperClient minesweeperClient;

    private GameBoardService gameBoardService;

    @BeforeEach
    void setUp() {
        minesweeperClient = mock(MinesweeperClient.class);
        gameBoardRender = mock(GameBoardRender.class);
        gameBoardService = new GameBoardService(minesweeperClient, gameBoardRender);
    }

    @AfterEach
    void tearDown() {
    }

    @Test
    void createGameBoard() {
        GameResponse gameResponse = new GameResponse();
        gameResponse.setGameId(gameId);
        gameResponse.setRows(rows);
        gameResponse.setColumns(columns);
        gameResponse.setMineAmount(mineAmount);
        GameRequest gameRequest = GameRequest.of(rows, columns, mineAmount);

        doReturn(gameResponse).when(minesweeperClient).create(gameRequest);

        doReturn(boardFront).when(gameBoardRender).generateEmptyBoard(rows, columns);

        //execute
        var result = gameBoardService.createGameBoard(rows, columns, mineAmount);

        assertEquals(gameId, result.getGameId());
        assertEquals(mineAmount, result.getMineAmount());
        assertEquals(3, result.getBoard().size());
        assertEquals(3, result.getBoard().get(0).size());
    }

    @Test
    void playMovementShouldNoInvokeAPIWhenTileWasPlayedOK() {
        var boardFront = Arrays.asList(Arrays.asList("C", "", ""), Arrays.asList("", "", ""), Arrays.asList("", "", ""));
        var gameBoard = new GameBoard();
        gameBoard.setBoard(boardFront);
        gameBoard.setBoard(boardFront);
        gameBoard.setGameId(gameId);
        gameBoard.setState(StateGameEnum.RUNNING);
        gameBoard.setMineAmount(1);

        //execute
        var result = gameBoardService.playMovement(gameBoard, rowZero, columnZero);

        GameBoard expected = new GameBoard();
        expected.setBoard(boardFront);
        expected.setGameId(gameId);
        expected.setState(StateGameEnum.RUNNING);
        expected.setMineAmount(1);

        assertEquals(expected, result);
    }

    @Test
    void playMovementShouldInvokeAPIWhenTileWasIsNewPlayMovement() {
        PlayRequest playRequest = PlayRequest.of(rowOne, rowOne);
        PlayResponse playResponse = new PlayResponse();
        playResponse.setStateGame(StateGameEnum.WON.name());

        GameDTO gameDTO = new GameDTO();
        TileDTO[][] board = {{new TileDTO(), new TileDTO(), new TileDTO()},
                {new TileDTO(), new TileDTO(), new TileDTO()},
                {new TileDTO(), new TileDTO(), new TileDTO()}};

        gameDTO.setBoard(board);
        gameDTO.setRows(rows);
        gameDTO.setColumns(columns);
        gameDTO.setFlagAmount(1);
        playResponse.setGame(gameDTO);
        doReturn(playResponse).when(minesweeperClient).play(gameId, playRequest);
        var gameBoard = new GameBoard();
        gameBoard.setBoard(boardFront);
        gameBoard.setBoard(boardFront);
        gameBoard.setGameId(gameId);
        gameBoard.setState(StateGameEnum.RUNNING);
        gameBoard.setMineAmount(1);

        var gameBoardFrontResult = new GameBoard();
        gameBoardFrontResult.setBoard(boardFront);
        gameBoardFrontResult.setGameId(gameId);
        gameBoardFrontResult.setState(StateGameEnum.WON);
        gameBoardFrontResult.setMineAmount(1);

        doReturn(gameBoardFrontResult).when(gameBoardRender).updateGameBoard(gameBoard, playResponse);

        //execute
        var result = gameBoardService.playMovement(gameBoard, rowOne, columnOne);

        GameBoard expected = new GameBoard();
        expected.setBoard(boardFront);
        expected.setGameId(gameId);
        expected.setState(StateGameEnum.WON);
        expected.setMineAmount(1);

        assertEquals(expected, result);
    }

    @Test
    void playMovementShouldInvokeAPIWhenTileWasMarked() {
        MarkRequest markRequest = MarkRequest.flagBuilder(rowZero, columnZero);
        doNothing().when(minesweeperClient).mark(gameId, markRequest);

        var boardFront = Arrays.asList(Arrays.asList("F", "", ""), Arrays.asList("", "", ""), Arrays.asList("", "", ""));

        var gameBoard = new GameBoard();
        gameBoard.setBoard(boardFront);
        gameBoard.setGameId(gameId);
        gameBoard.setState(StateGameEnum.RUNNING);
        gameBoard.setMineAmount(0);

        //execute
        var result = gameBoardService.playMovement(gameBoard, rowZero, columnZero);

        GameBoard expected = new GameBoard();
        expected.setBoard(boardFront);
        expected.setGameId(gameId);
        expected.setState(StateGameEnum.RUNNING);
        expected.setMineAmount(1);

        assertEquals(expected, result);
        verify(minesweeperClient, atLeastOnce()).mark(gameId, markRequest);
    }


    @Test
    void markTileShouldDecrementAndInvokeAPIWhenMarkIsFlag() {
        MarkRequest markRequest = MarkRequest.flagBuilder(rowZero, columnZero);
        doNothing().when(minesweeperClient).mark(gameId, markRequest);

        var gameBoard = new GameBoard();
        gameBoard.setBoard(boardFront);
        gameBoard.setGameId(gameId);
        gameBoard.setState(StateGameEnum.RUNNING);
        gameBoard.setMineAmount(1);


        //execute
        var result = gameBoardService.markTile(gameBoard, rowZero, columnZero, "FLAG");

        GameBoard expected = new GameBoard();
        expected.setBoard(boardFront);
        expected.setGameId(gameId);
        expected.setState(StateGameEnum.RUNNING);
        expected.setMineAmount(0);

        assertEquals(expected, result);
        verify(minesweeperClient, atLeastOnce()).mark(gameId, markRequest);
    }

    @Test
    void markTileShouldInclementAndInvokeAPIWhenMarkIsEmpty() {
        MarkRequest markRequest = MarkRequest.flagBuilder(rowZero, columnZero);
        doNothing().when(minesweeperClient).mark(gameId, markRequest);

        var gameBoard = new GameBoard();
        gameBoard.setBoard(boardFront);
        gameBoard.setGameId(gameId);
        gameBoard.setState(StateGameEnum.RUNNING);
        gameBoard.setMineAmount(1);


        //execute
        var result = gameBoardService.markTile(gameBoard, rowZero, columnZero, "FLAG");

        GameBoard expected = new GameBoard();
        expected.setBoard(boardFront);
        expected.setGameId(gameId);
        expected.setState(StateGameEnum.RUNNING);
        expected.setMineAmount(0);

        assertEquals(expected, result);
        verify(minesweeperClient, atLeastOnce()).mark(gameId, markRequest);
    }

}