TAG=$(git describe --tags)
TAG="0.0.0"
DIR="gcal_run-${TAG}"
rm -rf dist/${DIR}  && \
mkdir dist/${DIR} && \
cp README.md dist/${DIR}/README.md && \
go build -o dist/${DIR}/gcal_run cmd/gcal_run/gcal_run.go && \
go build -o dist/${DIR}/installer cmd/installer/installer.go && \
(cd dist && tar zcvf ${DIR}.tar.gz ${DIR})