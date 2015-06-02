package api

/*

func HermesSocket(c *serve.Conn, db data.DB, man autonomous.Manager) {
	socket, err := sock.DefaultUpgrader.Upgrade(c.ResponseWriter(), c.Request())
	if err != nil {
		log.Print(err)
	}

	protocol := c.Request().Header.Get(sock.WebSocketProtocolHeader)
	pTokens := strings.Split(protocol, "-")
	userID, publicKey := pTokens[0], pTokens[1]

	u := &models.User{}
	id, err := db.ParseID(userID)
	u.SetID(id)
	db.PopulateByID(u)

	var matchingKey string
	for _, key := range u.PublicKeys {
		if key == publicKey {
			matchingKey = key
		}
	}

	if matchingKey == "" {
		return
	}
}
*/
