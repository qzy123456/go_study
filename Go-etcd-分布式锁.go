package main
import (
	"context"
	"fmt"
	"time"
	//"github.com/etcd-io/etcd/clientv3" 这个引用会导致启动失败
	"go.etcd.io/etcd/clientv3"
)

//锁结构体
type EtcdMutex struct {
	Ttl int64  //租约时间
	Conf clientv3.Config  //etcd集群配置
	Key string   //etcd的key
	cancel context.CancelFunc  //关闭续租的func
	lease clientv3.Lease
	leaseID clientv3.LeaseID
	txn clientv3.Txn
}

//初始化锁
func (em *EtcdMutex) init() error {
	var err error
	var ctx context.Context
	client,err := clientv3.New(em.Conf)
	if err != nil{
		return err
	}
	em.txn = clientv3.NewKV(client).Txn(context.TODO())
	em.lease = clientv3.NewLease(client)
	leaseResp,err := em.lease.Grant(context.TODO(),em.Ttl)
	if err != nil{
		return err
	}
	ctx,em.cancel = context.WithCancel(context.TODO())
	em.leaseID = leaseResp.ID
	_,err = em.lease.KeepAlive(ctx,em.leaseID)
	return err
}

//获取锁
func (em *EtcdMutex) Lock() error {
	err := em.init()
	if err != nil{
		return err
	}
	//LOCK:
	em.txn.If(clientv3.Compare(clientv3.CreateRevision(em.Key),"=",0)).
		Then(clientv3.OpPut(em.Key,"",clientv3.WithLease(em.leaseID))).
		Else()
	txnResp,err := em.txn.Commit()
	if err != nil{
		return err
	}
	if !txnResp.Succeeded{   //判断txn.if条件是否成立
		return fmt.Errorf("抢锁失败")
	}
	return nil
}

//释放锁
func(em *EtcdMutex) UnLock() {
	em.cancel()
	em.lease.Revoke(context.TODO(),em.leaseID)
	fmt.Println("释放了锁")
}

func main()  {
	fmt.Println("start etcd connect ...")
	var conf = clientv3.Config{
		Endpoints:   []string{"192.168.16.51:2379"},
		DialTimeout: 5 * time.Second,
	}
	eMutex1 := &EtcdMutex{
		Conf:conf,
		Ttl:10,
		Key:"lock",
	}
	eMutex2 := &EtcdMutex{
		Conf:conf,
		Ttl:10,
		Key:"lock",
	}
	//groutine1
	go func() {
		fmt.Println("groutine1")
		err := eMutex1.Lock()
		fmt.Println(err)
		if err != nil {
			fmt.Println("groutine1抢锁失败")
			fmt.Println(err)
			return
		}
		fmt.Println("groutine1抢锁成功")
		time.Sleep(10*time.Second)
		defer eMutex1.UnLock()
	}()

	//groutine2
	go func() {
		fmt.Println("groutine2")
		err := eMutex2.Lock()
		fmt.Println(err)
		if err != nil {
			fmt.Println("groutine2抢锁失败")
			fmt.Println(err)
			return
		}
		fmt.Println("groutine2抢锁成功")
		defer eMutex2.UnLock()
	}()
	time.Sleep(300*time.Second)
}