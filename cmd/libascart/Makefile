all: main.go
	go build -buildmode=c-shared -o libascart.so ./main.go
install: libascart.so 
	sudo install ./libascart.so /usr/lib/
	sudo install ./libascart.h /usr/include
