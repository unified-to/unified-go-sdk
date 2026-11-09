# PropertyWorkspaceIntegrationAuthOutput


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ClientID`                                                         | `*string`                                                          | :heavy_minus_sign:                                                 | N/A                                                                |
| `ConsumerKey`                                                      | `*string`                                                          | :heavy_minus_sign:                                                 | N/A                                                                |
| `NeedsPkce`                                                        | `*bool`                                                            | :heavy_minus_sign:                                                 | opt-in PKCE for OAuth2 when the integration app or org requires it |
| `SamlIdpEntityID`                                                  | `*string`                                                          | :heavy_minus_sign:                                                 | the SAML IdP issuer/entity ID                                      |
| `SamlIdpURL`                                                       | `*string`                                                          | :heavy_minus_sign:                                                 | the SAML IdP SSO/login URL for end-user SAML login                 |
| `SamlPem`                                                          | `*string`                                                          | :heavy_minus_sign:                                                 | the SAML IdP signing certificate; PEM X.509 in Base64 ASCII format |