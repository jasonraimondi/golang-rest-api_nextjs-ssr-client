FROM golang:1.12
ENV GO111MODULE on
ENV API_PATH /go/src/git.jasonraimondi.com/jason/learn-with-tests
WORKDIR $API_PATH
COPY go.mod go.sum $API_PATH/
RUN go mod download

COPY . $API_PATH/

#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app src/main.go
#RUN chmod +x app
CMD ["go", "test", "./..."]

#FROM scratch
#LABEL maintainer="Jason Raimondi <jason@raimondi.us>"
#COPY templates /templates
#COPY --from=0 /go/src/git.jasonraimondi.com/jason/goserver/app /bin/app
#ENTRYPOINT ["app"]