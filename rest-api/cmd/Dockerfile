FROM alpine:latest
MAINTAINER niloofar.gheibi@gmail.com

# add config
ADD config.development.yaml /

# add binary
ADD build/linux/rest-api/ .

EXPOSE 8080

ENTRYPOINT ["/rest-api"]
