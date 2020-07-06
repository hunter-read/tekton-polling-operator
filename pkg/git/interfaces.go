package git

import (
	pollingv1alpha1 "github.com/bigkevmcd/tekton-polling-operator/pkg/apis/polling/v1alpha1"
)

// CommitPoller implementations can check with an upstream Git hosting service
// to determine the current SHA and ETag.
type CommitPoller interface {
	Poll(repo string, pr pollingv1alpha1.RepositoryStatus) (pollingv1alpha1.RepositoryStatus, error)
}