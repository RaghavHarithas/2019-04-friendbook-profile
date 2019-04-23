FROM golang:1.11 as builder

WORKDIR $GOPATH/src/github.com/rajch/2019-04-friendbook-profile

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN CGO_ENABLED=0 go build -o /main .


FROM alpine:3.5

WORKDIR /profilesvcapp
COPY --from=builder /main .
EXPOSE 8080
CMD ["/profilesvcapp/main"]
