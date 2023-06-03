-- KEYS[1] 锁名称
-- ARGV[1] 锁value
-- ARGV[2] 锁失效时间


-- 加锁
if  redis.call('GET', KEYS[1]) == ARGV[1] then
    redis.call('EXPIRE', KEYS[1], ARGV[2])
    -- 无论是否续费成功，都已经获得锁，固定返回1
    return 1
else
    return redis.call('SET', KEYS[1], ARGV[1], 'NX', 'PX', ARGV[2])
end

-- 解锁
if  redis.call('GET', KEYS[1]) == ARGV[1] then
    return redis.call('DEL', KEYS[1])
else
    return 0
end

-- 看门狗
if  redis.call('GET', KEYS[1]) == ARGV[1] then
    return redis.call('EXPIRE', KEYS[1], ARGV[2])
else
    return 0
end