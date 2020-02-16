CURRENT_DIRECTORY=`pwd`
ROOT_DIRECTORY="$( dirname "$(dirname "$CURRENT_DIRECTORY")")"

docker run --privileged -v \
$ROOT_DIRECTORY:/gopkg/src/github.com/actiontech/dtle \
-it nomad-e2e /bin/bash \
-c "cd gopkg/src/github.com/actiontech/dtle/e2e/migrations && go test --run \
TestJobMigrations -integration"

docker run --privileged \
-v $ROOT_DIRECTORY:/gopkg/src/github.com/actiontech/dtle \
-it nomad-e2e /bin/bash \
-c "cd gopkg/src/github.com/actiontech/dtle/e2e/migrations && go test --run \
TestMigrations_WithACLs -integration"
