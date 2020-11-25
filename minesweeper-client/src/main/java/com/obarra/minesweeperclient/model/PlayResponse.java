package com.obarra.minesweeperclient.model;

import java.util.Objects;

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
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        PlayResponse that = (PlayResponse) o;
        return Objects.equals(stateGame, that.stateGame) &&
                Objects.equals(game, that.game);
    }

    @Override
    public int hashCode() {
        return Objects.hash(stateGame, game);
    }

    @Override
    public String toString() {
        return "PlayResponse{" +
                "stateGame='" + stateGame + '\'' +
                ", game=" + game +
                '}';
    }
}
