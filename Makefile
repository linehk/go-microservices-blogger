.PHONY: pb
pb:
	@./script/pb.sh user
	@./script/pb.sh blog
	@./script/pb.sh post
	@./script/pb.sh page
	@./script/pb.sh comment

.PHONY: test
test:
	go test -count=1 -p=8 -parallel=8 -race ./...
	
.PHONY: model
model:
	@./script/model.sh user