// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a VPN gateway managed outside of GCP.
//
// To get more information about ExternalVpnGateway, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/externalVpnGateways)
//
// ## Example Usage
// ### External Vpn Gateway
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			network, err := compute.NewNetwork(ctx, "network", &compute.NetworkArgs{
//				RoutingMode:           pulumi.String("GLOBAL"),
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			haGateway, err := compute.NewHaVpnGateway(ctx, "haGateway", &compute.HaVpnGatewayArgs{
//				Region:  pulumi.String("us-central1"),
//				Network: network.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			externalGateway, err := compute.NewExternalVpnGateway(ctx, "externalGateway", &compute.ExternalVpnGatewayArgs{
//				RedundancyType: pulumi.String("SINGLE_IP_INTERNALLY_REDUNDANT"),
//				Description:    pulumi.String("An externally managed VPN gateway"),
//				Interfaces: compute.ExternalVpnGatewayInterfaceArray{
//					&compute.ExternalVpnGatewayInterfaceArgs{
//						Id:        pulumi.Int(0),
//						IpAddress: pulumi.String("8.8.8.8"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewSubnetwork(ctx, "networkSubnet1", &compute.SubnetworkArgs{
//				IpCidrRange: pulumi.String("10.0.1.0/24"),
//				Region:      pulumi.String("us-central1"),
//				Network:     network.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewSubnetwork(ctx, "networkSubnet2", &compute.SubnetworkArgs{
//				IpCidrRange: pulumi.String("10.0.2.0/24"),
//				Region:      pulumi.String("us-west1"),
//				Network:     network.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			router1, err := compute.NewRouter(ctx, "router1", &compute.RouterArgs{
//				Network: network.Name,
//				Bgp: &compute.RouterBgpArgs{
//					Asn: pulumi.Int(64514),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			tunnel1, err := compute.NewVPNTunnel(ctx, "tunnel1", &compute.VPNTunnelArgs{
//				Region:                       pulumi.String("us-central1"),
//				VpnGateway:                   haGateway.ID(),
//				PeerExternalGateway:          externalGateway.ID(),
//				PeerExternalGatewayInterface: pulumi.Int(0),
//				SharedSecret:                 pulumi.String("a secret message"),
//				Router:                       router1.ID(),
//				VpnGatewayInterface:          pulumi.Int(0),
//			})
//			if err != nil {
//				return err
//			}
//			tunnel2, err := compute.NewVPNTunnel(ctx, "tunnel2", &compute.VPNTunnelArgs{
//				Region:                       pulumi.String("us-central1"),
//				VpnGateway:                   haGateway.ID(),
//				PeerExternalGateway:          externalGateway.ID(),
//				PeerExternalGatewayInterface: pulumi.Int(0),
//				SharedSecret:                 pulumi.String("a secret message"),
//				Router: router1.ID().ApplyT(func(id string) (string, error) {
//					return fmt.Sprintf(" %v", id), nil
//				}).(pulumi.StringOutput),
//				VpnGatewayInterface: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			router1Interface1, err := compute.NewRouterInterface(ctx, "router1Interface1", &compute.RouterInterfaceArgs{
//				Router:    router1.Name,
//				Region:    pulumi.String("us-central1"),
//				IpRange:   pulumi.String("169.254.0.1/30"),
//				VpnTunnel: tunnel1.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewRouterPeer(ctx, "router1Peer1", &compute.RouterPeerArgs{
//				Router:                  router1.Name,
//				Region:                  pulumi.String("us-central1"),
//				PeerIpAddress:           pulumi.String("169.254.0.2"),
//				PeerAsn:                 pulumi.Int(64515),
//				AdvertisedRoutePriority: pulumi.Int(100),
//				Interface:               router1Interface1.Name,
//			})
//			if err != nil {
//				return err
//			}
//			router1Interface2, err := compute.NewRouterInterface(ctx, "router1Interface2", &compute.RouterInterfaceArgs{
//				Router:    router1.Name,
//				Region:    pulumi.String("us-central1"),
//				IpRange:   pulumi.String("169.254.1.1/30"),
//				VpnTunnel: tunnel2.Name,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewRouterPeer(ctx, "router1Peer2", &compute.RouterPeerArgs{
//				Router:                  router1.Name,
//				Region:                  pulumi.String("us-central1"),
//				PeerIpAddress:           pulumi.String("169.254.1.2"),
//				PeerAsn:                 pulumi.Int(64515),
//				AdvertisedRoutePriority: pulumi.Int(100),
//				Interface:               router1Interface2.Name,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # ExternalVpnGateway can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default projects/{{project}}/global/externalVpnGateways/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{name}}
//
// ```
type ExternalVpnGateway struct {
	pulumi.CustomResourceState

	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interfaces ExternalVpnGatewayInterfaceArrayOutput `pulumi:"interfaces"`
	// The fingerprint used for optimistic locking of this resource.  Used
	// internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels for the external VPN gateway resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway
	// Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
	RedundancyType pulumi.StringPtrOutput `pulumi:"redundancyType"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewExternalVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewExternalVpnGateway(ctx *pulumi.Context,
	name string, args *ExternalVpnGatewayArgs, opts ...pulumi.ResourceOption) (*ExternalVpnGateway, error) {
	if args == nil {
		args = &ExternalVpnGatewayArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExternalVpnGateway
	err := ctx.RegisterResource("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalVpnGateway gets an existing ExternalVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalVpnGatewayState, opts ...pulumi.ResourceOption) (*ExternalVpnGateway, error) {
	var resource ExternalVpnGateway
	err := ctx.ReadResource("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalVpnGateway resources.
type externalVpnGatewayState struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interfaces []ExternalVpnGatewayInterface `pulumi:"interfaces"`
	// The fingerprint used for optimistic locking of this resource.  Used
	// internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels for the external VPN gateway resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// ***
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway
	// Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
	RedundancyType *string `pulumi:"redundancyType"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type ExternalVpnGatewayState struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interfaces ExternalVpnGatewayInterfaceArrayInput
	// The fingerprint used for optimistic locking of this resource.  Used
	// internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels for the external VPN gateway resource.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// ***
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Indicates the redundancy type of this external VPN gateway
	// Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
	RedundancyType pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (ExternalVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalVpnGatewayState)(nil)).Elem()
}

type externalVpnGatewayArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interfaces []ExternalVpnGatewayInterface `pulumi:"interfaces"`
	// Labels for the external VPN gateway resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// ***
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway
	// Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
	RedundancyType *string `pulumi:"redundancyType"`
}

// The set of arguments for constructing a ExternalVpnGateway resource.
type ExternalVpnGatewayArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interfaces ExternalVpnGatewayInterfaceArrayInput
	// Labels for the external VPN gateway resource.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// ***
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Indicates the redundancy type of this external VPN gateway
	// Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
	RedundancyType pulumi.StringPtrInput
}

func (ExternalVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalVpnGatewayArgs)(nil)).Elem()
}

type ExternalVpnGatewayInput interface {
	pulumi.Input

	ToExternalVpnGatewayOutput() ExternalVpnGatewayOutput
	ToExternalVpnGatewayOutputWithContext(ctx context.Context) ExternalVpnGatewayOutput
}

func (*ExternalVpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalVpnGateway)(nil)).Elem()
}

func (i *ExternalVpnGateway) ToExternalVpnGatewayOutput() ExternalVpnGatewayOutput {
	return i.ToExternalVpnGatewayOutputWithContext(context.Background())
}

func (i *ExternalVpnGateway) ToExternalVpnGatewayOutputWithContext(ctx context.Context) ExternalVpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalVpnGatewayOutput)
}

// ExternalVpnGatewayArrayInput is an input type that accepts ExternalVpnGatewayArray and ExternalVpnGatewayArrayOutput values.
// You can construct a concrete instance of `ExternalVpnGatewayArrayInput` via:
//
//	ExternalVpnGatewayArray{ ExternalVpnGatewayArgs{...} }
type ExternalVpnGatewayArrayInput interface {
	pulumi.Input

	ToExternalVpnGatewayArrayOutput() ExternalVpnGatewayArrayOutput
	ToExternalVpnGatewayArrayOutputWithContext(context.Context) ExternalVpnGatewayArrayOutput
}

type ExternalVpnGatewayArray []ExternalVpnGatewayInput

func (ExternalVpnGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalVpnGateway)(nil)).Elem()
}

func (i ExternalVpnGatewayArray) ToExternalVpnGatewayArrayOutput() ExternalVpnGatewayArrayOutput {
	return i.ToExternalVpnGatewayArrayOutputWithContext(context.Background())
}

func (i ExternalVpnGatewayArray) ToExternalVpnGatewayArrayOutputWithContext(ctx context.Context) ExternalVpnGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalVpnGatewayArrayOutput)
}

// ExternalVpnGatewayMapInput is an input type that accepts ExternalVpnGatewayMap and ExternalVpnGatewayMapOutput values.
// You can construct a concrete instance of `ExternalVpnGatewayMapInput` via:
//
//	ExternalVpnGatewayMap{ "key": ExternalVpnGatewayArgs{...} }
type ExternalVpnGatewayMapInput interface {
	pulumi.Input

	ToExternalVpnGatewayMapOutput() ExternalVpnGatewayMapOutput
	ToExternalVpnGatewayMapOutputWithContext(context.Context) ExternalVpnGatewayMapOutput
}

type ExternalVpnGatewayMap map[string]ExternalVpnGatewayInput

func (ExternalVpnGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalVpnGateway)(nil)).Elem()
}

func (i ExternalVpnGatewayMap) ToExternalVpnGatewayMapOutput() ExternalVpnGatewayMapOutput {
	return i.ToExternalVpnGatewayMapOutputWithContext(context.Background())
}

func (i ExternalVpnGatewayMap) ToExternalVpnGatewayMapOutputWithContext(ctx context.Context) ExternalVpnGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalVpnGatewayMapOutput)
}

type ExternalVpnGatewayOutput struct{ *pulumi.OutputState }

func (ExternalVpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalVpnGateway)(nil)).Elem()
}

func (o ExternalVpnGatewayOutput) ToExternalVpnGatewayOutput() ExternalVpnGatewayOutput {
	return o
}

func (o ExternalVpnGatewayOutput) ToExternalVpnGatewayOutputWithContext(ctx context.Context) ExternalVpnGatewayOutput {
	return o
}

// An optional description of this resource.
func (o ExternalVpnGatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalVpnGateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of interfaces on this external VPN gateway.
// Structure is documented below.
func (o ExternalVpnGatewayOutput) Interfaces() ExternalVpnGatewayInterfaceArrayOutput {
	return o.ApplyT(func(v *ExternalVpnGateway) ExternalVpnGatewayInterfaceArrayOutput { return v.Interfaces }).(ExternalVpnGatewayInterfaceArrayOutput)
}

// The fingerprint used for optimistic locking of this resource.  Used
// internally during updates.
func (o ExternalVpnGatewayOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalVpnGateway) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for the external VPN gateway resource.
func (o ExternalVpnGatewayOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExternalVpnGateway) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035.  Specifically, the name must be 1-63 characters long and
// match the regular expression `a-z?` which means
// the first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the last
// character, which cannot be a dash.
//
// ***
func (o ExternalVpnGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalVpnGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o ExternalVpnGatewayOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalVpnGateway) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Indicates the redundancy type of this external VPN gateway
// Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
func (o ExternalVpnGatewayOutput) RedundancyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalVpnGateway) pulumi.StringPtrOutput { return v.RedundancyType }).(pulumi.StringPtrOutput)
}

// The URI of the created resource.
func (o ExternalVpnGatewayOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalVpnGateway) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

type ExternalVpnGatewayArrayOutput struct{ *pulumi.OutputState }

func (ExternalVpnGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalVpnGateway)(nil)).Elem()
}

func (o ExternalVpnGatewayArrayOutput) ToExternalVpnGatewayArrayOutput() ExternalVpnGatewayArrayOutput {
	return o
}

func (o ExternalVpnGatewayArrayOutput) ToExternalVpnGatewayArrayOutputWithContext(ctx context.Context) ExternalVpnGatewayArrayOutput {
	return o
}

func (o ExternalVpnGatewayArrayOutput) Index(i pulumi.IntInput) ExternalVpnGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalVpnGateway {
		return vs[0].([]*ExternalVpnGateway)[vs[1].(int)]
	}).(ExternalVpnGatewayOutput)
}

type ExternalVpnGatewayMapOutput struct{ *pulumi.OutputState }

func (ExternalVpnGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalVpnGateway)(nil)).Elem()
}

func (o ExternalVpnGatewayMapOutput) ToExternalVpnGatewayMapOutput() ExternalVpnGatewayMapOutput {
	return o
}

func (o ExternalVpnGatewayMapOutput) ToExternalVpnGatewayMapOutputWithContext(ctx context.Context) ExternalVpnGatewayMapOutput {
	return o
}

func (o ExternalVpnGatewayMapOutput) MapIndex(k pulumi.StringInput) ExternalVpnGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalVpnGateway {
		return vs[0].(map[string]*ExternalVpnGateway)[vs[1].(string)]
	}).(ExternalVpnGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalVpnGatewayInput)(nil)).Elem(), &ExternalVpnGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalVpnGatewayArrayInput)(nil)).Elem(), ExternalVpnGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalVpnGatewayMapInput)(nil)).Elem(), ExternalVpnGatewayMap{})
	pulumi.RegisterOutputType(ExternalVpnGatewayOutput{})
	pulumi.RegisterOutputType(ExternalVpnGatewayArrayOutput{})
	pulumi.RegisterOutputType(ExternalVpnGatewayMapOutput{})
}