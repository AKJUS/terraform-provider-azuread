package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamsAdministrationTeamsAdminRoot{}

type TeamsAdministrationTeamsAdminRoot struct {
	Policy *TeamsAdministrationTeamsPolicyAssignment `json:"policy,omitempty"`

	// Represents the configuration information of users who have accounts hosted on Microsoft Teams.
	UserConfigurations *[]TeamsAdministrationTeamsUserConfiguration `json:"userConfigurations,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TeamsAdministrationTeamsAdminRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamsAdministrationTeamsAdminRoot{}

func (s TeamsAdministrationTeamsAdminRoot) MarshalJSON() ([]byte, error) {
	type wrapper TeamsAdministrationTeamsAdminRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamsAdministrationTeamsAdminRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamsAdministrationTeamsAdminRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamsAdministration.teamsAdminRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamsAdministrationTeamsAdminRoot: %+v", err)
	}

	return encoded, nil
}
