package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/get_pods", func(w http.ResponseWriter, r *http.Request) {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		var podInfoList []interface{}
		for _, pod := range pods.Items {
			podInfoList = append(podInfoList, map[string]interface{}{
				"name":      pod.GetName(),
				"namespace": pod.GetNamespace(),
			})
		}
		resp, _ := json.Marshal(map[string]interface{}{
			"message": fmt.Sprintf("get pods successful, there are %d pods", len(pods.Items)),
			"pods":    podInfoList,
		})
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	})

	fmt.Println("Running server on :8080")
	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}
