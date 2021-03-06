/*
 * Quatrix API
 *
 * Download and upload files or folders, share them with predefined security options, manage your account and profile settings and a lot more functionalities can be easily integrated into your application using our Quatrix APIs. Learn more how to authenticate the Quatrix session, how to construct JSON formatted API calls and what responses to expect in our [API Guide](https://docs.maytech.net/display/MD/API+Guide).
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type FileRenameResp struct {

	// File ID
	Id string `json:"id,omitempty"`

	// File creation timestamp
	Created float32 `json:"created,omitempty"`

	// File modification timestamp
	Modified float32 `json:"modified,omitempty"`

	// File name
	Name string `json:"name,omitempty"`

	// File parent ID
	ParentId string `json:"parent_id,omitempty"`

	// File size in bytes
	Size float32 `json:"size,omitempty"`

	// File type
	Type_ string `json:"type,omitempty"`

	// File metadata json
	Metadata *interface{} `json:"metadata,omitempty"`

	// File operations bitmask
	Operations float32 `json:"operations,omitempty"`

	// File name before rename
	OldName string `json:"old_name,omitempty"`
}
