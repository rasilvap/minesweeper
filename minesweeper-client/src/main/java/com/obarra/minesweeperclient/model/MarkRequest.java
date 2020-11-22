package com.obarra.minesweeperclient.model;

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
    public String toString() {
        return "MarkRequest{" +
                "row=" + row +
                ", column=" + column +
                ", mark='" + mark + '\'' +
                '}';
    }
}
