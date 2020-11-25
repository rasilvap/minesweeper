package com.obarra.minesweeperclient.controller;

import com.obarra.minesweeperclient.model.GameBoard;
import com.obarra.minesweeperclient.model.GameContainer;
import com.obarra.minesweeperclient.service.GameBoardService;
import com.obarra.minesweeperclient.utils.GameBoardConst;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.SessionAttributes;

import javax.websocket.server.PathParam;

@Controller
@SessionAttributes("gameContainer")
public class GameBoardController {
    private final static Logger LOGGER = LoggerFactory.getLogger(GameBoardController.class);
    private final static Integer ROWS_DEFAULT = 25;
    private final static Integer COLUMNS_DEFAULT = 20;
    private final static Integer MINE_AMOUNT_DEFAULT = 15;

    private final GameBoardService gameBoardService;

    public GameBoardController(final GameBoardService gameBoardService) {
        this.gameBoardService = gameBoardService;
    }

    @GetMapping("/")
    public String index(final Model model, final @ModelAttribute("gameContainer") GameContainer gameContainer) {
        LOGGER.info("Init new game");
        final var gameBoard = gameBoardService
                .createGameBoard(ROWS_DEFAULT, COLUMNS_DEFAULT, MINE_AMOUNT_DEFAULT);

        gameContainer.setGameBoard(gameBoard);
        fillModel(model, gameContainer);
        return "index";
    }

    @GetMapping("/play")
    public String play(final @PathParam("row") Integer row,
                       final @PathParam("column") Integer column,
                       final Model model,
                       final @ModelAttribute("gameContainer") GameContainer gameContainer) {
        LOGGER.info(String.format("Playing row: %s, column: %s", row, column));

        final var gameBoard = gameBoardService.playMovement(gameContainer.getGameBoard(), row, column);
        gameContainer.setGameBoard(gameBoard);
        fillModel(model, gameContainer);
        return "index";
    }

    @GetMapping("/mark")
    public String mark(final @PathParam("row") Integer row,
                       final @PathParam("column") Integer column,
                       final Model model,
                       final @ModelAttribute("gameContainer") GameContainer gameContainer) {
        LOGGER.info(String.format("Marking row: %s, column: %s", row, column));
        final GameBoard gameBoard = gameBoardService.markTile(gameContainer.getGameBoard(), row, column, GameBoardConst.FLAG_MOVEMENT);
        gameContainer.setGameBoard(gameBoard);
        fillModel(model, gameContainer);
        return "index";
    }

    @ModelAttribute("gameContainer")
    public GameContainer gameContainer() {
        return new GameContainer();
    }

    private void fillModel(final Model model, final GameContainer gameContainer) {
        model.addAttribute("board", gameContainer.getGameBoard().getBoard());
        model.addAttribute("state", gameContainer.getGameBoard().getState());
        model.addAttribute("mineAmount", gameContainer.getGameBoard().getMineAmount());
    }
}
