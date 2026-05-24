# tmux

install 

```bash
sudo apt install tmux
```

run 

```
tmux
[run-commands]
ctrl-b > d
tmux list-sessions
tmux attach
```

naming sessions

```
tmux
[run-commands]
ctrl-b > $
[rename-session]
ctrl-b > d
tmux list-sessions
tmux attach-session -t [session-name]
```

split screen 

```
tmux 
ctrl-b > % or "
ctrl-b > o
man tmux
```