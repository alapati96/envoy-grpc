VERSION="1.1.0"
ENVIRONMENT="${1:-prod}"  # Use provided argument or default to 'prod'

if [ "$ENVIRONMENT" == "prod" ]; then # Set TIMESTAMP based on the environment
  BUILD_TIME=$(date +%FT%T%z)
else
  set -x
  BUILD_TIME="use-cache"
fi


IMAGE_NAME="demo"
DEFAULT_TAG="devel"
TAG=${2:-$DEFAULT_TAG}
FULL_IMAGE_NAME="${IMAGE_NAME}:${TAG}"
echo "Building demo-grpc docker image: $FULL_IMAGE_NAME"


docker build  \
    --ssh default  \
    --build-arg BUILD_VERSION="$VERSION"  \
    --build-arg BUILD_TIME="$BUILD_TIME"  \
    -t $IMAGE_NAME  \
    -f Dockerfile.dev \
     --platform linux/amd64 \
     .

if [ $? -eq 0 ]; then
  echo "Docker image $FULL_IMAGE_NAME built successfully."
else
  echo "Failed to build Docker image $FULL_IMAGE_NAME."
  exit 1
fi
