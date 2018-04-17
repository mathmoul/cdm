import React from 'react'

import PropTypes from 'prop-types'

import { connect } from 'react-redux'
import { login } from '../../actions/auth'
import LoginForm from '../forms/LoginForm'

class LoginPage extends React.Component {
  render () {
    return (<div>
      <h1>Page de Login</h1>
      <LoginForm submit={this.submit} />
    </div>)
  }

  submit = (data) => this.props.login(data).then(() => {
    this.props.history.push('/dashboard')
  })
}

LoginPage.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func.isRequired,
  }).isRequired,
  login: PropTypes.func.isRequired,
}

export default connect(null, {login})(LoginPage)
