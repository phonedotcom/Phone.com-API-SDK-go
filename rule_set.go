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

type RuleSet struct {

	Filter RuleSetFilter `json:"filter,omitempty"`

	// Array of Action Objects. Required. Routes have rule sets, and rule sets have one or more actions. The supported actions are described below:
	Actions []RuleSetAction `json:"actions,omitempty"`
}