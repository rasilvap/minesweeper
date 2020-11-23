package com.obarra.minesweeperclient.controller;

import com.obarra.minesweeperclient.model.BoardGame;
import com.obarra.minesweeperclient.service.GameBoardService;
import com.obarra.minesweeperclient.utils.GameBoardRenderUtil;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.SessionAttributes;

import javax.websocket.server.PathParam;
import java.util.ArrayList;
import java.util.List;

@Controller
@SessionAttributes("boardGame")
@RequestMapping("/mineswipeer")
public class GameBoardController {

    private GameBoardService gameBoardService;

    public GameBoardController(final GameBoardService gameBoardService) {
        this.gameBoardService = gameBoardService;
    }

    @GetMapping("/index")
    public String index(Model model, @ModelAttribute("boardGame") BoardGame boardGame) {
        var gameId = gameBoardService.createGame(3, 3, 1);
        boardGame.setGameId(gameId);
        System.out.println(gameId);

        //TODO use in other feature
        var res = gameBoardService.getGame(gameId);
        System.out.println(res);


        List<List<String>> newBoard = GameBoardRenderUtil.generateEmptyBoard(3, 3);
        boardGame.setBoard(newBoard);

        model.addAttribute("board", newBoard);
        return "index";
    }

    @GetMapping("/play")
    public String play(@PathParam("row") Integer row,
                       @PathParam("column") Integer column,
                       Model model,
                       @ModelAttribute("boardGame") BoardGame boardGame) {
        System.out.println(row + " play " + column);
        var playResponse =  gameBoardService.playMovement(boardGame.getGameId(), row, column);
        System.out.println(playResponse);
        var board = boardGame.getBoard();
        GameBoardRenderUtil.render(board, playResponse);

        //board.get(row).set(column, "C");

        model.addAttribute("board", board);
        return "index";
    }

    @GetMapping("/mark")
    public String mark(@PathParam("row") Integer row,
                       @PathParam("column") Integer column,
                       Model model,
                       @ModelAttribute("boardGame") BoardGame boardGame) {
        System.out.println(row + "mark " + column);
        gameBoardService.markTile(boardGame.getGameId(), row, column);

        var board = boardGame.getBoard();
        board.get(row).set(column, "F");

        model.addAttribute("board", board);
        return "index";
    }

    @ModelAttribute("boardGame")
    public BoardGame boardGame() {
        var boardGame = new BoardGame();
        boardGame.setBoard(new ArrayList<>());
        return boardGame;
    }
}
