FROM debian:jessie-slim

ENTRYPOINT ["/usr/local/bin/hog"]

COPY hog /usr/local/bin/hog

