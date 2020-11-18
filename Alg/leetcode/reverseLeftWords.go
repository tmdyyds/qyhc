func TestAlg1010(t *testing.T) {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseLeftWords(s, k))
}

//剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]

	//第二种 不利用切片
	var str string
	var str2 string
	for i := 0; i < len(s); i++ {
		if i < n {
			str += string(s[i])
		} else {
			str2 += string(s[i])
		}
	}

	return str2 + str
}

#### Redis数据不一致
主从数据不一致
原因:
	主从同步传输延迟
	即使从库及时收到了主库的命令,但是从库如果在执行其他复杂的逻辑(例如集合操作命令)而阻塞

解决方案:
	在硬件环境配置方面，我们要尽量保证主从库间的网络连接状况良好。例如，我们要避免把主从库部署在不同的机房，或者是避免把网络通信密集的应用（例如数据分析应用）和 Redis 主从库部署在一起

Redis集群中读到过期数据(删除数据只在主库,从库不会执行删除)
原因:
	Redis 的过期数据删除策略造成: Redis 同时使用了两种策略来删除过期的数据，分别是惰性删除策略和定期删除策略
		惰性删除策略: 数据过期后不会立即删除,等再次读取这个数据时才会删除
		定期删除策略: Redis每隔一段时间(默认100ms),随机选择一些数据检查是否过期
		注意:  Redis 3.2 之前的版本，那么，从库在服务读请求时，并不会判断数据是否过期，而是会返回过期数据。在 3.2 版本后，Redis 做了改进，如果读取的数据已经过期了，从库虽然不会删除，但是会返回空值，这就避免了客户端读到过期数据。所以，在应用主从集群时，尽量使用 Redis 3.2 及以上版本。









