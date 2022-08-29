FUNCTION_NAME:=go-read-file
ENTRY_POINT:=ListFiles
EXECUTION_TYPE:=http
SOURCE_DIR:=./serverless_function_source_code/
PORT:=9999

# .env を読み込む
ifneq (${OPTION},)
	include $(subst --env-file=,,${OPTION})
	export
else ifeq ($(wildcard .env), .env)
	include .env
	export
endif

.PHONY: run
run:
	FUNCTION_TARGET=${ENTRY_POINT} go run cmd/main.go

deploy:
	gcloud functions deploy ${FUNCTION_NAME} \
		--project kmt-lab \
		--entry-point ${ENTRY_POINT} \
		--runtime go116 \
		--trigger-http \
		--region=asia-northeast1 \
		--set-env-vars TZ=Asia/Tokyo \
		--set-env-vars SOURCE_DIR=${SOURCE_DIR} \
		--max-instances 1

build-run:
	make build
	make run-container

build:
	pack build -v ${FUNCTION_NAME} \
    --builder gcr.io/buildpacks/builder:v1 \
    --env GOOGLE_FUNCTION_SIGNATURE_TYPE=${EXECUTION_TYPE} \
    --env GOOGLE_FUNCTION_TARGET=${ENTRY_POINT}

run-container:
	docker run --rm -p ${PORT}:8080 ${FUNCTION_NAME}

request:
	curl -XPOST http://localhost:${PORT}/ -H 'Content-Type:application/json; charset=utf-8'