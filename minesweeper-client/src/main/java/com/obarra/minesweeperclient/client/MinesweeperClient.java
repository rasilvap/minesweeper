package com.obarra.minesweeperclient.client;

import com.obarra.minesweeperclient.model.GameRequest;
import com.obarra.minesweeperclient.model.GameResponse;
import com.obarra.minesweeperclient.model.MarkRequest;
import com.obarra.minesweeperclient.model.PlayRequest;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

@FeignClient(value = "minesweeperClient", url = "http://localhost:5000/api")
public interface MinesweeperClient {

    @PostMapping("/games")
    GameResponse create(@RequestBody GameRequest gameRequest);

    @GetMapping("/games/{id}")
    GameResponse getGame(@PathVariable Integer id);

    @PostMapping("/games/{id}/mark")
    void mark(@PathVariable Integer id, @RequestBody MarkRequest markRequest);

    @PostMapping("/games/{id}/play")
    Object play(@PathVariable Integer id, @RequestBody PlayRequest playRequest);
}
