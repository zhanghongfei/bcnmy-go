## Biconomy Golang SDK

## CHANGELOG

1. `RawTransact` and `Transact` change returns `(*types.Transaction, *types.Receipt, error)`.

2. Add `WithFieldTimeout` for change httpClient timeout, due to wait Txn mined.
