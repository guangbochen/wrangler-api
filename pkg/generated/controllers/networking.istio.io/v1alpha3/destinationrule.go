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

package v1alpha3

import (
	"context"

	v1alpha3 "github.com/knative/pkg/apis/istio/v1alpha3"
	clientset "github.com/knative/pkg/client/clientset/versioned/typed/istio/v1alpha3"
	informers "github.com/knative/pkg/client/informers/externalversions/istio/v1alpha3"
	listers "github.com/knative/pkg/client/listers/istio/v1alpha3"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type DestinationRuleHandler func(string, *v1alpha3.DestinationRule) (*v1alpha3.DestinationRule, error)

type DestinationRuleController interface {
	DestinationRuleClient

	OnChange(ctx context.Context, name string, sync DestinationRuleHandler)
	OnRemove(ctx context.Context, name string, sync DestinationRuleHandler)
	Enqueue(namespace, name string)

	Cache() DestinationRuleCache

	Informer() cache.SharedIndexInformer
	GroupVersionKind() schema.GroupVersionKind

	AddGenericHandler(ctx context.Context, name string, handler generic.Handler)
	AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler)
	Updater() generic.Updater
}

type DestinationRuleClient interface {
	Create(*v1alpha3.DestinationRule) (*v1alpha3.DestinationRule, error)
	Update(*v1alpha3.DestinationRule) (*v1alpha3.DestinationRule, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1alpha3.DestinationRule, error)
	List(namespace string, opts metav1.ListOptions) (*v1alpha3.DestinationRuleList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.DestinationRule, err error)
}

type DestinationRuleCache interface {
	Get(namespace, name string) (*v1alpha3.DestinationRule, error)
	List(namespace string, selector labels.Selector) ([]*v1alpha3.DestinationRule, error)

	AddIndexer(indexName string, indexer DestinationRuleIndexer)
	GetByIndex(indexName, key string) ([]*v1alpha3.DestinationRule, error)
}

type DestinationRuleIndexer func(obj *v1alpha3.DestinationRule) ([]string, error)

type destinationRuleController struct {
	controllerManager *generic.ControllerManager
	clientGetter      clientset.DestinationRulesGetter
	informer          informers.DestinationRuleInformer
	gvk               schema.GroupVersionKind
}

func NewDestinationRuleController(gvk schema.GroupVersionKind, controllerManager *generic.ControllerManager, clientGetter clientset.DestinationRulesGetter, informer informers.DestinationRuleInformer) DestinationRuleController {
	return &destinationRuleController{
		controllerManager: controllerManager,
		clientGetter:      clientGetter,
		informer:          informer,
		gvk:               gvk,
	}
}

func FromDestinationRuleHandlerToHandler(sync DestinationRuleHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1alpha3.DestinationRule
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1alpha3.DestinationRule))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *destinationRuleController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1alpha3.DestinationRule))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateDestinationRuleOnChange(updater generic.Updater, handler DestinationRuleHandler) DestinationRuleHandler {
	return func(key string, obj *v1alpha3.DestinationRule) (*v1alpha3.DestinationRule, error) {
		if obj == nil {
			return handler(key, nil)
		}

		copyObj := obj.DeepCopy()
		newObj, err := handler(key, copyObj)
		if newObj != nil {
			copyObj = newObj
		}
		if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
			newObj, err := updater(copyObj)
			if newObj != nil && err == nil {
				copyObj = newObj.(*v1alpha3.DestinationRule)
			}
		}

		return copyObj, err
	}
}

func (c *destinationRuleController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, handler)
}

func (c *destinationRuleController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), handler)
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, removeHandler)
}

func (c *destinationRuleController) OnChange(ctx context.Context, name string, sync DestinationRuleHandler) {
	c.AddGenericHandler(ctx, name, FromDestinationRuleHandlerToHandler(sync))
}

func (c *destinationRuleController) OnRemove(ctx context.Context, name string, sync DestinationRuleHandler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), FromDestinationRuleHandlerToHandler(sync))
	c.AddGenericHandler(ctx, name, removeHandler)
}

func (c *destinationRuleController) Enqueue(namespace, name string) {
	c.controllerManager.Enqueue(c.gvk, c.informer.Informer(), namespace, name)
}

func (c *destinationRuleController) Informer() cache.SharedIndexInformer {
	return c.informer.Informer()
}

func (c *destinationRuleController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *destinationRuleController) Cache() DestinationRuleCache {
	return &destinationRuleCache{
		lister:  c.informer.Lister(),
		indexer: c.informer.Informer().GetIndexer(),
	}
}

func (c *destinationRuleController) Create(obj *v1alpha3.DestinationRule) (*v1alpha3.DestinationRule, error) {
	return c.clientGetter.DestinationRules(obj.Namespace).Create(obj)
}

func (c *destinationRuleController) Update(obj *v1alpha3.DestinationRule) (*v1alpha3.DestinationRule, error) {
	return c.clientGetter.DestinationRules(obj.Namespace).Update(obj)
}

func (c *destinationRuleController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return c.clientGetter.DestinationRules(namespace).Delete(name, options)
}

func (c *destinationRuleController) Get(namespace, name string, options metav1.GetOptions) (*v1alpha3.DestinationRule, error) {
	return c.clientGetter.DestinationRules(namespace).Get(name, options)
}

func (c *destinationRuleController) List(namespace string, opts metav1.ListOptions) (*v1alpha3.DestinationRuleList, error) {
	return c.clientGetter.DestinationRules(namespace).List(opts)
}

func (c *destinationRuleController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientGetter.DestinationRules(namespace).Watch(opts)
}

func (c *destinationRuleController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.DestinationRule, err error) {
	return c.clientGetter.DestinationRules(namespace).Patch(name, pt, data, subresources...)
}

type destinationRuleCache struct {
	lister  listers.DestinationRuleLister
	indexer cache.Indexer
}

func (c *destinationRuleCache) Get(namespace, name string) (*v1alpha3.DestinationRule, error) {
	return c.lister.DestinationRules(namespace).Get(name)
}

func (c *destinationRuleCache) List(namespace string, selector labels.Selector) ([]*v1alpha3.DestinationRule, error) {
	return c.lister.DestinationRules(namespace).List(selector)
}

func (c *destinationRuleCache) AddIndexer(indexName string, indexer DestinationRuleIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1alpha3.DestinationRule))
		},
	}))
}

func (c *destinationRuleCache) GetByIndex(indexName, key string) (result []*v1alpha3.DestinationRule, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		result = append(result, obj.(*v1alpha3.DestinationRule))
	}
	return result, nil
}
