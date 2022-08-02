FROM golang:1.18-alpine AS build

WORKDIR /app

# prevent the re-installation of vendors at every change in the source code
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download && go mod verify

COPY . ./
# Build app
RUN CGO_ENABLED=0 go build -o /server

# Expose port and run
EXPOSE 8080

ENTRYPOINT [ "/server" ]





FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

USER nonroot:nonroot

CMD [ "/server" ]
