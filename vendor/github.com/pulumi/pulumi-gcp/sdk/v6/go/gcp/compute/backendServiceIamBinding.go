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

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/global/backendServices/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine backendservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/backendServiceIamBinding:BackendServiceIamBinding editor "projects/{{project}}/global/backendServices/{{backend_service}} roles/compute.admin user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/backendServiceIamBinding:BackendServiceIamBinding editor "projects/{{project}}/global/backendServices/{{backend_service}} roles/compute.admin"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/backendServiceIamBinding:BackendServiceIamBinding editor projects/{{project}}/global/backendServices/{{backend_service}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type BackendServiceIamBinding struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BackendServiceIamBindingConditionPtrOutput `pulumi:"condition"`
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
	// The role that should be applied. Only one
	// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewBackendServiceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewBackendServiceIamBinding(ctx *pulumi.Context,
	name string, args *BackendServiceIamBindingArgs, opts ...pulumi.ResourceOption) (*BackendServiceIamBinding, error) {
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
	var resource BackendServiceIamBinding
	err := ctx.RegisterResource("gcp:compute/backendServiceIamBinding:BackendServiceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendServiceIamBinding gets an existing BackendServiceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendServiceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendServiceIamBindingState, opts ...pulumi.ResourceOption) (*BackendServiceIamBinding, error) {
	var resource BackendServiceIamBinding
	err := ctx.ReadResource("gcp:compute/backendServiceIamBinding:BackendServiceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendServiceIamBinding resources.
type backendServiceIamBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *BackendServiceIamBindingCondition `pulumi:"condition"`
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
	// The role that should be applied. Only one
	// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type BackendServiceIamBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BackendServiceIamBindingConditionPtrInput
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
	// The role that should be applied. Only one
	// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (BackendServiceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceIamBindingState)(nil)).Elem()
}

type backendServiceIamBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *BackendServiceIamBindingCondition `pulumi:"condition"`
	Members   []string                           `pulumi:"members"`
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
	// The role that should be applied. Only one
	// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a BackendServiceIamBinding resource.
type BackendServiceIamBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BackendServiceIamBindingConditionPtrInput
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
	// The role that should be applied. Only one
	// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (BackendServiceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceIamBindingArgs)(nil)).Elem()
}

type BackendServiceIamBindingInput interface {
	pulumi.Input

	ToBackendServiceIamBindingOutput() BackendServiceIamBindingOutput
	ToBackendServiceIamBindingOutputWithContext(ctx context.Context) BackendServiceIamBindingOutput
}

func (*BackendServiceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceIamBinding)(nil)).Elem()
}

func (i *BackendServiceIamBinding) ToBackendServiceIamBindingOutput() BackendServiceIamBindingOutput {
	return i.ToBackendServiceIamBindingOutputWithContext(context.Background())
}

func (i *BackendServiceIamBinding) ToBackendServiceIamBindingOutputWithContext(ctx context.Context) BackendServiceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceIamBindingOutput)
}

// BackendServiceIamBindingArrayInput is an input type that accepts BackendServiceIamBindingArray and BackendServiceIamBindingArrayOutput values.
// You can construct a concrete instance of `BackendServiceIamBindingArrayInput` via:
//
//	BackendServiceIamBindingArray{ BackendServiceIamBindingArgs{...} }
type BackendServiceIamBindingArrayInput interface {
	pulumi.Input

	ToBackendServiceIamBindingArrayOutput() BackendServiceIamBindingArrayOutput
	ToBackendServiceIamBindingArrayOutputWithContext(context.Context) BackendServiceIamBindingArrayOutput
}

type BackendServiceIamBindingArray []BackendServiceIamBindingInput

func (BackendServiceIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendServiceIamBinding)(nil)).Elem()
}

func (i BackendServiceIamBindingArray) ToBackendServiceIamBindingArrayOutput() BackendServiceIamBindingArrayOutput {
	return i.ToBackendServiceIamBindingArrayOutputWithContext(context.Background())
}

func (i BackendServiceIamBindingArray) ToBackendServiceIamBindingArrayOutputWithContext(ctx context.Context) BackendServiceIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceIamBindingArrayOutput)
}

// BackendServiceIamBindingMapInput is an input type that accepts BackendServiceIamBindingMap and BackendServiceIamBindingMapOutput values.
// You can construct a concrete instance of `BackendServiceIamBindingMapInput` via:
//
//	BackendServiceIamBindingMap{ "key": BackendServiceIamBindingArgs{...} }
type BackendServiceIamBindingMapInput interface {
	pulumi.Input

	ToBackendServiceIamBindingMapOutput() BackendServiceIamBindingMapOutput
	ToBackendServiceIamBindingMapOutputWithContext(context.Context) BackendServiceIamBindingMapOutput
}

type BackendServiceIamBindingMap map[string]BackendServiceIamBindingInput

func (BackendServiceIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendServiceIamBinding)(nil)).Elem()
}

func (i BackendServiceIamBindingMap) ToBackendServiceIamBindingMapOutput() BackendServiceIamBindingMapOutput {
	return i.ToBackendServiceIamBindingMapOutputWithContext(context.Background())
}

func (i BackendServiceIamBindingMap) ToBackendServiceIamBindingMapOutputWithContext(ctx context.Context) BackendServiceIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceIamBindingMapOutput)
}

type BackendServiceIamBindingOutput struct{ *pulumi.OutputState }

func (BackendServiceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceIamBinding)(nil)).Elem()
}

func (o BackendServiceIamBindingOutput) ToBackendServiceIamBindingOutput() BackendServiceIamBindingOutput {
	return o
}

func (o BackendServiceIamBindingOutput) ToBackendServiceIamBindingOutputWithContext(ctx context.Context) BackendServiceIamBindingOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o BackendServiceIamBindingOutput) Condition() BackendServiceIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *BackendServiceIamBinding) BackendServiceIamBindingConditionPtrOutput { return v.Condition }).(BackendServiceIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o BackendServiceIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BackendServiceIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendServiceIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o BackendServiceIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
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
func (o BackendServiceIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o BackendServiceIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type BackendServiceIamBindingArrayOutput struct{ *pulumi.OutputState }

func (BackendServiceIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendServiceIamBinding)(nil)).Elem()
}

func (o BackendServiceIamBindingArrayOutput) ToBackendServiceIamBindingArrayOutput() BackendServiceIamBindingArrayOutput {
	return o
}

func (o BackendServiceIamBindingArrayOutput) ToBackendServiceIamBindingArrayOutputWithContext(ctx context.Context) BackendServiceIamBindingArrayOutput {
	return o
}

func (o BackendServiceIamBindingArrayOutput) Index(i pulumi.IntInput) BackendServiceIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendServiceIamBinding {
		return vs[0].([]*BackendServiceIamBinding)[vs[1].(int)]
	}).(BackendServiceIamBindingOutput)
}

type BackendServiceIamBindingMapOutput struct{ *pulumi.OutputState }

func (BackendServiceIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendServiceIamBinding)(nil)).Elem()
}

func (o BackendServiceIamBindingMapOutput) ToBackendServiceIamBindingMapOutput() BackendServiceIamBindingMapOutput {
	return o
}

func (o BackendServiceIamBindingMapOutput) ToBackendServiceIamBindingMapOutputWithContext(ctx context.Context) BackendServiceIamBindingMapOutput {
	return o
}

func (o BackendServiceIamBindingMapOutput) MapIndex(k pulumi.StringInput) BackendServiceIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendServiceIamBinding {
		return vs[0].(map[string]*BackendServiceIamBinding)[vs[1].(string)]
	}).(BackendServiceIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceIamBindingInput)(nil)).Elem(), &BackendServiceIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceIamBindingArrayInput)(nil)).Elem(), BackendServiceIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceIamBindingMapInput)(nil)).Elem(), BackendServiceIamBindingMap{})
	pulumi.RegisterOutputType(BackendServiceIamBindingOutput{})
	pulumi.RegisterOutputType(BackendServiceIamBindingArrayOutput{})
	pulumi.RegisterOutputType(BackendServiceIamBindingMapOutput{})
}