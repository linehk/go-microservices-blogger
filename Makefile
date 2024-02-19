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
	@./script/model.sh user app_user
	@./script/model.sh user locale

	@./script/model.sh blog blog
	@./script/model.sh blog blog_user_info
	@./script/model.sh blog page_views

	@./script/model.sh post post
	@./script/model.sh post location
	@./script/model.sh post label
	@./script/model.sh post image
	@./script/model.sh post post_user_info
	@./script/model.sh post author

	@./script/model.sh page page

	@./script/model.sh comment comment 