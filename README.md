# tttm-rango

## Description
[tttm-rango](https://github.com/perkzen/tttm-rango) is a random player server for
the [tttm](https://github.com/ogrodje/tttm) tic-tac-toe challenge.

## Usage

```bash
    # install dependencies
    go mod download
    # run the server
    go run cmd/server/main.go
    # or
    make run
```

## Player Server Protocol

Players need to implement an HTTP server that needs to have one endpoint.

### `GET /move`

The game server will pass the following URL `query parameters` to the player server.

- `gid` - `UUID` that represents the given game ID.
    - It can be used on player servers to stick together individual games.
- `size` - A number - size - of tic-tac-toe grid.
    - By default, the size is set to `3`, representing the grid of size `3x3`. To win, one has to have 3 symbols in
      either row/column/diagonal.
- `playing` - A symbol that the player server needs to play. Possible values are `X` or `O`
- `moves` - A string that represents the previous moves.
    - Moves are separated by `_` and positions by `-`
    - Example: `X-1-1_O-0-0` means that the `X` symbol was at location `1,1` (centre of grid) and `O` at `0,0` (top-left
      corner of the grid)

### Body

The content of the successful response (HTTP status 200) needs to be a string that should have the following structure:

```
Move:X-2-2
```

The player server replied by placing the symbol `X` in position `2,2` in the grid—in this case, the very bottom right.

If, for some reason, the server is unable to reply or parse, please use the following payload structure:

```
Error:Sorry. Can't do it bro.
```

