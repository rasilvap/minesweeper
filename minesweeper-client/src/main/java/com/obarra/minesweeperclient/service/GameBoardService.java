package com.obarra.minesweeperclient.service;

import com.obarra.minesweeperclient.client.MinesweeperClient;
import com.obarra.minesweeperclient.model.GameRequest;
import com.obarra.minesweeperclient.model.GameResponse;
import com.obarra.minesweeperclient.model.MarkRequest;
import com.obarra.minesweeperclient.model.PlayRequest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class GameBoardService {

    private final MinesweeperClient minesweeperService;

    @Autowired
    public GameBoardService(final MinesweeperClient minesweeperService) {
        this.minesweeperService = minesweeperService;
    }

    public Integer createGame(final Integer rows, final Integer columns, final Integer mineAmount) {
        var gameResponse = minesweeperService.create(GameRequest.of(rows, columns, mineAmount));
        return gameResponse.getGameId();
    }

    public GameResponse getGame(final Integer gameId) {
        return minesweeperService.getGame(gameId);
    }

    public void markTile(final Integer gameId, final Integer row, final Integer column) {
        minesweeperService.mark(gameId, MarkRequest.flagBuilder(row, column));
    }

    public void playMovement(final Integer gameId, final Integer row, final Integer column) {
        var res = minesweeperService.play(gameId, PlayRequest.of(row, column));
        System.out.println(res);
    }

}
