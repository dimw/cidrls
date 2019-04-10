# STEP 1
FROM golang:1.12 AS builder

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -race -ldflags "-extldflags '-static'" -o main .

CMD ["/app/main"]

# STEP 2
FROM scratch

COPY --from=builder /app/main /main
ENTRYPOINT ["/main"]
