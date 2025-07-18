# GetActivities200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Externalip** | **string** | The external IP for the entry. | 
**Internalip** | **string** | The internal IP for the entry. | 
**Policycategories** | [**[]Category**](Category.md) | The list of policy categories. | 
**Categories** | [**[]Category**](Category.md) | The list of categories. | 
**Verdict** | **string** | The verdict for entry. | 
**Domain** | **string** | The domain name for the entry. | 
**Timestamp** | **int64** | The timestamp represented in milliseconds. | 
**Identities** | [**[]Identity**](Identity.md) | The list of identities for the entry. | 
**Allapplications** | [**[]ActivityZTNAAllapplicationsInner**](ActivityZTNAAllapplicationsInner.md) | The list of private applications that are connected through Zero Trust Access. | 
**Threats** | [**[]Threat**](Threat.md) |  | 
**Type** | **string** | type of the request. Decryption | 
**Querytype** | **string** | The type of the DNS request. | 
**Date** | **string** | The date from the timestamp based on the timezone parameter. | 
**Time** | **string** | The time in 24-hour format based on the timezone parameter. | 
**Returncode** | **int64** | The DNS return code for this request. | 
**Allowedapplications** | [**[]Application**](Application.md) | The list of allowed applications for the entry. | 
**Blockedapplications** | [**[]Application**](Application.md) | The list of blocked applications for the entry. | 
**Destinationip** | **string** | The resolved IP for Zero Trust Network Access (ZTNA) client-based events. | 
**Sourceip** | **string** | The source IP for the entry. | 
**Sourceport** | **int64** | The source port for the entry. | 
**Destinationport** | **int64** | The destination port for entry. | 
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Rule** | [**Rule**](Rule.md) |  | 
**Applicationprotocols** | [**[]FirewallApplication**](FirewallApplication.md) | A list of firewall application protocols. | 
**Direction** | **string** | The direction of the packet. It is destined either towards the internet or to the customer&#39;s network. | 
**Packetsize** | **int64** | The size of the packet that was received. | 
**Classification** | **string** | The category of attack detected by a rule that is part of a more general type of attack class, such as trojan-activity, attempted-user, and unknown. | 
**Sessionid** | **int64** | The unique identifier of a session, which is used to group the correlated events between various services. | 
**Severity** | [**Severity**](Severity.md) |  | 
**Signature** | [**Signature**](Signature.md) |  | 
**Signaturelist** | [**SignatureList**](SignatureList.md) |  | 
**Responsefilename** | **string** | The response filename for the entry. | 
**Blockedfiletype** | **string** | The blocked file type for the entry. | 
**Bundleid** | **int64** | The ID of the bundle type. | 
**Amp** | [**CiscoAMP**](CiscoAMP.md) |  | 
**Tenantcontrols** | **bool** | Specifies whether the request is part of a tenant control policy. | 
**Port** | **NullableInt64** | The port used to make the request. | 
**Antivirusthreats** | [**AntivirusThreats**](AntivirusThreats.md) |  | 
**Policy** | [**PolicyZTNA**](PolicyZTNA.md) |  | 
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
**Egress** | [**Egress**](Egress.md) |  | 
**Datacenter** | [**DataCenter**](DataCenter.md) |  | 
**Url** | **string** | The URL that was requested. | 
**Clientfirewall** | Pointer to **string** | The type of the firewall that is used, either &#x60;SYS&#x60; or &#x60;NONE&#x60;. | [optional] 
**Diskencryption** | Pointer to **string** | Type of the disk encryption, either &#x60;SYS&#x60;, &#x60;THIRDPARTY&#x60;, &#x60;NONE&#x60;. | [optional] 
**Antimalwareagents** | Pointer to **[]string** | The list of anti-malware agents. | [optional] 
**Systempassword** | Pointer to **string** | The system password. | [optional] 
**Device** | Pointer to [**Device**](Device.md) |  | [optional] 
**DecryptAction** | **string** | Type of decryption action (Decrypt Inbound, Decrypt Outbound, Do Not Decrypt, Decrypt Error). | 

## Methods

### NewGetActivities200ResponseDataInner

`func NewGetActivities200ResponseDataInner(externalip string, internalip string, policycategories []Category, categories []Category, verdict string, domain string, timestamp int64, identities []Identity, allapplications []ActivityZTNAAllapplicationsInner, threats []Threat, type_ string, querytype string, date string, time string, returncode int64, allowedapplications []Application, blockedapplications []Application, destinationip string, sourceip string, sourceport int64, destinationport int64, protocol Protocol, rule Rule, applicationprotocols []FirewallApplication, direction string, packetsize int64, classification string, sessionid int64, severity Severity, signature Signature, signaturelist SignatureList, responsefilename string, blockedfiletype string, bundleid int64, amp CiscoAMP, tenantcontrols bool, port NullableInt64, antivirusthreats AntivirusThreats, policy PolicyZTNA, responsesize int64, requestsize int64, statuscode int64, useragent string, referer string, warnstatus string, sha256 string, isolated Isolated, datalossprevention DataLossPreventionState, securityoverridden bool, contenttype string, forwardingmethod string, httperrors []HttpError, egress Egress, datacenter DataCenter, url string, decryptAction string, ) *GetActivities200ResponseDataInner`

NewGetActivities200ResponseDataInner instantiates a new GetActivities200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivities200ResponseDataInnerWithDefaults

`func NewGetActivities200ResponseDataInnerWithDefaults() *GetActivities200ResponseDataInner`

NewGetActivities200ResponseDataInnerWithDefaults instantiates a new GetActivities200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalip

`func (o *GetActivities200ResponseDataInner) GetExternalip() string`

GetExternalip returns the Externalip field if non-nil, zero value otherwise.

### GetExternalipOk

`func (o *GetActivities200ResponseDataInner) GetExternalipOk() (*string, bool)`

GetExternalipOk returns a tuple with the Externalip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalip

`func (o *GetActivities200ResponseDataInner) SetExternalip(v string)`

SetExternalip sets Externalip field to given value.


### GetInternalip

`func (o *GetActivities200ResponseDataInner) GetInternalip() string`

GetInternalip returns the Internalip field if non-nil, zero value otherwise.

### GetInternalipOk

`func (o *GetActivities200ResponseDataInner) GetInternalipOk() (*string, bool)`

GetInternalipOk returns a tuple with the Internalip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalip

`func (o *GetActivities200ResponseDataInner) SetInternalip(v string)`

SetInternalip sets Internalip field to given value.


### GetPolicycategories

`func (o *GetActivities200ResponseDataInner) GetPolicycategories() []Category`

GetPolicycategories returns the Policycategories field if non-nil, zero value otherwise.

### GetPolicycategoriesOk

`func (o *GetActivities200ResponseDataInner) GetPolicycategoriesOk() (*[]Category, bool)`

GetPolicycategoriesOk returns a tuple with the Policycategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicycategories

`func (o *GetActivities200ResponseDataInner) SetPolicycategories(v []Category)`

SetPolicycategories sets Policycategories field to given value.


### GetCategories

`func (o *GetActivities200ResponseDataInner) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *GetActivities200ResponseDataInner) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *GetActivities200ResponseDataInner) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetVerdict

`func (o *GetActivities200ResponseDataInner) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *GetActivities200ResponseDataInner) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *GetActivities200ResponseDataInner) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetDomain

`func (o *GetActivities200ResponseDataInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetActivities200ResponseDataInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetActivities200ResponseDataInner) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetTimestamp

`func (o *GetActivities200ResponseDataInner) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetActivities200ResponseDataInner) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetActivities200ResponseDataInner) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetIdentities

`func (o *GetActivities200ResponseDataInner) GetIdentities() []Identity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *GetActivities200ResponseDataInner) GetIdentitiesOk() (*[]Identity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *GetActivities200ResponseDataInner) SetIdentities(v []Identity)`

SetIdentities sets Identities field to given value.


### GetAllapplications

`func (o *GetActivities200ResponseDataInner) GetAllapplications() []ActivityZTNAAllapplicationsInner`

GetAllapplications returns the Allapplications field if non-nil, zero value otherwise.

### GetAllapplicationsOk

`func (o *GetActivities200ResponseDataInner) GetAllapplicationsOk() (*[]ActivityZTNAAllapplicationsInner, bool)`

GetAllapplicationsOk returns a tuple with the Allapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllapplications

`func (o *GetActivities200ResponseDataInner) SetAllapplications(v []ActivityZTNAAllapplicationsInner)`

SetAllapplications sets Allapplications field to given value.


### GetThreats

`func (o *GetActivities200ResponseDataInner) GetThreats() []Threat`

GetThreats returns the Threats field if non-nil, zero value otherwise.

### GetThreatsOk

`func (o *GetActivities200ResponseDataInner) GetThreatsOk() (*[]Threat, bool)`

GetThreatsOk returns a tuple with the Threats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreats

`func (o *GetActivities200ResponseDataInner) SetThreats(v []Threat)`

SetThreats sets Threats field to given value.


### GetType

`func (o *GetActivities200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetActivities200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetActivities200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.


### GetQuerytype

`func (o *GetActivities200ResponseDataInner) GetQuerytype() string`

GetQuerytype returns the Querytype field if non-nil, zero value otherwise.

### GetQuerytypeOk

`func (o *GetActivities200ResponseDataInner) GetQuerytypeOk() (*string, bool)`

GetQuerytypeOk returns a tuple with the Querytype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerytype

`func (o *GetActivities200ResponseDataInner) SetQuerytype(v string)`

SetQuerytype sets Querytype field to given value.


### GetDate

`func (o *GetActivities200ResponseDataInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetActivities200ResponseDataInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetActivities200ResponseDataInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *GetActivities200ResponseDataInner) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetActivities200ResponseDataInner) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetActivities200ResponseDataInner) SetTime(v string)`

SetTime sets Time field to given value.


### GetReturncode

`func (o *GetActivities200ResponseDataInner) GetReturncode() int64`

GetReturncode returns the Returncode field if non-nil, zero value otherwise.

### GetReturncodeOk

`func (o *GetActivities200ResponseDataInner) GetReturncodeOk() (*int64, bool)`

GetReturncodeOk returns a tuple with the Returncode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturncode

`func (o *GetActivities200ResponseDataInner) SetReturncode(v int64)`

SetReturncode sets Returncode field to given value.


### GetAllowedapplications

`func (o *GetActivities200ResponseDataInner) GetAllowedapplications() []Application`

GetAllowedapplications returns the Allowedapplications field if non-nil, zero value otherwise.

### GetAllowedapplicationsOk

`func (o *GetActivities200ResponseDataInner) GetAllowedapplicationsOk() (*[]Application, bool)`

GetAllowedapplicationsOk returns a tuple with the Allowedapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedapplications

`func (o *GetActivities200ResponseDataInner) SetAllowedapplications(v []Application)`

SetAllowedapplications sets Allowedapplications field to given value.


### GetBlockedapplications

`func (o *GetActivities200ResponseDataInner) GetBlockedapplications() []Application`

GetBlockedapplications returns the Blockedapplications field if non-nil, zero value otherwise.

### GetBlockedapplicationsOk

`func (o *GetActivities200ResponseDataInner) GetBlockedapplicationsOk() (*[]Application, bool)`

GetBlockedapplicationsOk returns a tuple with the Blockedapplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedapplications

`func (o *GetActivities200ResponseDataInner) SetBlockedapplications(v []Application)`

SetBlockedapplications sets Blockedapplications field to given value.


### GetDestinationip

`func (o *GetActivities200ResponseDataInner) GetDestinationip() string`

GetDestinationip returns the Destinationip field if non-nil, zero value otherwise.

### GetDestinationipOk

`func (o *GetActivities200ResponseDataInner) GetDestinationipOk() (*string, bool)`

GetDestinationipOk returns a tuple with the Destinationip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationip

`func (o *GetActivities200ResponseDataInner) SetDestinationip(v string)`

SetDestinationip sets Destinationip field to given value.


### GetSourceip

`func (o *GetActivities200ResponseDataInner) GetSourceip() string`

GetSourceip returns the Sourceip field if non-nil, zero value otherwise.

### GetSourceipOk

`func (o *GetActivities200ResponseDataInner) GetSourceipOk() (*string, bool)`

GetSourceipOk returns a tuple with the Sourceip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceip

`func (o *GetActivities200ResponseDataInner) SetSourceip(v string)`

SetSourceip sets Sourceip field to given value.


### GetSourceport

`func (o *GetActivities200ResponseDataInner) GetSourceport() int64`

GetSourceport returns the Sourceport field if non-nil, zero value otherwise.

### GetSourceportOk

`func (o *GetActivities200ResponseDataInner) GetSourceportOk() (*int64, bool)`

GetSourceportOk returns a tuple with the Sourceport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceport

`func (o *GetActivities200ResponseDataInner) SetSourceport(v int64)`

SetSourceport sets Sourceport field to given value.


### GetDestinationport

`func (o *GetActivities200ResponseDataInner) GetDestinationport() int64`

GetDestinationport returns the Destinationport field if non-nil, zero value otherwise.

### GetDestinationportOk

`func (o *GetActivities200ResponseDataInner) GetDestinationportOk() (*int64, bool)`

GetDestinationportOk returns a tuple with the Destinationport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationport

`func (o *GetActivities200ResponseDataInner) SetDestinationport(v int64)`

SetDestinationport sets Destinationport field to given value.


### GetProtocol

`func (o *GetActivities200ResponseDataInner) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetActivities200ResponseDataInner) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetActivities200ResponseDataInner) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetRule

`func (o *GetActivities200ResponseDataInner) GetRule() Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *GetActivities200ResponseDataInner) GetRuleOk() (*Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *GetActivities200ResponseDataInner) SetRule(v Rule)`

SetRule sets Rule field to given value.


### GetApplicationprotocols

`func (o *GetActivities200ResponseDataInner) GetApplicationprotocols() []FirewallApplication`

GetApplicationprotocols returns the Applicationprotocols field if non-nil, zero value otherwise.

### GetApplicationprotocolsOk

`func (o *GetActivities200ResponseDataInner) GetApplicationprotocolsOk() (*[]FirewallApplication, bool)`

GetApplicationprotocolsOk returns a tuple with the Applicationprotocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationprotocols

`func (o *GetActivities200ResponseDataInner) SetApplicationprotocols(v []FirewallApplication)`

SetApplicationprotocols sets Applicationprotocols field to given value.


### GetDirection

`func (o *GetActivities200ResponseDataInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetActivities200ResponseDataInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetActivities200ResponseDataInner) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetPacketsize

`func (o *GetActivities200ResponseDataInner) GetPacketsize() int64`

GetPacketsize returns the Packetsize field if non-nil, zero value otherwise.

### GetPacketsizeOk

`func (o *GetActivities200ResponseDataInner) GetPacketsizeOk() (*int64, bool)`

GetPacketsizeOk returns a tuple with the Packetsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsize

`func (o *GetActivities200ResponseDataInner) SetPacketsize(v int64)`

SetPacketsize sets Packetsize field to given value.


### GetClassification

`func (o *GetActivities200ResponseDataInner) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *GetActivities200ResponseDataInner) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *GetActivities200ResponseDataInner) SetClassification(v string)`

SetClassification sets Classification field to given value.


### GetSessionid

`func (o *GetActivities200ResponseDataInner) GetSessionid() int64`

GetSessionid returns the Sessionid field if non-nil, zero value otherwise.

### GetSessionidOk

`func (o *GetActivities200ResponseDataInner) GetSessionidOk() (*int64, bool)`

GetSessionidOk returns a tuple with the Sessionid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionid

`func (o *GetActivities200ResponseDataInner) SetSessionid(v int64)`

SetSessionid sets Sessionid field to given value.


### GetSeverity

`func (o *GetActivities200ResponseDataInner) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetActivities200ResponseDataInner) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetActivities200ResponseDataInner) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.


### GetSignature

`func (o *GetActivities200ResponseDataInner) GetSignature() Signature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *GetActivities200ResponseDataInner) GetSignatureOk() (*Signature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *GetActivities200ResponseDataInner) SetSignature(v Signature)`

SetSignature sets Signature field to given value.


### GetSignaturelist

`func (o *GetActivities200ResponseDataInner) GetSignaturelist() SignatureList`

GetSignaturelist returns the Signaturelist field if non-nil, zero value otherwise.

### GetSignaturelistOk

`func (o *GetActivities200ResponseDataInner) GetSignaturelistOk() (*SignatureList, bool)`

GetSignaturelistOk returns a tuple with the Signaturelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignaturelist

`func (o *GetActivities200ResponseDataInner) SetSignaturelist(v SignatureList)`

SetSignaturelist sets Signaturelist field to given value.


### GetResponsefilename

`func (o *GetActivities200ResponseDataInner) GetResponsefilename() string`

GetResponsefilename returns the Responsefilename field if non-nil, zero value otherwise.

### GetResponsefilenameOk

`func (o *GetActivities200ResponseDataInner) GetResponsefilenameOk() (*string, bool)`

GetResponsefilenameOk returns a tuple with the Responsefilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsefilename

`func (o *GetActivities200ResponseDataInner) SetResponsefilename(v string)`

SetResponsefilename sets Responsefilename field to given value.


### GetBlockedfiletype

`func (o *GetActivities200ResponseDataInner) GetBlockedfiletype() string`

GetBlockedfiletype returns the Blockedfiletype field if non-nil, zero value otherwise.

### GetBlockedfiletypeOk

`func (o *GetActivities200ResponseDataInner) GetBlockedfiletypeOk() (*string, bool)`

GetBlockedfiletypeOk returns a tuple with the Blockedfiletype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedfiletype

`func (o *GetActivities200ResponseDataInner) SetBlockedfiletype(v string)`

SetBlockedfiletype sets Blockedfiletype field to given value.


### GetBundleid

`func (o *GetActivities200ResponseDataInner) GetBundleid() int64`

GetBundleid returns the Bundleid field if non-nil, zero value otherwise.

### GetBundleidOk

`func (o *GetActivities200ResponseDataInner) GetBundleidOk() (*int64, bool)`

GetBundleidOk returns a tuple with the Bundleid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleid

`func (o *GetActivities200ResponseDataInner) SetBundleid(v int64)`

SetBundleid sets Bundleid field to given value.


### GetAmp

`func (o *GetActivities200ResponseDataInner) GetAmp() CiscoAMP`

GetAmp returns the Amp field if non-nil, zero value otherwise.

### GetAmpOk

`func (o *GetActivities200ResponseDataInner) GetAmpOk() (*CiscoAMP, bool)`

GetAmpOk returns a tuple with the Amp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmp

`func (o *GetActivities200ResponseDataInner) SetAmp(v CiscoAMP)`

SetAmp sets Amp field to given value.


### GetTenantcontrols

`func (o *GetActivities200ResponseDataInner) GetTenantcontrols() bool`

GetTenantcontrols returns the Tenantcontrols field if non-nil, zero value otherwise.

### GetTenantcontrolsOk

`func (o *GetActivities200ResponseDataInner) GetTenantcontrolsOk() (*bool, bool)`

GetTenantcontrolsOk returns a tuple with the Tenantcontrols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantcontrols

`func (o *GetActivities200ResponseDataInner) SetTenantcontrols(v bool)`

SetTenantcontrols sets Tenantcontrols field to given value.


### GetPort

`func (o *GetActivities200ResponseDataInner) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetActivities200ResponseDataInner) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetActivities200ResponseDataInner) SetPort(v int64)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *GetActivities200ResponseDataInner) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *GetActivities200ResponseDataInner) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetAntivirusthreats

`func (o *GetActivities200ResponseDataInner) GetAntivirusthreats() AntivirusThreats`

GetAntivirusthreats returns the Antivirusthreats field if non-nil, zero value otherwise.

### GetAntivirusthreatsOk

`func (o *GetActivities200ResponseDataInner) GetAntivirusthreatsOk() (*AntivirusThreats, bool)`

GetAntivirusthreatsOk returns a tuple with the Antivirusthreats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusthreats

`func (o *GetActivities200ResponseDataInner) SetAntivirusthreats(v AntivirusThreats)`

SetAntivirusthreats sets Antivirusthreats field to given value.


### GetPolicy

`func (o *GetActivities200ResponseDataInner) GetPolicy() PolicyZTNA`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetActivities200ResponseDataInner) GetPolicyOk() (*PolicyZTNA, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetActivities200ResponseDataInner) SetPolicy(v PolicyZTNA)`

SetPolicy sets Policy field to given value.


### GetRequestmethod

`func (o *GetActivities200ResponseDataInner) GetRequestmethod() string`

GetRequestmethod returns the Requestmethod field if non-nil, zero value otherwise.

### GetRequestmethodOk

`func (o *GetActivities200ResponseDataInner) GetRequestmethodOk() (*string, bool)`

GetRequestmethodOk returns a tuple with the Requestmethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestmethod

`func (o *GetActivities200ResponseDataInner) SetRequestmethod(v string)`

SetRequestmethod sets Requestmethod field to given value.

### HasRequestmethod

`func (o *GetActivities200ResponseDataInner) HasRequestmethod() bool`

HasRequestmethod returns a boolean if a field has been set.

### GetResponsesize

`func (o *GetActivities200ResponseDataInner) GetResponsesize() int64`

GetResponsesize returns the Responsesize field if non-nil, zero value otherwise.

### GetResponsesizeOk

`func (o *GetActivities200ResponseDataInner) GetResponsesizeOk() (*int64, bool)`

GetResponsesizeOk returns a tuple with the Responsesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsesize

`func (o *GetActivities200ResponseDataInner) SetResponsesize(v int64)`

SetResponsesize sets Responsesize field to given value.


### GetRequestsize

`func (o *GetActivities200ResponseDataInner) GetRequestsize() int64`

GetRequestsize returns the Requestsize field if non-nil, zero value otherwise.

### GetRequestsizeOk

`func (o *GetActivities200ResponseDataInner) GetRequestsizeOk() (*int64, bool)`

GetRequestsizeOk returns a tuple with the Requestsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsize

`func (o *GetActivities200ResponseDataInner) SetRequestsize(v int64)`

SetRequestsize sets Requestsize field to given value.


### GetStatuscode

`func (o *GetActivities200ResponseDataInner) GetStatuscode() int64`

GetStatuscode returns the Statuscode field if non-nil, zero value otherwise.

### GetStatuscodeOk

`func (o *GetActivities200ResponseDataInner) GetStatuscodeOk() (*int64, bool)`

GetStatuscodeOk returns a tuple with the Statuscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuscode

`func (o *GetActivities200ResponseDataInner) SetStatuscode(v int64)`

SetStatuscode sets Statuscode field to given value.


### GetUseragent

`func (o *GetActivities200ResponseDataInner) GetUseragent() string`

GetUseragent returns the Useragent field if non-nil, zero value otherwise.

### GetUseragentOk

`func (o *GetActivities200ResponseDataInner) GetUseragentOk() (*string, bool)`

GetUseragentOk returns a tuple with the Useragent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseragent

`func (o *GetActivities200ResponseDataInner) SetUseragent(v string)`

SetUseragent sets Useragent field to given value.


### GetReferer

`func (o *GetActivities200ResponseDataInner) GetReferer() string`

GetReferer returns the Referer field if non-nil, zero value otherwise.

### GetRefererOk

`func (o *GetActivities200ResponseDataInner) GetRefererOk() (*string, bool)`

GetRefererOk returns a tuple with the Referer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferer

`func (o *GetActivities200ResponseDataInner) SetReferer(v string)`

SetReferer sets Referer field to given value.


### GetWarnstatus

`func (o *GetActivities200ResponseDataInner) GetWarnstatus() string`

GetWarnstatus returns the Warnstatus field if non-nil, zero value otherwise.

### GetWarnstatusOk

`func (o *GetActivities200ResponseDataInner) GetWarnstatusOk() (*string, bool)`

GetWarnstatusOk returns a tuple with the Warnstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnstatus

`func (o *GetActivities200ResponseDataInner) SetWarnstatus(v string)`

SetWarnstatus sets Warnstatus field to given value.


### GetSha256

`func (o *GetActivities200ResponseDataInner) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *GetActivities200ResponseDataInner) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *GetActivities200ResponseDataInner) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetIsolated

`func (o *GetActivities200ResponseDataInner) GetIsolated() Isolated`

GetIsolated returns the Isolated field if non-nil, zero value otherwise.

### GetIsolatedOk

`func (o *GetActivities200ResponseDataInner) GetIsolatedOk() (*Isolated, bool)`

GetIsolatedOk returns a tuple with the Isolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolated

`func (o *GetActivities200ResponseDataInner) SetIsolated(v Isolated)`

SetIsolated sets Isolated field to given value.


### GetDatalossprevention

`func (o *GetActivities200ResponseDataInner) GetDatalossprevention() DataLossPreventionState`

GetDatalossprevention returns the Datalossprevention field if non-nil, zero value otherwise.

### GetDatalosspreventionOk

`func (o *GetActivities200ResponseDataInner) GetDatalosspreventionOk() (*DataLossPreventionState, bool)`

GetDatalosspreventionOk returns a tuple with the Datalossprevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatalossprevention

`func (o *GetActivities200ResponseDataInner) SetDatalossprevention(v DataLossPreventionState)`

SetDatalossprevention sets Datalossprevention field to given value.


### GetSecurityoverridden

`func (o *GetActivities200ResponseDataInner) GetSecurityoverridden() bool`

GetSecurityoverridden returns the Securityoverridden field if non-nil, zero value otherwise.

### GetSecurityoverriddenOk

`func (o *GetActivities200ResponseDataInner) GetSecurityoverriddenOk() (*bool, bool)`

GetSecurityoverriddenOk returns a tuple with the Securityoverridden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityoverridden

`func (o *GetActivities200ResponseDataInner) SetSecurityoverridden(v bool)`

SetSecurityoverridden sets Securityoverridden field to given value.


### GetContenttype

`func (o *GetActivities200ResponseDataInner) GetContenttype() string`

GetContenttype returns the Contenttype field if non-nil, zero value otherwise.

### GetContenttypeOk

`func (o *GetActivities200ResponseDataInner) GetContenttypeOk() (*string, bool)`

GetContenttypeOk returns a tuple with the Contenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContenttype

`func (o *GetActivities200ResponseDataInner) SetContenttype(v string)`

SetContenttype sets Contenttype field to given value.


### GetForwardingmethod

`func (o *GetActivities200ResponseDataInner) GetForwardingmethod() string`

GetForwardingmethod returns the Forwardingmethod field if non-nil, zero value otherwise.

### GetForwardingmethodOk

`func (o *GetActivities200ResponseDataInner) GetForwardingmethodOk() (*string, bool)`

GetForwardingmethodOk returns a tuple with the Forwardingmethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingmethod

`func (o *GetActivities200ResponseDataInner) SetForwardingmethod(v string)`

SetForwardingmethod sets Forwardingmethod field to given value.


### GetHttperrors

`func (o *GetActivities200ResponseDataInner) GetHttperrors() []HttpError`

GetHttperrors returns the Httperrors field if non-nil, zero value otherwise.

### GetHttperrorsOk

`func (o *GetActivities200ResponseDataInner) GetHttperrorsOk() (*[]HttpError, bool)`

GetHttperrorsOk returns a tuple with the Httperrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttperrors

`func (o *GetActivities200ResponseDataInner) SetHttperrors(v []HttpError)`

SetHttperrors sets Httperrors field to given value.


### GetEgress

`func (o *GetActivities200ResponseDataInner) GetEgress() Egress`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *GetActivities200ResponseDataInner) GetEgressOk() (*Egress, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *GetActivities200ResponseDataInner) SetEgress(v Egress)`

SetEgress sets Egress field to given value.


### GetDatacenter

`func (o *GetActivities200ResponseDataInner) GetDatacenter() DataCenter`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *GetActivities200ResponseDataInner) GetDatacenterOk() (*DataCenter, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *GetActivities200ResponseDataInner) SetDatacenter(v DataCenter)`

SetDatacenter sets Datacenter field to given value.


### GetUrl

`func (o *GetActivities200ResponseDataInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetActivities200ResponseDataInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetActivities200ResponseDataInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetClientfirewall

`func (o *GetActivities200ResponseDataInner) GetClientfirewall() string`

GetClientfirewall returns the Clientfirewall field if non-nil, zero value otherwise.

### GetClientfirewallOk

`func (o *GetActivities200ResponseDataInner) GetClientfirewallOk() (*string, bool)`

GetClientfirewallOk returns a tuple with the Clientfirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientfirewall

`func (o *GetActivities200ResponseDataInner) SetClientfirewall(v string)`

SetClientfirewall sets Clientfirewall field to given value.

### HasClientfirewall

`func (o *GetActivities200ResponseDataInner) HasClientfirewall() bool`

HasClientfirewall returns a boolean if a field has been set.

### GetDiskencryption

`func (o *GetActivities200ResponseDataInner) GetDiskencryption() string`

GetDiskencryption returns the Diskencryption field if non-nil, zero value otherwise.

### GetDiskencryptionOk

`func (o *GetActivities200ResponseDataInner) GetDiskencryptionOk() (*string, bool)`

GetDiskencryptionOk returns a tuple with the Diskencryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskencryption

`func (o *GetActivities200ResponseDataInner) SetDiskencryption(v string)`

SetDiskencryption sets Diskencryption field to given value.

### HasDiskencryption

`func (o *GetActivities200ResponseDataInner) HasDiskencryption() bool`

HasDiskencryption returns a boolean if a field has been set.

### GetAntimalwareagents

`func (o *GetActivities200ResponseDataInner) GetAntimalwareagents() []string`

GetAntimalwareagents returns the Antimalwareagents field if non-nil, zero value otherwise.

### GetAntimalwareagentsOk

`func (o *GetActivities200ResponseDataInner) GetAntimalwareagentsOk() (*[]string, bool)`

GetAntimalwareagentsOk returns a tuple with the Antimalwareagents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntimalwareagents

`func (o *GetActivities200ResponseDataInner) SetAntimalwareagents(v []string)`

SetAntimalwareagents sets Antimalwareagents field to given value.

### HasAntimalwareagents

`func (o *GetActivities200ResponseDataInner) HasAntimalwareagents() bool`

HasAntimalwareagents returns a boolean if a field has been set.

### GetSystempassword

`func (o *GetActivities200ResponseDataInner) GetSystempassword() string`

GetSystempassword returns the Systempassword field if non-nil, zero value otherwise.

### GetSystempasswordOk

`func (o *GetActivities200ResponseDataInner) GetSystempasswordOk() (*string, bool)`

GetSystempasswordOk returns a tuple with the Systempassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystempassword

`func (o *GetActivities200ResponseDataInner) SetSystempassword(v string)`

SetSystempassword sets Systempassword field to given value.

### HasSystempassword

`func (o *GetActivities200ResponseDataInner) HasSystempassword() bool`

HasSystempassword returns a boolean if a field has been set.

### GetDevice

`func (o *GetActivities200ResponseDataInner) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetActivities200ResponseDataInner) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetActivities200ResponseDataInner) SetDevice(v Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *GetActivities200ResponseDataInner) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDecryptAction

`func (o *GetActivities200ResponseDataInner) GetDecryptAction() string`

GetDecryptAction returns the DecryptAction field if non-nil, zero value otherwise.

### GetDecryptActionOk

`func (o *GetActivities200ResponseDataInner) GetDecryptActionOk() (*string, bool)`

GetDecryptActionOk returns a tuple with the DecryptAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptAction

`func (o *GetActivities200ResponseDataInner) SetDecryptAction(v string)`

SetDecryptAction sets DecryptAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


