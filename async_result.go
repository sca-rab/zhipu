package zhipu

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// AsyncResultService creates a new async result get service
type AsyncResultService struct {
	client *Client

	id string
}

// AsyncResultVideo is the video result of the AsyncResultService
type AsyncResultVideo struct {
	URL           string `json:"url"`
	CoverImageURL string `json:"cover_image_url"`
}

// AsyncResultResponse is the response of the AsyncResultService
type AsyncResultResponse struct {
	Model       string             `json:"model"`
	TaskStatus  string             `json:"task_status"`
	RequestID   string             `json:"request_id"`
	ID          string             `json:"id"`
	VideoResult []AsyncResultVideo `json:"video_result"`
}

// AsyncResultGLM4Response is the GLM4 response of the AsyncResultService
type AsyncResultGLM4Response struct {
	Model      string                 `json:"model"`
	Choices    []ChatCompletionChoice `json:"choices"`
	TaskStatus string                 `json:"task_status"`
	RequestID  string                 `json:"request_id"`
	ID         string                 `json:"id"`
	Usage      ChatCompletionUsage    `json:"usage"`
}

// NewAsyncResultService creates a new async result get service
func NewAsyncResultService(client *Client) *AsyncResultService {
	return &AsyncResultService{
		client: client,
	}
}

// SetID sets the id parameter
func (s *AsyncResultService) SetID(id string) *AsyncResultService {
	s.id = id
	return s
}

func (s *AsyncResultService) Do(ctx context.Context) (res AsyncResultResponse, err error) {
	var (
		resp     *resty.Response
		apiError APIErrorResponse
	)

	if resp, err = s.client.request(ctx).
		SetResult(&res).
		SetError(&apiError).
		Get("async-result/" + s.id); err != nil {
		return
	}

	if resp.IsError() {
		err = apiError
		return
	}

	return
}

func (s *AsyncResultService) DoGLM4(ctx context.Context) (res AsyncResultGLM4Response, err error) {
	var (
		resp     *resty.Response
		apiError APIErrorResponse
	)

	if resp, err = s.client.request(ctx).
		SetResult(&res).
		SetError(&apiError).
		Get("async-result/" + s.id); err != nil {
		return
	}

	if resp.IsError() {
		err = apiError
		return
	}

	return
}
