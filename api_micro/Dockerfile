# Stage 1
FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git && apk add --update go git build-base
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go get -u
RUN go mod download
RUN go build -o api_micro .

# Stage 2
FROM alpine
COPY --from=builder /build/ /app/
WORKDIR /app
EXPOSE 3030
CMD ["./api_micro"]