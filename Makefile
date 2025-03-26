.PHONY: run

lambda-build:
	for dir in $(shell find ./lambda -type d); do \
		echo "Processing directory: $$dir"; \
		GOOS=linux GOARCH=amd64 go build -o $$dir/bootstrap $$dir/main.go; \
		cd $$dir && zip function.zip bootstrap && cd -; \
	done

clean:
	for dir in $(shell find ./lambda -type d); do \
		rm $$dir/bootstrap; \
		rm $$dir/function.zip; \
	done

cdk-bootstrap:
	cdk bootstrap --profile $(AWS_PROFILE)

cdk-deploy:
	cdk deploy --profile $(AWS_PROFILE)

cdk-destroy:
	cdk destroy --profile $(AWS_PROFILE)