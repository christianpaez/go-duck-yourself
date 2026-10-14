.PHONY: install run test

install:
	bundle install

run:
	bundle exec puma -p 8080 

test:
	curl -i localhost:8080/health
	curl -i localhost:8080/file.txt
