App has to be run in GOPATH mode.
```
mkdir -p ~/tmp/del/go/src/keployswagger
export GOPATH=~/tmp/del/go
# get app from github.com/subtleseeker/keploy_swagger_app

# start keploy with sudo (uses /keploy path at the moment)
sudo keploy

# start server
KEPLOY_MODE="record" go run .

# send client request
curl 0:8089/upload

# rerun server in test mode.
KEPLOY_MODE="test" go run .
```

