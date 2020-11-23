package com.obarra.minesweeperclient.service;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.stream.Collectors;

class GameBoardRenderTest {

    private GameBoardRender gameBoardRender;

    @BeforeEach
    void setUp() {
        gameBoardRender = new GameBoardRender();
    }

    @AfterEach
    void tearDown() {
    }

    @Test
    void generateEmptyBoardWhenIsTwentyFourPositions() {
        var result = gameBoardRender.generateEmptyBoard(3, 8);
        Assertions.assertEquals(3, result.size());
        Assertions.assertEquals(8, result.get(0).size());
        var twentyFourPositions = ",,,,,,,,,,,,,,,,,,,,,,,";
        Assertions.assertEquals(twentyFourPositions,
                result.stream().map(s -> String.join(",", s)).collect(Collectors.joining(",")));

    }

    @Test
    void generateEmptyBoardWhenIsNinePositions() {
        var result = gameBoardRender.generateEmptyBoard(3, 3);
        Assertions.assertEquals(3, result.size());
        Assertions.assertEquals(3, result.get(0).size());
        var ninePositions = ",,,,,,,,";
        Assertions.assertEquals(ninePositions,
                result.stream().map(s -> String.join(",", s)).collect(Collectors.joining(",")));

    }

    @Test
    void updateGameBoard() {
    }
}