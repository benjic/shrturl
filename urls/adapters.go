// Copyright 2015 Benjamin Campbell <benji@benjica.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package urls

import "github.com/benjic/shrturl/faststore"

// A URLFastStorerAdapter provides a interface wrapper for a faststore.FastStorer
// mapping client calls to differing backing stores.
type URLFastStorerAdapter struct {
	store faststore.FastStorer
}

// createURLFastStorerAdapter provides a convience factory function for creating
// URLFastStorerAdapters.
func createURLFastStorerAdapter(store faststore.FastStorer) *URLFastStorerAdapter {
	return &URLFastStorerAdapter{store: store}
}

// list implements the urls.Store interface by providing an aggregation of
// models for available urls.
func (a *URLFastStorerAdapter) list() (models []urlModel) {
	fastModels, _ := a.store.AllURLs()
	for _, model := range fastModels {
		models = append(models, urlModel{modelID(model.Slug), model.URL})
	}

	return models
}

// find implements the urls.Store interface by providing a way of getting
// a single model for a given id.
func (a *URLFastStorerAdapter) find(id modelID) (model urlModel, err error) {
	fastModel, err := a.store.GetURL(string(id))

	if err != nil {
		return urlModel{}, errorModelNotFound
	}

	return urlModel{ID: modelID(fastModel.Slug), URL: fastModel.URL}, err
}

// add implements the urls.Store interface by providing a way to add a new model
// to the underlying store.
func (a *URLFastStorerAdapter) add(model urlModel) {
	a.store.AddURL(faststore.URLModel{Slug: string(model.ID), URL: model.URL})
}
