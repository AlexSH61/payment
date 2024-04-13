FROM golang AS Builder
COPY go.mod go.sum
RUN go mod download
EXPOSE 8001
CMD [ "./cmd" ]