FROM golang:alpine

WORKDIR /app

COPY frame_up_bot .

ENTRYPOINT ["/app/frame_up_bot"]