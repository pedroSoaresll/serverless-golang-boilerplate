.PHONY: run

build-lambda:
	for dir in $(shell find ./lambda -type d); do \
		echo "Processing directory: $$dir"; \
		GOOS=linux GOARCH=amd64 go build -o $$dir/bootstrap $$dir/main.go; \
		cd $$dir && zip function.zip bootstrap && cd -; \
	done

clean:
	for dir in $(shell find ./lambda -type d); do \
		rm -f $$dir/bootstrap; \
		rm -f $$dir/function.zip; \
	done

cdk-bootstrap:
	cdk bootstrap --all --profile $(AWS_PROFILE)

cdk-deploy:
	cdk deploy --all --profile $(AWS_PROFILE)

cdk-destroy:
	cdk destroy --all --profile $(AWS_PROFILE)

ci-synth:
	# Run CDK synth to validate the stack in CI
	cdk synth --all

ci-deploy:
	# Deploy the stack in CI without requiring manual approval
	cdk deploy --all --require-approval never