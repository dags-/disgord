package disgord

// Code generated - This file has been automatically generated by generate/events/main.go - DO NOT EDIT.
// Warning: This file is overwritten at "go generate", instead adapt events.go and event/events.go and run go generate

import (
	"github.com/andersfylling/disgord/internal/event"
)

// helpers

func AllEvents() []string {
	return AllEventsExcept()
}

func AllEventsExcept(except ...string) []string {
	evtsMap := map[string]int8{

		EvtChannelCreate: 0,

		EvtChannelDelete: 0,

		EvtChannelPinsUpdate: 0,

		EvtChannelUpdate: 0,

		EvtGuildBanAdd: 0,

		EvtGuildBanRemove: 0,

		EvtGuildCreate: 0,

		EvtGuildDelete: 0,

		EvtGuildEmojisUpdate: 0,

		EvtGuildIntegrationsUpdate: 0,

		EvtGuildMemberAdd: 0,

		EvtGuildMemberRemove: 0,

		EvtGuildMemberUpdate: 0,

		EvtGuildMembersChunk: 0,

		EvtGuildRoleCreate: 0,

		EvtGuildRoleDelete: 0,

		EvtGuildRoleUpdate: 0,

		EvtGuildUpdate: 0,

		EvtInviteCreate: 0,

		EvtInviteDelete: 0,

		EvtMessageCreate: 0,

		EvtMessageDelete: 0,

		EvtMessageDeleteBulk: 0,

		EvtMessageReactionAdd: 0,

		EvtMessageReactionRemove: 0,

		EvtMessageReactionRemoveAll: 0,

		EvtMessageReactionRemoveEmoji: 0,

		EvtMessageUpdate: 0,

		EvtPresenceUpdate: 0,

		EvtReady: 0,

		EvtResumed: 0,

		EvtTypingStart: 0,

		EvtUserUpdate: 0,

		EvtVoiceServerUpdate: 0,

		EvtVoiceStateUpdate: 0,

		EvtWebhooksUpdate: 0,
	}

	for i := range except {
		delete(evtsMap, except[i])
	}

	evts := make([]string, 0, len(evtsMap))
	for k := range evtsMap {
		evts = append(evts, k)
	}
	return evts
}

// ---------------------------

// EvtChannelCreate Sent when a new channel is created, relevant to the current user. The inner payload is a DM channel or
// guild channel object.
//
const EvtChannelCreate = event.ChannelCreate

func (h *ChannelCreate) setShardID(id uint) { h.ShardID = id }

type HandlerChannelCreate = func(Session, *ChannelCreate)

// ---------------------------

// EvtChannelDelete Sent when a channel relevant to the current user is deleted. The inner payload is a DM or Guild channel object.
//
const EvtChannelDelete = event.ChannelDelete

func (h *ChannelDelete) setShardID(id uint) { h.ShardID = id }

type HandlerChannelDelete = func(Session, *ChannelDelete)

// ---------------------------

// EvtChannelPinsUpdate Sent when a message is pinned or unpinned in a text channel. This is not sent when a pinned message is deleted.
//  Fields:
//  - ChannelID int64 or Snowflake
//  - LastPinTimestamp time.Now().UTC().Format(time.RFC3339)
// TODO fix.
//
const EvtChannelPinsUpdate = event.ChannelPinsUpdate

func (h *ChannelPinsUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerChannelPinsUpdate = func(Session, *ChannelPinsUpdate)

// ---------------------------

// EvtChannelUpdate Sent when a channel is updated. The inner payload is a guild channel object.
//
const EvtChannelUpdate = event.ChannelUpdate

func (h *ChannelUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerChannelUpdate = func(Session, *ChannelUpdate)

// ---------------------------

// EvtGuildBanAdd Sent when a user is banned from a guild. The inner payload is a user object, with an extra guild_id key.
//
const EvtGuildBanAdd = event.GuildBanAdd

func (h *GuildBanAdd) setShardID(id uint) { h.ShardID = id }

type HandlerGuildBanAdd = func(Session, *GuildBanAdd)

// ---------------------------

// EvtGuildBanRemove Sent when a user is unbanned from a guild. The inner payload is a user object, with an extra guild_id key.
//
const EvtGuildBanRemove = event.GuildBanRemove

func (h *GuildBanRemove) setShardID(id uint) { h.ShardID = id }

type HandlerGuildBanRemove = func(Session, *GuildBanRemove)

// ---------------------------

// EvtGuildCreate This event can be sent in three different scenarios:
//  1. When a user is initially connecting, to lazily load and backfill information for all unavailable guilds
//     sent in the Ready event.
// 	2. When a Guild becomes available again to the client.
// 	3. When the current user joins a new Guild.
//
const EvtGuildCreate = event.GuildCreate

func (h *GuildCreate) setShardID(id uint) { h.ShardID = id }

type HandlerGuildCreate = func(Session, *GuildCreate)

// ---------------------------

// EvtGuildDelete Sent when a guild becomes unavailable during a guild outage, or when the user leaves or is removed from a guild.
// The inner payload is an unavailable guild object. If the unavailable field is not set, the user was removed
// from the guild.
//
const EvtGuildDelete = event.GuildDelete

func (h *GuildDelete) setShardID(id uint) { h.ShardID = id }

type HandlerGuildDelete = func(Session, *GuildDelete)

// ---------------------------

// EvtGuildEmojisUpdate Sent when a guild's emojis have been updated.
//  Fields:
//  - GuildID Snowflake
//  - Emojis []*Emoji
//
const EvtGuildEmojisUpdate = event.GuildEmojisUpdate

func (h *GuildEmojisUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerGuildEmojisUpdate = func(Session, *GuildEmojisUpdate)

// ---------------------------

// EvtGuildIntegrationsUpdate Sent when a guild integration is updated.
//  Fields:
//  - GuildID Snowflake
//
const EvtGuildIntegrationsUpdate = event.GuildIntegrationsUpdate

func (h *GuildIntegrationsUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerGuildIntegrationsUpdate = func(Session, *GuildIntegrationsUpdate)

// ---------------------------

// EvtGuildMemberAdd Sent when a new user joins a guild. The inner payload is a guild member object with these extra fields:
//  - GuildID Snowflake
//
//  Fields:
//  - Member *Member
//
const EvtGuildMemberAdd = event.GuildMemberAdd

func (h *GuildMemberAdd) setShardID(id uint) { h.ShardID = id }

type HandlerGuildMemberAdd = func(Session, *GuildMemberAdd)

// ---------------------------

// EvtGuildMemberRemove Sent when a user is removed from a guild (leave/kick/ban).
//  Fields:
//  - GuildID   Snowflake
//  - User      *User
//
const EvtGuildMemberRemove = event.GuildMemberRemove

func (h *GuildMemberRemove) setShardID(id uint) { h.ShardID = id }

type HandlerGuildMemberRemove = func(Session, *GuildMemberRemove)

// ---------------------------

// EvtGuildMemberUpdate Sent when a guild member is updated.
//  Fields:
//  - GuildID   Snowflake
//  - Roles     []Snowflake
//  - User      *User
//  - Nick      string
//
const EvtGuildMemberUpdate = event.GuildMemberUpdate

func (h *GuildMemberUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerGuildMemberUpdate = func(Session, *GuildMemberUpdate)

// ---------------------------

// EvtGuildMembersChunk Sent in response to Gateway Request Guild Members.
//  Fields:
//  - GuildID Snowflake
//  - Members []*Member
//
const EvtGuildMembersChunk = event.GuildMembersChunk

func (h *GuildMembersChunk) setShardID(id uint) { h.ShardID = id }

type HandlerGuildMembersChunk = func(Session, *GuildMembersChunk)

// ---------------------------

// EvtGuildRoleCreate Sent when a guild role is created.
//  Fields:
//  - GuildID   Snowflake
//  - Role      *Role
//
const EvtGuildRoleCreate = event.GuildRoleCreate

func (h *GuildRoleCreate) setShardID(id uint) { h.ShardID = id }

type HandlerGuildRoleCreate = func(Session, *GuildRoleCreate)

// ---------------------------

// EvtGuildRoleDelete Sent when a guild role is created.
//  Fields:
//  - GuildID Snowflake
//  - RoleID  Snowflake
//
const EvtGuildRoleDelete = event.GuildRoleDelete

func (h *GuildRoleDelete) setShardID(id uint) { h.ShardID = id }

type HandlerGuildRoleDelete = func(Session, *GuildRoleDelete)

// ---------------------------

// EvtGuildRoleUpdate Sent when a guild role is created.
//  Fields:
//  - GuildID Snowflake
//  - Role    *Role
//
const EvtGuildRoleUpdate = event.GuildRoleUpdate

func (h *GuildRoleUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerGuildRoleUpdate = func(Session, *GuildRoleUpdate)

// ---------------------------

// EvtGuildUpdate Sent when a guild is updated. The inner payload is a guild object.
//
const EvtGuildUpdate = event.GuildUpdate

func (h *GuildUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerGuildUpdate = func(Session, *GuildUpdate)

// ---------------------------

// EvtInviteCreate Sent when a guild's invite is created.
//  Fields:
//  - Code String
//  - GuildID   Snowflake
//  - ChannelID Snowflake
//  - Inviter *User
//  - Inviter *User
//  - Target *User
//  - TargetType int
//  - CreatedAt Time
//  - MaxAge int
//  - MaxUses int
//  - Temporary bool
//  - Uses int
//  - Revoked bool
//  - Unique bool
//  - ApproximatePresenceCount int
//  - ApproximateMemberCount int
//
const EvtInviteCreate = event.InviteCreate

func (h *InviteCreate) setShardID(id uint) { h.ShardID = id }

type HandlerInviteCreate = func(Session, *InviteCreate)

// ---------------------------

// EvtInviteDelete Sent when an invite is deleted.
//
const EvtInviteDelete = event.InviteDelete

func (h *InviteDelete) setShardID(id uint) { h.ShardID = id }

type HandlerInviteDelete = func(Session, *InviteDelete)

// ---------------------------

// EvtMessageCreate Sent when a message is created. The inner payload is a message object.
//
const EvtMessageCreate = event.MessageCreate

func (h *MessageCreate) setShardID(id uint) { h.ShardID = id }

type HandlerMessageCreate = func(Session, *MessageCreate)

// ---------------------------

// EvtMessageDelete Sent when a message is deleted.
//  Fields:
//  - ID        Snowflake
//  - ChannelID Snowflake
//
const EvtMessageDelete = event.MessageDelete

func (h *MessageDelete) setShardID(id uint) { h.ShardID = id }

type HandlerMessageDelete = func(Session, *MessageDelete)

// ---------------------------

// EvtMessageDeleteBulk Sent when multiple messages are deleted at once.
//  Fields:
//  - IDs       []Snowflake
//  - ChannelID Snowflake
//
const EvtMessageDeleteBulk = event.MessageDeleteBulk

func (h *MessageDeleteBulk) setShardID(id uint) { h.ShardID = id }

type HandlerMessageDeleteBulk = func(Session, *MessageDeleteBulk)

// ---------------------------

// EvtMessageReactionAdd Sent when a user adds a reaction to a message.
//  Fields:
//  - UserID     Snowflake
//  - ChannelID  Snowflake
//  - MessageID  Snowflake
//  - Emoji      *Emoji
//
const EvtMessageReactionAdd = event.MessageReactionAdd

func (h *MessageReactionAdd) setShardID(id uint) { h.ShardID = id }

type HandlerMessageReactionAdd = func(Session, *MessageReactionAdd)

// ---------------------------

// EvtMessageReactionRemove Sent when a user removes a reaction from a message.
//  Fields:
//  - UserID     Snowflake
//  - ChannelID  Snowflake
//  - MessageID  Snowflake
//  - Emoji      *Emoji
//
const EvtMessageReactionRemove = event.MessageReactionRemove

func (h *MessageReactionRemove) setShardID(id uint) { h.ShardID = id }

type HandlerMessageReactionRemove = func(Session, *MessageReactionRemove)

// ---------------------------

// EvtMessageReactionRemoveAll Sent when a user explicitly removes all reactions from a message.
//  Fields:
//  - ChannelID Snowflake
//  - MessageID Snowflake
//
const EvtMessageReactionRemoveAll = event.MessageReactionRemoveAll

func (h *MessageReactionRemoveAll) setShardID(id uint) { h.ShardID = id }

type HandlerMessageReactionRemoveAll = func(Session, *MessageReactionRemoveAll)

// ---------------------------

// EvtMessageReactionRemoveEmoji Sent when a bot removes all instances of a given emoji from the reactions of a message.
//
const EvtMessageReactionRemoveEmoji = event.MessageReactionRemoveEmoji

func (h *MessageReactionRemoveEmoji) setShardID(id uint) { h.ShardID = id }

type HandlerMessageReactionRemoveEmoji = func(Session, *MessageReactionRemoveEmoji)

// ---------------------------

// EvtMessageUpdate Sent when a message is updated. The inner payload is a message object.
//
// NOTE! Has _at_least_ the GuildID and ChannelID fields.
//
const EvtMessageUpdate = event.MessageUpdate

func (h *MessageUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerMessageUpdate = func(Session, *MessageUpdate)

// ---------------------------

// EvtPresenceUpdate A user's presence is their current state on a guild. This event is sent when a user's presence is updated for a guild.
//  Fields:
//  - User    *User
//  - Roles   []Snowflake
//  - Game    *Activity
//  - GuildID Snowflake
//  - Status  string
//
const EvtPresenceUpdate = event.PresenceUpdate

func (h *PresenceUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerPresenceUpdate = func(Session, *PresenceUpdate)

// ---------------------------

// EvtReady The ready event is dispatched when a client has completed the initial handshake with the gateway (for new sessions).
// // The ready event can be the largest and most complex event the gateway will send, as it contains all the state
// // required for a client to begin interacting with the rest of the platform.
// //  Fields:
// //  - V int
// //  - User *User
// //  - PrivateChannels []*Channel
// //  - Guilds []*GuildUnavailable
// //  - SessionID string
// //  - Trace []string
//
const EvtReady = event.Ready

func (h *Ready) setShardID(id uint) { h.ShardID = id }

type HandlerReady = func(Session, *Ready)

// ---------------------------

// EvtResumed The resumed event is dispatched when a client has sent a resume payload to the gateway
// (for resuming existing sessions).
//  Fields:
//  - Trace []string
//
const EvtResumed = event.Resumed

func (h *Resumed) setShardID(id uint) { h.ShardID = id }

type HandlerResumed = func(Session, *Resumed)

// ---------------------------

// EvtTypingStart Sent when a user starts typing in a channel.
//  Fields:
//  - ChannelID     Snowflake
//  - UserID        Snowflake
//  - TimestampUnix int
//
const EvtTypingStart = event.TypingStart

func (h *TypingStart) setShardID(id uint) { h.ShardID = id }

type HandlerTypingStart = func(Session, *TypingStart)

// ---------------------------

// EvtUserUpdate Sent when properties about the user change. Inner payload is a user object.
//
const EvtUserUpdate = event.UserUpdate

func (h *UserUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerUserUpdate = func(Session, *UserUpdate)

// ---------------------------

// EvtVoiceServerUpdate Sent when a guild's voice server is updated. This is sent when initially connecting to voice, and when the current
// voice instance fails over to a new server.
//  Fields:
//  - Token     string
//  - ChannelID Snowflake
//  - Endpoint  string
//
const EvtVoiceServerUpdate = event.VoiceServerUpdate

func (h *VoiceServerUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerVoiceServerUpdate = func(Session, *VoiceServerUpdate)

// ---------------------------

// EvtVoiceStateUpdate Sent when someone joins/leaves/moves voice channels. Inner payload is a voice state object.
//
const EvtVoiceStateUpdate = event.VoiceStateUpdate

func (h *VoiceStateUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerVoiceStateUpdate = func(Session, *VoiceStateUpdate)

// ---------------------------

// EvtWebhooksUpdate Sent when a guild channel's WebHook is created, updated, or deleted.
//  Fields:
//  - GuildID   Snowflake
//  - ChannelID Snowflake
//
const EvtWebhooksUpdate = event.WebhooksUpdate

func (h *WebhooksUpdate) setShardID(id uint) { h.ShardID = id }

type HandlerWebhooksUpdate = func(Session, *WebhooksUpdate)

// ---------------------------

type socketHandlerRegister struct {
	evtName     string
	middlewares []Middleware
	handlers    []Handler
	ctrl        HandlerCtrl

	register OnSocketEventer
}

func (shr *socketHandlerRegister) build() {
	inputs := make([]interface{}, 0, len(shr.middlewares)+len(shr.handlers)+1)
	for _, mdlw := range shr.middlewares {
		inputs = append(inputs, mdlw)
	}
	for _, handler := range shr.handlers {
		inputs = append(inputs, handler)
	}
	if shr.ctrl != nil {
		inputs = append(inputs, shr.ctrl)
	}

	shr.register.On(shr.evtName, inputs...)
}

func (shr *socketHandlerRegister) WithCtrl(ctrl HandlerCtrl) SocketHandlerRegistrator {
	if ctrl != nil {
		panic("a controller was already registered for this handler specification")
	}
	shr.ctrl = ctrl
	return shr
}

func (shr *socketHandlerRegister) WithMdlw(mdlws ...Middleware) SocketHandlerRegistrator {
	shr.middlewares = append(shr.middlewares, mdlws...)
	return shr
}

func (shr *socketHandlerRegister) ChannelCreate(handlers ...HandlerChannelCreate) {
	shr.evtName = EvtChannelCreate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) ChannelDelete(handlers ...HandlerChannelDelete) {
	shr.evtName = EvtChannelDelete
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) ChannelPinsUpdate(handlers ...HandlerChannelPinsUpdate) {
	shr.evtName = EvtChannelPinsUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) ChannelUpdate(handlers ...HandlerChannelUpdate) {
	shr.evtName = EvtChannelUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildBanAdd(handlers ...HandlerGuildBanAdd) {
	shr.evtName = EvtGuildBanAdd
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildBanRemove(handlers ...HandlerGuildBanRemove) {
	shr.evtName = EvtGuildBanRemove
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildCreate(handlers ...HandlerGuildCreate) {
	shr.evtName = EvtGuildCreate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildDelete(handlers ...HandlerGuildDelete) {
	shr.evtName = EvtGuildDelete
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildEmojisUpdate(handlers ...HandlerGuildEmojisUpdate) {
	shr.evtName = EvtGuildEmojisUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildIntegrationsUpdate(handlers ...HandlerGuildIntegrationsUpdate) {
	shr.evtName = EvtGuildIntegrationsUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildMemberAdd(handlers ...HandlerGuildMemberAdd) {
	shr.evtName = EvtGuildMemberAdd
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildMemberRemove(handlers ...HandlerGuildMemberRemove) {
	shr.evtName = EvtGuildMemberRemove
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildMemberUpdate(handlers ...HandlerGuildMemberUpdate) {
	shr.evtName = EvtGuildMemberUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildMembersChunk(handlers ...HandlerGuildMembersChunk) {
	shr.evtName = EvtGuildMembersChunk
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildRoleCreate(handlers ...HandlerGuildRoleCreate) {
	shr.evtName = EvtGuildRoleCreate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildRoleDelete(handlers ...HandlerGuildRoleDelete) {
	shr.evtName = EvtGuildRoleDelete
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildRoleUpdate(handlers ...HandlerGuildRoleUpdate) {
	shr.evtName = EvtGuildRoleUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) GuildUpdate(handlers ...HandlerGuildUpdate) {
	shr.evtName = EvtGuildUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) InviteCreate(handlers ...HandlerInviteCreate) {
	shr.evtName = EvtInviteCreate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) InviteDelete(handlers ...HandlerInviteDelete) {
	shr.evtName = EvtInviteDelete
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) MessageCreate(handlers ...HandlerMessageCreate) {
	shr.evtName = EvtMessageCreate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) MessageDelete(handlers ...HandlerMessageDelete) {
	shr.evtName = EvtMessageDelete
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) MessageDeleteBulk(handlers ...HandlerMessageDeleteBulk) {
	shr.evtName = EvtMessageDeleteBulk
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) MessageReactionAdd(handlers ...HandlerMessageReactionAdd) {
	shr.evtName = EvtMessageReactionAdd
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) MessageReactionRemove(handlers ...HandlerMessageReactionRemove) {
	shr.evtName = EvtMessageReactionRemove
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) MessageReactionRemoveAll(handlers ...HandlerMessageReactionRemoveAll) {
	shr.evtName = EvtMessageReactionRemoveAll
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) MessageReactionRemoveEmoji(handlers ...HandlerMessageReactionRemoveEmoji) {
	shr.evtName = EvtMessageReactionRemoveEmoji
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) MessageUpdate(handlers ...HandlerMessageUpdate) {
	shr.evtName = EvtMessageUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) PresenceUpdate(handlers ...HandlerPresenceUpdate) {
	shr.evtName = EvtPresenceUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) Ready(handlers ...HandlerReady) {
	shr.evtName = EvtReady
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) Resumed(handlers ...HandlerResumed) {
	shr.evtName = EvtResumed
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) TypingStart(handlers ...HandlerTypingStart) {
	shr.evtName = EvtTypingStart
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) UserUpdate(handlers ...HandlerUserUpdate) {
	shr.evtName = EvtUserUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) VoiceServerUpdate(handlers ...HandlerVoiceServerUpdate) {
	shr.evtName = EvtVoiceServerUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) VoiceStateUpdate(handlers ...HandlerVoiceStateUpdate) {
	shr.evtName = EvtVoiceStateUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}
func (shr *socketHandlerRegister) WebhooksUpdate(handlers ...HandlerWebhooksUpdate) {
	shr.evtName = EvtWebhooksUpdate
	for _, handler := range handlers {
		shr.handlers = append(shr.handlers, handler)
	}
	shr.build()
}

type SocketHandlerRegistrator interface {
	ChannelCreate(...HandlerChannelCreate)
	ChannelDelete(...HandlerChannelDelete)
	ChannelPinsUpdate(...HandlerChannelPinsUpdate)
	ChannelUpdate(...HandlerChannelUpdate)
	GuildBanAdd(...HandlerGuildBanAdd)
	GuildBanRemove(...HandlerGuildBanRemove)
	GuildCreate(...HandlerGuildCreate)
	GuildDelete(...HandlerGuildDelete)
	GuildEmojisUpdate(...HandlerGuildEmojisUpdate)
	GuildIntegrationsUpdate(...HandlerGuildIntegrationsUpdate)
	GuildMemberAdd(...HandlerGuildMemberAdd)
	GuildMemberRemove(...HandlerGuildMemberRemove)
	GuildMemberUpdate(...HandlerGuildMemberUpdate)
	GuildMembersChunk(...HandlerGuildMembersChunk)
	GuildRoleCreate(...HandlerGuildRoleCreate)
	GuildRoleDelete(...HandlerGuildRoleDelete)
	GuildRoleUpdate(...HandlerGuildRoleUpdate)
	GuildUpdate(...HandlerGuildUpdate)
	InviteCreate(...HandlerInviteCreate)
	InviteDelete(...HandlerInviteDelete)
	MessageCreate(...HandlerMessageCreate)
	MessageDelete(...HandlerMessageDelete)
	MessageDeleteBulk(...HandlerMessageDeleteBulk)
	MessageReactionAdd(...HandlerMessageReactionAdd)
	MessageReactionRemove(...HandlerMessageReactionRemove)
	MessageReactionRemoveAll(...HandlerMessageReactionRemoveAll)
	MessageReactionRemoveEmoji(...HandlerMessageReactionRemoveEmoji)
	MessageUpdate(...HandlerMessageUpdate)
	PresenceUpdate(...HandlerPresenceUpdate)
	Ready(...HandlerReady)
	Resumed(...HandlerResumed)
	TypingStart(...HandlerTypingStart)
	UserUpdate(...HandlerUserUpdate)
	VoiceServerUpdate(...HandlerVoiceServerUpdate)
	VoiceStateUpdate(...HandlerVoiceStateUpdate)
	WebhooksUpdate(...HandlerWebhooksUpdate)
	WithCtrl(HandlerCtrl) SocketHandlerRegistrator
	WithMdlw(...Middleware) SocketHandlerRegistrator
}
