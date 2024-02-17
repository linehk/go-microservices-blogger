.PHONY: pb
pb:
	@./script/pb.sh user
	@./script/pb.sh blog

.PHONY: test
test:
	go test -count=1 -p=8 -parallel=8 -race ./...
	
.PHONY: model
model:
	@./script/model.sh user