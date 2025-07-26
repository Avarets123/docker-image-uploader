FROM golang:1.24-alpine AS build
  
WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .

RUN mkdir -p /opt/dimage_uploader

RUN CGO_ENABLED=0 GOOS=linux go build  -o /opt/dimage_uploader/main


# RUN CGO_ENABLED=0 GOOS=linux go build  -o ./main


# FROM alpine:3.18

# RUN mkdir -p /opt/dimage_uploader

# # COPY --from=build  /build/entrypoint.sh /opt/dimage_uploader/entrypoint.sh
# # RUN chmod +x /opt/dimage_uploader/entrypoint.sh

# COPY --from=build  /build/main /opt/dimage_uploader/main
# RUN chmod +x /opt/dimage_uploader/main


ENTRYPOINT ["/opt/dimage_uploader/main"]