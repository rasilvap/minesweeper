package com.obarra.minesweeperclient.model;

import java.util.Objects;

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
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        PlayRequest that = (PlayRequest) o;
        return Objects.equals(row, that.row) &&
                Objects.equals(column, that.column);
    }

    @Override
    public int hashCode() {
        return Objects.hash(row, column);
    }

    @Override
    public String toString() {
        return "PlayRequest{" +
                "row=" + row +
                ", column=" + column +
                '}';
    }
}
