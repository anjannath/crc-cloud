// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// `ec2.getVpcIpamPoolCidrs` provides details about an IPAM pool.
//
// This resource can prove useful when an ipam pool was shared to your account and you want to know all (or a filtered list) of the CIDRs that are provisioned into the pool.
func GetVpcIpamPoolCidrs(ctx *pulumi.Context, args *GetVpcIpamPoolCidrsArgs, opts ...pulumi.InvokeOption) (*GetVpcIpamPoolCidrsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpcIpamPoolCidrsResult
	err := ctx.Invoke("aws:ec2/getVpcIpamPoolCidrs:getVpcIpamPoolCidrs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcIpamPoolCidrs.
type GetVpcIpamPoolCidrsArgs struct {
	// Custom filter block as described below.
	Filters []GetVpcIpamPoolCidrsFilter `pulumi:"filters"`
	// ID of the IPAM pool you would like the list of provisioned CIDRs.
	IpamPoolId string `pulumi:"ipamPoolId"`
}

// A collection of values returned by getVpcIpamPoolCidrs.
type GetVpcIpamPoolCidrsResult struct {
	Filters []GetVpcIpamPoolCidrsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The CIDRs provisioned into the IPAM pool, described below.
	IpamPoolCidrs []GetVpcIpamPoolCidrsIpamPoolCidr `pulumi:"ipamPoolCidrs"`
	IpamPoolId    string                            `pulumi:"ipamPoolId"`
}

func GetVpcIpamPoolCidrsOutput(ctx *pulumi.Context, args GetVpcIpamPoolCidrsOutputArgs, opts ...pulumi.InvokeOption) GetVpcIpamPoolCidrsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpcIpamPoolCidrsResult, error) {
			args := v.(GetVpcIpamPoolCidrsArgs)
			r, err := GetVpcIpamPoolCidrs(ctx, &args, opts...)
			var s GetVpcIpamPoolCidrsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpcIpamPoolCidrsResultOutput)
}

// A collection of arguments for invoking getVpcIpamPoolCidrs.
type GetVpcIpamPoolCidrsOutputArgs struct {
	// Custom filter block as described below.
	Filters GetVpcIpamPoolCidrsFilterArrayInput `pulumi:"filters"`
	// ID of the IPAM pool you would like the list of provisioned CIDRs.
	IpamPoolId pulumi.StringInput `pulumi:"ipamPoolId"`
}

func (GetVpcIpamPoolCidrsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcIpamPoolCidrsArgs)(nil)).Elem()
}

// A collection of values returned by getVpcIpamPoolCidrs.
type GetVpcIpamPoolCidrsResultOutput struct{ *pulumi.OutputState }

func (GetVpcIpamPoolCidrsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcIpamPoolCidrsResult)(nil)).Elem()
}

func (o GetVpcIpamPoolCidrsResultOutput) ToGetVpcIpamPoolCidrsResultOutput() GetVpcIpamPoolCidrsResultOutput {
	return o
}

func (o GetVpcIpamPoolCidrsResultOutput) ToGetVpcIpamPoolCidrsResultOutputWithContext(ctx context.Context) GetVpcIpamPoolCidrsResultOutput {
	return o
}

func (o GetVpcIpamPoolCidrsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetVpcIpamPoolCidrsResult] {
	return pulumix.Output[GetVpcIpamPoolCidrsResult]{
		OutputState: o.OutputState,
	}
}

func (o GetVpcIpamPoolCidrsResultOutput) Filters() GetVpcIpamPoolCidrsFilterArrayOutput {
	return o.ApplyT(func(v GetVpcIpamPoolCidrsResult) []GetVpcIpamPoolCidrsFilter { return v.Filters }).(GetVpcIpamPoolCidrsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcIpamPoolCidrsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcIpamPoolCidrsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The CIDRs provisioned into the IPAM pool, described below.
func (o GetVpcIpamPoolCidrsResultOutput) IpamPoolCidrs() GetVpcIpamPoolCidrsIpamPoolCidrArrayOutput {
	return o.ApplyT(func(v GetVpcIpamPoolCidrsResult) []GetVpcIpamPoolCidrsIpamPoolCidr { return v.IpamPoolCidrs }).(GetVpcIpamPoolCidrsIpamPoolCidrArrayOutput)
}

func (o GetVpcIpamPoolCidrsResultOutput) IpamPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcIpamPoolCidrsResult) string { return v.IpamPoolId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcIpamPoolCidrsResultOutput{})
}