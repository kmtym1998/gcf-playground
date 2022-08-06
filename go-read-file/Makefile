FUNCTION_NAME:=ListFiles
SOURCE_DIR:=./serverless_function_source_code

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
	FUNCTION_TARGET=${FUNCTION_NAME} go run cmd/main.go

deploy:
	gcloud beta functions deploy ${FUNCTION_NAME} \
		--project kmt-lab \
		--entry-point ${FUNCTION_NAME} \
		--runtime go116 \
		--trigger-http \
		--region=asia-northeast1 \
		--set-env-vars TZ=Asia/Tokyo \
		--set-env-vars SOURCE_DIR=${SOURCE_DIR} \
		--max-instances 1
