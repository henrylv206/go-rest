#! /bin/bash

# 1. go build
#cd cmd
docker run --rm -v "$PWD":/usr/src/myapp -v /Users/lvhonglei/works/go/src:/go/src -w /usr/src/myapp golang:1.10.3-alpine3.8  go build -o ./build/bin/jd-pms-api ./cmd

# 2. copy binary file
#mv jd-pms-api ../

# 3. docker build
cd build
docker build -t j-hub.jd.com/jdevops/jd-pms-api:v1.0 .

# 4. docker push
docker push j-hub.jd.com/jdevops/jd-pms-api:v1.0

# 5. restart api server
curl http://jd-pms-api.jdevops.okstack.com/shutdown
