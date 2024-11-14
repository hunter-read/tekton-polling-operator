package pipelines

import (
	"context"

	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelineRunner executes a pipeline by name, creating a PipelineRun with the
// correct params and resource bindings.
type PipelineRunner interface {
	Run(ctx context.Context, pipelineName, ns, serviceAccountName string, params []pipelinev1.Param, res []pipelinev1.PipelineResourceBinding, ws []pipelinev1.WorkspaceBinding, to *metav1.Duration) (*pipelinev1.PipelineRun, error)
}
