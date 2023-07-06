# GRPC #
## GRPC? ##
    - binary serialization
    - smaller payload size when compared to XML & JSON
    - communication patterns
        - request / response pattern
        - client streaming (many requests & one response)
        - server streaming (one request & many responses)
        - bidirectional streaming (many requests & many responses)
    - support for limited number of languages
    - share the service & payload schema beforehand

## Steps: ##
    - Create service / operations / data contracts using protocol buffers
    - Share the contract between the client & server
    - Generate proxy & stub using the contracts
    - Server
        - implement the service (with the business logic) based on the contract
        - host the service
    - Client
        - Use the proxy to communicate to the service


## Tools Installation ##
    1. Protocol Buffers Compiler (protoc tool)
        Windows:
            Download the file, extract and keep in a folder (PATH) accessble through the command line
            https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-win64.zip
        Mac:
            brew install protobuf

        Verification:
            protoc --version

    2. Go plugins (installed in the GOPATH/bin folder)
        go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

