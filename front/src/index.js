import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter } from 'react-router-dom'

import { createStore, applyMiddleware } from 'redux'
import { Provider } from 'react-redux'
import thunk from 'redux-thunk'
import rootReducer from './rootReducer'

import { composeWithDevTools } from 'redux-devtools-extension'

import App from './App'

import 'semantic-ui-css/semantic.min.css'

import registerServiceWorker from './registerServiceWorker'
import { userLoggedIn } from './actions/auth'

const store = createStore(
  rootReducer,
  composeWithDevTools(applyMiddleware(thunk)),
)

if (localStorage.cdmJWT) {
  const user = { token: localStorage.cdmJWT}
  store.dispatch(userLoggedIn(user))
}

ReactDOM.render(
  <BrowserRouter>
    <Provider store={store}>
      <App />
    </Provider>
  </BrowserRouter>
  , document.getElementById('app'))
registerServiceWorker()
