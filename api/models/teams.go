package models

//Team describes a sports team
type Team struct {
	ID   int
	Name string
	City string
}

//AllTeams returns a list of all teams in DB
func AllTeams() ([]*Team, error) {
	rows, err := db.Query("SELECT * FROM public.teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := make([]*Team, 0)
	for rows.Next() {
		team := new(Team)
		err := rows.Scan(&team.ID, &team.Name, &team.City)
		if err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

//SingleTeam returns a single team by TeamID
func SingleTeam(ID int) (*Team, error) {
	rows, err := db.Query("SELECT * FROM public.teams WHERE ID = $1", ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	team := new(Team)
	for rows.Next() {

		err := rows.Scan(&team.ID, &team.Name, &team.City)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return team, nil
}
