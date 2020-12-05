
abigen --sol dFedIndex.sol  --pkg index --out sdk/index/index.go
abigen --sol dFedPair.sol  --pkg pair --out sdk/pair/pair.go
abigen --sol dFedFactory.sol  --pkg factory --out sdk/factory/factory.go
abigen --sol test/TestERC20.sol  --pkg testtoken --out sdk/testtoken/TestERC20.go
abigen --sol dFedUsdd.sol  --pkg usdd --out sdk/usdd/usdd.go
abigen --sol FedToken.sol  --pkg fed --out sdk/fed/fed.go
abigen --sol dFedCreator.sol  --pkg creator --out sdk/creator/creator.go

solc --optimize --abi *.sol -o abi --overwrite

solc --optimize --bin *.sol -o bin --overwrite

