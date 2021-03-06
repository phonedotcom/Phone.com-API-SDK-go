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
type RuleSetAction struct {

	// Required.
	Action string `json:"action,omitempty"`

	// Extension that this action refers to. Output is an Extension Summary Object. Input must be an Extension Lookup Object. Required.
	Extension ExtensionSummary `json:"extension,omitempty"`

	// This action is for forwarding calls to any number of extensions or phone numbers. The forwarding is handled in parallel, meaning that all phone numbers and/or extensions will ring simultaneously. When the call is answered by any single phone number or extension, ringing will stop for all of them. Subsequent actions in this rule set will be performed if the call is not answered before the timeout period is reached, or if it is forwarded to an extension that has its own route and that route does not result in any actions that disconnect the call or take over call handling.
	Items []RuleSetForwardItem `json:"items,omitempty"`

	// Seconds that our routing engine should wait until moving on. Optional. Must be an integer between 5 and 90. Default is 5 seconds.
	Timeout int32 `json:"timeout,omitempty"`

	// Hold Music to be played while callers are waiting. Output is a Media Summary Object. Input must be a Media Lookup Object. Optional. Must refer to a media recording that has is_hold_music set to TRUE. Default is to play a standard ring tone.
	HoldMusic MediaSummary `json:"hold_music,omitempty"`

	// Greeting that this action refers to. Output is a Media Summary Object. Input must be a Media Lookup Object. Required. Must refer to a media recording that has is_hold_music set to FALSE.
	Greeting MediaSummary `json:"greeting,omitempty"`

	// Required. Seconds that the caller should be placed on hold before being moved onto the next action. Must be an integer between 1 and 60 seconds.
	Duration int32 `json:"duration,omitempty"`

	// Menu that this action refers to. Required. Output is a Menu Summary Object. Input must be a Menu Lookup Object.
	Menu MenuSummary `json:"menu,omitempty"`

	// Queue that this action refers to. Required. Output is a Queue Summary Object. Input must be a Queue Lookup Object.
	Queue QueueSummary `json:"queue,omitempty"`

	// Trunk that this action refers to. Required. Output is a Trunk Summary Object. Input must be a Trunk Lookup Object.
	Trunk TrunkSummary `json:"trunk,omitempty"`
}
