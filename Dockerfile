FROM golang:1.16-alpine AS build
WORKDIR /app
COPY /go.mod .
COPY /go.sum .
RUN go mod download
COPY *.go .

RUN go build -o /out


FROM scratch AS bin
COPY --from=build /out /example
EXPOSE 8080
ENTRYPOINT ["/example"]




