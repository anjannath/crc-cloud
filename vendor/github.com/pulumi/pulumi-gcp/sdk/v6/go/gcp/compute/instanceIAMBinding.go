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

// Three different resources help you manage your IAM policy for Compute Engine Instance. Each of these resources serves a different use case:
//
// * `compute.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
// * `compute.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `compute.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `compute.InstanceIAMPolicy`: Retrieves the IAM policy for the instance
//
// > **Note:** `compute.InstanceIAMPolicy` **cannot** be used in conjunction with `compute.InstanceIAMBinding` and `compute.InstanceIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.InstanceIAMBinding` resources **can be** used in conjunction with `compute.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## google\_compute\_instance\_iam\_policy
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
//						Role: "roles/compute.osLogin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInstanceIAMPolicy(ctx, "policy", &compute.InstanceIAMPolicyArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				PolicyData:   *pulumi.String(admin.PolicyData),
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
// With IAM Conditions:
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
//						Role: "roles/compute.osLogin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Condition: {
//							Title:       "expires_after_2019_12_31",
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInstanceIAMPolicy(ctx, "policy", &compute.InstanceIAMPolicyArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				PolicyData:   *pulumi.String(admin.PolicyData),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## google\_compute\_instance\_iam\_binding
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
//			_, err := compute.NewInstanceIAMBinding(ctx, "binding", &compute.InstanceIAMBindingArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				Role:         pulumi.String("roles/compute.osLogin"),
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
// With IAM Conditions:
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
//			_, err := compute.NewInstanceIAMBinding(ctx, "binding", &compute.InstanceIAMBindingArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				Role:         pulumi.String("roles/compute.osLogin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &compute.InstanceIAMBindingConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## google\_compute\_instance\_iam\_member
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
//			_, err := compute.NewInstanceIAMMember(ctx, "member", &compute.InstanceIAMMemberArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				Role:         pulumi.String("roles/compute.osLogin"),
//				Member:       pulumi.String("user:jane@example.com"),
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
// With IAM Conditions:
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
//			_, err := compute.NewInstanceIAMMember(ctx, "member", &compute.InstanceIAMMemberArgs{
//				Project:      pulumi.Any(google_compute_instance.Default.Project),
//				Zone:         pulumi.Any(google_compute_instance.Default.Zone),
//				InstanceName: pulumi.Any(google_compute_instance.Default.Name),
//				Role:         pulumi.String("roles/compute.osLogin"),
//				Member:       pulumi.String("user:jane@example.com"),
//				Condition: &compute.InstanceIAMMemberConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/zones/{{zone}}/instances/{{name}} * {{project}}/{{zone}}/{{name}} * {{zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine instance IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceIAMBinding:InstanceIAMBinding editor "projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceIAMBinding:InstanceIAMBinding editor "projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceIAMBinding:InstanceIAMBinding editor projects/{{project}}/zones/{{zone}}/instances/{{instance}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIAMBinding struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition InstanceIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringOutput      `pulumi:"instanceName"`
	Members      pulumi.StringArrayOutput `pulumi:"members"`
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
	// The role that should be applied. Only one
	// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewInstanceIAMBinding(ctx *pulumi.Context,
	name string, args *InstanceIAMBindingArgs, opts ...pulumi.ResourceOption) (*InstanceIAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceIAMBinding
	err := ctx.RegisterResource("gcp:compute/instanceIAMBinding:InstanceIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIAMBinding gets an existing InstanceIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIAMBindingState, opts ...pulumi.ResourceOption) (*InstanceIAMBinding, error) {
	var resource InstanceIAMBinding
	err := ctx.ReadResource("gcp:compute/instanceIAMBinding:InstanceIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIAMBinding resources.
type instanceIAMBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *InstanceIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName *string  `pulumi:"instanceName"`
	Members      []string `pulumi:"members"`
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
	// The role that should be applied. Only one
	// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone *string `pulumi:"zone"`
}

type InstanceIAMBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition InstanceIAMBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringPtrInput
	Members      pulumi.StringArrayInput
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
	// The role that should be applied. Only one
	// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringPtrInput
}

func (InstanceIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMBindingState)(nil)).Elem()
}

type instanceIAMBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *InstanceIAMBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName string   `pulumi:"instanceName"`
	Members      []string `pulumi:"members"`
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
	// The role that should be applied. Only one
	// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceIAMBinding resource.
type InstanceIAMBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition InstanceIAMBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringInput
	Members      pulumi.StringArrayInput
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
	// The role that should be applied. Only one
	// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone pulumi.StringPtrInput
}

func (InstanceIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMBindingArgs)(nil)).Elem()
}

type InstanceIAMBindingInput interface {
	pulumi.Input

	ToInstanceIAMBindingOutput() InstanceIAMBindingOutput
	ToInstanceIAMBindingOutputWithContext(ctx context.Context) InstanceIAMBindingOutput
}

func (*InstanceIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIAMBinding)(nil)).Elem()
}

func (i *InstanceIAMBinding) ToInstanceIAMBindingOutput() InstanceIAMBindingOutput {
	return i.ToInstanceIAMBindingOutputWithContext(context.Background())
}

func (i *InstanceIAMBinding) ToInstanceIAMBindingOutputWithContext(ctx context.Context) InstanceIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMBindingOutput)
}

// InstanceIAMBindingArrayInput is an input type that accepts InstanceIAMBindingArray and InstanceIAMBindingArrayOutput values.
// You can construct a concrete instance of `InstanceIAMBindingArrayInput` via:
//
//	InstanceIAMBindingArray{ InstanceIAMBindingArgs{...} }
type InstanceIAMBindingArrayInput interface {
	pulumi.Input

	ToInstanceIAMBindingArrayOutput() InstanceIAMBindingArrayOutput
	ToInstanceIAMBindingArrayOutputWithContext(context.Context) InstanceIAMBindingArrayOutput
}

type InstanceIAMBindingArray []InstanceIAMBindingInput

func (InstanceIAMBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIAMBinding)(nil)).Elem()
}

func (i InstanceIAMBindingArray) ToInstanceIAMBindingArrayOutput() InstanceIAMBindingArrayOutput {
	return i.ToInstanceIAMBindingArrayOutputWithContext(context.Background())
}

func (i InstanceIAMBindingArray) ToInstanceIAMBindingArrayOutputWithContext(ctx context.Context) InstanceIAMBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMBindingArrayOutput)
}

// InstanceIAMBindingMapInput is an input type that accepts InstanceIAMBindingMap and InstanceIAMBindingMapOutput values.
// You can construct a concrete instance of `InstanceIAMBindingMapInput` via:
//
//	InstanceIAMBindingMap{ "key": InstanceIAMBindingArgs{...} }
type InstanceIAMBindingMapInput interface {
	pulumi.Input

	ToInstanceIAMBindingMapOutput() InstanceIAMBindingMapOutput
	ToInstanceIAMBindingMapOutputWithContext(context.Context) InstanceIAMBindingMapOutput
}

type InstanceIAMBindingMap map[string]InstanceIAMBindingInput

func (InstanceIAMBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIAMBinding)(nil)).Elem()
}

func (i InstanceIAMBindingMap) ToInstanceIAMBindingMapOutput() InstanceIAMBindingMapOutput {
	return i.ToInstanceIAMBindingMapOutputWithContext(context.Background())
}

func (i InstanceIAMBindingMap) ToInstanceIAMBindingMapOutputWithContext(ctx context.Context) InstanceIAMBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMBindingMapOutput)
}

type InstanceIAMBindingOutput struct{ *pulumi.OutputState }

func (InstanceIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIAMBinding)(nil)).Elem()
}

func (o InstanceIAMBindingOutput) ToInstanceIAMBindingOutput() InstanceIAMBindingOutput {
	return o
}

func (o InstanceIAMBindingOutput) ToInstanceIAMBindingOutputWithContext(ctx context.Context) InstanceIAMBindingOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o InstanceIAMBindingOutput) Condition() InstanceIAMBindingConditionPtrOutput {
	return o.ApplyT(func(v *InstanceIAMBinding) InstanceIAMBindingConditionPtrOutput { return v.Condition }).(InstanceIAMBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o InstanceIAMBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIAMBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o InstanceIAMBindingOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIAMBinding) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

func (o InstanceIAMBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceIAMBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
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
func (o InstanceIAMBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIAMBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `compute.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o InstanceIAMBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIAMBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
// zone is specified, it is taken from the provider configuration.
func (o InstanceIAMBindingOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIAMBinding) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstanceIAMBindingArrayOutput struct{ *pulumi.OutputState }

func (InstanceIAMBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIAMBinding)(nil)).Elem()
}

func (o InstanceIAMBindingArrayOutput) ToInstanceIAMBindingArrayOutput() InstanceIAMBindingArrayOutput {
	return o
}

func (o InstanceIAMBindingArrayOutput) ToInstanceIAMBindingArrayOutputWithContext(ctx context.Context) InstanceIAMBindingArrayOutput {
	return o
}

func (o InstanceIAMBindingArrayOutput) Index(i pulumi.IntInput) InstanceIAMBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceIAMBinding {
		return vs[0].([]*InstanceIAMBinding)[vs[1].(int)]
	}).(InstanceIAMBindingOutput)
}

type InstanceIAMBindingMapOutput struct{ *pulumi.OutputState }

func (InstanceIAMBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIAMBinding)(nil)).Elem()
}

func (o InstanceIAMBindingMapOutput) ToInstanceIAMBindingMapOutput() InstanceIAMBindingMapOutput {
	return o
}

func (o InstanceIAMBindingMapOutput) ToInstanceIAMBindingMapOutputWithContext(ctx context.Context) InstanceIAMBindingMapOutput {
	return o
}

func (o InstanceIAMBindingMapOutput) MapIndex(k pulumi.StringInput) InstanceIAMBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceIAMBinding {
		return vs[0].(map[string]*InstanceIAMBinding)[vs[1].(string)]
	}).(InstanceIAMBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIAMBindingInput)(nil)).Elem(), &InstanceIAMBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIAMBindingArrayInput)(nil)).Elem(), InstanceIAMBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIAMBindingMapInput)(nil)).Elem(), InstanceIAMBindingMap{})
	pulumi.RegisterOutputType(InstanceIAMBindingOutput{})
	pulumi.RegisterOutputType(InstanceIAMBindingArrayOutput{})
	pulumi.RegisterOutputType(InstanceIAMBindingMapOutput{})
}