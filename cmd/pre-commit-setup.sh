# Install pre-commit and required dependencies.
set -m
set -e

function handle_exit(){
    kill $!
    exit 2
}

trap "handle_exit" SIGINT

brew info pre-commit | grep -i "bottled"

if [[ $? != "0" ]]; then
    echo "Installing pre-commit"
    brew install pre-commit
    INSTALL_PRECOMMIT=$(jobs -p)
fi

wait $INSTALL_PRECOMMIT
if [ $? != "0" ]; then
    echo "--- setup failed ---"
    exit 1
else
    pre-commit install
fi

echo "--Installing dependencies---"
GO111MODULE=off go get github.com/fzipp/gocyclo
GO111MODULE=off go get golang.org/x/tools/cmd/goimports
GO111MODULE=off go get -v -u github.com/go-critic/go-critic/cmd/gocritic
GO111MODULE=off go get -u golang.org/x/lint/golint


if [ "$?" = "0" ]; then
    echo "--- setup successful ---"
else 
    echo "--- setup failed ---"
    exit 1
fi

exit 0
