FROM golang:latest
COPY /statics ./
COPY go.mod ./
COPY hangman.html ./
COPY home.html ./
COPY webmotif.go ./
COPY words.txt ./
ENTRYPOINT ["go", "run", "webmotif.go", "words.txt"]