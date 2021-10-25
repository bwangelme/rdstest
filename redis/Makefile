.PHONY: m s init


init:
	$([[ ! -d "run_6379" ]] && mkdir run_6379)
	$([[ ! -d "run_6380" ]] && mkdir run_6380)

m:init
	redis-server ./redis_6379.conf

s:init
	redis-server ./redis_6380.conf
