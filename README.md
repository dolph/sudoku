# Sudoku

Renders a Sudoku board by representing the viable choices for any given square.
For example, this is the equivalent of a "blank" Sudoku board:

```bash
┏━━━━━━━┯━━━━━━━┯━━━━━━━┳━━━━━━━┯━━━━━━━┯━━━━━━━┳━━━━━━━┯━━━━━━━┯━━━━━━━┓
┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃
┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃
┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃
┠───────┼───────┼───────╂───────┼───────┼───────╂───────┼───────┼───────┨
┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃
┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃
┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃
┠───────┼───────┼───────╂───────┼───────┼───────╂───────┼───────┼───────┨
┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃
┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃
┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃
┣━━━━━━━┿━━━━━━━┿━━━━━━━╋━━━━━━━┿━━━━━━━┿━━━━━━━╋━━━━━━━┿━━━━━━━┿━━━━━━━┫
┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃
┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃
┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃
┠───────┼───────┼───────╂───────┼───────┼───────╂───────┼───────┼───────┨
┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃
┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃
┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃
┠───────┼───────┼───────╂───────┼───────┼───────╂───────┼───────┼───────┨
┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃
┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃
┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃
┣━━━━━━━┿━━━━━━━┿━━━━━━━╋━━━━━━━┿━━━━━━━┿━━━━━━━╋━━━━━━━┿━━━━━━━┿━━━━━━━┫
┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃
┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃
┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃
┠───────┼───────┼───────╂───────┼───────┼───────╂───────┼───────┼───────┨
┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃
┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃
┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃
┠───────┼───────┼───────╂───────┼───────┼───────╂───────┼───────┼───────┨
┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃ 1 2 3 │ 1 2 3 │ 1 2 3 ┃
┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃ 4 5 6 │ 4 5 6 │ 4 5 6 ┃
┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃ 7 8 9 │ 7 8 9 │ 7 8 9 ┃
┗━━━━━━━┷━━━━━━━┷━━━━━━━┻━━━━━━━┷━━━━━━━┷━━━━━━━┻━━━━━━━┷━━━━━━━┷━━━━━━━┛
```