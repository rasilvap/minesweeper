package com.obarra.minesweeperclient.model;

public class PlayRequest {
    private Integer row;
    private Integer column;

    public static PlayRequest of(final Integer row, final Integer column) {
        final var playRequest = new PlayRequest();
        playRequest.setRow(row);
        playRequest.setColumn(column);
        return playRequest;
    }

    public Integer getRow() {
        return row;
    }

    public void setRow(Integer row) {
        this.row = row;
    }

    public Integer getColumn() {
        return column;
    }

    public void setColumn(Integer column) {
        this.column = column;
    }

    @Override
    public String toString() {
        return "PlayRequest{" +
                "row=" + row +
                ", column=" + column +
                '}';
    }
}
