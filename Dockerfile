FROM golang:1.24-alpine AS build
  
WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build  -o ./main


FROM docker:cli

RUN mkdir -p /opt/dimage_uploader

COPY --from=build  /build/main /opt/dimage_uploader/main


ENTRYPOINT ["/opt/dimage_uploader/main"]