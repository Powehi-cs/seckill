local key = KEYS[1]
local requestId = ARGV[1]
local value = redis.call('get', key)
if value == requestId then
    redis.call('del', key)
    return 1;
end
return 0