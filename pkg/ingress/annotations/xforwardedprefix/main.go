/*
Copyright 2017 The Kubernetes Authors.

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

package xforwardedprefix

import (
	extensions "k8s.io/api/extensions/v1beta1"

	"github.ibm.com/IBMPrivateCloud/icp-management-ingress/pkg/ingress/annotations/parser"
	"github.ibm.com/IBMPrivateCloud/icp-management-ingress/pkg/ingress/resolver"
)

type xforwardedprefix struct {
	r resolver.Resolver
}

// NewParser creates a new xforwardedprefix annotation parser
func NewParser(r resolver.Resolver) parser.IngressAnnotation {
	return xforwardedprefix{r}
}

// Parse parses the annotations contained in the ingress rule
// used to add an x-forwarded-prefix header to the request
func (cbbs xforwardedprefix) Parse(ing *extensions.Ingress) (interface{}, error) {
	return parser.GetBoolAnnotation("x-forwarded-prefix", ing)
}