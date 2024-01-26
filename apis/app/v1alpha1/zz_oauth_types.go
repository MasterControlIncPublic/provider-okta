/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GroupsClaimObservation struct {

	// Groups claim filter. Can only be set if type is FILTER.
	FilterType *string `json:"filterType,omitempty" tf:"filter_type,omitempty"`

	// Issuer mode inherited from OAuth App
	IssuerMode *string `json:"issuerMode,omitempty" tf:"issuer_mode,omitempty"`

	// Name of the claim that will be used in the token.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Groups claim type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Value of the claim. Can be an Okta Expression Language statement that evaluates at the time the token is minted.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GroupsClaimParameters struct {

	// Groups claim filter. Can only be set if type is FILTER.
	// +kubebuilder:validation:Optional
	FilterType *string `json:"filterType,omitempty" tf:"filter_type,omitempty"`

	// Name of the claim that will be used in the token.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Groups claim type.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Value of the claim. Can be an Okta Expression Language statement that evaluates at the time the token is minted.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type JwksObservation struct {

	// RSA Exponent
	E *string `json:"e,omitempty" tf:"e,omitempty"`

	// Key ID
	Kid *string `json:"kid,omitempty" tf:"kid,omitempty"`

	// Key type
	Kty *string `json:"kty,omitempty" tf:"kty,omitempty"`

	// RSA Modulus
	N *string `json:"n,omitempty" tf:"n,omitempty"`

	// X coordinate of the elliptic curve point
	X *string `json:"x,omitempty" tf:"x,omitempty"`

	// Y coordinate of the elliptic curve point
	Y *string `json:"y,omitempty" tf:"y,omitempty"`
}

type JwksParameters struct {

	// RSA Exponent
	// +kubebuilder:validation:Optional
	E *string `json:"e,omitempty" tf:"e,omitempty"`

	// Key ID
	// +kubebuilder:validation:Required
	Kid *string `json:"kid" tf:"kid,omitempty"`

	// Key type
	// +kubebuilder:validation:Required
	Kty *string `json:"kty" tf:"kty,omitempty"`

	// RSA Modulus
	// +kubebuilder:validation:Optional
	N *string `json:"n,omitempty" tf:"n,omitempty"`

	// X coordinate of the elliptic curve point
	// +kubebuilder:validation:Optional
	X *string `json:"x,omitempty" tf:"x,omitempty"`

	// Y coordinate of the elliptic curve point
	// +kubebuilder:validation:Optional
	Y *string `json:"y,omitempty" tf:"y,omitempty"`
}

type OauthObservation struct {

	// Custom error page URL
	AccessibilityErrorRedirectURL *string `json:"accessibilityErrorRedirectUrl,omitempty" tf:"accessibility_error_redirect_url,omitempty"`

	// Custom login page URL
	AccessibilityLoginRedirectURL *string `json:"accessibilityLoginRedirectUrl,omitempty" tf:"accessibility_login_redirect_url,omitempty"`

	// Enable self service
	AccessibilitySelfService *bool `json:"accessibilitySelfService,omitempty" tf:"accessibility_self_service,omitempty"`

	// Application notes for admins.
	AdminNote *string `json:"adminNote,omitempty" tf:"admin_note,omitempty"`

	// Displays specific appLinks for the app
	AppLinksJSON *string `json:"appLinksJson,omitempty" tf:"app_links_json,omitempty"`

	// Application settings in JSON format
	AppSettingsJSON *string `json:"appSettingsJson,omitempty" tf:"app_settings_json,omitempty"`

	// Id of this apps authentication policy
	AuthenticationPolicy *string `json:"authenticationPolicy,omitempty" tf:"authentication_policy,omitempty"`

	// Requested key rotation mode.
	AutoKeyRotation *bool `json:"autoKeyRotation,omitempty" tf:"auto_key_rotation,omitempty"`

	// Display auto submit toolbar
	AutoSubmitToolbar *bool `json:"autoSubmitToolbar,omitempty" tf:"auto_submit_toolbar,omitempty"`

	// OAuth client ID. If set during creation, app is created with this id.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// URI to a web page providing information about the client.
	ClientURI *string `json:"clientUri,omitempty" tf:"client_uri,omitempty"`

	// *Early Access Property*. Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED
	ConsentMethod *string `json:"consentMethod,omitempty" tf:"consent_method,omitempty"`

	// Application notes for end users.
	EnduserNote *string `json:"enduserNote,omitempty" tf:"enduser_note,omitempty"`

	// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
	GrantTypes []*string `json:"grantTypes,omitempty" tf:"grant_types,omitempty"`

	// Groups claim for an OpenID Connect client application (argument is ignored when API auth is done with OAuth 2.0 credentials)
	GroupsClaim []GroupsClaimObservation `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	// Do not display application icon on mobile app
	HideIos *bool `json:"hideIos,omitempty" tf:"hide_ios,omitempty"`

	// Do not display application icon to users
	HideWeb *bool `json:"hideWeb,omitempty" tf:"hide_web,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// *Early Access Property*. Enable Federation Broker Mode.
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty" tf:"implicit_assignment,omitempty"`

	// *Early Access Property*. Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
	IssuerMode *string `json:"issuerMode,omitempty" tf:"issuer_mode,omitempty"`

	Jwks []JwksObservation `json:"jwks,omitempty" tf:"jwks,omitempty"`

	// URL reference to JWKS
	JwksURI *string `json:"jwksUri,omitempty" tf:"jwks_uri,omitempty"`

	// Pretty name of app.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The type of Idp-Initiated login that the client supports, if any
	LoginMode *string `json:"loginMode,omitempty" tf:"login_mode,omitempty"`

	// List of scopes to use for the request
	LoginScopes []*string `json:"loginScopes,omitempty" tf:"login_scopes,omitempty"`

	// URI that initiates login.
	LoginURI *string `json:"loginUri,omitempty" tf:"login_uri,omitempty"`

	// Local path to logo of the application.
	Logo *string `json:"logo,omitempty" tf:"logo,omitempty"`

	// URI that references a logo for the client.
	LogoURI *string `json:"logoUri,omitempty" tf:"logo_uri,omitempty"`

	// URL of the application's logo
	LogoURL *string `json:"logoUrl,omitempty" tf:"logo_url,omitempty"`

	// Name of the app.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
	OmitSecret *bool `json:"omitSecret,omitempty" tf:"omit_secret,omitempty"`

	// Require Proof Key for Code Exchange (PKCE) for additional verification key rotation mode. See: https://developer.okta.com/docs/reference/api/apps/#oauth-credential-object
	PkceRequired *bool `json:"pkceRequired,omitempty" tf:"pkce_required,omitempty"`

	// URI to web page providing client policy document.
	PolicyURI *string `json:"policyUri,omitempty" tf:"policy_uri,omitempty"`

	// List of URIs for redirection after logout. Note: see okta_app_oauth_post_logout_redirect_uri for appending to this list in a decentralized way.
	PostLogoutRedirectUris []*string `json:"postLogoutRedirectUris,omitempty" tf:"post_logout_redirect_uris,omitempty"`

	// Custom JSON that represents an OAuth application's profile
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// List of URIs for use in the redirect-based flow. This is required for all application types except service. Note: see okta_app_oauth_redirect_uri for appending to this list in a decentralized way.
	RedirectUris []*string `json:"redirectUris,omitempty" tf:"redirect_uris,omitempty"`

	// *Early Access Property* Grace period for token rotation, required with grant types refresh_token
	RefreshTokenLeeway *float64 `json:"refreshTokenLeeway,omitempty" tf:"refresh_token_leeway,omitempty"`

	// *Early Access Property* Refresh token rotation behavior, required with grant types refresh_token
	RefreshTokenRotation *string `json:"refreshTokenRotation,omitempty" tf:"refresh_token_rotation,omitempty"`

	// List of OAuth 2.0 response type strings.
	ResponseTypes []*string `json:"responseTypes,omitempty" tf:"response_types,omitempty"`

	// Sign on mode of application.
	SignOnMode *string `json:"signOnMode,omitempty" tf:"sign_on_mode,omitempty"`

	// Status of application.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Requested authentication method for the token endpoint.
	TokenEndpointAuthMethod *string `json:"tokenEndpointAuthMethod,omitempty" tf:"token_endpoint_auth_method,omitempty"`

	// URI to web page providing client tos (terms of service).
	TosURI *string `json:"tosUri,omitempty" tf:"tos_uri,omitempty"`

	// The type of client application.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Username template
	UserNameTemplate *string `json:"userNameTemplate,omitempty" tf:"user_name_template,omitempty"`

	// Push username on update
	UserNameTemplatePushStatus *string `json:"userNameTemplatePushStatus,omitempty" tf:"user_name_template_push_status,omitempty"`

	// Username template suffix
	UserNameTemplateSuffix *string `json:"userNameTemplateSuffix,omitempty" tf:"user_name_template_suffix,omitempty"`

	// Username template type
	UserNameTemplateType *string `json:"userNameTemplateType,omitempty" tf:"user_name_template_type,omitempty"`

	// *Early Access Property*. Indicates if the client is allowed to use wildcard matching of redirect_uris
	WildcardRedirect *string `json:"wildcardRedirect,omitempty" tf:"wildcard_redirect,omitempty"`
}

type OauthParameters struct {

	// Custom error page URL
	// +kubebuilder:validation:Optional
	AccessibilityErrorRedirectURL *string `json:"accessibilityErrorRedirectUrl,omitempty" tf:"accessibility_error_redirect_url,omitempty"`

	// Custom login page URL
	// +kubebuilder:validation:Optional
	AccessibilityLoginRedirectURL *string `json:"accessibilityLoginRedirectUrl,omitempty" tf:"accessibility_login_redirect_url,omitempty"`

	// Enable self service
	// +kubebuilder:validation:Optional
	AccessibilitySelfService *bool `json:"accessibilitySelfService,omitempty" tf:"accessibility_self_service,omitempty"`

	// Application notes for admins.
	// +kubebuilder:validation:Optional
	AdminNote *string `json:"adminNote,omitempty" tf:"admin_note,omitempty"`

	// Displays specific appLinks for the app
	// +kubebuilder:validation:Optional
	AppLinksJSON *string `json:"appLinksJson,omitempty" tf:"app_links_json,omitempty"`

	// Application settings in JSON format
	// +kubebuilder:validation:Optional
	AppSettingsJSON *string `json:"appSettingsJson,omitempty" tf:"app_settings_json,omitempty"`

	// Id of this apps authentication policy
	// +kubebuilder:validation:Optional
	AuthenticationPolicy *string `json:"authenticationPolicy,omitempty" tf:"authentication_policy,omitempty"`

	// Requested key rotation mode.
	// +kubebuilder:validation:Optional
	AutoKeyRotation *bool `json:"autoKeyRotation,omitempty" tf:"auto_key_rotation,omitempty"`

	// Display auto submit toolbar
	// +kubebuilder:validation:Optional
	AutoSubmitToolbar *bool `json:"autoSubmitToolbar,omitempty" tf:"auto_submit_toolbar,omitempty"`

	// OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
	// +kubebuilder:validation:Optional
	ClientBasicSecretSecretRef *v1.SecretKeySelector `json:"clientBasicSecretSecretRef,omitempty" tf:"-"`

	// OAuth client ID. If set during creation, app is created with this id.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// URI to a web page providing information about the client.
	// +kubebuilder:validation:Optional
	ClientURI *string `json:"clientUri,omitempty" tf:"client_uri,omitempty"`

	// *Early Access Property*. Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED
	// +kubebuilder:validation:Optional
	ConsentMethod *string `json:"consentMethod,omitempty" tf:"consent_method,omitempty"`

	// Application notes for end users.
	// +kubebuilder:validation:Optional
	EnduserNote *string `json:"enduserNote,omitempty" tf:"enduser_note,omitempty"`

	// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
	// +kubebuilder:validation:Optional
	GrantTypes []*string `json:"grantTypes,omitempty" tf:"grant_types,omitempty"`

	// Groups claim for an OpenID Connect client application (argument is ignored when API auth is done with OAuth 2.0 credentials)
	// +kubebuilder:validation:Optional
	GroupsClaim []GroupsClaimParameters `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	// Do not display application icon on mobile app
	// +kubebuilder:validation:Optional
	HideIos *bool `json:"hideIos,omitempty" tf:"hide_ios,omitempty"`

	// Do not display application icon to users
	// +kubebuilder:validation:Optional
	HideWeb *bool `json:"hideWeb,omitempty" tf:"hide_web,omitempty"`

	// *Early Access Property*. Enable Federation Broker Mode.
	// +kubebuilder:validation:Optional
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty" tf:"implicit_assignment,omitempty"`

	// *Early Access Property*. Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
	// +kubebuilder:validation:Optional
	IssuerMode *string `json:"issuerMode,omitempty" tf:"issuer_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Jwks []JwksParameters `json:"jwks,omitempty" tf:"jwks,omitempty"`

	// URL reference to JWKS
	// +kubebuilder:validation:Optional
	JwksURI *string `json:"jwksUri,omitempty" tf:"jwks_uri,omitempty"`

	// Pretty name of app.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The type of Idp-Initiated login that the client supports, if any
	// +kubebuilder:validation:Optional
	LoginMode *string `json:"loginMode,omitempty" tf:"login_mode,omitempty"`

	// List of scopes to use for the request
	// +kubebuilder:validation:Optional
	LoginScopes []*string `json:"loginScopes,omitempty" tf:"login_scopes,omitempty"`

	// URI that initiates login.
	// +kubebuilder:validation:Optional
	LoginURI *string `json:"loginUri,omitempty" tf:"login_uri,omitempty"`

	// Local path to logo of the application.
	// +kubebuilder:validation:Optional
	Logo *string `json:"logo,omitempty" tf:"logo,omitempty"`

	// URI that references a logo for the client.
	// +kubebuilder:validation:Optional
	LogoURI *string `json:"logoUri,omitempty" tf:"logo_uri,omitempty"`

	// This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.
	// +kubebuilder:validation:Optional
	OmitSecret *bool `json:"omitSecret,omitempty" tf:"omit_secret,omitempty"`

	// Require Proof Key for Code Exchange (PKCE) for additional verification key rotation mode. See: https://developer.okta.com/docs/reference/api/apps/#oauth-credential-object
	// +kubebuilder:validation:Optional
	PkceRequired *bool `json:"pkceRequired,omitempty" tf:"pkce_required,omitempty"`

	// URI to web page providing client policy document.
	// +kubebuilder:validation:Optional
	PolicyURI *string `json:"policyUri,omitempty" tf:"policy_uri,omitempty"`

	// List of URIs for redirection after logout. Note: see okta_app_oauth_post_logout_redirect_uri for appending to this list in a decentralized way.
	// +kubebuilder:validation:Optional
	PostLogoutRedirectUris []*string `json:"postLogoutRedirectUris,omitempty" tf:"post_logout_redirect_uris,omitempty"`

	// Custom JSON that represents an OAuth application's profile
	// +kubebuilder:validation:Optional
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// List of URIs for use in the redirect-based flow. This is required for all application types except service. Note: see okta_app_oauth_redirect_uri for appending to this list in a decentralized way.
	// +kubebuilder:validation:Optional
	RedirectUris []*string `json:"redirectUris,omitempty" tf:"redirect_uris,omitempty"`

	// *Early Access Property* Grace period for token rotation, required with grant types refresh_token
	// +kubebuilder:validation:Optional
	RefreshTokenLeeway *float64 `json:"refreshTokenLeeway,omitempty" tf:"refresh_token_leeway,omitempty"`

	// *Early Access Property* Refresh token rotation behavior, required with grant types refresh_token
	// +kubebuilder:validation:Optional
	RefreshTokenRotation *string `json:"refreshTokenRotation,omitempty" tf:"refresh_token_rotation,omitempty"`

	// List of OAuth 2.0 response type strings.
	// +kubebuilder:validation:Optional
	ResponseTypes []*string `json:"responseTypes,omitempty" tf:"response_types,omitempty"`

	// Status of application.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Requested authentication method for the token endpoint.
	// +kubebuilder:validation:Optional
	TokenEndpointAuthMethod *string `json:"tokenEndpointAuthMethod,omitempty" tf:"token_endpoint_auth_method,omitempty"`

	// URI to web page providing client tos (terms of service).
	// +kubebuilder:validation:Optional
	TosURI *string `json:"tosUri,omitempty" tf:"tos_uri,omitempty"`

	// The type of client application.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Username template
	// +kubebuilder:validation:Optional
	UserNameTemplate *string `json:"userNameTemplate,omitempty" tf:"user_name_template,omitempty"`

	// Push username on update
	// +kubebuilder:validation:Optional
	UserNameTemplatePushStatus *string `json:"userNameTemplatePushStatus,omitempty" tf:"user_name_template_push_status,omitempty"`

	// Username template suffix
	// +kubebuilder:validation:Optional
	UserNameTemplateSuffix *string `json:"userNameTemplateSuffix,omitempty" tf:"user_name_template_suffix,omitempty"`

	// Username template type
	// +kubebuilder:validation:Optional
	UserNameTemplateType *string `json:"userNameTemplateType,omitempty" tf:"user_name_template_type,omitempty"`

	// *Early Access Property*. Indicates if the client is allowed to use wildcard matching of redirect_uris
	// +kubebuilder:validation:Optional
	WildcardRedirect *string `json:"wildcardRedirect,omitempty" tf:"wildcard_redirect,omitempty"`
}

// OauthSpec defines the desired state of Oauth
type OauthSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OauthParameters `json:"forProvider"`
}

// OauthStatus defines the observed state of Oauth.
type OauthStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OauthObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Oauth is the Schema for the Oauths API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,okta}
type Oauth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.label)",message="label is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.type)",message="type is a required parameter"
	Spec   OauthSpec   `json:"spec"`
	Status OauthStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OauthList contains a list of Oauths
type OauthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Oauth `json:"items"`
}

// Repository type metadata.
var (
	Oauth_Kind             = "Oauth"
	Oauth_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Oauth_Kind}.String()
	Oauth_KindAPIVersion   = Oauth_Kind + "." + CRDGroupVersion.String()
	Oauth_GroupVersionKind = CRDGroupVersion.WithKind(Oauth_Kind)
)

func init() {
	SchemeBuilder.Register(&Oauth{}, &OauthList{})
}
