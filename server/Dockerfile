FROM golang

WORKDIR /outstagram

COPY go.mod .

COPY go.sum .

RUN go mod download

# Comment the following line for production
# RUN go get github.com/codegangsta/gin

COPY . .

# Uncomment the following line for production
#CMD ["go", "run", "app.go"]

# Comment the following line for production
#CMD ["gin", "-p", "5000", "-a", "5001", "run", "app.go"]