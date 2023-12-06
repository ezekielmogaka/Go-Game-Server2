# This file should be populated with the necessary configuration to build your web service's container image


# import base image for building our Go application using the Alpine Linux distribution.

FROM golang:alpine AS BUILDER 

#set the working directory inside the container to /app

WORKDIR /app

#copy only the go.mod file from the host machine to the container's working directory

COPY go.mod .

#executes the go mod download command, which downloads the dependencies specified in go.mod

RUN go mod download

#copy the entire source code from the host machine to the container's working directory

COPY . .

#build the Go application with the CGO_ENABLED environment variable set to 0 (disabling CGO) and GOOS set to linux. 
#The resulting binary is named server and is placed in the current directory (.)

RUN CGO_ENABLED=0 GOOS=linux go build -o server .


## Runtime stage
#set the base image for the runtime stage as alpine
FROM alpine AS RUNNER

#set the working directory inside the container to /app.

WORKDIR /app

#copy the server binary from the build stage container (specified by --from=BUILDER) to the runtime stage container.
# Copy games.json file from the build stage container to the runtime stage container.

COPY --from=BUILDER /app/server  .
COPY --from=BUILDER /app/games.json .

#Listen on port 8080 for the server

EXPOSE 8080


#set the default command to run when the container starts.

CMD ["./server"]
