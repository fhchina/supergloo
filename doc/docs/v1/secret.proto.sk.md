<!-- Code generated by protoc-gen-solo-kit. DO NOT EDIT. -->

## Package:
encryption.istio.io

## Source File:
secret.proto 

## Description:  

## Contents:
- Messages:  
	- [IstioCacertsSecret](#IstioCacertsSecret)

---
  
### <a name="IstioCacertsSecret">IstioCacertsSecret</a>

Description: @solo-kit:resource.short_name=ics
@solo-kit:resource.plural_name=istiocerts
@solo-kit:resource.resource_groups=translator.supergloo.solo.io,install.supergloo.solo.io
Secret containing CA Certs for Istio
Structured TLS Secret that istio uses for non-default root certificates

```yaml
"root_cert": string
"cert_chain": string
"ca_cert": string
"ca_key": string
"metadata": .core.solo.io.Metadata

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| root_cert | string |  |  |
| cert_chain | string |  |  |
| ca_cert | string |  |  |
| ca_key | string |  |  |
| metadata | [.core.solo.io.Metadata](secret.proto.sk.md#IstioCacertsSecret) | Metadata contains the object metadata for this resource |  |

