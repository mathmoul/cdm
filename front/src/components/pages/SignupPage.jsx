import React from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { signup } from '../../actions/users'

import SignupForm from '../forms/SignupForm'

class SignupPage extends React.Component {

  submit = data =>
    this.props.signup(data).then(() => this.props.history.push('/dashboard'))

  render () {
    return <div>
      <h1>Signup</h1>
      <SignupForm submit={this.submit}/>
    </div>
  }
}

SignupPage.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func.isRequired
  }).isRequired,
  signup: PropTypes.func.isRequired
}

export default connect(null, {signup})(SignupPage)
