package kubernetes

import (
	log "github.com/sirupsen/logrus"
	api_core "k8s.io/api/core/v1"
	api_meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//GetEndpointList returns endpoints list
func (k *Kube) GetEndpointList(namespace string) (*api_core.EndpointsList, error) {
	endpointsAfter, err := k.CoreV1().Endpoints(namespace).List(api_meta.ListOptions{})
	if err != nil {
		log.WithFields(log.Fields{
			"Namespace": namespace,
		}).Error(err)
		return nil, err
	}
	return endpointsAfter, nil
}

//GetEndpoint returns endpoint
func (k *Kube) GetEndpoint(namespace, endpoint string) (*api_core.Endpoints, error) {
	endpointAfter, err := k.CoreV1().Endpoints(namespace).Get(endpoint, api_meta.GetOptions{})
	if err != nil {
		log.WithFields(log.Fields{
			"Namespace": namespace,
			"Endpoint":  endpoint,
		}).Error(err)
		return nil, err
	}
	return endpointAfter, nil
}

//CreateEndpoint creates endpoint
func (k *Kube) CreateEndpoint(endpoint *api_core.Endpoints) (*api_core.Endpoints, error) {
	endpointAfter, err := k.CoreV1().Endpoints(endpoint.Namespace).Create(endpoint)
	if err != nil {
		log.WithFields(log.Fields{
			"Namespace": endpoint.Namespace,
			"Endpoint":  endpoint.Name,
		}).Error(err)
		return nil, err
	}
	return endpointAfter, nil
}

//UpdateEndpoint updates endpoint
func (k *Kube) UpdateEndpoint(endpoint *api_core.Endpoints) (*api_core.Endpoints, error) {
	endpointAfter, err := k.CoreV1().Endpoints(endpoint.Namespace).Update(endpoint)
	if err != nil {
		log.WithFields(log.Fields{
			"Namespace": endpoint.Namespace,
			"Endpoint":  endpoint.Name,
		}).Error(err)
		return nil, err
	}
	return endpointAfter, nil
}

//DeleteEndpoint deletes endpoint
func (k *Kube) DeleteEndpoint(namespace, endpoint string) error {
	err := k.CoreV1().Endpoints(namespace).Delete(endpoint, &api_meta.DeleteOptions{})
	if err != nil {
		log.WithFields(log.Fields{
			"Namespace": namespace,
			"Endpoint":  endpoint,
		}).Error(err)
		return err
	}
	return nil
}
