FROM golang:1.18-buster

RUN wget -q https://raw.githubusercontent.com/dapr/cli/master/install/install.sh -O - | /bin/bash

# Install daprd
ARG DAPR_BUILD_DIR
COPY $DAPR_BUILD_DIR /opt/dapr
ENV PATH="/opt/dapr/:${PATH}"
RUN dapr init --slim

WORKDIR /app

COPY ./src/go.mod .
COPY ./src/go.sum . 

COPY ./src/ .
RUN go mod download

EXPOSE 8080
ENTRYPOINT ["go", "run", "."]
# ENTRYPOINT ["dapr"]
# CMD ["run", "--app-id", "go-rest", "--app-port", "8080", "go", "run", "."]

