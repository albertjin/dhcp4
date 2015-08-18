package dhcp4

// Get message from options. Return `MessageTypeNil` when this option is not found.
func (o Options) GetMessageType() MessageType {
    if t := o[OptionDHCPMessageType]; len(t) == 1 {
        return MessageType(t[0])
    }

    return MessageTypeNil
}
