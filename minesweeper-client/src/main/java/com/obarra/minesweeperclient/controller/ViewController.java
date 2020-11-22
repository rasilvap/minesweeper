package com.obarra.minesweeperclient.controller;

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
public class ViewController {

    @GetMapping("/index")
    public String index(Model model, @ModelAttribute("boardGame") BoardGame boardGame) {
        var board = boardGame.getBoard();
        board.add(List.of("Omar", "barradev", "dev"));
        board.add(List.of("Tesla", "dev", "dev"));
        board.add(List.of("Newton", "newdev", "dev"));
        board.add(List.of("Newton", "newdev", "dev"));
        board.add(List.of("Newton", "newdev", "dev"));
        board.add(List.of("Newton", "newdev", "dev"));

        model.addAttribute("board", board);
        return "index";
    }

    @GetMapping("/play")
    public String play(@PathParam("row") Integer row,
                       @PathParam("column") Integer column,
                       Model model,
                       @ModelAttribute("boardGame") BoardGame boardGame) {
        System.out.println(row + " play " + column);
        var board = boardGame.getBoard();
        board.add(List.of("Play", "merge", "new"));
        model.addAttribute("board", board);
        return "index";
    }

    @GetMapping("/mark")
    public String mark(@PathParam("row") Integer row,
                       @PathParam("column") Integer column,
                       Model model,
                       @ModelAttribute("boardGame") BoardGame boardGame) {
        System.out.println(row + "mark " + column);
        var board = boardGame.getBoard();
        board.add(List.of("Mark", "merge", "new"));
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
