// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package replicatedbucket

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Bucket struct {
	pulumi.ResourceState

	// Bucket to which data should be replicated.
	DestinationBucket s3.BucketOutput `pulumi:"destinationBucket"`
	// Bucket to which objects are written.
	SourceBucket s3.BucketOutput `pulumi:"sourceBucket"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationRegion == nil {
		return nil, errors.New("invalid value for required argument 'DestinationRegion'")
	}
	var resource Bucket
	err := ctx.RegisterRemoteComponentResource("replicatedbucket:index:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type bucketArgs struct {
	// Region to which data should be replicated.
	DestinationRegion string `pulumi:"destinationRegion"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// Region to which data should be replicated.
	DestinationRegion pulumi.StringInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((*Bucket)(nil))
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

func (i *Bucket) ToBucketPtrOutput() BucketPtrOutput {
	return i.ToBucketPtrOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketPtrOutputWithContext(ctx context.Context) BucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPtrOutput)
}

type BucketPtrInput interface {
	pulumi.Input

	ToBucketPtrOutput() BucketPtrOutput
	ToBucketPtrOutputWithContext(ctx context.Context) BucketPtrOutput
}

type bucketPtrType BucketArgs

func (*bucketPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil))
}

func (i *bucketPtrType) ToBucketPtrOutput() BucketPtrOutput {
	return i.ToBucketPtrOutputWithContext(context.Background())
}

func (i *bucketPtrType) ToBucketPtrOutputWithContext(ctx context.Context) BucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPtrOutput)
}

// BucketArrayInput is an input type that accepts BucketArray and BucketArrayOutput values.
// You can construct a concrete instance of `BucketArrayInput` via:
//
//          BucketArray{ BucketArgs{...} }
type BucketArrayInput interface {
	pulumi.Input

	ToBucketArrayOutput() BucketArrayOutput
	ToBucketArrayOutputWithContext(context.Context) BucketArrayOutput
}

type BucketArray []BucketInput

func (BucketArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Bucket)(nil))
}

func (i BucketArray) ToBucketArrayOutput() BucketArrayOutput {
	return i.ToBucketArrayOutputWithContext(context.Background())
}

func (i BucketArray) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketArrayOutput)
}

// BucketMapInput is an input type that accepts BucketMap and BucketMapOutput values.
// You can construct a concrete instance of `BucketMapInput` via:
//
//          BucketMap{ "key": BucketArgs{...} }
type BucketMapInput interface {
	pulumi.Input

	ToBucketMapOutput() BucketMapOutput
	ToBucketMapOutputWithContext(context.Context) BucketMapOutput
}

type BucketMap map[string]BucketInput

func (BucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Bucket)(nil))
}

func (i BucketMap) ToBucketMapOutput() BucketMapOutput {
	return i.ToBucketMapOutputWithContext(context.Background())
}

func (i BucketMap) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMapOutput)
}

type BucketOutput struct {
	*pulumi.OutputState
}

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Bucket)(nil))
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

func (o BucketOutput) ToBucketPtrOutput() BucketPtrOutput {
	return o.ToBucketPtrOutputWithContext(context.Background())
}

func (o BucketOutput) ToBucketPtrOutputWithContext(ctx context.Context) BucketPtrOutput {
	return o.ApplyT(func(v Bucket) *Bucket {
		return &v
	}).(BucketPtrOutput)
}

type BucketPtrOutput struct {
	*pulumi.OutputState
}

func (BucketPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil))
}

func (o BucketPtrOutput) ToBucketPtrOutput() BucketPtrOutput {
	return o
}

func (o BucketPtrOutput) ToBucketPtrOutputWithContext(ctx context.Context) BucketPtrOutput {
	return o
}

type BucketArrayOutput struct{ *pulumi.OutputState }

func (BucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Bucket)(nil))
}

func (o BucketArrayOutput) ToBucketArrayOutput() BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) Index(i pulumi.IntInput) BucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Bucket {
		return vs[0].([]Bucket)[vs[1].(int)]
	}).(BucketOutput)
}

type BucketMapOutput struct{ *pulumi.OutputState }

func (BucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Bucket)(nil))
}

func (o BucketMapOutput) ToBucketMapOutput() BucketMapOutput {
	return o
}

func (o BucketMapOutput) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return o
}

func (o BucketMapOutput) MapIndex(k pulumi.StringInput) BucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Bucket {
		return vs[0].(map[string]Bucket)[vs[1].(string)]
	}).(BucketOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketOutput{})
	pulumi.RegisterOutputType(BucketPtrOutput{})
	pulumi.RegisterOutputType(BucketArrayOutput{})
	pulumi.RegisterOutputType(BucketMapOutput{})
}