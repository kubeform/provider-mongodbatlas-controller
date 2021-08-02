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
	triggersBasePath = appsBasePath + "/%s/triggers"
)

// EventTriggersService provides access to the event triggers related functions in the Realm API.
//
// See more: https://docs.mongodb.com/realm/admin/api/v3/#event-trigger-apis
type EventTriggersService interface {
	Create(context.Context, string, string, *EventTriggerRequest) (*EventTrigger, *Response, error)
	Get(context.Context, string, string, string) (*EventTrigger, *Response, error)
	List(context.Context, string, string) ([]EventTrigger, *Response, error)
	Update(context.Context, string, string, string, *EventTriggerRequest) (*EventTrigger, *Response, error)
	Delete(context.Context, string, string, string) (*Response, error)
}

// EventTriggersServiceOp provides an implementation of the EventTriggersService interface.
type EventTriggersServiceOp service

var _ EventTriggersService = &EventTriggersServiceOp{}

// Create one trigger.
//
// See more: https://docs.mongodb.com/realm/admin/api/v3/#post-/groups/%7Bgroupid%7D/apps/%7Bappid%7D/triggers
func (s *EventTriggersServiceOp) Create(ctx context.Context, groupID, appID string, createRequest *EventTriggerRequest) (*EventTrigger, *Response, error) {
	if groupID == "" {
		return nil, nil, atlas.NewArgError("groupId", "must be set")
	}
	if appID == "" {
		return nil, nil, atlas.NewArgError("appID", "must be set")
	}

	path := fmt.Sprintf(triggersBasePath, groupID, appID)

	req, err := s.Client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(EventTrigger)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Get Retrieve the configuration for a specific trigger.
//
// See more: https://docs.mongodb.com/realm/admin/api/v3/#get-/groups/%7Bgroupid%7D/apps/%7Bappid%7D/triggers/%7Btriggerid%7D
func (s *EventTriggersServiceOp) Get(ctx context.Context, groupID, appID, triggerID string) (*EventTrigger, *Response, error) {
	if groupID == "" {
		return nil, nil, atlas.NewArgError("groupId", "must be set")
	}
	if appID == "" {
		return nil, nil, atlas.NewArgError("appID", "must be set")
	}

	basePath := fmt.Sprintf(triggersBasePath, groupID, appID)
	path := fmt.Sprintf("%s/%s", basePath, triggerID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(EventTrigger)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// List all triggers.
//
// See more: https://docs.mongodb.com/realm/admin/api/v3/#responses-96
func (s EventTriggersServiceOp) List(ctx context.Context, groupID, appID string) ([]EventTrigger, *Response, error) {
	if groupID == "" {
		return nil, nil, atlas.NewArgError("groupId", "must be set")
	}
	if appID == "" {
		return nil, nil, atlas.NewArgError("appID", "must be set")
	}

	path := fmt.Sprintf(triggersBasePath, groupID, appID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var root []EventTrigger
	resp, err := s.Client.Do(ctx, req, &root)

	return root, resp, err
}

// Update updates a trigger.
//
// See more: https://docs.mongodb.com/realm/admin/api/v3/#put-/groups/%7Bgroupid%7D/apps/%7Bappid%7D/triggers/%7Btriggerid%7D
func (s *EventTriggersServiceOp) Update(ctx context.Context, groupID, appID, triggerID string, updateRequest *EventTriggerRequest) (*EventTrigger, *Response, error) {
	if groupID == "" {
		return nil, nil, atlas.NewArgError("groupId", "must be set")
	}
	if appID == "" {
		return nil, nil, atlas.NewArgError("appID", "must be set")
	}

	basePath := fmt.Sprintf(triggersBasePath, groupID, appID)
	path := fmt.Sprintf("%s/%s", basePath, triggerID)

	req, err := s.Client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(EventTrigger)
	resp, err := s.Client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Delete one trigger.
//
// See more https://docs.mongodb.com/realm/admin/api/v3/#delete-/groups/%7Bgroupid%7D/apps/%7Bappid%7D/triggers/%7Btriggerid%7D
func (s *EventTriggersServiceOp) Delete(ctx context.Context, groupID, appID, triggerID string) (*Response, error) {
	if groupID == "" {
		return nil, atlas.NewArgError("groupId", "must be set")
	}
	if appID == "" {
		return nil, atlas.NewArgError("appID", "must be set")
	}

	basePath := fmt.Sprintf(triggersBasePath, groupID, appID)
	path := fmt.Sprintf("%s/%s", basePath, triggerID)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.Client.Do(ctx, req, nil)
}

// EventTrigger Represents a response of a trigger
type EventTrigger struct {
	ID              string                 `json:"_id,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Type            string                 `json:"type,omitempty"`
	FunctionID      string                 `json:"function_id,omitempty"`
	FunctionName    string                 `json:"function_name,omitempty"`
	Disabled        *bool                  `json:"disabled,omitempty"`
	Config          EventTriggerConfig     `json:"config,omitempty"`
	EventProcessors map[string]interface{} `json:"event_processors,omitempty"`
	LastModified    *int64                 `json:"last_modified,omitempty"`
}

// EventTriggerRequest Represents a request of create a trigger
type EventTriggerRequest struct {
	Name            string                 `json:"name,omitempty"`
	Type            string                 `json:"type,omitempty"`
	FunctionID      string                 `json:"function_id,omitempty"`
	Disabled        *bool                  `json:"disabled,omitempty"`
	Config          *EventTriggerConfig    `json:"config,omitempty"`
	EventProcessors map[string]interface{} `json:"event_processors,omitempty"`
}

// EventTriggerConfig Represents a request of a trigger config
type EventTriggerConfig struct {
	OperationTypes           []string    `json:"operation_types,omitempty"`
	OperationType            string      `json:"operation_type,omitempty"`
	Providers                []string    `json:"providers,omitempty"`
	Database                 string      `json:"database,omitempty"`
	Collection               string      `json:"collection,omitempty"`
	ServiceID                string      `json:"service_id,omitempty"`
	Match                    interface{} `json:"match,omitempty"`
	Project                  interface{} `json:"project,omitempty"`
	FullDocument             *bool       `json:"full_document,omitempty"`
	FullDocumentBeforeChange *bool       `json:"full_document_before_change,omitempty"`
	Schedule                 string      `json:"schedule,omitempty"`
	ScheduleType             string      `json:"schedule_type,omitempty"`
	Unordered                *bool       `json:"unordered,omitempty"`
	ClusterName              string      `json:"clusterName,omitempty"`
}
