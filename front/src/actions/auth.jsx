import api from '../api'
import { USER_LOGGED_IN, USER_LOGGED_OUT } from '../types'

export const userLoggedIn = (user) => ({
  type: USER_LOGGED_IN,
  user
})

export const userLoggedOut = () => ({
  type: USER_LOGGED_OUT
})

export const login = (credentials) => (dispatch) =>
  api.user.login(credentials).then((user) => {
    localStorage.cdmJWT = user.token
    dispatch(userLoggedIn(user))
  })

export const logout = () => (dispatch) => {
  localStorage.removeItem('cdmJWT')
  dispatch(userLoggedOut())
}

export const confirm = (token) => (dispatch) =>
  api.user.confirm(token).then(user => {
    localStorage.cdmJwt = user.token
    dispatch(userLoggedIn(user))
  })
