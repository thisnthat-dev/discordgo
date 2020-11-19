package discordgo

// GuildMemberEditCustom updates a members roles and nick
// guildID  : The ID of a Guild.
// userID   : The ID of a User.
// nick     : The new nickname to set on the
// roles    : A list of role ID's to set on the member.
func (s *Session) GuildMemberEditCustom(guildID, userID, nick string, roles []string) (err error) {

	data := struct {
		Nick  string   `json:"nick,omitempty"`
		Roles []string `json:"roles,omitempty"`
	}{nick, roles}

	_, err = s.RequestWithBucketID("PATCH", EndpointGuildMember(guildID, userID), data, EndpointGuildMember(guildID, ""))
	if err != nil {
		return
	}

	return
}
