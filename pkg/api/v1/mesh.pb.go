// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/api/v1/mesh.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//Meshes represent a currently registered service mesh.
type Mesh struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by supergloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// Types that are valid to be assigned to MeshType:
	//	*Mesh_Istio
	//	*Mesh_AwsAppMesh
	//	*Mesh_Linkerd
	MeshType isMesh_MeshType `protobuf_oneof:"mesh_type"`
	// mtls config specifies configuration options for enabling mutual
	// tls between pods in this mesh
	MtlsConfig *MtlsConfig `protobuf:"bytes,2,opt,name=mtls_config,json=mtlsConfig,proto3" json:"mtls_config,omitempty"`
	// configuration for propagating stats and metrics from
	// mesh controllers and sidecars to a centralized datastore
	// such as prometheus
	MonitoringConfig *MonitoringConfig `protobuf:"bytes,3,opt,name=monitoring_config,json=monitoringConfig,proto3" json:"monitoring_config,omitempty"`
	// object which represents the data mesh discovery finds about a given mesh
	DiscoveryMetadata *DiscoveryMetadata `protobuf:"bytes,6,opt,name=discovery_metadata,json=discoveryMetadata,proto3" json:"discovery_metadata,omitempty"`
	// whether or not to use SMI to configure this mesh
	SmiEnabled           bool     `protobuf:"varint,7,opt,name=smi_enabled,json=smiEnabled,proto3" json:"smi_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mesh) Reset()         { *m = Mesh{} }
func (m *Mesh) String() string { return proto.CompactTextString(m) }
func (*Mesh) ProtoMessage()    {}
func (*Mesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{0}
}
func (m *Mesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mesh.Unmarshal(m, b)
}
func (m *Mesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mesh.Marshal(b, m, deterministic)
}
func (m *Mesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh.Merge(m, src)
}
func (m *Mesh) XXX_Size() int {
	return xxx_messageInfo_Mesh.Size(m)
}
func (m *Mesh) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh proto.InternalMessageInfo

type isMesh_MeshType interface {
	isMesh_MeshType()
	Equal(interface{}) bool
}

type Mesh_Istio struct {
	Istio *IstioMesh `protobuf:"bytes,1,opt,name=istio,proto3,oneof"`
}
type Mesh_AwsAppMesh struct {
	AwsAppMesh *AwsAppMesh `protobuf:"bytes,4,opt,name=aws_app_mesh,json=awsAppMesh,proto3,oneof"`
}
type Mesh_Linkerd struct {
	Linkerd *LinkerdMesh `protobuf:"bytes,5,opt,name=linkerd,proto3,oneof"`
}

func (*Mesh_Istio) isMesh_MeshType()      {}
func (*Mesh_AwsAppMesh) isMesh_MeshType() {}
func (*Mesh_Linkerd) isMesh_MeshType()    {}

func (m *Mesh) GetMeshType() isMesh_MeshType {
	if m != nil {
		return m.MeshType
	}
	return nil
}

func (m *Mesh) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Mesh) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Mesh) GetIstio() *IstioMesh {
	if x, ok := m.GetMeshType().(*Mesh_Istio); ok {
		return x.Istio
	}
	return nil
}

func (m *Mesh) GetAwsAppMesh() *AwsAppMesh {
	if x, ok := m.GetMeshType().(*Mesh_AwsAppMesh); ok {
		return x.AwsAppMesh
	}
	return nil
}

func (m *Mesh) GetLinkerd() *LinkerdMesh {
	if x, ok := m.GetMeshType().(*Mesh_Linkerd); ok {
		return x.Linkerd
	}
	return nil
}

func (m *Mesh) GetMtlsConfig() *MtlsConfig {
	if m != nil {
		return m.MtlsConfig
	}
	return nil
}

func (m *Mesh) GetMonitoringConfig() *MonitoringConfig {
	if m != nil {
		return m.MonitoringConfig
	}
	return nil
}

func (m *Mesh) GetDiscoveryMetadata() *DiscoveryMetadata {
	if m != nil {
		return m.DiscoveryMetadata
	}
	return nil
}

func (m *Mesh) GetSmiEnabled() bool {
	if m != nil {
		return m.SmiEnabled
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Mesh) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Mesh_OneofMarshaler, _Mesh_OneofUnmarshaler, _Mesh_OneofSizer, []interface{}{
		(*Mesh_Istio)(nil),
		(*Mesh_AwsAppMesh)(nil),
		(*Mesh_Linkerd)(nil),
	}
}

func _Mesh_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Mesh)
	// mesh_type
	switch x := m.MeshType.(type) {
	case *Mesh_Istio:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Istio); err != nil {
			return err
		}
	case *Mesh_AwsAppMesh:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AwsAppMesh); err != nil {
			return err
		}
	case *Mesh_Linkerd:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Linkerd); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Mesh.MeshType has unexpected type %T", x)
	}
	return nil
}

func _Mesh_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Mesh)
	switch tag {
	case 1: // mesh_type.istio
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IstioMesh)
		err := b.DecodeMessage(msg)
		m.MeshType = &Mesh_Istio{msg}
		return true, err
	case 4: // mesh_type.aws_app_mesh
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AwsAppMesh)
		err := b.DecodeMessage(msg)
		m.MeshType = &Mesh_AwsAppMesh{msg}
		return true, err
	case 5: // mesh_type.linkerd
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LinkerdMesh)
		err := b.DecodeMessage(msg)
		m.MeshType = &Mesh_Linkerd{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Mesh_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Mesh)
	// mesh_type
	switch x := m.MeshType.(type) {
	case *Mesh_Istio:
		s := proto.Size(x.Istio)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mesh_AwsAppMesh:
		s := proto.Size(x.AwsAppMesh)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mesh_Linkerd:
		s := proto.Size(x.Linkerd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Generic discovery data shared between different meshes
type DiscoveryMetadata struct {
	// list of namespaces which we know are being injected by a given mesh
	InjectedNamespaceLabel string `protobuf:"bytes,1,opt,name=injected_namespace_label,json=injectedNamespaceLabel,proto3" json:"injected_namespace_label,omitempty"`
	// Whether or not auto-injection is enabled for a given mesh
	EnableAutoInject bool `protobuf:"varint,2,opt,name=enable_auto_inject,json=enableAutoInject,proto3" json:"enable_auto_inject,omitempty"`
	// version of the mesh which is installed
	MeshVersion string `protobuf:"bytes,3,opt,name=mesh_version,json=meshVersion,proto3" json:"mesh_version,omitempty"`
	// namespace which the mesh is installed into
	InstallationNamespace string `protobuf:"bytes,4,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// upstreams which point to injected pods in the mesh
	Upstreams []*core.ResourceRef `protobuf:"bytes,5,rep,name=upstreams,proto3" json:"upstreams,omitempty"`
	// discovered mtls config of the given mesh
	MtlsConfig           *MtlsConfig `protobuf:"bytes,6,opt,name=mtls_config,json=mtlsConfig,proto3" json:"mtls_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DiscoveryMetadata) Reset()         { *m = DiscoveryMetadata{} }
func (m *DiscoveryMetadata) String() string { return proto.CompactTextString(m) }
func (*DiscoveryMetadata) ProtoMessage()    {}
func (*DiscoveryMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{1}
}
func (m *DiscoveryMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryMetadata.Unmarshal(m, b)
}
func (m *DiscoveryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryMetadata.Marshal(b, m, deterministic)
}
func (m *DiscoveryMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryMetadata.Merge(m, src)
}
func (m *DiscoveryMetadata) XXX_Size() int {
	return xxx_messageInfo_DiscoveryMetadata.Size(m)
}
func (m *DiscoveryMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryMetadata proto.InternalMessageInfo

func (m *DiscoveryMetadata) GetInjectedNamespaceLabel() string {
	if m != nil {
		return m.InjectedNamespaceLabel
	}
	return ""
}

func (m *DiscoveryMetadata) GetEnableAutoInject() bool {
	if m != nil {
		return m.EnableAutoInject
	}
	return false
}

func (m *DiscoveryMetadata) GetMeshVersion() string {
	if m != nil {
		return m.MeshVersion
	}
	return ""
}

func (m *DiscoveryMetadata) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *DiscoveryMetadata) GetUpstreams() []*core.ResourceRef {
	if m != nil {
		return m.Upstreams
	}
	return nil
}

func (m *DiscoveryMetadata) GetMtlsConfig() *MtlsConfig {
	if m != nil {
		return m.MtlsConfig
	}
	return nil
}

// Mesh object representing an installed Istio control plane
type IstioMesh struct {
	// where the istio control plane has been installed
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// version of istio which has been installed
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioMesh) Reset()         { *m = IstioMesh{} }
func (m *IstioMesh) String() string { return proto.CompactTextString(m) }
func (*IstioMesh) ProtoMessage()    {}
func (*IstioMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{2}
}
func (m *IstioMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioMesh.Unmarshal(m, b)
}
func (m *IstioMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioMesh.Marshal(b, m, deterministic)
}
func (m *IstioMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioMesh.Merge(m, src)
}
func (m *IstioMesh) XXX_Size() int {
	return xxx_messageInfo_IstioMesh.Size(m)
}
func (m *IstioMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioMesh.DiscardUnknown(m)
}

var xxx_messageInfo_IstioMesh proto.InternalMessageInfo

func (m *IstioMesh) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *IstioMesh) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Mesh object representing AWS App Mesh
type AwsAppMesh struct {
	// Reference to the secret that holds the AWS credentials that will be used to access the AWS App Mesh service.
	AwsSecret *core.ResourceRef `protobuf:"bytes,1,opt,name=aws_secret,json=awsSecret,proto3" json:"aws_secret,omitempty"`
	// The AWS region the AWS App Mesh control plane resources (Virtual Nodes, Virtual Routers, etc.) will be created in.
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// Determines whether pods will be automatically injected with the AWS App Mesh Envoy sidecar proxy.
	//
	// If set to true, supergloo will ensure that a MutatingAdmissionWebhook server with the injection logic is deployed
	// to the cluster and that it has been registered with the Kubernetes API server via a MutatingWebhookConfiguration.
	// This will cause the webhook to be invoked on each pod creation event.
	EnableAutoInject bool `protobuf:"varint,3,opt,name=enable_auto_inject,json=enableAutoInject,proto3" json:"enable_auto_inject,omitempty"`
	// Pods matching this selector will be injected with the sidecar proxy at creation time.
	//
	// NOTE: the sidecar injector webhook currently supports only the NamespaceSelector and LabelSelector
	InjectionSelector *PodSelector `protobuf:"bytes,4,opt,name=injection_selector,json=injectionSelector,proto3" json:"injection_selector,omitempty"`
	// If auto-injection is enabled, the value of the pod label with this key will be used to calculate the value of
	// APPMESH_VIRTUAL_NODE_NAME environment variable that is set on the injected sidecar proxy container.
	VirtualNodeLabel string `protobuf:"bytes,5,opt,name=virtual_node_label,json=virtualNodeLabel,proto3" json:"virtual_node_label,omitempty"`
	// Reference to the config map that contains the patch that will be applied to the spec of the pods matching the
	// injection_selector.
	SidecarPatchConfigMap *core.ResourceRef `protobuf:"bytes,6,opt,name=sidecar_patch_config_map,json=sidecarPatchConfigMap,proto3" json:"sidecar_patch_config_map,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *AwsAppMesh) Reset()         { *m = AwsAppMesh{} }
func (m *AwsAppMesh) String() string { return proto.CompactTextString(m) }
func (*AwsAppMesh) ProtoMessage()    {}
func (*AwsAppMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{3}
}
func (m *AwsAppMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsAppMesh.Unmarshal(m, b)
}
func (m *AwsAppMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsAppMesh.Marshal(b, m, deterministic)
}
func (m *AwsAppMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsAppMesh.Merge(m, src)
}
func (m *AwsAppMesh) XXX_Size() int {
	return xxx_messageInfo_AwsAppMesh.Size(m)
}
func (m *AwsAppMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsAppMesh.DiscardUnknown(m)
}

var xxx_messageInfo_AwsAppMesh proto.InternalMessageInfo

func (m *AwsAppMesh) GetAwsSecret() *core.ResourceRef {
	if m != nil {
		return m.AwsSecret
	}
	return nil
}

func (m *AwsAppMesh) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *AwsAppMesh) GetEnableAutoInject() bool {
	if m != nil {
		return m.EnableAutoInject
	}
	return false
}

func (m *AwsAppMesh) GetInjectionSelector() *PodSelector {
	if m != nil {
		return m.InjectionSelector
	}
	return nil
}

func (m *AwsAppMesh) GetVirtualNodeLabel() string {
	if m != nil {
		return m.VirtualNodeLabel
	}
	return ""
}

func (m *AwsAppMesh) GetSidecarPatchConfigMap() *core.ResourceRef {
	if m != nil {
		return m.SidecarPatchConfigMap
	}
	return nil
}

// Mesh object representing an installed Linkerd control plane
type LinkerdMesh struct {
	// where the Linkerd control plane has been installed
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// version of istio which has been installed
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkerdMesh) Reset()         { *m = LinkerdMesh{} }
func (m *LinkerdMesh) String() string { return proto.CompactTextString(m) }
func (*LinkerdMesh) ProtoMessage()    {}
func (*LinkerdMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{4}
}
func (m *LinkerdMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkerdMesh.Unmarshal(m, b)
}
func (m *LinkerdMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkerdMesh.Marshal(b, m, deterministic)
}
func (m *LinkerdMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkerdMesh.Merge(m, src)
}
func (m *LinkerdMesh) XXX_Size() int {
	return xxx_messageInfo_LinkerdMesh.Size(m)
}
func (m *LinkerdMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkerdMesh.DiscardUnknown(m)
}

var xxx_messageInfo_LinkerdMesh proto.InternalMessageInfo

func (m *LinkerdMesh) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *LinkerdMesh) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// the encryption configuration that will be applied by the role
type MtlsConfig struct {
	// whether or not mutual TLS should be enabled between pods in this mesh
	MtlsEnabled bool `protobuf:"varint,1,opt,name=mtls_enabled,json=mtlsEnabled,proto3" json:"mtls_enabled,omitempty"`
	// if set, rootCertificate will override the root certificate used by the mesh
	// to encrypt mtls connections.
	//
	// The structure of the secret must be a standard kubernetes TLS secret
	// such as can be created via `kubectl create secret tls`
	//
	// if mtlsEnabled is false, this field is ignored
	// If deploying to Consul, Consul Connect requires that the cert and key are generated using ec, not rsa.
	RootCertificate      *core.ResourceRef `protobuf:"bytes,3,opt,name=root_certificate,json=rootCertificate,proto3" json:"root_certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MtlsConfig) Reset()         { *m = MtlsConfig{} }
func (m *MtlsConfig) String() string { return proto.CompactTextString(m) }
func (*MtlsConfig) ProtoMessage()    {}
func (*MtlsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{5}
}
func (m *MtlsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MtlsConfig.Unmarshal(m, b)
}
func (m *MtlsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MtlsConfig.Marshal(b, m, deterministic)
}
func (m *MtlsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MtlsConfig.Merge(m, src)
}
func (m *MtlsConfig) XXX_Size() int {
	return xxx_messageInfo_MtlsConfig.Size(m)
}
func (m *MtlsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MtlsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MtlsConfig proto.InternalMessageInfo

func (m *MtlsConfig) GetMtlsEnabled() bool {
	if m != nil {
		return m.MtlsEnabled
	}
	return false
}

func (m *MtlsConfig) GetRootCertificate() *core.ResourceRef {
	if m != nil {
		return m.RootCertificate
	}
	return nil
}

// Contains configuration options for monitoring a mesh
// Currently MonitoringConfig only contains options for configuring
// an in-cluster Prometheus instance to scrape a mesh for metrics
type MonitoringConfig struct {
	// indicates to supergloo that metrics should be propagated to one or more instances of prometheus.
	// add a [`core.solo.io.ResourceRef`](../../../../solo-kit/api/v1/ref.proto.sk#ResourceRef) for each
	// NAMESPACE.NAME of the configmap used to configure each prometheus instance.
	// assumes that the configmap contains a key named `prometheus.yml` or `prometheus.yaml` whose value
	// is the prometheus yaml config as an inline string
	PrometheusConfigmaps []core.ResourceRef `protobuf:"bytes,1,rep,name=prometheus_configmaps,json=prometheusConfigmaps,proto3" json:"prometheus_configmaps"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MonitoringConfig) Reset()         { *m = MonitoringConfig{} }
func (m *MonitoringConfig) String() string { return proto.CompactTextString(m) }
func (*MonitoringConfig) ProtoMessage()    {}
func (*MonitoringConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{6}
}
func (m *MonitoringConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitoringConfig.Unmarshal(m, b)
}
func (m *MonitoringConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitoringConfig.Marshal(b, m, deterministic)
}
func (m *MonitoringConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoringConfig.Merge(m, src)
}
func (m *MonitoringConfig) XXX_Size() int {
	return xxx_messageInfo_MonitoringConfig.Size(m)
}
func (m *MonitoringConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoringConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoringConfig proto.InternalMessageInfo

func (m *MonitoringConfig) GetPrometheusConfigmaps() []core.ResourceRef {
	if m != nil {
		return m.PrometheusConfigmaps
	}
	return nil
}

type MeshGroup struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by supergloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// the meshes contained in this group
	Meshes               []*core.ResourceRef `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MeshGroup) Reset()         { *m = MeshGroup{} }
func (m *MeshGroup) String() string { return proto.CompactTextString(m) }
func (*MeshGroup) ProtoMessage()    {}
func (*MeshGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{7}
}
func (m *MeshGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshGroup.Unmarshal(m, b)
}
func (m *MeshGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshGroup.Marshal(b, m, deterministic)
}
func (m *MeshGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshGroup.Merge(m, src)
}
func (m *MeshGroup) XXX_Size() int {
	return xxx_messageInfo_MeshGroup.Size(m)
}
func (m *MeshGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MeshGroup proto.InternalMessageInfo

func (m *MeshGroup) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *MeshGroup) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *MeshGroup) GetMeshes() []*core.ResourceRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func init() {
	proto.RegisterType((*Mesh)(nil), "supergloo.solo.io.Mesh")
	proto.RegisterType((*DiscoveryMetadata)(nil), "supergloo.solo.io.DiscoveryMetadata")
	proto.RegisterType((*IstioMesh)(nil), "supergloo.solo.io.IstioMesh")
	proto.RegisterType((*AwsAppMesh)(nil), "supergloo.solo.io.AwsAppMesh")
	proto.RegisterType((*LinkerdMesh)(nil), "supergloo.solo.io.LinkerdMesh")
	proto.RegisterType((*MtlsConfig)(nil), "supergloo.solo.io.MtlsConfig")
	proto.RegisterType((*MonitoringConfig)(nil), "supergloo.solo.io.MonitoringConfig")
	proto.RegisterType((*MeshGroup)(nil), "supergloo.solo.io.MeshGroup")
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/api/v1/mesh.proto", fileDescriptor_713281dd1a237b0d)
}

var fileDescriptor_713281dd1a237b0d = []byte{
	// 923 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xae, 0x77, 0x9d, 0x4d, 0xf6, 0x6c, 0x24, 0x76, 0x87, 0x24, 0x72, 0x2b, 0x48, 0xc2, 0x82,
	0x44, 0x2e, 0xda, 0xb5, 0x12, 0x40, 0x44, 0xb9, 0x40, 0x4a, 0x52, 0x54, 0x2a, 0x75, 0xab, 0xc8,
	0x41, 0x48, 0x20, 0x84, 0x35, 0x6b, 0xcf, 0x7a, 0x87, 0xd8, 0x9e, 0xd1, 0xcc, 0x38, 0x51, 0xef,
	0x50, 0x9e, 0x86, 0x47, 0xe1, 0x8a, 0x17, 0x40, 0xf4, 0x82, 0x07, 0x40, 0x5a, 0x9e, 0x00, 0xcd,
	0x78, 0x6c, 0xb7, 0xe9, 0x66, 0x9b, 0x4a, 0x5c, 0x70, 0x65, 0xcf, 0x7c, 0xdf, 0x77, 0xe6, 0xcc,
	0xf9, 0x1b, 0x18, 0x25, 0x54, 0xcd, 0x8a, 0xc9, 0x28, 0x62, 0x99, 0x2f, 0x59, 0xca, 0x1e, 0x51,
	0xe6, 0xcb, 0x82, 0x13, 0x91, 0xa4, 0x8c, 0xf9, 0x98, 0x53, 0xff, 0x72, 0xdf, 0xcf, 0x88, 0x9c,
	0x8d, 0xb8, 0x60, 0x8a, 0xa1, 0x41, 0x0d, 0x8e, 0x34, 0x7d, 0x44, 0xd9, 0x83, 0x8d, 0x84, 0x25,
	0xcc, 0xa0, 0xbe, 0xfe, 0x2b, 0x89, 0x0f, 0xb6, 0x13, 0xc6, 0x92, 0x94, 0xf8, 0x66, 0x35, 0x29,
	0xa6, 0x7e, 0x5c, 0x08, 0xac, 0x28, 0xcb, 0x6f, 0xc3, 0xaf, 0x04, 0xe6, 0x9c, 0x08, 0x69, 0xf1,
	0xfd, 0x45, 0x8e, 0xe9, 0xef, 0x05, 0x55, 0x8d, 0x5f, 0x0a, 0xc7, 0x58, 0x61, 0x2b, 0xf1, 0xef,
	0x20, 0x91, 0x0a, 0xab, 0xa2, 0x3a, 0xe3, 0xe1, 0x1d, 0x04, 0x82, 0x4c, 0xdf, 0xc1, 0xa3, 0x6a,
	0x6d, 0x25, 0x07, 0x77, 0x89, 0xae, 0x24, 0x29, 0x89, 0x14, 0x13, 0xa5, 0x66, 0xf8, 0xb7, 0x0b,
	0xee, 0x98, 0xc8, 0x19, 0x7a, 0x02, 0x9d, 0xd2, 0x5b, 0x2f, 0xde, 0x75, 0xf6, 0x7a, 0x07, 0x1b,
	0xa3, 0x88, 0x09, 0x52, 0x85, 0x7d, 0x74, 0x6e, 0xb0, 0x93, 0xfb, 0xbf, 0xbd, 0xdc, 0xb9, 0xf7,
	0xcf, 0xcb, 0x9d, 0x81, 0x22, 0x52, 0xc5, 0x74, 0x3a, 0x3d, 0x1a, 0xd2, 0x24, 0x67, 0x82, 0x0c,
	0x03, 0x2b, 0x47, 0x87, 0xb0, 0x56, 0x45, 0xca, 0x23, 0xc6, 0xd4, 0xd6, 0xeb, 0xa6, 0xc6, 0x16,
	0x3d, 0x71, 0xb5, 0xb1, 0xa0, 0x66, 0xa3, 0xcf, 0x61, 0x85, 0x4a, 0x45, 0x99, 0xe7, 0x18, 0xd9,
	0x07, 0xa3, 0x37, 0xb2, 0x3f, 0x7a, 0xaa, 0x71, 0xed, 0xef, 0x37, 0xf7, 0x82, 0x92, 0x8c, 0x8e,
	0x61, 0x1d, 0x5f, 0xc9, 0x10, 0x73, 0x1e, 0xea, 0xca, 0xf1, 0x5c, 0x23, 0xfe, 0x70, 0x81, 0xf8,
	0xf8, 0x4a, 0x1e, 0x73, 0x6e, 0xd5, 0x80, 0xeb, 0x15, 0x3a, 0x82, 0xd5, 0x94, 0xe6, 0x17, 0x44,
	0xc4, 0xde, 0x8a, 0x51, 0x6f, 0x2f, 0x50, 0x3f, 0x2b, 0x19, 0x56, 0x5e, 0x09, 0xd0, 0x57, 0xd0,
	0xcb, 0x54, 0x2a, 0xc3, 0x88, 0xe5, 0x53, 0x9a, 0x78, 0xad, 0x5b, 0x4f, 0x1f, 0xab, 0x54, 0x9e,
	0x1a, 0x52, 0x00, 0x59, 0xfd, 0x8f, 0xce, 0x60, 0x90, 0xb1, 0x9c, 0x2a, 0x26, 0x68, 0x9e, 0x54,
	0x56, 0xda, 0xc6, 0xca, 0xc7, 0x8b, 0xac, 0xd4, 0x5c, 0x6b, 0xab, 0x9f, 0xdd, 0xd8, 0x41, 0xdf,
	0x03, 0x8a, 0xa9, 0x8c, 0xd8, 0x25, 0x11, 0x2f, 0xc2, 0x3a, 0x15, 0x1d, 0x63, 0xf2, 0x93, 0x05,
	0x26, 0x1f, 0x57, 0xe4, 0x26, 0x31, 0xbf, 0xcc, 0x5d, 0x27, 0x18, 0xc4, 0x37, 0x01, 0xb4, 0x03,
	0x3d, 0x99, 0xd1, 0x90, 0xe4, 0x78, 0x92, 0x92, 0xd8, 0x5b, 0xdd, 0x75, 0xf6, 0xd6, 0x02, 0x90,
	0x19, 0xfd, 0xba, 0xdc, 0x39, 0x7a, 0xff, 0x7a, 0xee, 0xb6, 0xc1, 0xc9, 0xae, 0xe7, 0xee, 0x1a,
	0xea, 0xe8, 0x7c, 0x10, 0x79, 0xd2, 0x83, 0xae, 0xfe, 0x0b, 0xd5, 0x0b, 0x4e, 0x86, 0xbf, 0xb7,
	0x60, 0xf0, 0xc6, 0x89, 0xe8, 0x10, 0x3c, 0x9a, 0xff, 0x4c, 0x22, 0x45, 0xe2, 0x30, 0xc7, 0x19,
	0x91, 0x1c, 0x47, 0x24, 0x4c, 0xf1, 0x84, 0xa4, 0xa6, 0x1a, 0xba, 0xc1, 0x56, 0x85, 0x3f, 0xaf,
	0xe0, 0x67, 0x1a, 0x45, 0x0f, 0x01, 0x95, 0xee, 0x84, 0xb8, 0x50, 0x2c, 0x2c, 0x59, 0x26, 0x0d,
	0x6b, 0x41, 0xbf, 0x44, 0x8e, 0x0b, 0xc5, 0x9e, 0x9a, 0x7d, 0xf4, 0x11, 0xac, 0x1b, 0x57, 0x2e,
	0x89, 0x90, 0x94, 0xe5, 0x26, 0xd0, 0xdd, 0xa0, 0xa7, 0xf7, 0xbe, 0x2b, 0xb7, 0xd0, 0x17, 0xb0,
	0x45, 0x73, 0xa9, 0x70, 0x9a, 0x9a, 0x01, 0xd2, 0xb8, 0x63, 0x2a, 0xab, 0x1b, 0x6c, 0xbe, 0x8a,
	0xd6, 0xce, 0xa0, 0x2f, 0xa1, 0x5b, 0x70, 0xa9, 0x04, 0xc1, 0x99, 0xf4, 0x56, 0x76, 0xdb, 0x7b,
	0xbd, 0x83, 0xfb, 0xaf, 0xd7, 0x7d, 0x40, 0x24, 0x2b, 0x44, 0x44, 0x02, 0x32, 0x0d, 0x1a, 0xee,
	0xcd, 0x02, 0xea, 0xbc, 0x63, 0x01, 0x0d, 0x7f, 0x84, 0x6e, 0xdd, 0x15, 0x4b, 0x9c, 0x77, 0x96,
	0x39, 0xef, 0xc1, 0x6a, 0x15, 0x91, 0x96, 0xe1, 0x55, 0xcb, 0xe1, 0x9f, 0x2d, 0x80, 0xa6, 0x6f,
	0xd0, 0x21, 0xe8, 0xbe, 0x09, 0x25, 0x89, 0x04, 0x51, 0xb6, 0x4f, 0x97, 0x5d, 0x13, 0x5f, 0xc9,
	0x73, 0xc3, 0x45, 0x5b, 0xd0, 0x11, 0x24, 0x69, 0x4e, 0xb0, 0xab, 0x5b, 0xf2, 0xd7, 0xbe, 0x25,
	0x7f, 0x63, 0x40, 0x25, 0x43, 0x5f, 0xae, 0x1a, 0x65, 0xb6, 0xe5, 0x17, 0x35, 0xed, 0x19, 0x8b,
	0xcf, 0x2d, 0x2b, 0x18, 0xd4, 0xca, 0x6a, 0x4b, 0x1f, 0x7e, 0x49, 0x85, 0x2a, 0x70, 0x1a, 0xe6,
	0x2c, 0xae, 0x0a, 0x6e, 0xc5, 0x38, 0xd8, 0xb7, 0xc8, 0x73, 0x16, 0xdb, 0x52, 0x0b, 0xc0, 0x93,
	0x34, 0x26, 0x11, 0x16, 0x21, 0xc7, 0x2a, 0x9a, 0xd9, 0x94, 0x85, 0x19, 0xe6, 0x36, 0x6d, 0x4b,
	0x42, 0xb1, 0x69, 0xa5, 0x67, 0x5a, 0x59, 0xa6, 0x6e, 0x8c, 0xf9, 0xf0, 0x27, 0xe8, 0xbd, 0x32,
	0x58, 0xfe, 0xfb, 0xfc, 0x15, 0x00, 0x4d, 0xdd, 0x98, 0xf2, 0xd7, 0xb5, 0x56, 0x35, 0xb0, 0x63,
	0xc2, 0x6c, 0xea, 0xcf, 0x76, 0x30, 0x7a, 0x0c, 0x7d, 0xc1, 0x98, 0x0a, 0x23, 0x22, 0x14, 0x9d,
	0xd2, 0x08, 0x2b, 0x62, 0xc7, 0xd1, 0x92, 0xcb, 0xbd, 0xa7, 0x25, 0xa7, 0x8d, 0x62, 0x38, 0x83,
	0xfe, 0xcd, 0x49, 0x85, 0xbe, 0x85, 0x4d, 0x2e, 0x58, 0x46, 0xd4, 0x8c, 0x14, 0x55, 0xb9, 0x67,
	0x98, 0x4b, 0xcf, 0x79, 0x4b, 0xb7, 0xd8, 0x87, 0x62, 0xa3, 0x51, 0x9f, 0xd6, 0xe2, 0xe1, 0x1f,
	0x0e, 0x74, 0x75, 0xe8, 0x9e, 0x08, 0x56, 0xf0, 0xff, 0xc3, 0x2b, 0xb6, 0x0f, 0x76, 0xee, 0x79,
	0xed, 0xb7, 0x4d, 0x01, 0x4b, 0x3c, 0xf2, 0xae, 0xe7, 0xae, 0x0b, 0xad, 0x2c, 0xb9, 0x9e, 0xbb,
	0xeb, 0x08, 0xf4, 0x6e, 0xa2, 0xaf, 0x23, 0x4f, 0x1e, 0xfd, 0xfa, 0xd7, 0xb6, 0xf3, 0xc3, 0xa7,
	0x4b, 0x1f, 0x76, 0x7e, 0x91, 0xd8, 0xc7, 0x7d, 0xd2, 0x31, 0x8f, 0xfa, 0x67, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xf9, 0x59, 0x1d, 0x16, 0x68, 0x09, 0x00, 0x00,
}

func (this *Mesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh)
	if !ok {
		that2, ok := that.(Mesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if that1.MeshType == nil {
		if this.MeshType != nil {
			return false
		}
	} else if this.MeshType == nil {
		return false
	} else if !this.MeshType.Equal(that1.MeshType) {
		return false
	}
	if !this.MtlsConfig.Equal(that1.MtlsConfig) {
		return false
	}
	if !this.MonitoringConfig.Equal(that1.MonitoringConfig) {
		return false
	}
	if !this.DiscoveryMetadata.Equal(that1.DiscoveryMetadata) {
		return false
	}
	if this.SmiEnabled != that1.SmiEnabled {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Mesh_Istio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_Istio)
	if !ok {
		that2, ok := that.(Mesh_Istio)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Istio.Equal(that1.Istio) {
		return false
	}
	return true
}
func (this *Mesh_AwsAppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_AwsAppMesh)
	if !ok {
		that2, ok := that.(Mesh_AwsAppMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AwsAppMesh.Equal(that1.AwsAppMesh) {
		return false
	}
	return true
}
func (this *Mesh_Linkerd) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_Linkerd)
	if !ok {
		that2, ok := that.(Mesh_Linkerd)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Linkerd.Equal(that1.Linkerd) {
		return false
	}
	return true
}
func (this *DiscoveryMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryMetadata)
	if !ok {
		that2, ok := that.(DiscoveryMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InjectedNamespaceLabel != that1.InjectedNamespaceLabel {
		return false
	}
	if this.EnableAutoInject != that1.EnableAutoInject {
		return false
	}
	if this.MeshVersion != that1.MeshVersion {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if len(this.Upstreams) != len(that1.Upstreams) {
		return false
	}
	for i := range this.Upstreams {
		if !this.Upstreams[i].Equal(that1.Upstreams[i]) {
			return false
		}
	}
	if !this.MtlsConfig.Equal(that1.MtlsConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IstioMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IstioMesh)
	if !ok {
		that2, ok := that.(IstioMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AwsAppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsAppMesh)
	if !ok {
		that2, ok := that.(AwsAppMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AwsSecret.Equal(that1.AwsSecret) {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if this.EnableAutoInject != that1.EnableAutoInject {
		return false
	}
	if !this.InjectionSelector.Equal(that1.InjectionSelector) {
		return false
	}
	if this.VirtualNodeLabel != that1.VirtualNodeLabel {
		return false
	}
	if !this.SidecarPatchConfigMap.Equal(that1.SidecarPatchConfigMap) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LinkerdMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LinkerdMesh)
	if !ok {
		that2, ok := that.(LinkerdMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MtlsConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MtlsConfig)
	if !ok {
		that2, ok := that.(MtlsConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MtlsEnabled != that1.MtlsEnabled {
		return false
	}
	if !this.RootCertificate.Equal(that1.RootCertificate) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MonitoringConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MonitoringConfig)
	if !ok {
		that2, ok := that.(MonitoringConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.PrometheusConfigmaps) != len(that1.PrometheusConfigmaps) {
		return false
	}
	for i := range this.PrometheusConfigmaps {
		if !this.PrometheusConfigmaps[i].Equal(&that1.PrometheusConfigmaps[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshGroup) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshGroup)
	if !ok {
		that2, ok := that.(MeshGroup)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
