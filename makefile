BINARY_NAME=go-react
VERSION=v1
DIST_DIR=dist


clean:
	[ -f ${BINARY_NAME} ] && rm ${BINARY_NAME} ; [ -d ${DIST_DIR} ] && rm -r ${DIST_DIR}; mkdir ${DIST_DIR};

npm:
	cd app; npm run build; cd - ;
	cp -r app/build/* ${DIST_DIR}/ 

build: npm 
	go build; 

run: npm
	go run server.go

image:
	docker build -t imhshekhar47/${BINARY_NAME}:${VERSION} .

	
