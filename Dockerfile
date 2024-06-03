FROM ubuntu:latest
LABEL authors="joneco"

ENTRYPOINT ["top", "-b"]