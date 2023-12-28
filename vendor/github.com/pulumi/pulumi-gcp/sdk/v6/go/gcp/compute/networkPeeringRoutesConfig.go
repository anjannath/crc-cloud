// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a network peering's route settings without managing the peering as
// a whole. This resource is primarily intended for use with GCP-generated
// peerings that shouldn't otherwise be managed by other tools. Deleting this
// resource is a no-op and the peering will not be modified.
//
// To get more information about NetworkPeeringRoutesConfig, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/networks/updatePeering)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/vpc/docs/vpc-peering)
//
// ## Example Usage
// ### Network Peering Routes Config Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			networkPrimary, err := compute.NewNetwork(ctx, "networkPrimary", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			networkSecondary, err := compute.NewNetwork(ctx, "networkSecondary", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			peeringPrimary, err := compute.NewNetworkPeering(ctx, "peeringPrimary", &compute.NetworkPeeringArgs{
//				Network:            networkPrimary.ID(),
//				PeerNetwork:        networkSecondary.ID(),
//				ImportCustomRoutes: pulumi.Bool(true),
//				ExportCustomRoutes: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewNetworkPeeringRoutesConfig(ctx, "peeringPrimaryRoutes", &compute.NetworkPeeringRoutesConfigArgs{
//				Peering:            peeringPrimary.Name,
//				Network:            networkPrimary.Name,
//				ImportCustomRoutes: pulumi.Bool(true),
//				ExportCustomRoutes: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewNetworkPeering(ctx, "peeringSecondary", &compute.NetworkPeeringArgs{
//				Network:     networkSecondary.ID(),
//				PeerNetwork: networkPrimary.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Network Peering Routes Config Gke
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			containerNetwork, err := compute.NewNetwork(ctx, "containerNetwork", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			containerSubnetwork, err := compute.NewSubnetwork(ctx, "containerSubnetwork", &compute.SubnetworkArgs{
//				Region:                pulumi.String("us-central1"),
//				Network:               containerNetwork.Name,
//				IpCidrRange:           pulumi.String("10.0.36.0/24"),
//				PrivateIpGoogleAccess: pulumi.Bool(true),
//				SecondaryIpRanges: compute.SubnetworkSecondaryIpRangeArray{
//					&compute.SubnetworkSecondaryIpRangeArgs{
//						RangeName:   pulumi.String("pod"),
//						IpCidrRange: pulumi.String("10.0.0.0/19"),
//					},
//					&compute.SubnetworkSecondaryIpRangeArgs{
//						RangeName:   pulumi.String("svc"),
//						IpCidrRange: pulumi.String("10.0.32.0/22"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			privateCluster, err := container.NewCluster(ctx, "privateCluster", &container.ClusterArgs{
//				Location:         pulumi.String("us-central1-a"),
//				InitialNodeCount: pulumi.Int(1),
//				Network:          containerNetwork.Name,
//				Subnetwork:       containerSubnetwork.Name,
//				PrivateClusterConfig: &container.ClusterPrivateClusterConfigArgs{
//					EnablePrivateEndpoint: pulumi.Bool(true),
//					EnablePrivateNodes:    pulumi.Bool(true),
//					MasterIpv4CidrBlock:   pulumi.String("10.42.0.0/28"),
//				},
//				MasterAuthorizedNetworksConfig: nil,
//				IpAllocationPolicy: &container.ClusterIpAllocationPolicyArgs{
//					ClusterSecondaryRangeName: containerSubnetwork.SecondaryIpRanges.ApplyT(func(secondaryIpRanges []compute.SubnetworkSecondaryIpRange) (*string, error) {
//						return &secondaryIpRanges[0].RangeName, nil
//					}).(pulumi.StringPtrOutput),
//					ServicesSecondaryRangeName: containerSubnetwork.SecondaryIpRanges.ApplyT(func(secondaryIpRanges []compute.SubnetworkSecondaryIpRange) (*string, error) {
//						return &secondaryIpRanges[1].RangeName, nil
//					}).(pulumi.StringPtrOutput),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewNetworkPeeringRoutesConfig(ctx, "peeringGkeRoutes", &compute.NetworkPeeringRoutesConfigArgs{
//				Peering: privateCluster.PrivateClusterConfig.ApplyT(func(privateClusterConfig container.ClusterPrivateClusterConfig) (*string, error) {
//					return &privateClusterConfig.PeeringName, nil
//				}).(pulumi.StringPtrOutput),
//				Network:            containerNetwork.Name,
//				ImportCustomRoutes: pulumi.Bool(true),
//				ExportCustomRoutes: pulumi.Bool(true),
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
// # NetworkPeeringRoutesConfig can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default {{project}}/{{network}}/{{peering}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default {{network}}/{{peering}}
//
// ```
type NetworkPeeringRoutesConfig struct {
	pulumi.CustomResourceState

	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes pulumi.BoolOutput `pulumi:"exportCustomRoutes"`
	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes pulumi.BoolOutput `pulumi:"importCustomRoutes"`
	// The name of the primary network for the peering.
	//
	// ***
	Network pulumi.StringOutput `pulumi:"network"`
	// Name of the peering.
	Peering pulumi.StringOutput `pulumi:"peering"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewNetworkPeeringRoutesConfig registers a new resource with the given unique name, arguments, and options.
func NewNetworkPeeringRoutesConfig(ctx *pulumi.Context,
	name string, args *NetworkPeeringRoutesConfigArgs, opts ...pulumi.ResourceOption) (*NetworkPeeringRoutesConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExportCustomRoutes == nil {
		return nil, errors.New("invalid value for required argument 'ExportCustomRoutes'")
	}
	if args.ImportCustomRoutes == nil {
		return nil, errors.New("invalid value for required argument 'ImportCustomRoutes'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.Peering == nil {
		return nil, errors.New("invalid value for required argument 'Peering'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkPeeringRoutesConfig
	err := ctx.RegisterResource("gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPeeringRoutesConfig gets an existing NetworkPeeringRoutesConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPeeringRoutesConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPeeringRoutesConfigState, opts ...pulumi.ResourceOption) (*NetworkPeeringRoutesConfig, error) {
	var resource NetworkPeeringRoutesConfig
	err := ctx.ReadResource("gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPeeringRoutesConfig resources.
type networkPeeringRoutesConfigState struct {
	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes *bool `pulumi:"exportCustomRoutes"`
	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes *bool `pulumi:"importCustomRoutes"`
	// The name of the primary network for the peering.
	//
	// ***
	Network *string `pulumi:"network"`
	// Name of the peering.
	Peering *string `pulumi:"peering"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type NetworkPeeringRoutesConfigState struct {
	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes pulumi.BoolPtrInput
	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes pulumi.BoolPtrInput
	// The name of the primary network for the peering.
	//
	// ***
	Network pulumi.StringPtrInput
	// Name of the peering.
	Peering pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (NetworkPeeringRoutesConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPeeringRoutesConfigState)(nil)).Elem()
}

type networkPeeringRoutesConfigArgs struct {
	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes bool `pulumi:"exportCustomRoutes"`
	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes bool `pulumi:"importCustomRoutes"`
	// The name of the primary network for the peering.
	//
	// ***
	Network string `pulumi:"network"`
	// Name of the peering.
	Peering string `pulumi:"peering"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a NetworkPeeringRoutesConfig resource.
type NetworkPeeringRoutesConfigArgs struct {
	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes pulumi.BoolInput
	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes pulumi.BoolInput
	// The name of the primary network for the peering.
	//
	// ***
	Network pulumi.StringInput
	// Name of the peering.
	Peering pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (NetworkPeeringRoutesConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPeeringRoutesConfigArgs)(nil)).Elem()
}

type NetworkPeeringRoutesConfigInput interface {
	pulumi.Input

	ToNetworkPeeringRoutesConfigOutput() NetworkPeeringRoutesConfigOutput
	ToNetworkPeeringRoutesConfigOutputWithContext(ctx context.Context) NetworkPeeringRoutesConfigOutput
}

func (*NetworkPeeringRoutesConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPeeringRoutesConfig)(nil)).Elem()
}

func (i *NetworkPeeringRoutesConfig) ToNetworkPeeringRoutesConfigOutput() NetworkPeeringRoutesConfigOutput {
	return i.ToNetworkPeeringRoutesConfigOutputWithContext(context.Background())
}

func (i *NetworkPeeringRoutesConfig) ToNetworkPeeringRoutesConfigOutputWithContext(ctx context.Context) NetworkPeeringRoutesConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPeeringRoutesConfigOutput)
}

// NetworkPeeringRoutesConfigArrayInput is an input type that accepts NetworkPeeringRoutesConfigArray and NetworkPeeringRoutesConfigArrayOutput values.
// You can construct a concrete instance of `NetworkPeeringRoutesConfigArrayInput` via:
//
//	NetworkPeeringRoutesConfigArray{ NetworkPeeringRoutesConfigArgs{...} }
type NetworkPeeringRoutesConfigArrayInput interface {
	pulumi.Input

	ToNetworkPeeringRoutesConfigArrayOutput() NetworkPeeringRoutesConfigArrayOutput
	ToNetworkPeeringRoutesConfigArrayOutputWithContext(context.Context) NetworkPeeringRoutesConfigArrayOutput
}

type NetworkPeeringRoutesConfigArray []NetworkPeeringRoutesConfigInput

func (NetworkPeeringRoutesConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPeeringRoutesConfig)(nil)).Elem()
}

func (i NetworkPeeringRoutesConfigArray) ToNetworkPeeringRoutesConfigArrayOutput() NetworkPeeringRoutesConfigArrayOutput {
	return i.ToNetworkPeeringRoutesConfigArrayOutputWithContext(context.Background())
}

func (i NetworkPeeringRoutesConfigArray) ToNetworkPeeringRoutesConfigArrayOutputWithContext(ctx context.Context) NetworkPeeringRoutesConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPeeringRoutesConfigArrayOutput)
}

// NetworkPeeringRoutesConfigMapInput is an input type that accepts NetworkPeeringRoutesConfigMap and NetworkPeeringRoutesConfigMapOutput values.
// You can construct a concrete instance of `NetworkPeeringRoutesConfigMapInput` via:
//
//	NetworkPeeringRoutesConfigMap{ "key": NetworkPeeringRoutesConfigArgs{...} }
type NetworkPeeringRoutesConfigMapInput interface {
	pulumi.Input

	ToNetworkPeeringRoutesConfigMapOutput() NetworkPeeringRoutesConfigMapOutput
	ToNetworkPeeringRoutesConfigMapOutputWithContext(context.Context) NetworkPeeringRoutesConfigMapOutput
}

type NetworkPeeringRoutesConfigMap map[string]NetworkPeeringRoutesConfigInput

func (NetworkPeeringRoutesConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPeeringRoutesConfig)(nil)).Elem()
}

func (i NetworkPeeringRoutesConfigMap) ToNetworkPeeringRoutesConfigMapOutput() NetworkPeeringRoutesConfigMapOutput {
	return i.ToNetworkPeeringRoutesConfigMapOutputWithContext(context.Background())
}

func (i NetworkPeeringRoutesConfigMap) ToNetworkPeeringRoutesConfigMapOutputWithContext(ctx context.Context) NetworkPeeringRoutesConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPeeringRoutesConfigMapOutput)
}

type NetworkPeeringRoutesConfigOutput struct{ *pulumi.OutputState }

func (NetworkPeeringRoutesConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPeeringRoutesConfig)(nil)).Elem()
}

func (o NetworkPeeringRoutesConfigOutput) ToNetworkPeeringRoutesConfigOutput() NetworkPeeringRoutesConfigOutput {
	return o
}

func (o NetworkPeeringRoutesConfigOutput) ToNetworkPeeringRoutesConfigOutputWithContext(ctx context.Context) NetworkPeeringRoutesConfigOutput {
	return o
}

// Whether to export the custom routes to the peer network.
func (o NetworkPeeringRoutesConfigOutput) ExportCustomRoutes() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkPeeringRoutesConfig) pulumi.BoolOutput { return v.ExportCustomRoutes }).(pulumi.BoolOutput)
}

// Whether to import the custom routes to the peer network.
func (o NetworkPeeringRoutesConfigOutput) ImportCustomRoutes() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkPeeringRoutesConfig) pulumi.BoolOutput { return v.ImportCustomRoutes }).(pulumi.BoolOutput)
}

// The name of the primary network for the peering.
//
// ***
func (o NetworkPeeringRoutesConfigOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPeeringRoutesConfig) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// Name of the peering.
func (o NetworkPeeringRoutesConfigOutput) Peering() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPeeringRoutesConfig) pulumi.StringOutput { return v.Peering }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o NetworkPeeringRoutesConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPeeringRoutesConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type NetworkPeeringRoutesConfigArrayOutput struct{ *pulumi.OutputState }

func (NetworkPeeringRoutesConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPeeringRoutesConfig)(nil)).Elem()
}

func (o NetworkPeeringRoutesConfigArrayOutput) ToNetworkPeeringRoutesConfigArrayOutput() NetworkPeeringRoutesConfigArrayOutput {
	return o
}

func (o NetworkPeeringRoutesConfigArrayOutput) ToNetworkPeeringRoutesConfigArrayOutputWithContext(ctx context.Context) NetworkPeeringRoutesConfigArrayOutput {
	return o
}

func (o NetworkPeeringRoutesConfigArrayOutput) Index(i pulumi.IntInput) NetworkPeeringRoutesConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkPeeringRoutesConfig {
		return vs[0].([]*NetworkPeeringRoutesConfig)[vs[1].(int)]
	}).(NetworkPeeringRoutesConfigOutput)
}

type NetworkPeeringRoutesConfigMapOutput struct{ *pulumi.OutputState }

func (NetworkPeeringRoutesConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPeeringRoutesConfig)(nil)).Elem()
}

func (o NetworkPeeringRoutesConfigMapOutput) ToNetworkPeeringRoutesConfigMapOutput() NetworkPeeringRoutesConfigMapOutput {
	return o
}

func (o NetworkPeeringRoutesConfigMapOutput) ToNetworkPeeringRoutesConfigMapOutputWithContext(ctx context.Context) NetworkPeeringRoutesConfigMapOutput {
	return o
}

func (o NetworkPeeringRoutesConfigMapOutput) MapIndex(k pulumi.StringInput) NetworkPeeringRoutesConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkPeeringRoutesConfig {
		return vs[0].(map[string]*NetworkPeeringRoutesConfig)[vs[1].(string)]
	}).(NetworkPeeringRoutesConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPeeringRoutesConfigInput)(nil)).Elem(), &NetworkPeeringRoutesConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPeeringRoutesConfigArrayInput)(nil)).Elem(), NetworkPeeringRoutesConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPeeringRoutesConfigMapInput)(nil)).Elem(), NetworkPeeringRoutesConfigMap{})
	pulumi.RegisterOutputType(NetworkPeeringRoutesConfigOutput{})
	pulumi.RegisterOutputType(NetworkPeeringRoutesConfigArrayOutput{})
	pulumi.RegisterOutputType(NetworkPeeringRoutesConfigMapOutput{})
}