build_img:
	docker build -t snippets-app .

run_cnt:
	docker run -p 4000:4000 -d --rm --name snippets snippets-app

stop_cnt:
	docker stop snippets

del_img:
	docker image rm snippets