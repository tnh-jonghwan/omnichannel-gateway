package router

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/nats-io/nats.go"
)

// DTOs (Simplified for example)
type SendMessageRequestDto struct {
	ChannelType string `json:"channelType"`
	// Add other fields as needed
}

type BaseChatRequestDto struct {
	ChannelType string `json:"channelType"`
	// Add other fields as needed
}

type NatsRouter struct {
	nc *nats.Conn
	// Add services here, e.g.:
	// respondIoService *services.RespondIoService
}

func NewNatsRouter(nc *nats.Conn) *NatsRouter {
	return &NatsRouter{
		nc: nc,
	}
}

func (r *NatsRouter) RegisterSubscriptions() {
	r.subscribe("omnichannel.starfruit.req.send.org.*", r.handleSendMessage)
	r.subscribe("omnichannel.starfruit.req.open.org.*", r.handleOpenChat)
	r.subscribe("omnichannel.starfruit.req.close.org.*", r.handleCloseChat)
	r.subscribe("omnichannel.starfruit.req.channels.org.*", r.handleListSpaceChannels)
}

func (r *NatsRouter) subscribe(subject string, handler func(*nats.Msg)) {
	_, err := r.nc.Subscribe(subject, handler)
	if err != nil {
		log.Printf("Failed to subscribe to %s: %v", subject, err)
	} else {
		log.Printf("Subscribed to %s", subject)
	}
}

func (r *NatsRouter) handleSendMessage(msg *nats.Msg) {
	var req SendMessageRequestDto
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		r.replyError(msg, err)
		return
	}

	orgId := extractOrgId(msg.Subject)
	log.Printf("Handling SendMessage for Org: %s, Data: %+v", orgId, req)

	// logic to call service...
	// res := r.someService.SendMessage(orgId, req)

	// Mock response
	res := map[string]interface{}{"success": true}
	r.replyJSON(msg, res)
}

func (r *NatsRouter) handleOpenChat(msg *nats.Msg) {
	var req BaseChatRequestDto
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		r.replyError(msg, err)
		return
	}

	orgId := extractOrgId(msg.Subject)
	log.Printf("Handling OpenChat for Org: %s", orgId)

	// Mock response
	res := map[string]interface{}{"success": true}
	r.replyJSON(msg, res)
}

func (r *NatsRouter) handleCloseChat(msg *nats.Msg) {
	var req BaseChatRequestDto
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		r.replyError(msg, err)
		return
	}

	orgId := extractOrgId(msg.Subject)
	log.Printf("Handling CloseChat for Org: %s", orgId)

	// Mock response
	res := map[string]interface{}{"success": true}
	r.replyJSON(msg, res)
}

func (r *NatsRouter) handleListSpaceChannels(msg *nats.Msg) {
	orgId := extractOrgId(msg.Subject)
	log.Printf("Handling ListSpaceChannels for Org: %s", orgId)

	// Mock response
	res := map[string]interface{}{"items": []string{}}
	r.replyJSON(msg, res)
}

func (r *NatsRouter) replyJSON(msg *nats.Msg, data interface{}) {
	bytes, _ := json.Marshal(data)
	msg.Respond(bytes)
}

func (r *NatsRouter) replyError(msg *nats.Msg, err error) {
	errResp := map[string]string{"error": err.Error()}
	bytes, _ := json.Marshal(errResp)
	msg.Respond(bytes)
}

func extractOrgId(subject string) string {
	parts := strings.Split(subject, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}
