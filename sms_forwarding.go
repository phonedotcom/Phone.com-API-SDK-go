/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// SMS Forwarding Object, or NULL
type SmsForwarding struct {

	// Can be \"extension\" or \"application\"
	Type_ string `json:"type,omitempty"`

	// Required if type = \"extension\". Extension that messages should be directed to. Output is an Extension Summary Object. Input must be an Extension Lookup Object.
	Extension ExtensionSummary `json:"extension,omitempty"`

	// Required if type = \"application\". Application that messages should be directed to. Output is an Application Summary Object. Input must be an Application Lookup Object.
	Application ApplicationSummary `json:"application,omitempty"`
}