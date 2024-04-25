FROM golang:alpine3.9 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build 

FROM alpine AS production
COPY --from=builder /app .
CMD [ "./app" ]