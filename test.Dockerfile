FROM golang:1.12
ENV GO111MODULE on
ENV API_PATH /go/src/git.jasonraimondi.com/jason/jasontest
WORKDIR $API_PATH
COPY go.mod go.sum $API_PATH/
RUN go mod download
COPY . $API_PATH/
CMD ["go", "test", "./..."]
