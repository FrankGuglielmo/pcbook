# PCBook

PCBook is a Go project made from the dev.to tutorial: [dev.to tutorial](https://dev.to/techschoolguru/the-complete-grpc-course-protobuf-go-java-2af6), built with protocol buffers.

## Features

- **Protocol Buffers:** Uses Protocol Buffers for serializing structured data.
- **CRUD Operations:** Supports create, read, update, and delete operations for laptop entities.

## Requirements

- Go (1.13+)
- Protocol Buffers Compiler (protoc)
- gRPC plugin for Protocol Buffers Compiler

## Installation

1. **Clone the Repository:**

   ```sh
   git clone https://github.com/FrankGuglielmo/pcbook.git
   cd pcbook
   ```

2. **Install Dependencies:**

   ```sh
   go mod download
   ```

3. **Install Protocol Buffers Compiler:**

   - **MacOS:**

     ```sh
     brew install protobuf
     ```

   - **Linux:**

     ```sh
     sudo apt-get install -y protobuf-compiler
     ```

   - **Windows:**

     Download and install the precompiled binaries from the [official website](https://github.com/protocolbuffers/protobuf/releases).

4. **Install gRPC Plugin for Protocol Buffers Compiler:**

   ```sh
   go get -u google.golang.org/grpc
   go get -u github.com/golang/protobuf/protoc-gen-go
   ```

## Generate .proto Files

To generate the Go code from the .proto files, use the Makefile provided in the repository.

1. **Navigate to the Project Directory:**

   ```sh
   cd pcbook
   ```

2. **Run the Makefile:**

   ```sh
   make
   ```

This will generate the necessary Go files from the .proto definitions.

## Project Structure

- `pb/`: Contains the Protocol Buffers definitions and the generated Go files.
- `serializer/`: Contains utility functions for serializing and deserializing data.
- `sample/`: Contains sample data generation code.

## Contributing

Feel free to open issues or submit pull requests for any improvements or bug fixes. For major changes, please open an issue first to discuss what you would like to change.
