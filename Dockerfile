FROM golang:1.18-alpine as build
RUN apk --no-cache add gcc g++ make git

ENV GOOS=linux \
    GOARCH=amd64 
    
ARG VERSION=dev
WORKDIR /go/src/github.com/Test/CovidApp
COPY . .

# RUN go mod download
RUN go build -tags musl -o ./bin/CovidApp -mod vendor -ldflags "-X main.version=${VERSION} -w -s"

FROM alpine:3.9

# Add timezone data
RUN apk add --no-cache tzdata

WORKDIR /usr/bin
COPY --from=build /go/src/github.com/Test/CovidApp/bin /go/bin
COPY --from=build /go/src/github.com/Test/CovidApp/config /usr
COPY --from=build /go/src/github.com/Test/CovidApp/migrations /usr/bin/migrations
COPY --from=build /go/src/github.com/Test/CovidApp/config/covid-app.yaml /tmp/config/
COPY --from=build /go/src/github.com/Test/CovidApp/covid_19_data.csv /tmp/config/
EXPOSE 3001
ENTRYPOINT /go/bin/CovidApp serve
