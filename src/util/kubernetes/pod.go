package kubernetes

import (
	"bufio"
	"context"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetLogs(instance *entity.Instance) (string, error) {
	client, err := NewKubernetesClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return "", err
	}

	req := client.CoreV1().Pods(instance.Namespace).GetLogs(instance.Name, &v1.PodLogOptions{Follow: true})
	readCloser, err := req.Stream(context.TODO())
	if err != nil {
		return "", err
	}
	defer func(readCloser io.ReadCloser) {
		err := readCloser.Close()
		if err != nil {
			return
		}
	}(readCloser)

	read := bufio.NewReader(readCloser)
	for {
		bytes, err := read.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			return string(bytes), nil
		}
	}
}

func GetPods(instance *entity.Instance) ([]string, error) {
	client, err := NewKubernetesClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}

	req, err := client.CoreV1().Pods(instance.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var podNames []string
	for _, item := range req.Items {
		podNames = append(podNames, item.Name)
	}
	return podNames, nil
}

func GetPodStatus(instance *entity.Instance) (*entity.PodStatus, error) {
	client, err := NewKubernetesClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	pod, err := client.CoreV1().Pods(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	podStatus := entity.PodStatus{}
	podStatus.Name = pod.Name
	podStatus.Phase = string(pod.Status.Phase)
	var podConditions []entity.PodCondition
	for _, condition := range pod.Status.Conditions {
		podCondition := entity.PodCondition{}
		podCondition.Status = string(condition.Status)
		podCondition.Message = condition.Message
		podCondition.Reason = condition.Reason
		podConditions = append(podConditions, podCondition)
	}
	podStatus.Conditions = podConditions
	return &podStatus, nil
}

func GetPodsStatus(instance *entity.Instance) ([]entity.PodStatus, error) {
	client, err := NewKubernetesClient(&Config{
		ApiServer: instance.Cluster.ApiServer,
		Token:     instance.Cluster.Token,
	})
	if err != nil {
		return nil, err
	}
	podList, err := client.CoreV1().Pods(instance.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var podsStatus []entity.PodStatus
	for _, pod := range podList.Items {
		podStatus := entity.PodStatus{}
		podStatus.Name = pod.Name
		podStatus.Phase = string(pod.Status.Phase)
		var podConditions []entity.PodCondition
		for _, condition := range pod.Status.Conditions {
			podCondition := entity.PodCondition{}
			podCondition.Status = string(condition.Status)
			podCondition.Message = condition.Message
			podCondition.Reason = condition.Reason
			podConditions = append(podConditions, podCondition)
		}
		podStatus.Conditions = podConditions
		podsStatus = append(podsStatus, podStatus)
	}
	return podsStatus, nil
}
