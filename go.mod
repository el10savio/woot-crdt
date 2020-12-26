module github.com/el10savio/woot-crdt

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.6.1
)

replace github.com/el10savio/woot-crdt/handlers => ../handlers

replace github.com/el10savio/woot-crdt/woot => ../woot
