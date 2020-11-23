package com.obarra.minesweeperclient.model;

public class PlayResponse {
    private String stateGame;
    private GameDTO game;

    public String getStateGame() {
        return stateGame;
    }

    public void setStateGame(String stateGame) {
        this.stateGame = stateGame;
    }

    public GameDTO getGame() {
        return game;
    }

    public void setGame(GameDTO game) {
        this.game = game;
    }

    @Override
    public String toString() {
        return "PlayResponse{" +
                "stateGame='" + stateGame + '\'' +
                ", game=" + game +
                '}';
    }
}
