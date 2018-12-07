/*
Copyright 2018 Pusher Ltd.

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

package v1alpha1

import (
	v1alpha1 "github.com/pusher/faros/pkg/apis/faros/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterGitTrackObjectLister helps list ClusterGitTrackObjects.
type ClusterGitTrackObjectLister interface {
	// List lists all ClusterGitTrackObjects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterGitTrackObject, err error)
	// Get retrieves the ClusterGitTrackObject from the index for a given name.
	Get(name string) (*v1alpha1.ClusterGitTrackObject, error)
	ClusterGitTrackObjectListerExpansion
}

// clusterGitTrackObjectLister implements the ClusterGitTrackObjectLister interface.
type clusterGitTrackObjectLister struct {
	indexer cache.Indexer
}

// NewClusterGitTrackObjectLister returns a new ClusterGitTrackObjectLister.
func NewClusterGitTrackObjectLister(indexer cache.Indexer) ClusterGitTrackObjectLister {
	return &clusterGitTrackObjectLister{indexer: indexer}
}

// List lists all ClusterGitTrackObjects in the indexer.
func (s *clusterGitTrackObjectLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterGitTrackObject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterGitTrackObject))
	})
	return ret, err
}

// Get retrieves the ClusterGitTrackObject from the index for a given name.
func (s *clusterGitTrackObjectLister) Get(name string) (*v1alpha1.ClusterGitTrackObject, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustergittrackobject"), name)
	}
	return obj.(*v1alpha1.ClusterGitTrackObject), nil
}