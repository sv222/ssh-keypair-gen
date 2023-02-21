# SSH Keypair Generator

This Go utility automatically generates a cryptographically stable RSA SSH key pair and outputs the public and private keys.

## Installation

To install and use the SSH keygen utility, follow these steps:

1. Install Go and set up your Go environment. See the [Go installation guide](https://golang.org/doc/install) for more information.

2. Clone this repository:

    ```sh
    git clone https://github.com/sv222/ssh-keypair-gen.git
    cd ssh-keypair-gen
    ```

3. Build the binary:

    ```sh
    go build -o ssh-keypair-gen
    ```

## Usage

### Local installation

1. Run the binary with the desired flags. For example, to generate a 4096-bit key and save the private key as `mykey.pem` and the public key as `mykey.pub`, run the following command:

    ```sh
    ./ssh-keypair-gen --size 4096 --private-key mykey.pem --public-key mykey.pub
    ```

You can choose not to pass values, then the names will be given by default and the files will be saved in the current folder.

```sh
./ssh-keypair-gen
```

### Docker installation

1. Build the Docker image:

    ```sh
    docker build -t ssh-keypair-gen .
    ```

2. Run the Docker container with the desired flags. For example, to generate a 4096-bit key and save the private key as `mykey.pem` and the public key as `mykey.pub`, run the following command:

    ```sh
    docker run -v $(pwd):/keys ssh-keypair-gen --private-key /keys/id_rsa.pem --public-key /keys/id_rsa.pub
    ```

The above command uses the `-v` flag to mount the current directory as a volume inside the container, and the `-w` flag to set the working directory inside the container to `/app`, which is where the binary and output files are located.

Replace "my-go-app" with the name you gave to the Docker image when you built it.

## Contribution

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
