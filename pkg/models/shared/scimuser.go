// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ScimUser struct {
	Active            *bool                 `json:"active,omitempty"`
	Addresses         []ScimAddress         `json:"addresses,omitempty"`
	DisplayName       *string               `json:"displayName,omitempty"`
	Emails            []ScimEmail           `json:"emails,omitempty"`
	Entitlements      []ScimEntitlement     `json:"entitlements,omitempty"`
	ExternalID        *string               `json:"externalId,omitempty"`
	Groups            []ScimUserGroups      `json:"groups,omitempty"`
	ID                *string               `json:"id,omitempty"`
	Ims               []ScimIms             `json:"ims,omitempty"`
	Locale            *string               `json:"locale,omitempty"`
	Meta              *PropertyScimUserMeta `json:"meta,omitempty"`
	Name              *PropertyScimUserName `json:"name,omitempty"`
	NickName          *string               `json:"nickName,omitempty"`
	Password          *string               `json:"password,omitempty"`
	PhoneNumbers      []ScimPhoneNumber     `json:"phoneNumbers,omitempty"`
	Photos            []ScimPhoto           `json:"photos,omitempty"`
	PreferredLanguage *string               `json:"preferredLanguage,omitempty"`
	ProfileURL        *string               `json:"profileUrl,omitempty"`
	// Student, Faculty, ...
	Roles    []ScimRole                `json:"roles,omitempty"`
	Schemas  []PropertyScimUserSchemas `json:"schemas,omitempty"`
	Timezone *string                   `json:"timezone,omitempty"`
	Title    *string                   `json:"title,omitempty"`
	// an organization.
	UrnIetfParamsScimSchemasExtensionEnterprise20User        *PropertyScimUserUrnIetfParamsScimSchemasExtensionEnterprise20User        `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`
	UrnIetfParamsScimSchemasExtensionLatticeAttributes10User *PropertyScimUserUrnIetfParamsScimSchemasExtensionLatticeAttributes10User `json:"urn:ietf:params:scim:schemas:extension:lattice:attributes:1.0:User,omitempty"`
	UrnIetfParamsScimSchemasExtensionPeakon20User            *PropertyScimUserUrnIetfParamsScimSchemasExtensionPeakon20User            `json:"urn:ietf:params:scim:schemas:extension:peakon:2.0:User,omitempty"`
	UserName                                                 *string                                                                   `json:"userName,omitempty"`
	UserType                                                 *string                                                                   `json:"userType,omitempty"`
	X509Certificates                                         []ScimRole                                                                `json:"x509Certificates,omitempty"`
}

func (o *ScimUser) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ScimUser) GetAddresses() []ScimAddress {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *ScimUser) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *ScimUser) GetEmails() []ScimEmail {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *ScimUser) GetEntitlements() []ScimEntitlement {
	if o == nil {
		return nil
	}
	return o.Entitlements
}

func (o *ScimUser) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *ScimUser) GetGroups() []ScimUserGroups {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *ScimUser) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ScimUser) GetIms() []ScimIms {
	if o == nil {
		return nil
	}
	return o.Ims
}

func (o *ScimUser) GetLocale() *string {
	if o == nil {
		return nil
	}
	return o.Locale
}

func (o *ScimUser) GetMeta() *PropertyScimUserMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *ScimUser) GetName() *PropertyScimUserName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ScimUser) GetNickName() *string {
	if o == nil {
		return nil
	}
	return o.NickName
}

func (o *ScimUser) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ScimUser) GetPhoneNumbers() []ScimPhoneNumber {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *ScimUser) GetPhotos() []ScimPhoto {
	if o == nil {
		return nil
	}
	return o.Photos
}

func (o *ScimUser) GetPreferredLanguage() *string {
	if o == nil {
		return nil
	}
	return o.PreferredLanguage
}

func (o *ScimUser) GetProfileURL() *string {
	if o == nil {
		return nil
	}
	return o.ProfileURL
}

func (o *ScimUser) GetRoles() []ScimRole {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *ScimUser) GetSchemas() []PropertyScimUserSchemas {
	if o == nil {
		return nil
	}
	return o.Schemas
}

func (o *ScimUser) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *ScimUser) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *ScimUser) GetUrnIetfParamsScimSchemasExtensionEnterprise20User() *PropertyScimUserUrnIetfParamsScimSchemasExtensionEnterprise20User {
	if o == nil {
		return nil
	}
	return o.UrnIetfParamsScimSchemasExtensionEnterprise20User
}

func (o *ScimUser) GetUrnIetfParamsScimSchemasExtensionLatticeAttributes10User() *PropertyScimUserUrnIetfParamsScimSchemasExtensionLatticeAttributes10User {
	if o == nil {
		return nil
	}
	return o.UrnIetfParamsScimSchemasExtensionLatticeAttributes10User
}

func (o *ScimUser) GetUrnIetfParamsScimSchemasExtensionPeakon20User() *PropertyScimUserUrnIetfParamsScimSchemasExtensionPeakon20User {
	if o == nil {
		return nil
	}
	return o.UrnIetfParamsScimSchemasExtensionPeakon20User
}

func (o *ScimUser) GetUserName() *string {
	if o == nil {
		return nil
	}
	return o.UserName
}

func (o *ScimUser) GetUserType() *string {
	if o == nil {
		return nil
	}
	return o.UserType
}

func (o *ScimUser) GetX509Certificates() []ScimRole {
	if o == nil {
		return nil
	}
	return o.X509Certificates
}
