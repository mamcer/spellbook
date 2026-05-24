# vim

[https://thevaluable.dev/vim-advanced](https://thevaluable.dev/vim-advanced)

## normal and insert modes

You can toogle between normal and insert modes with: esc (normal mode) I (insert mode) 

## basic movement

You can use the keys h, j, k and l

## word movement

e: moves to the end of the word  
b: moves to the begining of the word  
w: moves to the start of next word  

## number powered movements

e.g. 3w is the same as pressing w three times

## show line numbers

:set number

:set nonumber

## insert text repeatedly

30i- esc  
insert 30 times - 

## find character

Use f and F to move to the next or previous ocurrence of a word. e.g. fo finds next o

## go to matching parenthesis

% and shift-% to go to next or previous ( [ or {
	
## go to start/end of the line

0 to go to the beginning of the line  
\$ to go to the end of the line

## find word under cursor
 
"*" next occurrence # previous
 
## goto line
 
gg to go to the beginning of the file. G to go to the end. [number]G to go to a specific line
 
## search
 
You  can search by pressing /
 
n and N to navigate to next and previous results

## insert new line

o will insert a new line below the current line

O will insert a new line after the current line

## delete a character

x and X removes the character to the left and right respectively

## delete command

d is the delete command. You can combine it with w for example to delete the next word.

## repeat last command

with .

## visual mode

v to enter visual mode. You can select a text with e and delete it with d

## useful commands

:w to save

:q to quit

:q! to quit without saving

u for undo

ctrl+r for redo

:help to get some help 

## file history

	:browse oldfiles
> it will show a list of files and you can select by number

or..

you can directly type the last 10 files

	:'0

> 0 is the most useful as this is the previous file

## file explorer

You can directly type 

	:Explore 

to show the explorer window and select a file

or.. 

you can directly specify a file

	:e /home/mario/projects/stash/tech/vim.go

you can return to the explorer window with ctrl-6

you can also list files

	:e [ctrl-D]

## multiple cursors

```
ctrl-v
shift-i
esc
:x	
```	