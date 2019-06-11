FROM golang:1.12 as builder
ENV GO111MODULE on
ENV API_PATH /go/src/git.jasonraimondi.com/jason/jasontest
WORKDIR $API_PATH
COPY go.mod go.sum $API_PATH/
RUN go mod download
COPY . $API_PATH/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app src/main.go
RUN chmod +x app

FROM scratch
LABEL maintainer="Jason Raimondi <jason@raimondi.us>"
COPY --from=builder /go/src/git.jasonraimondi.com/jason/jasontest/app /bin/app
ENTRYPOINT ["app"]
