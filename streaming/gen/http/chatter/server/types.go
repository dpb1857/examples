// Code generated by goa v3.1.1, DO NOT EDIT.
//
// chatter HTTP server types
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o
// $(GOPATH)/src/goa.design/examples/streaming

package server

import (
	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
)

// ChatSummaryResponseCollection is the type of the "chatter" service "summary"
// endpoint HTTP response body.
type ChatSummaryResponseCollection []*ChatSummaryResponse

// SubscribeResponseBody is the type of the "chatter" service "subscribe"
// endpoint HTTP response body.
type SubscribeResponseBody struct {
	// Message sent to the server
	Message string `form:"message" json:"message" xml:"message"`
	Action  string `form:"action" json:"action" xml:"action"`
	// Time at which the message was added
	AddedAt string `form:"added_at" json:"added_at" xml:"added_at"`
}

// HistoryResponseBodyTiny is the type of the "chatter" service "history"
// endpoint HTTP response body.
type HistoryResponseBodyTiny struct {
	// Message sent to the server
	Message string `form:"message" json:"message" xml:"message"`
}

// HistoryResponseBody is the type of the "chatter" service "history" endpoint
// HTTP response body.
type HistoryResponseBody struct {
	// Message sent to the server
	Message string `form:"message" json:"message" xml:"message"`
	// Length of the message sent
	Length *int `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Time at which the message was sent
	SentAt string `form:"sent_at" json:"sent_at" xml:"sent_at"`
}

// LoginUnauthorizedResponseBody is the type of the "chatter" service "login"
// endpoint HTTP response body for the "unauthorized" error.
type LoginUnauthorizedResponseBody string

// EchoerInvalidScopesResponseBody is the type of the "chatter" service
// "echoer" endpoint HTTP response body for the "invalid-scopes" error.
type EchoerInvalidScopesResponseBody string

// EchoerUnauthorizedResponseBody is the type of the "chatter" service "echoer"
// endpoint HTTP response body for the "unauthorized" error.
type EchoerUnauthorizedResponseBody string

// ListenerInvalidScopesResponseBody is the type of the "chatter" service
// "listener" endpoint HTTP response body for the "invalid-scopes" error.
type ListenerInvalidScopesResponseBody string

// ListenerUnauthorizedResponseBody is the type of the "chatter" service
// "listener" endpoint HTTP response body for the "unauthorized" error.
type ListenerUnauthorizedResponseBody string

// SummaryInvalidScopesResponseBody is the type of the "chatter" service
// "summary" endpoint HTTP response body for the "invalid-scopes" error.
type SummaryInvalidScopesResponseBody string

// SummaryUnauthorizedResponseBody is the type of the "chatter" service
// "summary" endpoint HTTP response body for the "unauthorized" error.
type SummaryUnauthorizedResponseBody string

// SubscribeInvalidScopesResponseBody is the type of the "chatter" service
// "subscribe" endpoint HTTP response body for the "invalid-scopes" error.
type SubscribeInvalidScopesResponseBody string

// SubscribeUnauthorizedResponseBody is the type of the "chatter" service
// "subscribe" endpoint HTTP response body for the "unauthorized" error.
type SubscribeUnauthorizedResponseBody string

// HistoryInvalidScopesResponseBody is the type of the "chatter" service
// "history" endpoint HTTP response body for the "invalid-scopes" error.
type HistoryInvalidScopesResponseBody string

// HistoryUnauthorizedResponseBody is the type of the "chatter" service
// "history" endpoint HTTP response body for the "unauthorized" error.
type HistoryUnauthorizedResponseBody string

// ChatSummaryResponse is used to define fields on response body types.
type ChatSummaryResponse struct {
	// Message sent to the server
	Message string `form:"message" json:"message" xml:"message"`
	// Length of the message sent
	Length *int `form:"length,omitempty" json:"length,omitempty" xml:"length,omitempty"`
	// Time at which the message was sent
	SentAt string `form:"sent_at" json:"sent_at" xml:"sent_at"`
}

// NewChatSummaryResponseCollection builds the HTTP response body from the
// result of the "summary" endpoint of the "chatter" service.
func NewChatSummaryResponseCollection(res chatterviews.ChatSummaryCollectionView) ChatSummaryResponseCollection {
	body := make([]*ChatSummaryResponse, len(res))
	for i, val := range res {
		body[i] = marshalChatterviewsChatSummaryViewToChatSummaryResponse(val)
	}
	return body
}

// NewSubscribeResponseBody builds the HTTP response body from the result of
// the "subscribe" endpoint of the "chatter" service.
func NewSubscribeResponseBody(res *chatter.Event) *SubscribeResponseBody {
	body := &SubscribeResponseBody{
		Message: res.Message,
		Action:  res.Action,
		AddedAt: res.AddedAt,
	}
	return body
}

// NewHistoryResponseBodyTiny builds the HTTP response body from the result of
// the "history" endpoint of the "chatter" service.
func NewHistoryResponseBodyTiny(res *chatterviews.ChatSummaryView) *HistoryResponseBodyTiny {
	body := &HistoryResponseBodyTiny{
		Message: *res.Message,
	}
	return body
}

// NewHistoryResponseBody builds the HTTP response body from the result of the
// "history" endpoint of the "chatter" service.
func NewHistoryResponseBody(res *chatterviews.ChatSummaryView) *HistoryResponseBody {
	body := &HistoryResponseBody{
		Message: *res.Message,
		Length:  res.Length,
		SentAt:  *res.SentAt,
	}
	return body
}

// NewLoginUnauthorizedResponseBody builds the HTTP response body from the
// result of the "login" endpoint of the "chatter" service.
func NewLoginUnauthorizedResponseBody(res chatter.Unauthorized) LoginUnauthorizedResponseBody {
	body := LoginUnauthorizedResponseBody(res)
	return body
}

// NewEchoerInvalidScopesResponseBody builds the HTTP response body from the
// result of the "echoer" endpoint of the "chatter" service.
func NewEchoerInvalidScopesResponseBody(res chatter.InvalidScopes) EchoerInvalidScopesResponseBody {
	body := EchoerInvalidScopesResponseBody(res)
	return body
}

// NewEchoerUnauthorizedResponseBody builds the HTTP response body from the
// result of the "echoer" endpoint of the "chatter" service.
func NewEchoerUnauthorizedResponseBody(res chatter.Unauthorized) EchoerUnauthorizedResponseBody {
	body := EchoerUnauthorizedResponseBody(res)
	return body
}

// NewListenerInvalidScopesResponseBody builds the HTTP response body from the
// result of the "listener" endpoint of the "chatter" service.
func NewListenerInvalidScopesResponseBody(res chatter.InvalidScopes) ListenerInvalidScopesResponseBody {
	body := ListenerInvalidScopesResponseBody(res)
	return body
}

// NewListenerUnauthorizedResponseBody builds the HTTP response body from the
// result of the "listener" endpoint of the "chatter" service.
func NewListenerUnauthorizedResponseBody(res chatter.Unauthorized) ListenerUnauthorizedResponseBody {
	body := ListenerUnauthorizedResponseBody(res)
	return body
}

// NewSummaryInvalidScopesResponseBody builds the HTTP response body from the
// result of the "summary" endpoint of the "chatter" service.
func NewSummaryInvalidScopesResponseBody(res chatter.InvalidScopes) SummaryInvalidScopesResponseBody {
	body := SummaryInvalidScopesResponseBody(res)
	return body
}

// NewSummaryUnauthorizedResponseBody builds the HTTP response body from the
// result of the "summary" endpoint of the "chatter" service.
func NewSummaryUnauthorizedResponseBody(res chatter.Unauthorized) SummaryUnauthorizedResponseBody {
	body := SummaryUnauthorizedResponseBody(res)
	return body
}

// NewSubscribeInvalidScopesResponseBody builds the HTTP response body from the
// result of the "subscribe" endpoint of the "chatter" service.
func NewSubscribeInvalidScopesResponseBody(res chatter.InvalidScopes) SubscribeInvalidScopesResponseBody {
	body := SubscribeInvalidScopesResponseBody(res)
	return body
}

// NewSubscribeUnauthorizedResponseBody builds the HTTP response body from the
// result of the "subscribe" endpoint of the "chatter" service.
func NewSubscribeUnauthorizedResponseBody(res chatter.Unauthorized) SubscribeUnauthorizedResponseBody {
	body := SubscribeUnauthorizedResponseBody(res)
	return body
}

// NewHistoryInvalidScopesResponseBody builds the HTTP response body from the
// result of the "history" endpoint of the "chatter" service.
func NewHistoryInvalidScopesResponseBody(res chatter.InvalidScopes) HistoryInvalidScopesResponseBody {
	body := HistoryInvalidScopesResponseBody(res)
	return body
}

// NewHistoryUnauthorizedResponseBody builds the HTTP response body from the
// result of the "history" endpoint of the "chatter" service.
func NewHistoryUnauthorizedResponseBody(res chatter.Unauthorized) HistoryUnauthorizedResponseBody {
	body := HistoryUnauthorizedResponseBody(res)
	return body
}

// NewLoginPayload builds a chatter service login endpoint payload.
func NewLoginPayload() *chatter.LoginPayload {
	v := &chatter.LoginPayload{}

	return v
}

// NewEchoerPayload builds a chatter service echoer endpoint payload.
func NewEchoerPayload(token string) *chatter.EchoerPayload {
	v := &chatter.EchoerPayload{}
	v.Token = token

	return v
}

// NewListenerPayload builds a chatter service listener endpoint payload.
func NewListenerPayload(token string) *chatter.ListenerPayload {
	v := &chatter.ListenerPayload{}
	v.Token = token

	return v
}

// NewSummaryPayload builds a chatter service summary endpoint payload.
func NewSummaryPayload(token string) *chatter.SummaryPayload {
	v := &chatter.SummaryPayload{}
	v.Token = token

	return v
}

// NewSubscribePayload builds a chatter service subscribe endpoint payload.
func NewSubscribePayload(token string) *chatter.SubscribePayload {
	v := &chatter.SubscribePayload{}
	v.Token = token

	return v
}

// NewHistoryPayload builds a chatter service history endpoint payload.
func NewHistoryPayload(view *string, token string) *chatter.HistoryPayload {
	v := &chatter.HistoryPayload{}
	v.View = view
	v.Token = token

	return v
}
