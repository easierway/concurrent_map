# ConcurrentMap for GO
## The better performance thread-safe map in GO

After v1.9, normally, programmers have two options for thread-safe map. One is to build the thread-safe solution with syn.RWMutex. But, in many cases, especially, in the case the number of CPU cores is larger than 2, this option's performance is quite poor.

Another is to use the sync.map, which has been added to the sync package from v.1.9. Unfortunately, sync.map can not work well for all the cases, especially, the case of multi-threads writing. For more info, please, check the great video https://www.youtube.com/watch?v=C1EtfDnsdDs.

This project is to provide a thread-safe map which is Java ConcurrentMap's GO version. From the following benchmark you can see it is better in the multi-thread writing cases.

The following test is about 100 Goroutines writing and 100 Groutines reading. The test is executed on Macbook (macOS 10.13.2, 2 core (2.3G Intel Core i5), 8G LPDDR3)

![image](https://github.com/easierway/concurrent_map/blob/master/map_benchmark.png)

### FAQ
1 Why not provide the default hash function for partition?

Ans: As you known, the partition solution would impact the performance significantly. The proper partition solution balances the access to the different partitions and avoid of the hot partition. The access mode highly relates to your business. So, the better partition solution would just be designed according to your business.
