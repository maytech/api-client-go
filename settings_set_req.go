/*
 * Quatrix API
 *
 * The Quatrix API enables you to automate your interaction with Quatrix using any scripting or programming language. You can test API calls on [Swagger](https://docs.maytech.net/swagger/), the interactive environment we use as the repository for the Quatrix API. In order to authorize to Quatrix API you will need to have a Quatrix account https://yourcompanyname.quatrix.it/ (e.g. https://test-api.quatrix.it/). You can set up a free trial account [here](https://www.maytech.net/freetrial.html#trial=qtrx). Below you'll find information on how to authenticate in a Quatrix session, how to construct JSON formatted API call and what JSON responses to expect.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package quatrix

type SettingsSetReq struct {

	Title string `json:"title,omitempty"`

	Bcc []string `json:"bcc,omitempty"`

	BillingEmails []string `json:"billing_emails,omitempty"`

	EmailFooter string `json:"email_footer,omitempty"`

	Language string `json:"language,omitempty"`

	EnablePgp bool `json:"enable_pgp,omitempty"`

	VirusScan bool `json:"virus_scan,omitempty"`

	ShareTypes *interface{} `json:"share_types,omitempty"`

	AuthMethods []string `json:"auth_methods,omitempty"`
}
