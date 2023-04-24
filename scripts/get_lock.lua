local key = KEYS[1]
local value = ARGV[1]
local ttl = tonumber(ARGV[2])
local result = redis.call('SET', key, value, NX, PX, ttl)
--- 1代表获取到锁并且扣减成功 -1代表扣减失败 0 代表没有获取到锁
if result == 1 then  --- 如果获取到锁
    --- 查找库存并进行库存扣减
    local inventory_key = "inventory_" .. key  --- 拼接字符串,真正的库存key
    local inventory = redis.GET(inventory_key)
    if inventory == nil or inventory == 0 then
        result = -1  --- 如果是-1证明没有该商品或者库存为空
    else
        redis.SET(inventory_key, inventory-1)
    end
end

return result