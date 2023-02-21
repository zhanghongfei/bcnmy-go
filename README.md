## Biconomy Golang SDK

## CHANGELOG

### 2023-02-21
1. Add `timeout` parameter in `NewBcnmy`

### 2023-02-20
1. Add `defer close channel` to response channel and error channel.

### 2023-02-16
1. `RawTransact` and `Transact` change returns `(*types.Transaction, *types.Receipt, error)`.

2. Add `WithFieldTimeout` for change httpClient timeout, due to wait Txn mined.
