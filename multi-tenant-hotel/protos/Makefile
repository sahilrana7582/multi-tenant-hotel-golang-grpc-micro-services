PROTO_DIRS := $(shell find . -type f -name "*.proto" -exec dirname {} \; | sort -u)

all: generate

generate:
	@echo "Generating all proto files..."
	@for dir in $(PROTO_DIRS); do \
		protoc --go_out=$$dir --go-grpc_out=$$dir \
			--go_opt=paths=source_relative \
			--go-grpc_opt=paths=source_relative \
			$$dir/*.proto ; \
	done
	@echo "✅ All proto files generated."

clean:
	@find . -type f \( -name "*.pb.go" -o -name "*_grpc.pb.go" \) -delete
	@echo "🧹 Cleaned generated files."
