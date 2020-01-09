/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package batch

import (
	v1 "github.com/rancher/wrangler-api/pkg/generated/controllers/batch/v1"
	v1beta1 "github.com/rancher/wrangler-api/pkg/generated/controllers/batch/v1beta1"
	"github.com/rancher/wrangler/pkg/generic"
	informers "k8s.io/client-go/informers/batch"
	clientset "k8s.io/client-go/kubernetes"
)

type Interface interface {
	V1() v1.Interface
	V1beta1() v1beta1.Interface
}

type group struct {
	controllerManager *generic.ControllerManager
	informers         informers.Interface
	client            clientset.Interface
}

// New returns a new Interface.
func New(controllerManager *generic.ControllerManager, informers informers.Interface,
	client clientset.Interface) Interface {
	return &group{
		controllerManager: controllerManager,
		informers:         informers,
		client:            client,
	}
}

func (g *group) V1() v1.Interface {
	return v1.New(g.controllerManager, g.client.BatchV1(), g.informers.V1())
}

func (g *group) V1beta1() v1beta1.Interface {
	return v1beta1.New(g.controllerManager, g.client.BatchV1beta1(), g.informers.V1beta1())
}
