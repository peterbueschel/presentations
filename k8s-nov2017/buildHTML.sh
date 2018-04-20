#!/usr/bin/env bash
FILENAME=sure-but
pandoc -t revealjs \
    -V revealjs-url="$GOPATH/src/github.com/hakimel/reveal.js" \
    -V theme=moon \
    -V transition=fade \
    -V slideNumber=true \
    --include-in-header=custom.css \
    --section-divs \
    ${FILENAME}.md -o index.html
