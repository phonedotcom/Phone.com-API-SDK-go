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

// Filter Object. Optional. See below for details.
type RuleSetFilter struct {

	// Required.
	Type_ string `json:"type,omitempty"`

	// Schedule that this filter refers to. Output is a Schedule Summary Object. Input must be a Schedule Lookup Object. Required.
	Schedule ScheduleSummary `json:"schedule,omitempty"`

	// Address Book Contact that this filter refers to. Output is a Contact Summary Object. Input must be a Contact Lookup Object. Required.
	Contact ContactSummary `json:"contact,omitempty"`

	// Address Book Group that this filter refers to. Output is a Contact Group Summary Object. Input must be a Contact Group Lookup Object. Required.
	Group GroupSummary `json:"group,omitempty"`
}
