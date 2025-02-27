package relayer

import (
	"fmt"
	"strings"

	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v2/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
)

// ParseClientIDFromEvents parses events emitted from a MsgCreateClient and returns the
// client identifier.
func ParseClientIDFromEvents(events map[string]string) (string, error) {
	for k, v := range events {
		eventType, attrKey := splitEventKey(k)
		if eventType == clienttypes.EventTypeCreateClient {
			if attrKey == clienttypes.AttributeKeyClientID {
				return v, nil
			}
		}
	}
	return "", fmt.Errorf("client identifier event attribute not found")
}

// ParseConnectionIDFromEvents parses events emitted from a MsgConnectionOpenInit or
// MsgConnectionOpenTry and returns the connection identifier.
func ParseConnectionIDFromEvents(events map[string]string) (string, error) {
	for k, v := range events {
		eventType, attrKey := splitEventKey(k)
		if eventType == connectiontypes.EventTypeConnectionOpenInit ||
			eventType == connectiontypes.EventTypeConnectionOpenTry {
			if attrKey == connectiontypes.AttributeKeyConnectionID {
				return v, nil
			}
		}
	}
	return "", fmt.Errorf("connection identifier event attribute not found")
}

// ParseChannelIDFromEvents parses events emitted from a MsgChannelOpenInit or
// MsgChannelOpenTry and returns the channel identifier.
func ParseChannelIDFromEvents(events map[string]string) (string, error) {
	for k, v := range events {
		eventType, attrKey := splitEventKey(k)
		if eventType == channeltypes.EventTypeChannelOpenInit || eventType == channeltypes.EventTypeChannelOpenTry {
			if attrKey == channeltypes.AttributeKeyChannelID {
				return v, nil
			}
		}
	}
	return "", fmt.Errorf("channel identifier event attribute not found")
}

func splitEventKey(key string) (string, string) {
	stuff := strings.Split(key, ".")
	eventType := stuff[0]
	attrKey := stuff[1]
	return eventType, attrKey
}

// ParseClientIDFromEvents parses events emitted from a MsgCreateClient and returns the
// client identifier.
//func ParseClientIDFromEvents(events sdk.StringEvents) (string, error) {
//	for _, ev := range events {
//		if ev.Type == clienttypes.EventTypeCreateClient {
//			for _, attr := range ev.Attributes {
//				if attr.Key == clienttypes.AttributeKeyClientID {
//					return attr.Value, nil
//				}
//			}
//		}
//	}
//	return "", fmt.Errorf("client identifier event attribute not found")
//}

//
//func ParseConnectionIDFromEvents(events sdk.StringEvents) (string, error) {
//	for _, ev := range events {
//		if ev.Type == connectiontypes.EventTypeConnectionOpenInit ||
//			ev.Type == connectiontypes.EventTypeConnectionOpenTry {
//			for _, attr := range ev.Attributes {
//				if attr.Key == connectiontypes.AttributeKeyConnectionID {
//					return attr.Value, nil
//				}
//			}
//		}
//	}
//	return "", fmt.Errorf("connection identifier event attribute not found")
//}

//func ParseChannelIDFromEvents(events sdk.StringEvents) (string, error) {
//	for _, ev := range events {
//		if ev.Type == channeltypes.EventTypeChannelOpenInit || ev.Type == channeltypes.EventTypeChannelOpenTry {
//			for _, attr := range ev.Attributes {
//				if attr.Key == channeltypes.AttributeKeyChannelID {
//					return attr.Value, nil
//				}
//			}
//		}
//	}
//	return "", fmt.Errorf("channel identifier event attribute not found")
//}
