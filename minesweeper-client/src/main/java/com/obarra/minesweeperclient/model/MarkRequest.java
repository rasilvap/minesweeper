package com.obarra.minesweeperclient.model;

import java.util.Objects;

public class MarkRequest {
    private Integer row;
    private Integer column;
    private String mark;

    public static MarkRequest flagBuilder(final Integer row, final Integer column) {
        final var makeRequest = new MarkRequest();
        makeRequest.setRow(row);
        makeRequest.setColumn(column);
        makeRequest.setMark("FLAG");
        return makeRequest;
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

    public String getMark() {
        return mark;
    }

    public void setMark(String mark) {
        this.mark = mark;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        MarkRequest that = (MarkRequest) o;
        return Objects.equals(row, that.row) &&
                Objects.equals(column, that.column) &&
                Objects.equals(mark, that.mark);
    }

    @Override
    public int hashCode() {
        return Objects.hash(row, column, mark);
    }

    @Override
    public String toString() {
        return "MarkRequest{" +
                "row=" + row +
                ", column=" + column +
                ", mark='" + mark + '\'' +
                '}';
    }
}
