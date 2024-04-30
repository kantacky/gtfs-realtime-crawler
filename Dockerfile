FROM golang:1.22.2 as builder
WORKDIR /app
COPY . .
RUN go mod tidy && \
    CGO_ENABLED=0 \
    GOOS=linux \
    go build \
    -a \
    -installsuffix cgo \
    -o main \
    .

FROM alpine:latest as runner
ENV TZ 'Asia/Tokyo'
WORKDIR /root
COPY --from=builder /app/main .
EXPOSE 8080
CMD [ "./main" ]
