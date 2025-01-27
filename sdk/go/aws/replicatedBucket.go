// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReplicatedBucket struct {
	pulumi.ResourceState

	// Bucket to which data should be replicated.
	DestinationBucket s3.BucketOutput `pulumi:"destinationBucket"`
	// Bucket to which objects are written.
	SourceBucket s3.BucketOutput `pulumi:"sourceBucket"`
}

// NewReplicatedBucket registers a new resource with the given unique name, arguments, and options.
func NewReplicatedBucket(ctx *pulumi.Context,
	name string, args *ReplicatedBucketArgs, opts ...pulumi.ResourceOption) (*ReplicatedBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationRegion == nil {
		return nil, errors.New("invalid value for required argument 'DestinationRegion'")
	}
	var resource ReplicatedBucket
	err := ctx.RegisterRemoteComponentResource("aws-s3-replicated-bucket:index:ReplicatedBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type replicatedBucketArgs struct {
	// Region to which data should be replicated.
	DestinationRegion string `pulumi:"destinationRegion"`
}

// The set of arguments for constructing a ReplicatedBucket resource.
type ReplicatedBucketArgs struct {
	// Region to which data should be replicated.
	DestinationRegion pulumi.StringInput
}

func (ReplicatedBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicatedBucketArgs)(nil)).Elem()
}

type ReplicatedBucketInput interface {
	pulumi.Input

	ToReplicatedBucketOutput() ReplicatedBucketOutput
	ToReplicatedBucketOutputWithContext(ctx context.Context) ReplicatedBucketOutput
}

func (*ReplicatedBucket) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicatedBucket)(nil))
}

func (i *ReplicatedBucket) ToReplicatedBucketOutput() ReplicatedBucketOutput {
	return i.ToReplicatedBucketOutputWithContext(context.Background())
}

func (i *ReplicatedBucket) ToReplicatedBucketOutputWithContext(ctx context.Context) ReplicatedBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedBucketOutput)
}

func (i *ReplicatedBucket) ToReplicatedBucketPtrOutput() ReplicatedBucketPtrOutput {
	return i.ToReplicatedBucketPtrOutputWithContext(context.Background())
}

func (i *ReplicatedBucket) ToReplicatedBucketPtrOutputWithContext(ctx context.Context) ReplicatedBucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedBucketPtrOutput)
}

type ReplicatedBucketPtrInput interface {
	pulumi.Input

	ToReplicatedBucketPtrOutput() ReplicatedBucketPtrOutput
	ToReplicatedBucketPtrOutputWithContext(ctx context.Context) ReplicatedBucketPtrOutput
}

type replicatedBucketPtrType ReplicatedBucketArgs

func (*replicatedBucketPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicatedBucket)(nil))
}

func (i *replicatedBucketPtrType) ToReplicatedBucketPtrOutput() ReplicatedBucketPtrOutput {
	return i.ToReplicatedBucketPtrOutputWithContext(context.Background())
}

func (i *replicatedBucketPtrType) ToReplicatedBucketPtrOutputWithContext(ctx context.Context) ReplicatedBucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedBucketPtrOutput)
}

// ReplicatedBucketArrayInput is an input type that accepts ReplicatedBucketArray and ReplicatedBucketArrayOutput values.
// You can construct a concrete instance of `ReplicatedBucketArrayInput` via:
//
//          ReplicatedBucketArray{ ReplicatedBucketArgs{...} }
type ReplicatedBucketArrayInput interface {
	pulumi.Input

	ToReplicatedBucketArrayOutput() ReplicatedBucketArrayOutput
	ToReplicatedBucketArrayOutputWithContext(context.Context) ReplicatedBucketArrayOutput
}

type ReplicatedBucketArray []ReplicatedBucketInput

func (ReplicatedBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicatedBucket)(nil)).Elem()
}

func (i ReplicatedBucketArray) ToReplicatedBucketArrayOutput() ReplicatedBucketArrayOutput {
	return i.ToReplicatedBucketArrayOutputWithContext(context.Background())
}

func (i ReplicatedBucketArray) ToReplicatedBucketArrayOutputWithContext(ctx context.Context) ReplicatedBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedBucketArrayOutput)
}

// ReplicatedBucketMapInput is an input type that accepts ReplicatedBucketMap and ReplicatedBucketMapOutput values.
// You can construct a concrete instance of `ReplicatedBucketMapInput` via:
//
//          ReplicatedBucketMap{ "key": ReplicatedBucketArgs{...} }
type ReplicatedBucketMapInput interface {
	pulumi.Input

	ToReplicatedBucketMapOutput() ReplicatedBucketMapOutput
	ToReplicatedBucketMapOutputWithContext(context.Context) ReplicatedBucketMapOutput
}

type ReplicatedBucketMap map[string]ReplicatedBucketInput

func (ReplicatedBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicatedBucket)(nil)).Elem()
}

func (i ReplicatedBucketMap) ToReplicatedBucketMapOutput() ReplicatedBucketMapOutput {
	return i.ToReplicatedBucketMapOutputWithContext(context.Background())
}

func (i ReplicatedBucketMap) ToReplicatedBucketMapOutputWithContext(ctx context.Context) ReplicatedBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedBucketMapOutput)
}

type ReplicatedBucketOutput struct{ *pulumi.OutputState }

func (ReplicatedBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicatedBucket)(nil))
}

func (o ReplicatedBucketOutput) ToReplicatedBucketOutput() ReplicatedBucketOutput {
	return o
}

func (o ReplicatedBucketOutput) ToReplicatedBucketOutputWithContext(ctx context.Context) ReplicatedBucketOutput {
	return o
}

func (o ReplicatedBucketOutput) ToReplicatedBucketPtrOutput() ReplicatedBucketPtrOutput {
	return o.ToReplicatedBucketPtrOutputWithContext(context.Background())
}

func (o ReplicatedBucketOutput) ToReplicatedBucketPtrOutputWithContext(ctx context.Context) ReplicatedBucketPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicatedBucket) *ReplicatedBucket {
		return &v
	}).(ReplicatedBucketPtrOutput)
}

type ReplicatedBucketPtrOutput struct{ *pulumi.OutputState }

func (ReplicatedBucketPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicatedBucket)(nil))
}

func (o ReplicatedBucketPtrOutput) ToReplicatedBucketPtrOutput() ReplicatedBucketPtrOutput {
	return o
}

func (o ReplicatedBucketPtrOutput) ToReplicatedBucketPtrOutputWithContext(ctx context.Context) ReplicatedBucketPtrOutput {
	return o
}

func (o ReplicatedBucketPtrOutput) Elem() ReplicatedBucketOutput {
	return o.ApplyT(func(v *ReplicatedBucket) ReplicatedBucket {
		if v != nil {
			return *v
		}
		var ret ReplicatedBucket
		return ret
	}).(ReplicatedBucketOutput)
}

type ReplicatedBucketArrayOutput struct{ *pulumi.OutputState }

func (ReplicatedBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicatedBucket)(nil))
}

func (o ReplicatedBucketArrayOutput) ToReplicatedBucketArrayOutput() ReplicatedBucketArrayOutput {
	return o
}

func (o ReplicatedBucketArrayOutput) ToReplicatedBucketArrayOutputWithContext(ctx context.Context) ReplicatedBucketArrayOutput {
	return o
}

func (o ReplicatedBucketArrayOutput) Index(i pulumi.IntInput) ReplicatedBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicatedBucket {
		return vs[0].([]ReplicatedBucket)[vs[1].(int)]
	}).(ReplicatedBucketOutput)
}

type ReplicatedBucketMapOutput struct{ *pulumi.OutputState }

func (ReplicatedBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReplicatedBucket)(nil))
}

func (o ReplicatedBucketMapOutput) ToReplicatedBucketMapOutput() ReplicatedBucketMapOutput {
	return o
}

func (o ReplicatedBucketMapOutput) ToReplicatedBucketMapOutputWithContext(ctx context.Context) ReplicatedBucketMapOutput {
	return o
}

func (o ReplicatedBucketMapOutput) MapIndex(k pulumi.StringInput) ReplicatedBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReplicatedBucket {
		return vs[0].(map[string]ReplicatedBucket)[vs[1].(string)]
	}).(ReplicatedBucketOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicatedBucketOutput{})
	pulumi.RegisterOutputType(ReplicatedBucketPtrOutput{})
	pulumi.RegisterOutputType(ReplicatedBucketArrayOutput{})
	pulumi.RegisterOutputType(ReplicatedBucketMapOutput{})
}
