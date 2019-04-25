FROM golang
VOLUME ["/bin"]
RUN go get github.com/jmoiron/sqlx &&  go get github.com/go-sql-driver/mysql
COPY intoSQL.go $GOPATH/src/
RUN go build -o / $GOPATH/src/intoSQL.go $GOPATH/src/selectSQL.go
ENTRYPOINT /intoSQL
