local key = KEYS[1]
local value = ARGV[1]
local result = redis.call('set', key, value, 'EX', 1, 'NX')

if result then
    --- 1: 获取到锁 0: 没有获取到锁 -1: 没有商品或者库存为空
    local inventory_key = "inventory_" .. key  --- 拼接字符串,真正的库存key
    local inventory = redis.call('get', inventory_key)
    if inventory == nil or inventory == 0 then
        return -1
    else
        redis.call('set', inventory_key, inventory - 1)
        return 1
    end
end

return 0



