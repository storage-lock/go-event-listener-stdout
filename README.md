# Event Listener Stdout

# 一、这是什么

事件机制的最简单的实现，只是在监听到事件的时候把事件输出到标准输出流中，打印的时候会把事件对象序列化为JSON字符串打印。

# 二、安装依赖

```bash
go get -u github.com/storage-lock/go-event-listener-stdout
```

# 三、使用示例

```go
package main

import (
	"context"
	events_listener_stdout "github.com/storage-lock/go-event-listener-stdout"
	memory_locks "github.com/storage-lock/go-memory-locks"
	storage_lock "github.com/storage-lock/go-storage-lock"
	"github.com/storage-lock/go-utils"
)

func main() {

	lockId := utils.RandomID()
	// 在创建锁的时候把监听器传入进去
	options := storage_lock.NewStorageLockOptions().SetLockId(lockId).AddEventListeners(events_listener_stdout.NewEventListenerStdout())
	lock, err := memory_locks.NewLockWithOptions(context.Background(), options)
	if err != nil {
		panic(err)
	}

	ownerId := "CC11001100"
	err = lock.Lock(context.Background(), ownerId)
	if err != nil {
		panic(err)
	}

	err = lock.UnLock(context.Background(), ownerId)
	if err != nil {
		panic(err)
	}

	// Output:
	// {"id":"storage-lock-event-bedb8311c3354906b372ee2871904be8","root_id":"storage-lock-event-bedb8311c3354906b372ee2871904be8","parent_id":"","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1444007+08:00","end_time":"2023-08-31T00:50:04.1610456+08:00","event_type":1,"actions"
	//:null,"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1620868+08:00","end_time":"2023-08-31T00:50:04.1620868+08:00","event_type":2
	//,"actions":[{"start_time":"2023-08-31T00:50:04.1620868+08:00","end_time":null,"name":"StorageLock.Lock.Begin","err":null,"payload_map":null}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-a8c60d2a16924192a208c64427e705cb","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1620868+08:00","end_time"
	//:"2023-08-31T00:50:04.1620868+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1620868+08:00","end_time":null,"name":"StorageLock.Lock.Try.Begin","err":null,"payload_map":null}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-a8c60d2a16924192a208c64427e705cb","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1620868+08:00","end_time"
	//:"2023-08-31T00:50:04.1620868+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1620868+08:00","end_time":null,"name":"StorageLock.Lock.Try.Begin","err":null,"payload_map":null},{"start_time":"2023-08-31T00:50:04.1620868+08:00","end_time":null,"name":"StorageLock.getLockInformation.Begin","err":null,"payload_map":null}],"watch_dog
	//_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-da6f67672a944e519adf4e417c090422","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-a8c60d2a16924192a208c64427e705cb","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1626025+08:00","end_time"
	//:"2023-08-31T00:50:04.1626025+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1626025+08:00","end_time":"2023-08-31T00:50:04.1626025+08:00","name":"Storage.Get","err":{},"payload_map":{"lockId":"163e1fe9ab0047fa94f0694b55cebbbe","lockInformationJsonString":""}}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-1c0031d29cd741c094e12e4bdad82d18","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-a8c60d2a16924192a208c64427e705cb","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1626025+08:00","end_time"
	//:"2023-08-31T00:50:04.1626025+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1626025+08:00","end_time":null,"name":"Storage.Get.Error","err":{},"payload_map":{"lockId":"163e1fe9ab0047fa94f0694b55cebbbe"}}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-2f00684732c2494cb240aa1972e44be7","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-a8c60d2a16924192a208c64427e705cb","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1626025+08:00","end_time"
	//:"2023-08-31T00:50:04.1626025+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1626025+08:00","end_time":null,"name":"StorageLock.Lock.NotExists","err":null,"payload_map":null}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-7ba32e6e696b44c8af9d0679ecf191a6","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-2f00684732c2494cb240aa1972e44be7","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1626025+08:00","end_time"
	//:"2023-08-31T00:50:04.1626025+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1626025+08:00","end_time":"2023-08-31T00:50:04.1626025+08:00","name":"Storage.GetTime","err":null,"payload_map":{"time":"2023-08-31T00:50:04.1626025+08:00"}}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-56020636aaa14e0a9210c687310b7af9","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-2f00684732c2494cb240aa1972e44be7","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1631182+08:00","end_time"
	//:"2023-08-31T00:50:04.1631182+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1631182+08:00","end_time":"2023-08-31T00:50:04.1631182+08:00","name":"Storage.CreateWithVersion","err":null,"payload_map":{"lockId":"163e1fe9ab0047fa94f0694b55cebbbe","lockInformation":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC1100110
	//0","version":1,"lock_count":1,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"version":1}}],"watch_dog_id":"","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","version":1,"lock_count":1,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease
	//_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-b4c7b00499f24b9ab042b474e9cdfdf2","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-2f00684732c2494cb240aa1972e44be7","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1636388+08:00","end_time"
	//:"2023-08-31T00:50:04.1636388+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1636388+08:00","end_time":null,"name":"Storage.CreateWithVersion.Success","err":null,"payload_map":null}],"watch_dog_id":"","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","version":1,"lock_count":1,"lock_begin_
	//time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-ad39e9f97577431887ead931739d9f1c","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-2f00684732c2494cb240aa1972e44be7","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1636388+08:00","end_time"
	//:"2023-08-31T00:50:04.1636388+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1636388+08:00","end_time":null,"name":"WatchDog.Create","err":null,"payload_map":null}],"watch_dog_id":"storage-lock-watch-dog-4ef792488e94488db2ae004efcd87408","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","v
	//ersion":1,"lock_count":1,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-564d7de2022e413ea9ba25991332cf1d","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-2f00684732c2494cb240aa1972e44be7","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1636388+08:00","end_time"
	//:"2023-08-31T00:50:04.1636388+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1636388+08:00","end_time":null,"name":"WatchDog.Create.Success","err":null,"payload_map":null}],"watch_dog_id":"storage-lock-watch-dog-4ef792488e94488db2ae004efcd87408","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC1100
	//1100","version":1,"lock_count":1,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-0dc92504508945c7b8cbbaf16400526f","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-ad39e9f97577431887ead931739d9f1c","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1641572+08:00","end_time"
	//:"2023-08-31T00:50:04.1641572+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1641572+08:00","end_time":null,"name":"WatchDog.Start","err":null,"payload_map":null}],"watch_dog_id":"storage-lock-watch-dog-4ef792488e94488db2ae004efcd87408","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","ve
	//rsion":1,"lock_count":1,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-f83cd8b7a4f5417388f2d2675390f2da","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-2f00684732c2494cb240aa1972e44be7","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1641572+08:00","end_time"
	//:"2023-08-31T00:50:04.1641572+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1641572+08:00","end_time":null,"name":"WatchDog.Start.Success","err":null,"payload_map":null}],"watch_dog_id":"storage-lock-watch-dog-4ef792488e94488db2ae004efcd87408","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001
	//100","version":1,"lock_count":1,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-8489bbeaea074cac96457d667fb38b62","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1646756+08:00","end_time"
	//:"2023-08-31T00:50:04.1646756+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1646756+08:00","end_time":null,"name":"StorageLock.Lock.Begin.Success","err":null,"payload_map":{"lockBusyCount":0,"versionMissCount":0}}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-f7423a14d4704bfa93a9ec711d047f25","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1646756+08:00","end_time"
	//:"2023-08-31T00:50:04.1646756+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1646756+08:00","end_time":null,"name":"StorageLock.Lock.Finish","err":null,"payload_map":{"versionMissCount":0}}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-d3346449d68645a09ce917ee8faa08e2","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-ad39e9f97577431887ead931739d9f1c","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1641572+08:00","end_time"
	//:"2023-08-31T00:50:04.1641572+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1641572+08:00","end_time":null,"name":"WatchDog.Exit","err":null,"payload_map":{"continueErrorCount":0,"refreshSuccessCount":0}}],"watch_dog_id":"storage-lock-watch-dog-4ef792488e94488db2ae004efcd87408","lock_information":{"lock_id":"163e1fe9ab0047fa94
	//f0694b55cebbbe","owner_id":"CC11001100","version":1,"lock_count":1,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-6aef6f5c8d7e4751baa7a284570d2ee0","root_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","parent_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1646756+08:00","end_time"
	//:"2023-08-31T00:50:04.1646756+08:00","event_type":3,"actions":[{"start_time":"2023-08-31T00:50:04.1646756+08:00","end_time":null,"name":"StorageLock.Unlock","err":null,"payload_map":null}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-fbb46c10dd924fa7b797a5f7ed0a2334","root_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","parent_id":"storage-lock-event-6aef6f5c8d7e4751baa7a284570d2ee0","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1651932+08:00","end_time"
	//:"2023-08-31T00:50:04.1651932+08:00","event_type":3,"actions":[{"start_time":"2023-08-31T00:50:04.1651932+08:00","end_time":null,"name":"StorageLock.getLockInformation.Begin","err":null,"payload_map":null}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-8f9c41a72e5f49898afc34cbcec489dc","root_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","parent_id":"storage-lock-event-fbb46c10dd924fa7b797a5f7ed0a2334","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1651932+08:00","end_time"
	//:"2023-08-31T00:50:04.1651932+08:00","event_type":3,"actions":[{"start_time":"2023-08-31T00:50:04.1651932+08:00","end_time":"2023-08-31T00:50:04.1651932+08:00","name":"Storage.Get","err":null,"payload_map":{"lockId":"163e1fe9ab0047fa94f0694b55cebbbe","lockInformationJsonString":"{\"lock_id\":\"163e1fe9ab0047fa94f0694b55cebbbe\",\"owner_id\":\"CC110
	//01100\",\"version\":1,\"lock_count\":1,\"lock_begin_time\":\"2023-08-31T00:50:04.1626025+08:00\",\"lease_expire_time\":\"2023-08-31T00:55:04.1626025+08:00\"}"}}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-fadb746454894c878088f154ff0bafa9","root_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","parent_id":"storage-lock-event-fbb46c10dd924fa7b797a5f7ed0a2334","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1651932+08:00","end_time"
	//:"2023-08-31T00:50:04.1651932+08:00","event_type":3,"actions":[{"start_time":"2023-08-31T00:50:04.1651932+08:00","end_time":null,"name":"Storage.Get.Success","err":null,"payload_map":{"lockId":"163e1fe9ab0047fa94f0694b55cebbbe","lockInformationJsonString":"{\"lock_id\":\"163e1fe9ab0047fa94f0694b55cebbbe\",\"owner_id\":\"CC11001100\",\"version\":1,\
	//"lock_count\":1,\"lock_begin_time\":\"2023-08-31T00:50:04.1626025+08:00\",\"lease_expire_time\":\"2023-08-31T00:55:04.1626025+08:00\"}"}}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-568d33f091ad498d85bd542699c04c70","root_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","parent_id":"storage-lock-event-6aef6f5c8d7e4751baa7a284570d2ee0","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1657133+08:00","end_time"
	//:"2023-08-31T00:50:04.1657133+08:00","event_type":3,"actions":[{"start_time":"2023-08-31T00:50:04.1657133+08:00","end_time":null,"name":"StorageLock.Unlock.Release","err":null,"payload_map":{"lastVersion":1,"lockInformation":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","version":2,"lock_count":0,"lock_begin_time":"2023-08-3
	//1T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"}}}],"watch_dog_id":"","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","version":2,"lock_count":0,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-568d33f091ad498d85bd542699c04c70","root_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","parent_id":"storage-lock-event-6aef6f5c8d7e4751baa7a284570d2ee0","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1657133+08:00","end_time"
	//:"2023-08-31T00:50:04.1657133+08:00","event_type":3,"actions":[{"start_time":"2023-08-31T00:50:04.1657133+08:00","end_time":null,"name":"StorageLock.Unlock.Release","err":null,"payload_map":{"lastVersion":1,"lockInformation":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","version":2,"lock_count":0,"lock_begin_time":"2023-08-3
	//1T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"}}},{"start_time":"2023-08-31T00:50:04.1662374+08:00","end_time":"2023-08-31T00:50:04.1662374+08:00","name":"Storage.UpdateWithVersion","err":null,"payload_map":{"exceptedVersion":1,"lockId":"163e1fe9ab0047fa94f0694b55cebbbe","lockInformation":{"lock_id":"163e1fe9ab004
	//7fa94f0694b55cebbbe","owner_id":"CC11001100","version":2,"lock_count":0,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"newVersion":2}}],"watch_dog_id":"","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","version":2,"lock_count":0,"lock_begin_tim
	//e":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-7bde1c46f7e2497cabd289033a33c2b7","root_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","parent_id":"storage-lock-event-568d33f091ad498d85bd542699c04c70","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1662374+08:00","end_time"
	//:"2023-08-31T00:50:04.1662374+08:00","event_type":3,"actions":[{"start_time":"2023-08-31T00:50:04.1662374+08:00","end_time":null,"name":"Storage.UpdateWithVersion.Success","err":null,"payload_map":null}],"watch_dog_id":"","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","version":2,"lock_count":0,"lock_begin_
	//time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-cd5b342ac63942528eab79b8d0b6e6d2","root_id":"storage-lock-event-d0ae6d47e63b46ed93b1f3b2431f2624","parent_id":"storage-lock-event-ad39e9f97577431887ead931739d9f1c","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1667489+08:00","end_time"
	//:"2023-08-31T00:50:04.1667489+08:00","event_type":2,"actions":[{"start_time":"2023-08-31T00:50:04.1667489+08:00","end_time":null,"name":"WatchDog.Stop","err":null,"payload_map":null}],"watch_dog_id":"storage-lock-watch-dog-4ef792488e94488db2ae004efcd87408","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","ver
	//sion":1,"lock_count":1,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-ac30bfbc42e24a9588c2b127aa241bd3","root_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","parent_id":"storage-lock-event-568d33f091ad498d85bd542699c04c70","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1667489+08:00","end_time"
	//:"2023-08-31T00:50:04.1667489+08:00","event_type":3,"actions":[{"start_time":"2023-08-31T00:50:04.1667489+08:00","end_time":null,"name":"WatchDog.Stop","err":null,"payload_map":null},{"start_time":"2023-08-31T00:50:04.1667489+08:00","end_time":null,"name":"WatchDog.Stop.success","err":null,"payload_map":null}],"watch_dog_id":"storage-lock-watch-dog
	//-4ef792488e94488db2ae004efcd87408","lock_information":{"lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","version":2,"lock_count":0,"lock_begin_time":"2023-08-31T00:50:04.1626025+08:00","lease_expire_time":"2023-08-31T00:55:04.1626025+08:00"},"err":null}
	//{"id":"storage-lock-event-1e7e64e88a9548fdaa83e63e7c8d7472","root_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","parent_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1667489+08:00","end_time"
	//:"2023-08-31T00:50:04.1667489+08:00","event_type":3,"actions":[{"start_time":"2023-08-31T00:50:04.1667489+08:00","end_time":null,"name":"StorageLock.Unlock.Success","err":null,"payload_map":{"versionMissCount":0}}],"watch_dog_id":"","lock_information":null,"err":null}
	//{"id":"storage-lock-event-1abd5284df444578ba2292950f5df793","root_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","parent_id":"storage-lock-event-15a6fd944ac945c88b5aca097a4e98ec","lock_id":"163e1fe9ab0047fa94f0694b55cebbbe","owner_id":"CC11001100","storage_name":"memory-storage","start_time":"2023-08-31T00:50:04.1672699+08:00","end_time"
	//:"2023-08-31T00:50:04.1672699+08:00","event_type":3,"actions":[{"start_time":"2023-08-31T00:50:04.1672699+08:00","end_time":null,"name":"StorageLock.Unlock.Finish","err":null,"payload_map":{"versionMissCount":0}}],"watch_dog_id":"","lock_information":null,"err":null}

}
```

