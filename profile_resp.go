/*
 * Quatrix API
 *
 * Download and upload files or folders, share them with predefined security options, manage your account and profile settings and a lot more functionalities can be easily integrated into your application using our Quatrix APIs. Learn more how to authenticate the Quatrix session, how to construct JSON formatted API calls and what responses to expect in our [API Guide](https://docs.maytech.net/display/MD/API+Guide).
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type ProfileResp struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Email string `json:"email,omitempty"`

	SuperAdmin string `json:"super_admin,omitempty"`

	HomeId string `json:"home_id,omitempty"`

	Operations float32 `json:"operations,omitempty"`

	MaxFileSize float32 `json:"max_file_size,omitempty"`

	MaxFilesPerShare float32 `json:"max_files_per_share,omitempty"`

	Plan string `json:"plan,omitempty"`

	IsContact bool `json:"is_contact,omitempty"`

	Language string `json:"language,omitempty"`

	ChannelId string `json:"channel_id,omitempty"`

	AccountId string `json:"account_id,omitempty"`

	StorageId string `json:"storage_id,omitempty"`

	Created float32 `json:"created,omitempty"`

	Modified float32 `json:"modified,omitempty"`

	Quota float32 `json:"quota,omitempty"`

	Status string `json:"status,omitempty"`

	Uid float32 `json:"uid,omitempty"`

	Gid float32 `json:"gid,omitempty"`

	HasKey bool `json:"has_key,omitempty"`

	PgpEnabled bool `json:"pgp_enabled,omitempty"`

	MessageSignature string `json:"message_signature,omitempty"`

	OutgoingId string `json:"outgoing_id,omitempty"`

	IncomingId string `json:"incoming_id,omitempty"`

	Services []ShortUserService `json:"services,omitempty"`

	ShareTypes *ProfileRespShareTypes `json:"share_types,omitempty"`

	UniqueLogin string `json:"unique_login,omitempty"`

	AuthMethods []string `json:"auth_methods,omitempty"`

	AccountStatus string `json:"account_status,omitempty"`

	SftpUrl string `json:"sftp_url,omitempty"`

	ForcedAuth []string `json:"forced_auth,omitempty"`
}
