# xkcd-pass

xkcd-pass is a password generator to be able to generate [XKCD style passwords](https://xkcd.com/936/) in any language using the words from [Letterpress](http://www.atebits.com/letterpress/).

# Installation

```
go install github.com/dmartzol/xkcd-pass/cmd/xkcd-pass@latest
```

# Usage

```
Usage: xkcd-pass [OPTIONS]
  -M int
    	max word length (default 5)
  -c int
    	number of words to use (default 4)
  -d string
    	path to file with dictionary of words
  -m int
    	min word length (default 2)
  -s string
    	separator (default "-")
  -v	logs password strength information on screen
```

# Examples

```
xkcd-pass -d ~/code/xkcd-pass/wordlists/es.txt 
casta-giben-crino-ver⏎ 

xkcd-pass -d ~/code/xkcd-pass/wordlists/es.txt -v
paxte-apuno-canse-dure
all words 636599
sample space 13898
entropy 55

xkcd-pass -d ~/code/xkcd-pass/wordlists/es.txt -m 4 -M 10 -v
conciliare-agrupable-atusadora-tañere
all words 636599
sample space 362935
entropy 74

xkcd-pass -d ~/code/xkcd-pass/wordlists/es.txt | pbcopy
```
