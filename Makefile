fmt:
	go fmt ./...

abigen:
	abigen --abi=./abi/forwarder/Forwarder.json --pkg=forwarder --type=Forwarder --out=./abi/forwarder/Forwarder.go
	abigen --abi=./abi/token/TestToken.json --pkg=token --type=TestToken --out=./abi/token/TestToken.go
	abigen --abi=./abi/demo/UniswapDemo.json --pkg=demo --type=UniswapDemo --out=./abi/demo/UniswapDemo.go
	abigen --abi=./abi/demo/TransferDemo.json --pkg=demo --type=TransferDemo --out=./abi/demo/TransferDemo.go
