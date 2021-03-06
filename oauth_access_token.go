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

type OauthAccessToken struct {

	AccessToken string `json:"access_token,omitempty"`

	TokenType string `json:"token_type,omitempty"`

	Scope string `json:"scope,omitempty"`

	RefreshToken string `json:"refresh_token,omitempty"`

	ExpiresIn int32 `json:"expires_in,omitempty"`
}
