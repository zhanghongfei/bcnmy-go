fmt:
	go fmt ./...

abigen:
	abigen --abi=./forwarder/Forwarder.json --pkg=forwarder --type=Forwarder --out=./forwarder/Forwarder.go
