.PHONY: build-ScheduleFunction clean

clean:
	rm -rf .aws-sam
	rm -f bootstrap

build-ScheduleFunction:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap lambda/main.go
	mkdir -p $(ARTIFACTS_DIR)
	cp bootstrap $(ARTIFACTS_DIR)/
