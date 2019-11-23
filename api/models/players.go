package models

//Player describes a tracked player
type Player struct {
	ID        int
	FirstName string
	LastName  string
	Team      int
}

//Team describes a sports team
type Team struct {
	ID   int
	Name string
	City string
}

func AllPlayers() ([]*Player, error) {
	rows, err := db.Query("SELECT * FROM public.players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	players := make([]*Player, 0)
	for rows.Next() {
		player := new(Player)
		err := rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Team)
		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
}

func SinglePlayer(ID int) (*Player, error) {
	rows, err := db.Query("SELECT * FROM public.players WHERE ID = $1", ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	player := new(Player)
	for rows.Next() {

		err := rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Team)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return player, nil
}
