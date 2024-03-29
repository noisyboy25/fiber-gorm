# Build client
FROM node AS client_builder

WORKDIR /client

COPY vite-project/package.json .
COPY vite-project/yarn.lock .
RUN yarn install
COPY vite-project .
RUN yarn build

# Build server
FROM golang:alpine AS server_builder

RUN apk add build-base

WORKDIR /server

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /fiber-gorm

# Deploy
FROM alpine

WORKDIR /root/

COPY --from=server_builder /fiber-gorm ./
COPY --from=client_builder /client/dist ./vite-project/dist

EXPOSE 3000

CMD [ "./fiber-gorm" ]