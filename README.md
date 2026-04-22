# golf

My attempt to create a very very basic version of [lf](https://github.com/gokcehan/lf) in go (that is already written in go lol).

## Plan:

### explore how to recreate basic commands

- [ ] pwd
- [ ] ls
- [ ] cat
- [ ] touch
- [ ] mv
- [ ] mkdir
- [ ] rm
- [ ] cp

### explore how to draw tui in go

- [ ] try [tview](https://github.com/rivo/tview)
- [ ] try drawing 3 sections.

### planned desgin

```
~/projects/golf/ (pwd)
┌───────┬───────────┬──────────────────┐
│ ls .. │ ls $PWD   │    cat file1     │
│       │           │   ┌──────────┐   │
│ >golf │ folder1/  │   │ contents │   │
│  abkb │ folder2/  │   │    of    │   │
│       │ >file1    │   │  file1   │   │
│       │  file2    │   └──────────┘   │
└───────┴───────────┴──────────────────┘
| hjkl/←↓↑→ | cop[y] | [c]ut | [r]ename | [s]elect |
```
