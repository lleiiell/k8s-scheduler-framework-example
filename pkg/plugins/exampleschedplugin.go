package plugins

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)
import "context"

type ExampleSchedPlugin struct {
}

func (e *ExampleSchedPlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("filter pod: %v, node: %v", pod.Name, nodeInfo)
	return framework.NewStatus(framework.Success, "")
}

func (e *ExampleSchedPlugin) PreFilter(ctx context.Context, state *framework.CycleState, p *v1.Pod) *framework.Status {
	klog.V(3).Infof("prefilter pod: %v", p.Name)

	return framework.NewStatus(framework.Success, "")
}

func (e *ExampleSchedPlugin) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}

func (e *ExampleSchedPlugin) PreBind(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) *framework.Status {
	if p == nil {
		return framework.NewStatus(framework.Error, "p cannot be nil")
	}
	if p.Namespace != "default" {
		return framework.NewStatus(framework.Unschedulable, "only pods from 'default' namespace are allowed")
	}

	if v, e := state.Read(framework.StateKey(p.Name)); e == nil {
		klog.V(3).Infof("PreBind state: %+v", v)

		if value, ok := v.(*exampleStateData); ok && value.data == "never bind" {
			return framework.NewStatus(framework.Unschedulable, "pod is not permitted")
		}
	}
	return nil
}

func (e *ExampleSchedPlugin) Reserve(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) *framework.Status {
	if p == nil {
		return framework.NewStatus(framework.Error, "pod cannot be nil")
	}
	if p.Name == "my-test-pod" {
		state.Write(framework.StateKey(p.Name), &exampleStateData{data: "never bind"})
	}
	return nil
}

func (e *ExampleSchedPlugin) Unreserve(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) {
	if p.Name == "my-test-pod" {
		// The pod is at the end of its lifecycle -- let's clean up the allocated
		// resources. In this case, our clean up is simply deleting the key written
		// in the Reserve operation.
		state.Delete(framework.StateKey(p.Name))
	}
}

func (e *ExampleSchedPlugin) Name() string {
	return ExampleSchedName
}

var _ framework.ReservePlugin = &ExampleSchedPlugin{}
var _ framework.PreBindPlugin = &ExampleSchedPlugin{}
var _ framework.PreFilterPlugin = &ExampleSchedPlugin{}
var _ framework.FilterPlugin = &ExampleSchedPlugin{}

const ExampleSchedName = "example-sched-plugin"

type exampleStateData struct {
	data string
}

func (s *exampleStateData) Clone() framework.StateData {
	copy := &exampleStateData{
		data: s.data,
	}
	return copy
}

// NewExampleSchedPlugin initializes a new plugin and returns it.
func NewExampleSchedPlugin(_ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &ExampleSchedPlugin{}, nil
}
