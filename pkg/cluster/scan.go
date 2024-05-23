package cluster

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// ClusterScanner defines a struct for scanning the cluster
type ClusterScanner struct {
	clientset *kubernetes.Clientset
}

// NewClusterScanner creates a new ClusterScanner
func NewClusterScanner(kubeconfig string) (*ClusterScanner, error) {
	config, err := buildConfig(kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &ClusterScanner{clientset: clientset}, nil
}

// ScanNodes lists all nodes in the cluster
func (cs *ClusterScanner) ScanNodes(ctx context.Context) error {
	nodes, err := cs.clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, node := range nodes.Items {
		fmt.Printf("Node Name: %s\n", node.Name)
	}
	return nil
}

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}
