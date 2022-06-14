package kubernetes

import (
	"context"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	v1 "github.com/zalando/postgres-operator/pkg/apis/acid.zalan.do/v1"
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreatePostgres(instance *entity.Instance) (*v1.Postgresql, error) {
	client, err := NewPostgresqlClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	postgresInstance := &v1.Postgresql{
		TypeMeta: metav1.TypeMeta{
			Kind:       "postgresql",
			APIVersion: "acid.zalan.do/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		},
		Spec: v1.PostgresSpec{
			TeamID: instance.Namespace,
			Volume: v1.Volume{
				Size: instance.Volume,
			},
			Resources: &v1.Resources{
				ResourceRequests: v1.ResourceDescription{
					CPU:    instance.RequestCpu,
					Memory: instance.RequestMemory,
				},
				ResourceLimits: v1.ResourceDescription{
					CPU:    instance.LimitCpu,
					Memory: instance.LimitMemory,
				},
			},
			NumberOfInstances: int32(instance.Count),
			Users: map[string]v1.UserFlags{
				"super": {
					"superuser",
					"createdb",
				},
			},
			PostgresqlParam: v1.PostgresqlParam{
				PgVersion: "14",
			},
		},
	}
	postsqlObj, err := client.AcidV1().Postgresqls(instance.Namespace).Create(context.TODO(), postgresInstance, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return postsqlObj, nil
}

func CreatePostgresFromTemplate(instance *entity.Postgres, postgresql *v1.Postgresql) (*v1.Postgresql, error) {
	client, err := NewPostgresqlClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	postgresql.Name = instance.Name
	postgresql.Namespace = instance.Namespace
	postgresql.Spec.TeamID = instance.Namespace
	postsqlObj, err := client.AcidV1().Postgresqls(instance.Namespace).Create(context.TODO(), postgresql, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return postsqlObj, nil
}

func DeletePostgres(instance *entity.Instance) error {
	client, err := NewPostgresqlClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return err
	}
	err = client.AcidV1().Postgresqls(instance.Namespace).Delete(context.TODO(), instance.Name, metav1.DeleteOptions{})
	return err
}

func UpdatePostgres(instance *entity.Instance) (*v1.Postgresql, error) {
	client, err := NewPostgresqlClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	postgresInstance := &v1.Postgresql{
		TypeMeta: metav1.TypeMeta{
			Kind:       "postgresql",
			APIVersion: "acid.zalan.do/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		},
		Spec: v1.PostgresSpec{
			TeamID: instance.Namespace,
			Volume: v1.Volume{
				Size: instance.Volume,
			},
			Resources: &v1.Resources{
				ResourceRequests: v1.ResourceDescription{
					CPU:    instance.RequestCpu,
					Memory: instance.RequestMemory,
				},
				ResourceLimits: v1.ResourceDescription{
					CPU:    instance.LimitCpu,
					Memory: instance.LimitMemory,
				},
			},
			NumberOfInstances: int32(instance.Count),
			Users: map[string]v1.UserFlags{
				"super": v1.UserFlags{
					"superuser",
					"createdb",
				},
			},
			PostgresqlParam: v1.PostgresqlParam{
				PgVersion: "14",
			},
		},
	}
	postsqlObj, err := client.AcidV1().Postgresqls(instance.Namespace).Update(context.TODO(), postgresInstance, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return postsqlObj, nil
}

func EditPostgres(instance *entity.Instance) (*v1.Postgresql, error) {
	client, err := NewPostgresqlClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	var postgresql *v1.Postgresql
	err = yaml.Unmarshal([]byte(instance.Yaml), &postgresql)
	if err != nil {
		return nil, err
	}
	postgresql.ObjectMeta.ResourceVersion = ""
	postgresql.ObjectMeta.UID = ""
	postgresql.ObjectMeta.SelfLink = ""
	res, err := client.AcidV1().Postgresqls(instance.Namespace).Update(context.TODO(), postgresql, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetPostgres(instance *entity.Instance) (*v1.Postgresql, error) {
	client, err := NewPostgresqlClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	res, err := client.AcidV1().Postgresqls(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
