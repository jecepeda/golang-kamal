FROM golang:1.20.9 as base
FROM base as build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o app .
FROM base as final
WORKDIR /app
COPY --from=build /app/app /app/app
EXPOSE 3000
CMD [ "/app/app" ]