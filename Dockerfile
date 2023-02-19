FROM golang:latest

RUN apt update && apt -y install exa bat strace sysstat
# RUN go install golang.org/x/website/tour@latest
RUN echo "alias ll='exa -la'" >> /root/.bashrc

WORKDIR /root/atcoder_golang/src

