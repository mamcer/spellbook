neovim

| Command            | Description                             |
| ------------------ | --------------------------------------- |
| `:w`               | Save current file                       |
| `:q`               | Quit current window                     |
| `:wq`              | Save and quit                           |
| `:q!`              | Quit without saving                     |
| `:e filename`      | Open a file                             |
| `:vsp filename`    | Open file in vertical split             |
| `:sp filename`     | Open file in horizontal split           |
| `:tabnew filename` | Open file in a new tab                  |
| `:bn`              | Go to next buffer                       |
| `:bp`              | Go to previous buffer                   |
| `:ls`              | List all buffers                        |
| `:bd`              | Delete current buffer                   |
| `:noh`             | Remove search highlights                |
| `/pattern`         | Search forward for pattern              |
| `?pattern`         | Search backward for pattern             |
| `n`                | Repeat search in same direction         |
| `N`                | Repeat search in opposite direction     |
| `gg`               | Go to beginning of file                 |
| `G`                | Go to end of file                       |
| `0`                | Go to beginning of line                 |
| `^`                | Go to first non-blank character of line |
| `$`                | Go to end of line                       |
| `:`                | Enter command mode                      |
| `i`                | Enter insert mode                       |
| `a`                | Append after cursor                     |
| `o`                | Open new line below cursor              |
| `x`                | Delete character under cursor           |
| `dd`               | Delete current line                     |
| `yy`               | Copy current line                       |
| `p`                | Paste after cursor                      |
| `u`                | Undo                                    |
| `Ctrl-r`           | Redo                                    |

neovim golang

| Command / Key       | Description                                             |
| ------------------- | ------------------------------------------------------- |
| `gd`                | Go to definition of symbol                              |
| `K`                 | Show hover/documentation for symbol                     |
| `<leader>rn`        | Rename symbol under cursor                              |
| `<leader>ca`        | Show code actions (fixes, suggestions)                  |
| `:GoBuild`          | Build current package                                   |
| `:GoRun`            | Run current Go file                                     |
| `:GoTest`           | Run tests in current package                            |
| `:GoFmt`            | Format current file with `gofmt`                        |
| `:GoImports`        | Organize imports                                        |
| `:GoDoc`            | Show documentation for package or symbol                |
| `:GoLint`           | Run linter on current package                           |
| `:GoUpdateBinaries` | Update Go tools used by Go.nvim                         |
| Saving a `.go` file | Will automatically format and fix imports if configured |

```
Use gd and K often to navigate code quickly.
Use splits (:vsp) and tabs (:tabnew) for multiple files.
Format and fix imports automatically on save (:w) to keep code clean.
Run tests often: :GoTest or :!go test ./... from terminal.

:term terminal dentro de nvim
Modo de terminal: i para escribir, Ctrl-\ Ctrl-n para salir al modo normal.

file explorer
:Ex
j/k → mover arriba/abajo
Enter → abrir archivo
- → subir un nivel


Tips para workflow Go en este setup
Terminal
Abre terminal integrado dentro de Neovim:
:term
Escribe comandos Go (go run ., go test, dlv debug)
Salir al modo normal: Ctrl-\ Ctrl-n
File Explorer
Toggle Nvim-Tree: <leader>e
Navega carpetas con j/k y abre archivos con Enter
Crear archivo: a, Borrar: d, Renombrar: r
Fuzzy Finder
Find file: <leader>ff
Search text in project: <leader>fg
List buffers: <leader>fb
Help tags: <leader>fh
Go/LSP Shortcuts
Go to definition: gd
Hover docs: K
Rename symbol: <leader>rn
Code actions: <leader>ca
Save .go file → autoformat + organize imports
Basic navigation
Splits: :vsp filename, :sp filename
Tabs: :tabnew filename
Switch buffers: :bn, :bp
Move between windows: Ctrl-w h/j/k/l
```