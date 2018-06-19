This is a ethereum wallet implementation which uses Infura API.

Infura API key is required as an environment variable. Refer the
env-sample for adding a .env file to the project root.

To use this, you need to have an existing ethereum keystore json.

There is an option of creating one but it has not been integrated yet.

To execute the wallet:

>$ go run cmd/main.go -p < wallet_password > -f < path-to-keystore-file >

eg

>$ go run cmd/main.go -p odetojoy -f ./UTC--2018-06-16T18-41-19.615987160Z--c0b4ef9e6d2806f643be94b2434f5c3d6cecd255

Where the wallet password is odetojoy and the keystore file is in the same directory
as the code.