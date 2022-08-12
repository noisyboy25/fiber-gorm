build: install-server install-client build-client

install-server:
	go install

install-client:
	cd vite-project && yarn install --frozen-lockfile

build-client:
	cd vite-project && yarn build

run: 
	go run .