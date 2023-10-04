FROM golang:1.19-buster

ENV N_PATH "/go/src/transactions-history"
ENV N_USER "stori"
ENV WORKDIR ${N_PATH}

# Install dependencies
RUN apt update -q && apt install git autoconf automake libtool curl make g++ upx unzip sudo inotify-tools -qy

# Add user
RUN addgroup --gid 1000 ${N_USER} && adduser --gid 1000 --uid 1000 --disabled-password ${N_USER}
RUN echo "ALL ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers
# Set user envs
USER ${N_USER}
ENV GOPATH=/home/${N_USER}/go
ENV GOBIN=/home/${N_USER}/go/bin
ENV PATH=$PATH:/home/${N_USER}/.local/bin:$GOBIN

# Install dependencies
RUN go install github.com/codegangsta/gin@latest

# Workspace
WORKDIR ${N_PATH}
COPY . .
RUN sudo chown -R ${N_USER}:${N_USER} .
RUN sudo chmod -R 777 .
# Install project packages
RUN go mod tidy

CMD gin --immediate --all --buildArgs="-buildvcs=false" --laddr=0.0.0.0 --port=3030 --appPort=8080 --path=./ run main.go

EXPOSE 3030
