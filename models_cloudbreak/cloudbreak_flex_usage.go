package models_cloudbreak

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*CloudbreakFlexUsage cloudbreak flex usage

swagger:model CloudbreakFlexUsage
*/
type CloudbreakFlexUsage struct {

	/* account id of the resource owner that is provided by OAuth provider
	 */
	Account *string `json:"account,omitempty"`

	/* name of the blueprint
	 */
	BlueprintName *string `json:"blueprintName,omitempty"`

	/* the day the usage of resources happened
	 */
	Day *string `json:"day,omitempty"`

	/* flex subscription id
	 */
	FlexID *string `json:"flexId,omitempty"`

	/* number of instances running
	 */
	InstanceNum *int32 `json:"instanceNum,omitempty"`

	/* id of the resource owner that is provided by OAuth provider
	 */
	Owner *string `json:"owner,omitempty"`

	/* unique id of the controller instance
	 */
	ParentUUID *string `json:"parentUuid,omitempty"`

	/* maximum number of instances running
	 */
	Peak *int32 `json:"peak,omitempty"`

	/* region of the stack
	 */
	Region *string `json:"region,omitempty"`

	/* Smartsense subscription id
	 */
	SmartSenseID *string `json:"smartSenseId,omitempty"`

	/* name of the stack
	 */
	StackName *string `json:"stackName,omitempty"`

	/* unique id of the cluster
	 */
	StackUUID *string `json:"stackUuid,omitempty"`

	/* ambari username
	 */
	Username *string `json:"username,omitempty"`
}

// Validate validates this cloudbreak flex usage
func (m *CloudbreakFlexUsage) Validate(formats strfmt.Registry) error {
	return nil
}
