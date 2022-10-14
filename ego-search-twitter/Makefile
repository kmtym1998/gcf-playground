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
	gcloud functions deploy ego-search-twitter \
		--project kmt-lab \
		--runtime go119 \
		--entry-point ${ENTRY_POINT} \
		--trigger-resource ego-search-twitter-topic \
		--trigger-event google.pubsub.topic.publish \
		--region=asia-northeast1 \
		--service-account function-ego-search@kmt-lab.iam.gserviceaccount.com \
		--set-secrets HASURA_GRAPHQL_ADMIN_SECRET=hasura-admin-secret:latest,SLACK_WEBHOOK_URL=slack-webhook-url-alert:latest \
		--set-env-vars TZ=Asia/Tokyo \
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