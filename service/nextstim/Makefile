.PHONY: build run upload
#PWD := `pwd`
PWD := $(shell pwd)
build:
	docker build . -t bundle-pip-modules-for-aws-lambda-layers:python3.8

layer:
	docker run --rm \
        -v $(PWD):/dist \
        bundle-pip-modules-for-aws-lambda-layers:python3.8

upload:
	aws --profile tetsuzawa_el --region ap-northeast-1 lambda publish-layer-version \
        --layer-name questplus \
        --compatible-runtimes python3.8 \
        --zip-file fileb://layer.zip
