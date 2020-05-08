FROM golang:alpine
LABEL maintainer "Mathias Beke <git@denbeke.be>"

COPY . /build

WORKDIR /build/sshlogexporter
RUN go build

# Copy binary to root
RUN cp /build/sshlogexporter/sshlogexporter /


EXPOSE 9090

VOLUME ["/var/log/auth.log"]

CMD ["/sshlogexporter"]
