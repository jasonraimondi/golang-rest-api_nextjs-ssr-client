FROM golang:1.13-rc-alpine as builder
RUN apk add --update git
ENV GO111MODULE on
ENV API_PATH /go/src/git.jasonraimondi.com/jason/jasontest
WORKDIR $API_PATH
COPY ./go.* $API_PATH/
RUN go mod download
COPY ./app/ $API_PATH/app/
COPY ./server/ $API_PATH/server/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/app ./server/main.go
RUN chmod +x app

FROM alpine
LABEL maintainer="Jason Raimondi <jason@raimondi.us>"
COPY --from=builder /go/src/git.jasonraimondi.com/jason/jasontest/dist/app /bin/app
CMD ["/bin/app"]
