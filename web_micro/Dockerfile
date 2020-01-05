# Stage 1
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git && apk add --update go git build-base
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod download
RUN go build -o web_micro .

# Stage 2
FROM alpine
COPY --from=builder /build/ /app/
WORKDIR /app
EXPOSE 3000
CMD ["./web_micro"]