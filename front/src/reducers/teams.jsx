import {TEAMS_FETCHED} from "../types";

export default function teams(state = {}, action = {}) {
    switch (action.type) {
        case TEAMS_FETCHED:
            return ({...state, ...action.entities.teams});
        default:
            return state
    }
}