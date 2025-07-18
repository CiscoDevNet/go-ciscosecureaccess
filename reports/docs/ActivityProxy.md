# ActivityProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Externalip** | **string** | The external IP for the entry. | 
**Internalip** | **string** | The internal IP for the entry. | 
**Policycategories** | [**[]Category**](Category.md) | The list of policy categories. | 
**Categories** | [**[]Category**](Category.md) | The list of categories. | 
**Verdict** | [**Verdict**](Verdict.md) |  | 
**Timestamp** | **int64** | The timestamp represented in milliseconds. | 
**Identities** | [**[]Identity**](Identity.md) | The list of identities for the entry. | 
**Allapplications** | [**[]Application**](Application.md) | The list of applications for the entry. | 
**Allowedapplications** | [**[]Application**](Application.md) | The list of allowed applications for the entry. | 
**Blockedapplications** | [**[]Application**](Application.md) | The list of blocked applications for the entry. | 
**Responsefilename** | **string** | The response filename for the entry. | 
**Blockedfiletype** | **string** | The blocked file type for the entry. | 
**Bundleid** | **int64** | The ID of the bundle type. | 
**Amp** | [**CiscoAMP**](CiscoAMP.md) |  | 
**Type** | **string** | The type of the request. A proxy request is always of type &#39;proxy&#39;. | 
**Tenantcontrols** | **bool** | Specifies whether the request is part of a tenant control policy. | 
**Port** | **NullableInt64** | The port used to make the request. | 
**Antivirusthreats** | [**AntivirusThreats**](AntivirusThreats.md) |  | 
**Policy** | [**Policy**](Policy.md) |  | 
**Requestmethod** | Pointer to **string** | The HTTP request method. | [optional] 
**Responsesize** | **int64** | The response size in bytes. | 
**Requestsize** | **int64** | The response size in bytes. | 
**Statuscode** | **int64** | The HTTP status code (&#x60;200&#x60; or &#x60;201&#x60;). | 
**Useragent** | **string** | The name of the browser that made the request. | 
**Referer** | **string** | The referring domain or URL. | 
**Warnstatus** | **string** | The warning status. | 
**Sha256** | **string** | The hex digest of the response content. | 
**Isolated** | [**Isolated**](Isolated.md) |  | 
**Datalossprevention** | [**DataLossPreventionState**](DataLossPreventionState.md) |  | 
**Securityoverridden** | **bool** | Specifies whether security overrides are configured. | 
**Contenttype** | **string** | The type of web content, typically text/html. | 
**Forwardingmethod** | **string** | The request method (GET, POST, HEAD, etc.) | 
**Httperrors** | [**[]HttpError**](HttpError.md) | Certificate &amp; TLS Errors | 
**Threats** | [**[]Threat**](Threat.md) |  | 
**Egress** | [**Egress**](Egress.md) |  | 
**Datacenter** | [**DataCenter**](DataCenter.md) |  | 
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 
**Destinationip** | **string** | The destination IP for the entry. | 
**Url** | **string** | The URL that was requested. | 

## Methods

### NewActivityProxy

`func NewActivityProxy(externalip string, internalip string, policycategories []Category, categories []Category, verdict Verdict, timestamp int64, identities []Identity, allapplications []Application, allowedapplications []Application, blockedapplications []Application, responsefilename string, blockedfiletype string, bundleid int64, amp CiscoAMP, type_ string, tenantcontrols bool, port NullableInt64, antivirusthreats AntivirusThreats, policy Policy, responsesize int64, requestsize int64, statuscode int64, useragent string, referer string, warnstatus string, sha256 string, isolated Isolated, datalossprevention DataLossPreventionState, securityoverridden bool, contenttype string, forwardingmethod string, httperrors []HttpError, threats []Threat, egress Egress, datacenter DataCenter, date string, time string, destinationip string, url string, ) *ActivityProxy`

NewActivityProxy instantiates a new ActivityProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityProxyWithDefaults

`func NewActivityProxyWithDefaults() *ActivityProxy`

NewActivityProxyWithDefaults instantiates a new ActivityProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalip

`func (o *ActivityProxy) GetExternalip() string`

GetExternalip returns the Externalip field if non-nil, zero value otherwise.

### GetExternalipOk

`func (o *ActivityProxy) GetExternalipOk() (*string, bool)`

GetExternalipOk returns a tuple with the Externalip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalip

`func (o *ActivityProxy) SetExternalip(v string)`

SetExternalip sets Externalip field to given value.


### GetInternalip

`func (o *ActivityProxy) GetInternalip() string`

GetInternalip returns the Internalip field if non-nil, zero value otherwise.

### GetInternalipOk

`func (o *ActivityProxy) GetInternalipOk() (*string, bool)`

GetInternalipOk returns a tuple with the Internalip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalip

`func (o *ActivityProxy) SetInternalip(v string)`

SetInternalip sets Internalip field to given value.


### GetPolicycategories

`func (o *ActivityProxy) GetPolicycategories() []Category`

GetPolicycategories returns the Policycategories field if non-nil, zero value otherwise.

### GetPolicycategoriesOk

`func (o *ActivityProxy) GetPolicycategoriesOk() (*[]Category, bool)`

GetPolicycategoriesOk returns a tuple with the Policycategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicycategories

`func (o *ActivityProxy) SetPolicycategories(v []Category)`

SetPolicycategories sets Policycategories field to given value.


### GetCategories

`func (o *ActivityProxy) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ActivityProxy) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ActivityProxy) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetVerdict

`func (o *ActivityProxy) GetVerdict() Verdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ActivityProxy) GetVerdictOk() (*Verdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ActivityProxy) SetVerdict(v Verdict)`

SetVerdict sets Verdict field to given value.


### GetTimestamp

`func (o *ActivityProxy) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityProxy) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityProxy) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetIdentities

`func (o *ActivityProxy) GetIdentities() []Identity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ActivityProxy) GetIdentitiesOk() (*[]Identity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ActivityProxy) SetIdentities(v []Identity)`

SetIdentities sets Identities field to given value.


### GetAllapplications

`func (o *ActivityProxy) GetAllapplications() []Application`

GetAllapplications returns the Allapplications field if non-nil, zero value otherwise.

### GetAllapplicationsOk

`func (o *ActivityProxy) GetAllapplicationsOk() (*[]Application, bool)`

GetAllapplicationsOk returns a tuple with the Allapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllapplications

`func (o *ActivityProxy) SetAllapplications(v []Application)`

SetAllapplications sets Allapplications field to given value.


### GetAllowedapplications

`func (o *ActivityProxy) GetAllowedapplications() []Application`

GetAllowedapplications returns the Allowedapplications field if non-nil, zero value otherwise.

### GetAllowedapplicationsOk

`func (o *ActivityProxy) GetAllowedapplicationsOk() (*[]Application, bool)`

GetAllowedapplicationsOk returns a tuple with the Allowedapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedapplications

`func (o *ActivityProxy) SetAllowedapplications(v []Application)`

SetAllowedapplications sets Allowedapplications field to given value.


### GetBlockedapplications

`func (o *ActivityProxy) GetBlockedapplications() []Application`

GetBlockedapplications returns the Blockedapplications field if non-nil, zero value otherwise.

### GetBlockedapplicationsOk

`func (o *ActivityProxy) GetBlockedapplicationsOk() (*[]Application, bool)`

GetBlockedapplicationsOk returns a tuple with the Blockedapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedapplications

`func (o *ActivityProxy) SetBlockedapplications(v []Application)`

SetBlockedapplications sets Blockedapplications field to given value.


### GetResponsefilename

`func (o *ActivityProxy) GetResponsefilename() string`

GetResponsefilename returns the Responsefilename field if non-nil, zero value otherwise.

### GetResponsefilenameOk

`func (o *ActivityProxy) GetResponsefilenameOk() (*string, bool)`

GetResponsefilenameOk returns a tuple with the Responsefilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsefilename

`func (o *ActivityProxy) SetResponsefilename(v string)`

SetResponsefilename sets Responsefilename field to given value.


### GetBlockedfiletype

`func (o *ActivityProxy) GetBlockedfiletype() string`

GetBlockedfiletype returns the Blockedfiletype field if non-nil, zero value otherwise.

### GetBlockedfiletypeOk

`func (o *ActivityProxy) GetBlockedfiletypeOk() (*string, bool)`

GetBlockedfiletypeOk returns a tuple with the Blockedfiletype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedfiletype

`func (o *ActivityProxy) SetBlockedfiletype(v string)`

SetBlockedfiletype sets Blockedfiletype field to given value.


### GetBundleid

`func (o *ActivityProxy) GetBundleid() int64`

GetBundleid returns the Bundleid field if non-nil, zero value otherwise.

### GetBundleidOk

`func (o *ActivityProxy) GetBundleidOk() (*int64, bool)`

GetBundleidOk returns a tuple with the Bundleid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleid

`func (o *ActivityProxy) SetBundleid(v int64)`

SetBundleid sets Bundleid field to given value.


### GetAmp

`func (o *ActivityProxy) GetAmp() CiscoAMP`

GetAmp returns the Amp field if non-nil, zero value otherwise.

### GetAmpOk

`func (o *ActivityProxy) GetAmpOk() (*CiscoAMP, bool)`

GetAmpOk returns a tuple with the Amp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmp

`func (o *ActivityProxy) SetAmp(v CiscoAMP)`

SetAmp sets Amp field to given value.


### GetType

`func (o *ActivityProxy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityProxy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityProxy) SetType(v string)`

SetType sets Type field to given value.


### GetTenantcontrols

`func (o *ActivityProxy) GetTenantcontrols() bool`

GetTenantcontrols returns the Tenantcontrols field if non-nil, zero value otherwise.

### GetTenantcontrolsOk

`func (o *ActivityProxy) GetTenantcontrolsOk() (*bool, bool)`

GetTenantcontrolsOk returns a tuple with the Tenantcontrols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantcontrols

`func (o *ActivityProxy) SetTenantcontrols(v bool)`

SetTenantcontrols sets Tenantcontrols field to given value.


### GetPort

`func (o *ActivityProxy) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ActivityProxy) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ActivityProxy) SetPort(v int64)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *ActivityProxy) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ActivityProxy) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetAntivirusthreats

`func (o *ActivityProxy) GetAntivirusthreats() AntivirusThreats`

GetAntivirusthreats returns the Antivirusthreats field if non-nil, zero value otherwise.

### GetAntivirusthreatsOk

`func (o *ActivityProxy) GetAntivirusthreatsOk() (*AntivirusThreats, bool)`

GetAntivirusthreatsOk returns a tuple with the Antivirusthreats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusthreats

`func (o *ActivityProxy) SetAntivirusthreats(v AntivirusThreats)`

SetAntivirusthreats sets Antivirusthreats field to given value.


### GetPolicy

`func (o *ActivityProxy) GetPolicy() Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ActivityProxy) GetPolicyOk() (*Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ActivityProxy) SetPolicy(v Policy)`

SetPolicy sets Policy field to given value.


### GetRequestmethod

`func (o *ActivityProxy) GetRequestmethod() string`

GetRequestmethod returns the Requestmethod field if non-nil, zero value otherwise.

### GetRequestmethodOk

`func (o *ActivityProxy) GetRequestmethodOk() (*string, bool)`

GetRequestmethodOk returns a tuple with the Requestmethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestmethod

`func (o *ActivityProxy) SetRequestmethod(v string)`

SetRequestmethod sets Requestmethod field to given value.

### HasRequestmethod

`func (o *ActivityProxy) HasRequestmethod() bool`

HasRequestmethod returns a boolean if a field has been set.

### GetResponsesize

`func (o *ActivityProxy) GetResponsesize() int64`

GetResponsesize returns the Responsesize field if non-nil, zero value otherwise.

### GetResponsesizeOk

`func (o *ActivityProxy) GetResponsesizeOk() (*int64, bool)`

GetResponsesizeOk returns a tuple with the Responsesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsesize

`func (o *ActivityProxy) SetResponsesize(v int64)`

SetResponsesize sets Responsesize field to given value.


### GetRequestsize

`func (o *ActivityProxy) GetRequestsize() int64`

GetRequestsize returns the Requestsize field if non-nil, zero value otherwise.

### GetRequestsizeOk

`func (o *ActivityProxy) GetRequestsizeOk() (*int64, bool)`

GetRequestsizeOk returns a tuple with the Requestsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsize

`func (o *ActivityProxy) SetRequestsize(v int64)`

SetRequestsize sets Requestsize field to given value.


### GetStatuscode

`func (o *ActivityProxy) GetStatuscode() int64`

GetStatuscode returns the Statuscode field if non-nil, zero value otherwise.

### GetStatuscodeOk

`func (o *ActivityProxy) GetStatuscodeOk() (*int64, bool)`

GetStatuscodeOk returns a tuple with the Statuscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuscode

`func (o *ActivityProxy) SetStatuscode(v int64)`

SetStatuscode sets Statuscode field to given value.


### GetUseragent

`func (o *ActivityProxy) GetUseragent() string`

GetUseragent returns the Useragent field if non-nil, zero value otherwise.

### GetUseragentOk

`func (o *ActivityProxy) GetUseragentOk() (*string, bool)`

GetUseragentOk returns a tuple with the Useragent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseragent

`func (o *ActivityProxy) SetUseragent(v string)`

SetUseragent sets Useragent field to given value.


### GetReferer

`func (o *ActivityProxy) GetReferer() string`

GetReferer returns the Referer field if non-nil, zero value otherwise.

### GetRefererOk

`func (o *ActivityProxy) GetRefererOk() (*string, bool)`

GetRefererOk returns a tuple with the Referer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferer

`func (o *ActivityProxy) SetReferer(v string)`

SetReferer sets Referer field to given value.


### GetWarnstatus

`func (o *ActivityProxy) GetWarnstatus() string`

GetWarnstatus returns the Warnstatus field if non-nil, zero value otherwise.

### GetWarnstatusOk

`func (o *ActivityProxy) GetWarnstatusOk() (*string, bool)`

GetWarnstatusOk returns a tuple with the Warnstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnstatus

`func (o *ActivityProxy) SetWarnstatus(v string)`

SetWarnstatus sets Warnstatus field to given value.


### GetSha256

`func (o *ActivityProxy) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ActivityProxy) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ActivityProxy) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetIsolated

`func (o *ActivityProxy) GetIsolated() Isolated`

GetIsolated returns the Isolated field if non-nil, zero value otherwise.

### GetIsolatedOk

`func (o *ActivityProxy) GetIsolatedOk() (*Isolated, bool)`

GetIsolatedOk returns a tuple with the Isolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolated

`func (o *ActivityProxy) SetIsolated(v Isolated)`

SetIsolated sets Isolated field to given value.


### GetDatalossprevention

`func (o *ActivityProxy) GetDatalossprevention() DataLossPreventionState`

GetDatalossprevention returns the Datalossprevention field if non-nil, zero value otherwise.

### GetDatalosspreventionOk

`func (o *ActivityProxy) GetDatalosspreventionOk() (*DataLossPreventionState, bool)`

GetDatalosspreventionOk returns a tuple with the Datalossprevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatalossprevention

`func (o *ActivityProxy) SetDatalossprevention(v DataLossPreventionState)`

SetDatalossprevention sets Datalossprevention field to given value.


### GetSecurityoverridden

`func (o *ActivityProxy) GetSecurityoverridden() bool`

GetSecurityoverridden returns the Securityoverridden field if non-nil, zero value otherwise.

### GetSecurityoverriddenOk

`func (o *ActivityProxy) GetSecurityoverriddenOk() (*bool, bool)`

GetSecurityoverriddenOk returns a tuple with the Securityoverridden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityoverridden

`func (o *ActivityProxy) SetSecurityoverridden(v bool)`

SetSecurityoverridden sets Securityoverridden field to given value.


### GetContenttype

`func (o *ActivityProxy) GetContenttype() string`

GetContenttype returns the Contenttype field if non-nil, zero value otherwise.

### GetContenttypeOk

`func (o *ActivityProxy) GetContenttypeOk() (*string, bool)`

GetContenttypeOk returns a tuple with the Contenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContenttype

`func (o *ActivityProxy) SetContenttype(v string)`

SetContenttype sets Contenttype field to given value.


### GetForwardingmethod

`func (o *ActivityProxy) GetForwardingmethod() string`

GetForwardingmethod returns the Forwardingmethod field if non-nil, zero value otherwise.

### GetForwardingmethodOk

`func (o *ActivityProxy) GetForwardingmethodOk() (*string, bool)`

GetForwardingmethodOk returns a tuple with the Forwardingmethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingmethod

`func (o *ActivityProxy) SetForwardingmethod(v string)`

SetForwardingmethod sets Forwardingmethod field to given value.


### GetHttperrors

`func (o *ActivityProxy) GetHttperrors() []HttpError`

GetHttperrors returns the Httperrors field if non-nil, zero value otherwise.

### GetHttperrorsOk

`func (o *ActivityProxy) GetHttperrorsOk() (*[]HttpError, bool)`

GetHttperrorsOk returns a tuple with the Httperrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttperrors

`func (o *ActivityProxy) SetHttperrors(v []HttpError)`

SetHttperrors sets Httperrors field to given value.


### GetThreats

`func (o *ActivityProxy) GetThreats() []Threat`

GetThreats returns the Threats field if non-nil, zero value otherwise.

### GetThreatsOk

`func (o *ActivityProxy) GetThreatsOk() (*[]Threat, bool)`

GetThreatsOk returns a tuple with the Threats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreats

`func (o *ActivityProxy) SetThreats(v []Threat)`

SetThreats sets Threats field to given value.


### GetEgress

`func (o *ActivityProxy) GetEgress() Egress`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *ActivityProxy) GetEgressOk() (*Egress, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *ActivityProxy) SetEgress(v Egress)`

SetEgress sets Egress field to given value.


### GetDatacenter

`func (o *ActivityProxy) GetDatacenter() DataCenter`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *ActivityProxy) GetDatacenterOk() (*DataCenter, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *ActivityProxy) SetDatacenter(v DataCenter)`

SetDatacenter sets Datacenter field to given value.


### GetDate

`func (o *ActivityProxy) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityProxy) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityProxy) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *ActivityProxy) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ActivityProxy) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ActivityProxy) SetTime(v string)`

SetTime sets Time field to given value.


### GetDestinationip

`func (o *ActivityProxy) GetDestinationip() string`

GetDestinationip returns the Destinationip field if non-nil, zero value otherwise.

### GetDestinationipOk

`func (o *ActivityProxy) GetDestinationipOk() (*string, bool)`

GetDestinationipOk returns a tuple with the Destinationip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationip

`func (o *ActivityProxy) SetDestinationip(v string)`

SetDestinationip sets Destinationip field to given value.


### GetUrl

`func (o *ActivityProxy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ActivityProxy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ActivityProxy) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


