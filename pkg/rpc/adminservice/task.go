package adminservice

import (
	"context"
	"time"

	"github.com/lyft/flyteadmin/pkg/audit"

	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/lyft/flytestdlib/logger"

	"github.com/lyft/flyteadmin/pkg/rpc/adminservice/util"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *AdminService) CreateTask(
	ctx context.Context,
	request *admin.TaskCreateRequest) (*admin.TaskCreateResponse, error) {
	defer m.interceptPanic(ctx, request)
	requestedAt := time.Now()
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Incorrect request, nil requests not allowed")
	}
	var response *admin.TaskCreateResponse
	var err error
	m.Metrics.taskEndpointMetrics.create.Time(func() {
		response, err = m.TaskManager.CreateTask(ctx, *request)
	})
	audit.NewLogBuilder().WithAuthenticatedCtx(ctx).WithRequest(
		"CreateTask",
		audit.ParametersFromIdentifier(request.Id),
		audit.ReadWrite,
		requestedAt,
	).WithResponse(time.Now(), err).Log(ctx)
	if err != nil {
		return nil, util.TransformAndRecordError(err, &m.Metrics.taskEndpointMetrics.create)
	}
	m.Metrics.taskEndpointMetrics.create.Success()
	return response, nil
}

func (m *AdminService) GetTask(ctx context.Context, request *admin.ObjectGetRequest) (*admin.Task, error) {
	defer m.interceptPanic(ctx, request)
	requestedAt := time.Now()
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Incorrect request, nil requests not allowed")
	}
	// NOTE: When the Get HTTP endpoint is called the resource type is implicit (from the URL) so we must add it
	// to the request.
	if request.Id != nil && request.Id.ResourceType == core.ResourceType_UNSPECIFIED {
		logger.Info(ctx, "Adding resource type for unspecified value in request: [%+v]", request)
		request.Id.ResourceType = core.ResourceType_TASK
	}
	var response *admin.Task
	var err error
	m.Metrics.taskEndpointMetrics.get.Time(func() {
		response, err = m.TaskManager.GetTask(ctx, *request)
	})
	audit.NewLogBuilder().WithAuthenticatedCtx(ctx).WithRequest(
		"GetTask",
		audit.ParametersFromIdentifier(request.Id),
		audit.ReadOnly,
		requestedAt,
	).WithResponse(time.Now(), err).Log(ctx)
	if err != nil {
		return nil, util.TransformAndRecordError(err, &m.Metrics.taskEndpointMetrics.get)
	}
	m.Metrics.taskEndpointMetrics.get.Success()
	return response, nil
}

func (m *AdminService) ListTaskIds(
	ctx context.Context, request *admin.NamedEntityIdentifierListRequest) (*admin.NamedEntityIdentifierList, error) {
	defer m.interceptPanic(ctx, request)
	requestedAt := time.Now()
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Incorrect request, nil requests not allowed")
	}
	var response *admin.NamedEntityIdentifierList
	var err error
	m.Metrics.taskEndpointMetrics.listIds.Time(func() {
		response, err = m.TaskManager.ListUniqueTaskIdentifiers(ctx, *request)
	})
	audit.NewLogBuilder().WithAuthenticatedCtx(ctx).WithRequest(
		"ListTaskIds",
		map[string]string{
			audit.Project: request.Project,
			audit.Domain:  request.Domain,
		},
		audit.ReadOnly,
		requestedAt,
	).WithResponse(time.Now(), err).Log(ctx)
	if err != nil {
		return nil, util.TransformAndRecordError(err, &m.Metrics.taskEndpointMetrics.listIds)
	}

	m.Metrics.taskEndpointMetrics.listIds.Success()
	return response, nil
}

func (m *AdminService) ListTasks(ctx context.Context, request *admin.ResourceListRequest) (*admin.TaskList, error) {
	defer m.interceptPanic(ctx, request)
	requestedAt := time.Now()
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Incorrect request, nil requests not allowed")
	}
	var response *admin.TaskList
	var err error
	m.Metrics.taskEndpointMetrics.list.Time(func() {
		response, err = m.TaskManager.ListTasks(ctx, *request)
	})
	audit.NewLogBuilder().WithAuthenticatedCtx(ctx).WithRequest(
		"ListTasks",
		audit.ParametersFromNamedEntityIdentifier(request.Id),
		audit.ReadOnly,
		requestedAt,
	).WithResponse(time.Now(), err).Log(ctx)
	if err != nil {
		return nil, util.TransformAndRecordError(err, &m.Metrics.taskEndpointMetrics.list)
	}

	m.Metrics.taskEndpointMetrics.list.Success()
	return response, nil
}
