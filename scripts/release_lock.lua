local key = KEYS[1]
local requestId = ARGV[1]
local value = redis.GET(key)
if value == requestId then
    redis.DEL(key)
    return 1;
end
return 0