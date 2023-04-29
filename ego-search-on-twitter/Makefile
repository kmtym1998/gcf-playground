ENTRY_POINT:=Run

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
	gcloud functions deploy ego-search-on-twitter \
		--project kmtym1998 \
		--runtime go119 \
		--trigger-http \
		--entry-point ${ENTRY_POINT} \
		--region=asia-northeast1 \
		--service-account gcf-ego-search@kmtym1998.iam.gserviceaccount.com \
		--set-env-vars TZ=Asia/Tokyo \
		--set-secrets BEARER_TOKEN=TWITTER_BEARER_TOKEN:latest,SLACK_WEBHOOK_URL=SLACK_WEBHOOK_URL_EGO_SEARCH:latest \
		--min-instances 0 \
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