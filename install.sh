# Clone siglens repository
git clone https://github.com/siglens/siglens.git

# Install dependencies (Vector and Vector-Agent)
curl --proto '=https' --tlsv1.2 -sSfL https://sh.vector.dev | bash -s -- -y  && \
helm repo add vector https://helm.vector.dev && \
helm repo update && \
helm install vector-agent vector/vector --values=vectordep-values.yaml

# Run siglens server
# Enter siglens directory
cd siglens
go run cmd/siglens/main.go --config server.yaml
