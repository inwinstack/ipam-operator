/*
Copyright © 2018 inwinSTACK.inc

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

package pool

import (
	"reflect"

	"github.com/golang/glog"
	inwinalphav1 "github.com/inwinstack/ipam-operator/pkg/apis/inwinstack/v1alpha1"
	inwinclientset "github.com/inwinstack/ipam-operator/pkg/client/clientset/versioned/typed/inwinstack/v1alpha1"
	opkit "github.com/inwinstack/operator-kit"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/client-go/tools/cache"
)

const (
	customResourceName       = "pool"
	customResourceNamePlural = "pools"
)

var Resource = opkit.CustomResource{
	Name:    customResourceName,
	Plural:  customResourceNamePlural,
	Group:   inwinalphav1.CustomResourceGroup,
	Version: inwinalphav1.Version,
	Scope:   apiextensionsv1beta1.ClusterScoped,
	Kind:    reflect.TypeOf(inwinalphav1.Pool{}).Name(),
}

type PoolController struct {
	ctx       *opkit.Context
	clientset inwinclientset.InwinstackV1alpha1Interface
}

func NewController(ctx *opkit.Context, clientset inwinclientset.InwinstackV1alpha1Interface) *PoolController {
	return &PoolController{ctx: ctx, clientset: clientset}
}

func (c *PoolController) StartWatch(namespace string, stopCh chan struct{}) error {
	resourceHandlerFuncs := cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	}

	glog.Infof("Start watching pool resources.")
	watcher := opkit.NewWatcher(Resource, namespace, resourceHandlerFuncs, c.clientset.RESTClient())
	go watcher.Watch(&inwinalphav1.Pool{}, stopCh)
	return nil
}

func (c *PoolController) onAdd(obj interface{}) {
	glog.Infof("Pool resource onAdd: %s.", obj.(*inwinalphav1.Pool).Name)
}

func (c *PoolController) onUpdate(oldObj, newObj interface{}) {
	glog.Infof("Pool resource onUpdate: %s.", newObj.(*inwinalphav1.Pool).Name)
}

func (c *PoolController) onDelete(obj interface{}) {
	glog.Infof("Pool resource onDelete: %s.", obj.(*inwinalphav1.Pool).Name)
}
