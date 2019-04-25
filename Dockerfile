FROM golang
VOLUME ["/bin"]
RUN go get github.com/jmoiron/sqlx &&  go get github.com/go-sql-driver/mysql
COPY ./* $GOPATH/src/
RUN go build -o /intoSQL $GOPATH/src/intoSQL.go
RUN go build -o /selectSQL $GOPATH/src/selectSQL.go
ENTRYPOINT /intoSQL
