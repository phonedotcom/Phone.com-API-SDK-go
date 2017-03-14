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

// Here are all of the properties you can expect for Caller ID objects:
type CallerIdPhoneNumber struct {

	// Name that this number will be associated with. Must be no longer than 15 characters, and can only contain English letters, numbers, spaces, and commas.
	Name string `json:"name,omitempty"`

	// Can be \"business\" or \"personal\"
	Type_ string `json:"type,omitempty"`
}