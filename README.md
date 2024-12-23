# GRPC Learning

This repository demonstrates a gRPC learning project built with **Go 1.22**.

[Next Steps](#next-steps)

---

## Test Coverage

56.2% test coverage.

---

## Protobuf

To generate protobuf files, run:

```bash
$ make generate
```

---

## Docker

To simplify development, a `Makefile` is provided for Docker operations. Below are the key commands:

- **Start Docker Compose with all instances (DB, Client, and Server):**
  ```bash
  $ make up
  ```
  Test the setup by running:
  ```bash
  $ make send    # Load ports to the database
  $ make get     # Fetch ports from the database
  ```

- **Stop and remove Docker Compose:**
  ```bash
  $ make down
  ```

---

## Kubernetes (KinD)

We use **KinD (Kubernetes in Docker)** for a local Kubernetes cluster. Here are the available commands:

- **Cluster management:**
  ```bash
  $ make kind-create       # Create a local KinD cluster
  $ make kind-delete       # Delete all resources in ./deploy/k8s
  $ make kind-clean        # Delete the entire KinD cluster
  ```

- **Cluster setup:**
  ```bash
  $ make kind-init         # Initialize KinD with necessary images
  $ make kind-load         # Load Docker-built images into the cluster
  ```

- **Resource management:**
  ```bash
  $ make kind-apply        # Apply the Kubernetes YAML files in ./deploy/k8s
  $ make kind-apply-secrets # Apply secrets from the .env file
  ```

- **Monitoring:**
  ```bash
  $ make kind-watch        # Watch the pods' status
  ```

- **Database session:**
  ```bash
  $ make db                # Connect to the database session
  ```

---

## Running Locally

- **Start the gRPC server:**
  ```bash
  $ make run-server
  ```

- **Start the gRPC client to send messages to the server:**
  ```bash
  $ make run-client
  ```

---

## Building Locally

Build Go binaries for server and client:

- **Server binary:**
  ```bash
  $ make build-server-bin
  ```

- **Client binary:**
  ```bash
  $ make build-client-bin
  ```

The binaries will be exported to the `bin` folder.

---

## Running Tests

Run the test suite:

```bash
$ make test
```

---

## Linting

Run the linter to ensure code quality:

```bash
$ make lint
```

---

## Useful Tools

- **Expose API for local access:**
  ```bash
  $ make expose
  ```

---

## Next Steps

1. **Improve data structure:**  
56.2% test coverage.
  Write additional tests to increase the overall coverage.

2. **Add detailed gRPC status codes:**  
  Use `codes` and `status` from the following packages to improve client-server communication:
  - `google.golang.org/grpc/codes`
  - `google.golang.org/grpc/status`
