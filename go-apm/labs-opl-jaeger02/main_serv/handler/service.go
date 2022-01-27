package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/transport"
)

type ServRequest struct {
	WhoRequest string `json:"who_request"`
	Data       string `json:"data"`
}

type ServResponse struct {
	ServName string `json:"appname"`
	Message  string `json:"message"`
}

type MicroServResp struct {
	ServName string `json:"service_name"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

type MicroServRespErr struct {
	ServName string `json:"service_name"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

type Service struct {
	ServiceName string
	Tracer      trace.Tracer
}

func NewService(serviceName string, tracer trace.Tracer) *Service {
	return &Service{ServiceName: serviceName, Tracer: tracer}
}

func (serv *Service) RequestToOtherService(ctx context.Context, reqBody ServRequest) (*ServResponse, error) {
	_, span := serv.Tracer.Start(ctx, "handle service RequestToOtherService")
	// ctx, span := otel.Tracer("repository").Start(ctx, "RequestToOtherService")
	defer span.End()

	endpoint := "http://localhost:8021/serv-01"
	reqBody.WhoRequest = serv.ServiceName

	cli := gentleman.New()
	cli.Use(transport.Set(otelhttp.NewTransport(http.DefaultTransport)))
	cli.URL(endpoint)

	// cli.Context = ctx

	req := cli.Request()

	req.Method("POST")
	req.Use(body.JSON(reqBody))

	res, err := req.Send()
	if err != nil {
		message := fmt.Sprintf("request to %v not success , error :%v", endpoint, err.Error())
		span.RecordError(err)
		span.SetStatus(codes.Error, message)
		return nil, err
	}

	if !res.Ok {
		response := &MicroServRespErr{}
		err = res.JSON(response)
		if err != nil {
			message := fmt.Sprintf("mapping error response not success ,error : %v", err.Error())
			span.RecordError(err)
			span.SetStatus(codes.Error, message)
			return nil, err
		}

		message := fmt.Sprintf("microservice error %v Status: %v message : %v ", response.ServName, response.Status, response.ServName)
		span.RecordError(err)
		span.SetStatus(codes.Error, message)
		return nil, fmt.Errorf("microservice error %v Status: %v message : %v ", response.ServName, response.Status, response.ServName)
	}

	response := &MicroServResp{}
	err = res.JSON(response)
	if err != nil {
		message := fmt.Sprintf("mapping success response not success ,error : %v", err.Error())
		span.RecordError(err)
		span.SetStatus(codes.Error, message)
		return nil, err
	}

	return &ServResponse{ServName: serv.ServiceName, Message: response.Message}, nil
}

func (serv *Service) RequestToOtherServiceNative(ctx context.Context, reqBody ServRequest) (*ServResponse, error) {
	_, span := serv.Tracer.Start(ctx, "handle service RequestToOtherService")
	// ctx, span := otel.Tracer("repository").Start(ctx, "RequestToOtherService")
	defer span.End()

	endpoint := "http://localhost:8021/serv-01"
	reqBody.WhoRequest = serv.ServiceName

	var httpClient = http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(reqBody)

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		message := fmt.Sprintf("request to %v not success , error :%v", endpoint, err.Error())
		span.RecordError(err)
		span.SetStatus(codes.Error, message)
		return nil, err
	}

	if resp.StatusCode != 200 {
		response := &MicroServRespErr{}
		err = json.Unmarshal(body, response)
		if err != nil {
			message := fmt.Sprintf("mapping error response not success ,error : %v", err.Error())
			span.RecordError(err)
			span.SetStatus(codes.Error, message)
			return nil, err
		}

		message := fmt.Sprintf("microservice error %v Status: %v message : %v ", response.ServName, resp.StatusCode, response.ServName)
		span.RecordError(err)
		span.SetStatus(codes.Error, message)
		return nil, fmt.Errorf("microservice error %v Status: %v message : %v ", response.ServName, resp.StatusCode, response.ServName)
	}

	response := &MicroServResp{}
	err = json.Unmarshal(body, response)
	if err != nil {
		message := fmt.Sprintf("mapping success response not success ,error : %v", err.Error())
		span.RecordError(err)
		span.SetStatus(codes.Error, message)
		return nil, err
	}

	return &ServResponse{ServName: serv.ServiceName, Message: response.Message}, nil
}
