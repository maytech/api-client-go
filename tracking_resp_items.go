/*
 * Quatrix API
 *
 * The Quatrix API enables you to automate your interaction with Quatrix using any scripting or programming language. You can test API calls on [Swagger](https://docs.maytech.net/swagger/), the interactive environment we use as the repository for the Quatrix API. In order to authorize to Quatrix API you will need to have a Quatrix account https://yourcompanyname.quatrix.it/ (e.g. https://test-api.quatrix.it/). You can set up a free trial account [here](https://www.maytech.net/freetrial.html#trial=qtrx). Below you'll find information on how to authenticate in a Quatrix session, how to construct JSON formatted API call and what JSON responses to expect.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type TrackingRespItems struct {

	Id string `json:"id,omitempty"`

	Subject string `json:"subject,omitempty"`

	Type_ string `json:"type,omitempty"`

	IsReply bool `json:"is_reply,omitempty"`

	ReleaseDate float32 `json:"release_date,omitempty"`

	ExpireDate float32 `json:"expire_date,omitempty"`

	Created float32 `json:"created,omitempty"`

	Status string `json:"status,omitempty"`

	Email string `json:"email,omitempty"`

	Emails []string `json:"emails,omitempty"`

	OwnerId string `json:"owner_id,omitempty"`

	IsRequest bool `json:"is_request,omitempty"`

	CanReply bool `json:"can_reply,omitempty"`

	Folder string `json:"folder,omitempty"`
}