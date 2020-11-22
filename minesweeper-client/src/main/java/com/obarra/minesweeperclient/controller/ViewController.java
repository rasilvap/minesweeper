package com.obarra.minesweeperclient.controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

import java.util.ArrayList;
import java.util.List;

@Controller
public class ViewController {



    @GetMapping("/index")
    public String index(Model model) {
        var board = new ArrayList<List<String>>();
        board.add(List.of("Omar", "barradev", "dev"));
        board.add(List.of("Tesla", "dev", "dev"));
        board.add(List.of("Newton", "newdev", "dev"));
        board.add(List.of("Newton", "newdev", "dev"));
        board.add(List.of("Newton", "newdev", "dev"));
        board.add(List.of("Newton", "newdev", "dev"));

        model.addAttribute("board", board);
        return "index";
    }
}
