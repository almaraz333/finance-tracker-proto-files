# Finance Tracking Protobuf Definitions

This repository contains the Protocol Buffers (protobuf) definitions used across various Go microservices in the Finance Tracking application. It serves as a centralized schema definition platform to ensure consistency, type safety, and efficient wire format across our microservices architecture.

## Getting Started

### Requirements

Ensure you have the following tools installed:

- `protoc`, the protobuf compiler. It can be downloaded [here](https://github.com/protocolbuffers/protobuf/releases).
- `protoc-gen-go`, the Go protocol buffers plugin. Install it using: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
- Go 1.22.1 or higher.

### Generating Go Code from Protobuf

After cloning this repository and making sure all requirements are met, you can generate the Go code needed by your microservices as follows:

1. Navigate to the root of the cloned repository.
2. Run the following command:
   ```bash 
   make
   ```

This command instructs `protoc` to generate Go code in the same directories as your `.proto` files, preserving package structure.
