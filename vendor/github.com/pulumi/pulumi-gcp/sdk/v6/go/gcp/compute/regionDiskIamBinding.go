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

// Three different resources help you manage your IAM policy for Compute Engine Disk. Each of these resources serves a different use case:
//
// * `compute.DiskIamPolicy`: Authoritative. Sets the IAM policy for the disk and replaces any existing policy already attached.
// * `compute.DiskIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the disk are preserved.
// * `compute.DiskIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the disk are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `compute.DiskIamPolicy`: Retrieves the IAM policy for the disk
//
// > **Note:** `compute.DiskIamPolicy` **cannot** be used in conjunction with `compute.DiskIamBinding` and `compute.DiskIamMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.DiskIamBinding` resources **can be** used in conjunction with `compute.DiskIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_compute\_disk\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewDiskIamPolicy(ctx, "policy", &compute.DiskIamPolicyArgs{
//				Project:    pulumi.Any(google_compute_disk.Default.Project),
//				Zone:       pulumi.Any(google_compute_disk.Default.Zone),
//				PolicyData: *pulumi.String(admin.PolicyData),
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
// ## google\_compute\_disk\_iam\_binding
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
//			_, err := compute.NewDiskIamBinding(ctx, "binding", &compute.DiskIamBindingArgs{
//				Project: pulumi.Any(google_compute_disk.Default.Project),
//				Zone:    pulumi.Any(google_compute_disk.Default.Zone),
//				Role:    pulumi.String("roles/viewer"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
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
// ## google\_compute\_disk\_iam\_member
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
//			_, err := compute.NewDiskIamMember(ctx, "member", &compute.DiskIamMemberArgs{
//				Project: pulumi.Any(google_compute_disk.Default.Project),
//				Zone:    pulumi.Any(google_compute_disk.Default.Zone),
//				Role:    pulumi.String("roles/viewer"),
//				Member:  pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/zones/{{zone}}/disks/{{name}} * {{project}}/{{zone}}/{{name}} * {{zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine disk IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/regionDiskIamBinding:RegionDiskIamBinding editor "projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/regionDiskIamBinding:RegionDiskIamBinding editor "projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/regionDiskIamBinding:RegionDiskIamBinding editor projects/{{project}}/zones/{{zone}}/disks/{{disk}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RegionDiskIamBinding struct {
	pulumi.CustomResourceState

	Condition RegionDiskIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRegionDiskIamBinding registers a new resource with the given unique name, arguments, and options.
func NewRegionDiskIamBinding(ctx *pulumi.Context,
	name string, args *RegionDiskIamBindingArgs, opts ...pulumi.ResourceOption) (*RegionDiskIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegionDiskIamBinding
	err := ctx.RegisterResource("gcp:compute/regionDiskIamBinding:RegionDiskIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionDiskIamBinding gets an existing RegionDiskIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionDiskIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionDiskIamBindingState, opts ...pulumi.ResourceOption) (*RegionDiskIamBinding, error) {
	var resource RegionDiskIamBinding
	err := ctx.ReadResource("gcp:compute/regionDiskIamBinding:RegionDiskIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionDiskIamBinding resources.
type regionDiskIamBindingState struct {
	Condition *RegionDiskIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type RegionDiskIamBindingState struct {
	Condition RegionDiskIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (RegionDiskIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskIamBindingState)(nil)).Elem()
}

type regionDiskIamBindingArgs struct {
	Condition *RegionDiskIamBindingCondition `pulumi:"condition"`
	Members   []string                       `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a RegionDiskIamBinding resource.
type RegionDiskIamBindingArgs struct {
	Condition RegionDiskIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (RegionDiskIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionDiskIamBindingArgs)(nil)).Elem()
}

type RegionDiskIamBindingInput interface {
	pulumi.Input

	ToRegionDiskIamBindingOutput() RegionDiskIamBindingOutput
	ToRegionDiskIamBindingOutputWithContext(ctx context.Context) RegionDiskIamBindingOutput
}

func (*RegionDiskIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionDiskIamBinding)(nil)).Elem()
}

func (i *RegionDiskIamBinding) ToRegionDiskIamBindingOutput() RegionDiskIamBindingOutput {
	return i.ToRegionDiskIamBindingOutputWithContext(context.Background())
}

func (i *RegionDiskIamBinding) ToRegionDiskIamBindingOutputWithContext(ctx context.Context) RegionDiskIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskIamBindingOutput)
}

// RegionDiskIamBindingArrayInput is an input type that accepts RegionDiskIamBindingArray and RegionDiskIamBindingArrayOutput values.
// You can construct a concrete instance of `RegionDiskIamBindingArrayInput` via:
//
//	RegionDiskIamBindingArray{ RegionDiskIamBindingArgs{...} }
type RegionDiskIamBindingArrayInput interface {
	pulumi.Input

	ToRegionDiskIamBindingArrayOutput() RegionDiskIamBindingArrayOutput
	ToRegionDiskIamBindingArrayOutputWithContext(context.Context) RegionDiskIamBindingArrayOutput
}

type RegionDiskIamBindingArray []RegionDiskIamBindingInput

func (RegionDiskIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionDiskIamBinding)(nil)).Elem()
}

func (i RegionDiskIamBindingArray) ToRegionDiskIamBindingArrayOutput() RegionDiskIamBindingArrayOutput {
	return i.ToRegionDiskIamBindingArrayOutputWithContext(context.Background())
}

func (i RegionDiskIamBindingArray) ToRegionDiskIamBindingArrayOutputWithContext(ctx context.Context) RegionDiskIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskIamBindingArrayOutput)
}

// RegionDiskIamBindingMapInput is an input type that accepts RegionDiskIamBindingMap and RegionDiskIamBindingMapOutput values.
// You can construct a concrete instance of `RegionDiskIamBindingMapInput` via:
//
//	RegionDiskIamBindingMap{ "key": RegionDiskIamBindingArgs{...} }
type RegionDiskIamBindingMapInput interface {
	pulumi.Input

	ToRegionDiskIamBindingMapOutput() RegionDiskIamBindingMapOutput
	ToRegionDiskIamBindingMapOutputWithContext(context.Context) RegionDiskIamBindingMapOutput
}

type RegionDiskIamBindingMap map[string]RegionDiskIamBindingInput

func (RegionDiskIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionDiskIamBinding)(nil)).Elem()
}

func (i RegionDiskIamBindingMap) ToRegionDiskIamBindingMapOutput() RegionDiskIamBindingMapOutput {
	return i.ToRegionDiskIamBindingMapOutputWithContext(context.Background())
}

func (i RegionDiskIamBindingMap) ToRegionDiskIamBindingMapOutputWithContext(ctx context.Context) RegionDiskIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionDiskIamBindingMapOutput)
}

type RegionDiskIamBindingOutput struct{ *pulumi.OutputState }

func (RegionDiskIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionDiskIamBinding)(nil)).Elem()
}

func (o RegionDiskIamBindingOutput) ToRegionDiskIamBindingOutput() RegionDiskIamBindingOutput {
	return o
}

func (o RegionDiskIamBindingOutput) ToRegionDiskIamBindingOutputWithContext(ctx context.Context) RegionDiskIamBindingOutput {
	return o
}

func (o RegionDiskIamBindingOutput) Condition() RegionDiskIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *RegionDiskIamBinding) RegionDiskIamBindingConditionPtrOutput { return v.Condition }).(RegionDiskIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o RegionDiskIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionDiskIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o RegionDiskIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RegionDiskIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o RegionDiskIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionDiskIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
//   - **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
//   - **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
//   - **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
func (o RegionDiskIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionDiskIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o RegionDiskIamBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionDiskIamBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `compute.DiskIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o RegionDiskIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionDiskIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type RegionDiskIamBindingArrayOutput struct{ *pulumi.OutputState }

func (RegionDiskIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionDiskIamBinding)(nil)).Elem()
}

func (o RegionDiskIamBindingArrayOutput) ToRegionDiskIamBindingArrayOutput() RegionDiskIamBindingArrayOutput {
	return o
}

func (o RegionDiskIamBindingArrayOutput) ToRegionDiskIamBindingArrayOutputWithContext(ctx context.Context) RegionDiskIamBindingArrayOutput {
	return o
}

func (o RegionDiskIamBindingArrayOutput) Index(i pulumi.IntInput) RegionDiskIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegionDiskIamBinding {
		return vs[0].([]*RegionDiskIamBinding)[vs[1].(int)]
	}).(RegionDiskIamBindingOutput)
}

type RegionDiskIamBindingMapOutput struct{ *pulumi.OutputState }

func (RegionDiskIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionDiskIamBinding)(nil)).Elem()
}

func (o RegionDiskIamBindingMapOutput) ToRegionDiskIamBindingMapOutput() RegionDiskIamBindingMapOutput {
	return o
}

func (o RegionDiskIamBindingMapOutput) ToRegionDiskIamBindingMapOutputWithContext(ctx context.Context) RegionDiskIamBindingMapOutput {
	return o
}

func (o RegionDiskIamBindingMapOutput) MapIndex(k pulumi.StringInput) RegionDiskIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegionDiskIamBinding {
		return vs[0].(map[string]*RegionDiskIamBinding)[vs[1].(string)]
	}).(RegionDiskIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionDiskIamBindingInput)(nil)).Elem(), &RegionDiskIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionDiskIamBindingArrayInput)(nil)).Elem(), RegionDiskIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionDiskIamBindingMapInput)(nil)).Elem(), RegionDiskIamBindingMap{})
	pulumi.RegisterOutputType(RegionDiskIamBindingOutput{})
	pulumi.RegisterOutputType(RegionDiskIamBindingArrayOutput{})
	pulumi.RegisterOutputType(RegionDiskIamBindingMapOutput{})
}