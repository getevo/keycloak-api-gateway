# Dockerfile
FROM ubuntu:bionic
LABEL maintainer="allan.nava@hiway.media"
#
# RUN go get -u github.com/getevo/evo
RUN mkdir -p /go/src/keycloak-api-gateway
#
#
WORKDIR /go/src/keycloak-api-gateway/
# Only runtime
#FROM golang:1.14.4-buster
COPY --from=builder /go/src/keycloak-api-gateway/ /go/src/keycloak-api-gateway/
#
EXPOSE 8010
#
CMD ["/go/src/keycloak-api-gateway/keycloak-api-gateway","-c","/go/src/keycloak-api-gateway/config.yml"]
#