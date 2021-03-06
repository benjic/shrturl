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

// Package faststore is a collection of stores that are optimzed and not
// neccesarly of the same interface as a url.Store
package faststore

// A FastStorer is an interface that allows URL models to be quickly stored.
type FastStorer interface {
	AddURL(URLModel)
	GetURL(id string) (URLModel, error)
	AllURLs() ([]URLModel, error)
}
