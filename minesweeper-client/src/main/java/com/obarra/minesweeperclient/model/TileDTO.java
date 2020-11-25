package com.obarra.minesweeperclient.model;

import java.util.Objects;

public class TileDTO {
    private String state;
    private Integer row;
    private Integer column;
    private Integer surroundingMineCount;
    private Boolean mine;
    private Integer valueTest;

    public String getState() {
        return state;
    }

    public void setState(String state) {
        this.state = state;
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

    public Integer getSurroundingMineCount() {
        return surroundingMineCount;
    }

    public void setSurroundingMineCount(Integer surroundingMineCount) {
        this.surroundingMineCount = surroundingMineCount;
    }

    public Boolean getMine() {
        return mine;
    }

    public void setMine(Boolean mine) {
        this.mine = mine;
    }

    public Integer getValueTest() {
        return valueTest;
    }

    public void setValueTest(Integer valueTest) {
        this.valueTest = valueTest;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        TileDTO tileDTO = (TileDTO) o;
        return Objects.equals(state, tileDTO.state) &&
                Objects.equals(row, tileDTO.row) &&
                Objects.equals(column, tileDTO.column) &&
                Objects.equals(surroundingMineCount, tileDTO.surroundingMineCount) &&
                Objects.equals(mine, tileDTO.mine) &&
                Objects.equals(valueTest, tileDTO.valueTest);
    }

    @Override
    public int hashCode() {
        return Objects.hash(state, row, column, surroundingMineCount, mine, valueTest);
    }

    @Override
    public String toString() {
        return "TileDTO{" +
                "state='" + state + '\'' +
                ", row=" + row +
                ", column=" + column +
                ", surroundingMineCount=" + surroundingMineCount +
                ", mine=" + mine +
                ", valueTest=" + valueTest +
                '}';
    }
}
