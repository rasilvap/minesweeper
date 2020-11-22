package com.obarra.minesweeperclient.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class ViewController {

    @GetMapping("/play")
    public String index() {
        return "index";
    }
}
