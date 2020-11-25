package com.obarra.minesweeperclient.doubletest;

import java.util.Arrays;
import java.util.List;

public class GameBoardDT {
    protected Integer rows = 3;
    protected Integer columns = 3;
    protected Integer mineAmount = 3;

    protected Integer rowOne = 1;
    protected Integer columnOne = 1;

    protected Integer rowZero = 0;
    protected Integer columnZero = 0;

    protected Integer gameId = 0;
    protected List<List<String>> boardFront = Arrays.asList(Arrays.asList("", "", ""), Arrays.asList("", "", ""), Arrays.asList("", "", ""));


}
