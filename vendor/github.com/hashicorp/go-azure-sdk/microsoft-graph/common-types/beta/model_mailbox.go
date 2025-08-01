package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = Mailbox{}

type Mailbox struct {
	// The collection of folders in the mailbox.
	Folders *[]MailboxFolder `json:"folders,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s Mailbox) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s Mailbox) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Mailbox{}

func (s Mailbox) MarshalJSON() ([]byte, error) {
	type wrapper Mailbox
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Mailbox: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Mailbox: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mailbox"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Mailbox: %+v", err)
	}

	return encoded, nil
}
