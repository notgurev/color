# color

`color` is a simple command line wrapper for [fatih/color](https://github.com/fatih/color) that allows to colorize piped input using ANSI escape codes.

## Installation

Requires Go to be installed.

```
go install github.com/notgurev/color@latest
```

## Usage

```
echo "hello" | color [options]
```

Options (in any order):

- Colors: `red`, `green`, `blue`, `yellow`, `magenta`, `cyan`, `white`, `black`
- Hi Colors: `hired`, `higreen`, `hiblue`, `hiyellow`, `himagenta`, `hicyan`, `hiwhite`
- Background Colors: `bgred`, `bggreen`, `bgblue`, `bgyellow`, `bgmagenta`, `bgcyan`, `bgwhite`
- Hi Background Colors: `bghired`, `bghigreen`, `bghiblue`, `bghiyellow`, `bghimagenta`, `bghicyan`, `bghiwhite`
- Text Styles: `bold`, `underline`, `italic`, `faint`, `crossedout`, `blinkslow`, `blinkrapid`, `reverse`, `concealed`

Example:

```
echo "hello" | color red bold
```