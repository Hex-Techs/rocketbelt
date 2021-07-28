/*
Copyright 2021 The.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	rocketbeltv1alpha1 "github.com/hex-techs/rocketbelt/api/v1alpha1"
	versioned "github.com/hex-techs/rocketbelt/generated/clientset/versioned"
	internalinterfaces "github.com/hex-techs/rocketbelt/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/hex-techs/rocketbelt/generated/listers/rocketbelt/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HexRoleBindingInformer provides access to a shared informer and lister for
// HexRoleBindings.
type HexRoleBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HexRoleBindingLister
}

type hexRoleBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewHexRoleBindingInformer constructs a new informer for HexRoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHexRoleBindingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHexRoleBindingInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredHexRoleBindingInformer constructs a new informer for HexRoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHexRoleBindingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RocketbeltV1alpha1().HexRoleBindings().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RocketbeltV1alpha1().HexRoleBindings().Watch(context.TODO(), options)
			},
		},
		&rocketbeltv1alpha1.HexRoleBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *hexRoleBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHexRoleBindingInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hexRoleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rocketbeltv1alpha1.HexRoleBinding{}, f.defaultInformer)
}

func (f *hexRoleBindingInformer) Lister() v1alpha1.HexRoleBindingLister {
	return v1alpha1.NewHexRoleBindingLister(f.Informer().GetIndexer())
}