.PHONY: build-ScheduleFunction build-ScheduleFunction-prod zip-ScheduleFunction clean

clean:
	rm -rf .aws-sam
	rm -f bootstrap

build-ScheduleFunction:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap lambda/main.go
	mkdir -p $(ARTIFACTS_DIR)
	cp bootstrap $(ARTIFACTS_DIR)/

zip-ScheduleFunction: build-ScheduleFunction-prod
	zip -r latest.zip *

build-ScheduleFunction-prod:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap lambda/main.go
