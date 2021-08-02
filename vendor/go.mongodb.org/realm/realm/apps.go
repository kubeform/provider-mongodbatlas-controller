// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package realm

import (
	"context"
	"fmt"
	"net/http"

	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const (
	appsBasePath = "groups/%s/apps"
)

// AppsService provides access to the applications related functions in the Realm API.
//
// See more: https://docs.mongodb.com/realm/admin/api/v3/#application-level-apis
type AppsService interface {
	List(context.Context, string, *ApplicationListOptions) ([]Application, *Response, error)
}

// AppsServiceOp provides an implementation of the AppsService interface.
type AppsServiceOp service

var _ AppsService = &AppsServiceOp{}

// List all Realm apps within an Atlas project/group.
//
// See more: https://docs.mongodb.com/realm/admin/api/v3/#application-level-apis
func (s *AppsServiceOp) List(ctx context.Context, groupID string, opts *ApplicationListOptions) ([]Application, *Response, error) {
	if groupID == "" {
		return nil, nil, atlas.NewArgError("groupId", "must be set")
	}
	basePath := fmt.Sprintf(appsBasePath, groupID)
	path, err := setQueryParams(basePath, opts)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var root []Application
	resp, err := s.Client.Do(ctx, req, &root)

	return root, resp, err
}

type ApplicationListOptions struct {
	Product string `url:"product,omitempty"`
}

type Application struct {
	ID              string `json:"_id,omitempty"`
	ClientAppID     string `json:"client_app_id,omitempty"`
	Name            string `json:"name,omitempty"`
	Location        string `json:"location,omitempty"`
	DeploymentModel string `json:"deployment_model,omitempty"`
	DomainID        string `json:"domain_id,omitempty"`
	GroupID         string `json:"group_id,omitempty"`
}
