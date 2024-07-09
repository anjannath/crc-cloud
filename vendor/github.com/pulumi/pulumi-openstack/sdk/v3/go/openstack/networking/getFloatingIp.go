// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack floating IP.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.LookupFloatingIp(ctx, &networking.LookupFloatingIpArgs{
//				Address: pulumi.StringRef("192.168.0.4"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFloatingIp(ctx *pulumi.Context, args *LookupFloatingIpArgs, opts ...pulumi.InvokeOption) (*LookupFloatingIpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFloatingIpResult
	err := ctx.Invoke("openstack:networking/getFloatingIp:getFloatingIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFloatingIp.
type LookupFloatingIpArgs struct {
	// The IP address of the floating IP.
	Address *string `pulumi:"address"`
	// Human-readable description of the floating IP.
	Description *string `pulumi:"description"`
	// The specific IP address of the internal port which should be associated with the floating IP.
	FixedIp *string `pulumi:"fixedIp"`
	// The name of the pool from which the floating IP belongs to.
	Pool *string `pulumi:"pool"`
	// The ID of the port the floating IP is attached.
	PortId *string `pulumi:"portId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve floating IP ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// status of the floating IP (ACTIVE/DOWN).
	Status *string `pulumi:"status"`
	// The list of floating IP tags to filter.
	Tags []string `pulumi:"tags"`
	// The owner of the floating IP.
	TenantId *string `pulumi:"tenantId"`
}

// A collection of values returned by getFloatingIp.
type LookupFloatingIpResult struct {
	Address *string `pulumi:"address"`
	// A set of string tags applied on the floating IP.
	AllTags     []string `pulumi:"allTags"`
	Description *string  `pulumi:"description"`
	// The floating IP DNS domain. Available, when Neutron DNS
	// extension is enabled.
	DnsDomain string `pulumi:"dnsDomain"`
	// The floating IP DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName string  `pulumi:"dnsName"`
	FixedIp *string `pulumi:"fixedIp"`
	// The provider-assigned unique ID for this managed resource.
	Id       string   `pulumi:"id"`
	Pool     *string  `pulumi:"pool"`
	PortId   *string  `pulumi:"portId"`
	Region   *string  `pulumi:"region"`
	Status   *string  `pulumi:"status"`
	Tags     []string `pulumi:"tags"`
	TenantId *string  `pulumi:"tenantId"`
}

func LookupFloatingIpOutput(ctx *pulumi.Context, args LookupFloatingIpOutputArgs, opts ...pulumi.InvokeOption) LookupFloatingIpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFloatingIpResult, error) {
			args := v.(LookupFloatingIpArgs)
			r, err := LookupFloatingIp(ctx, &args, opts...)
			var s LookupFloatingIpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFloatingIpResultOutput)
}

// A collection of arguments for invoking getFloatingIp.
type LookupFloatingIpOutputArgs struct {
	// The IP address of the floating IP.
	Address pulumi.StringPtrInput `pulumi:"address"`
	// Human-readable description of the floating IP.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The specific IP address of the internal port which should be associated with the floating IP.
	FixedIp pulumi.StringPtrInput `pulumi:"fixedIp"`
	// The name of the pool from which the floating IP belongs to.
	Pool pulumi.StringPtrInput `pulumi:"pool"`
	// The ID of the port the floating IP is attached.
	PortId pulumi.StringPtrInput `pulumi:"portId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve floating IP ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// status of the floating IP (ACTIVE/DOWN).
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The list of floating IP tags to filter.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The owner of the floating IP.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (LookupFloatingIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFloatingIpArgs)(nil)).Elem()
}

// A collection of values returned by getFloatingIp.
type LookupFloatingIpResultOutput struct{ *pulumi.OutputState }

func (LookupFloatingIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFloatingIpResult)(nil)).Elem()
}

func (o LookupFloatingIpResultOutput) ToLookupFloatingIpResultOutput() LookupFloatingIpResultOutput {
	return o
}

func (o LookupFloatingIpResultOutput) ToLookupFloatingIpResultOutputWithContext(ctx context.Context) LookupFloatingIpResultOutput {
	return o
}

func (o LookupFloatingIpResultOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// A set of string tags applied on the floating IP.
func (o LookupFloatingIpResultOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) []string { return v.AllTags }).(pulumi.StringArrayOutput)
}

func (o LookupFloatingIpResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The floating IP DNS domain. Available, when Neutron DNS
// extension is enabled.
func (o LookupFloatingIpResultOutput) DnsDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) string { return v.DnsDomain }).(pulumi.StringOutput)
}

// The floating IP DNS name. Available, when Neutron DNS extension
// is enabled.
func (o LookupFloatingIpResultOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) string { return v.DnsName }).(pulumi.StringOutput)
}

func (o LookupFloatingIpResultOutput) FixedIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) *string { return v.FixedIp }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFloatingIpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFloatingIpResultOutput) Pool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) *string { return v.Pool }).(pulumi.StringPtrOutput)
}

func (o LookupFloatingIpResultOutput) PortId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) *string { return v.PortId }).(pulumi.StringPtrOutput)
}

func (o LookupFloatingIpResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupFloatingIpResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupFloatingIpResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupFloatingIpResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFloatingIpResultOutput{})
}