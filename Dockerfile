FROM golang:1.19-alpine AS build
WORKDIR /src/app
COPY . .
RUN go build -o /bin/app ./cmd/main.go

FROM alpine
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]