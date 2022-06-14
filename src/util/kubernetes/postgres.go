package kubernetes

import (
	"context"
	"fmt"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	v1 "github.com/zalando/postgres-operator/pkg/apis/acid.zalan.do/v1"
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"
	"time"
)

type GatherPostgresInfoFunc func(instance *entity.Instance, wg *sync.WaitGroup)

type GatherPostgresStatusFunc func(instance *entity.Instance, wg *sync.WaitGroup)

var funcPostgresList = []GatherPostgresInfoFunc{
	GetPostgresSVC,
	GetPostgresSecret,
}

var funcPostgresStatus = []GatherPostgresStatusFunc{
	GetPostgresStatus,
}

func GetPostgresStatus(instance *entity.Instance, wg *sync.WaitGroup) {
	defer wg.Done()
	postgres, err := GetPostgres(instance)
	if err != nil {
		return
	}
	instance.Status = postgres.Status.PostgresClusterStatus
	byteData, _ := yaml.Marshal(postgres)
	instance.Yaml = string(byteData)
}

func WatchPostgresStatus(instance *entity.Instance, wg *sync.WaitGroup) {
	defer wg.Done()
	client, err := NewPostgresqlClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return
	}
	watcher, err := client.AcidV1().Postgresqls(instance.Namespace).Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	go func() {
		for event := range watcher.ResultChan() {
			fmt.Printf("Type: %v\n", event.Type)
			obj, ok := event.Object.(*v1.Postgresql)
			if !ok {
				logger.Log.Errorf("Unexpected type: %v\n", event.Type)
				return
			}
			instance.Status = obj.Status.PostgresClusterStatus
		}
	}()
	time.Sleep(5 * time.Second)
}

func GetPostgresSVC(instance *entity.Instance, wg *sync.WaitGroup) {
	defer wg.Done()

	svcInfos, err := GetServiceInfo(instance)
	if err != nil {
		return
	}
	instance.ServiceInfo = *svcInfos
}

func GetPostgresSecret(instance *entity.Instance, wg *sync.WaitGroup) {
	defer wg.Done()
	secretInfos, err := GetSecretInfo(instance)
	if err != nil {
		return
	}
	instance.Secret = secretInfos
}

func GatherPostgresInfo(instance *entity.Instance) error {
	var wg sync.WaitGroup
	for _, f := range funcPostgresList {
		wg.Add(1)
		go f(instance, &wg)
	}
	wg.Wait()
	return nil
}

func GatherPostgresStatus(instance *entity.Instance) error {
	var wg sync.WaitGroup
	for _, f := range funcPostgresStatus {
		wg.Add(1)
		go f(instance, &wg)
	}
	wg.Wait()
	return nil
}
