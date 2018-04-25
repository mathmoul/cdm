import {normalize} from "normalizr"
import api from "../api"
import {TEAMS_FETCHED} from '../types'
import {teamsSchema } from '../schemas'

export const teamsFetched = (teams) => ({
    type: TEAMS_FETCHED,
    teams
});


//Can use normalize
export const fetchAllTeams = () => (dispatch) =>
    api.teams
        .fetchAllTeams()
        .then(teams => {
            dispatch(teamsFetched(normalize(teams, [teamsSchema])))
        })
