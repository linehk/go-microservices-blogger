.PHONY: pb
pb:
	@./script/pb.sh

.PHONY: test
test:
	go test -count=1 -p=8 -parallel=8 -race ./...
	
.PHONY: model
model:
	@./script/model.sh

.PHONY: start
start:
	@./script/start.sh