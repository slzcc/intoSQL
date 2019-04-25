FROM golang
VOLUME ["/bin"]
RUN go get github.com/jmoiron/sqlx &&  go get github.com/go-sql-driver/mysql
COPY intoSQL.go $GOPATH/src/
RUN go build -o /intoSQL $GOPATH/src/intoSQL.go
ENTRYPOINT /intoSQL
