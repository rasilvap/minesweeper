package com.obarra.minesweeperclient.controller;

import com.obarra.minesweeperclient.model.GameBoard;
import com.obarra.minesweeperclient.service.GameBoardService;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.SessionAttributes;

import javax.websocket.server.PathParam;
import java.util.ArrayList;

@Controller
@SessionAttributes("boardGame")
@RequestMapping("/mineswipeer")
public class GameBoardController {

    private final GameBoardService gameBoardService;

    public GameBoardController(final GameBoardService gameBoardService) {
        this.gameBoardService = gameBoardService;
    }

    @GetMapping("/index")
    public String index(final Model model, final @ModelAttribute("boardGame") GameBoard gameBoard) {
        final var newGameBoard = gameBoardService.createGameBoard(3, 8, 1);
        gameBoard.setGameId(newGameBoard.getGameId());
        gameBoard.setBoard(newGameBoard.getBoard());
        gameBoard.setState(newGameBoard.getState());

        model.addAttribute("board", gameBoard.getBoard());
        model.addAttribute("state", gameBoard.getState());
        return "index";
    }

    @GetMapping("/play")
    public String play(@PathParam("row") Integer row,
                       @PathParam("column") Integer column,
                       Model model,
                       @ModelAttribute("boardGame") GameBoard gameBoard) {
        System.out.println(row + " play " + column);

        final var updatedGameBoard = gameBoardService.playMovement(gameBoard, row, column);
        gameBoard.setGameId(updatedGameBoard.getGameId());
        gameBoard.setState(updatedGameBoard.getState());
        gameBoard.setBoard(updatedGameBoard.getBoard());
        gameBoard.setState(updatedGameBoard.getState());


        model.addAttribute("board", gameBoard.getBoard());
        model.addAttribute("state", gameBoard.getState());

        return "index";
    }

    @GetMapping("/mark")
    public String mark(@PathParam("row") Integer row,
                       @PathParam("column") Integer column,
                       Model model,
                       @ModelAttribute("boardGame") GameBoard gameBoard) {
        System.out.println(row + "mark " + column);
        final GameBoard updatedGameBoard = gameBoardService.markTile(gameBoard, row, column);
        gameBoard.setGameId(updatedGameBoard.getGameId());
        gameBoard.setState(updatedGameBoard.getState());
        gameBoard.setBoard(updatedGameBoard.getBoard());
        gameBoard.setState(updatedGameBoard.getState());


        model.addAttribute("board", gameBoard.getBoard());
        model.addAttribute("state", gameBoard.getState());

        return "index";
    }

    @ModelAttribute("boardGame")
    public GameBoard boardGame() {
        final var boardGame = new GameBoard();
        boardGame.setBoard(new ArrayList<>());
        return boardGame;
    }
}
