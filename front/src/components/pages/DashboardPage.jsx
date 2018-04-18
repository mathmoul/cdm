import React from 'react'
import { connect } from 'react-redux'
import PropType from 'prop-types'
import ConfirmEmailMessage from '../messages/ConfirmEmailMessage'
import { Button } from 'semantic-ui-react'
import {Link} from 'react-router-dom'

import * as actions from '../../actions/auth'

const DashboardPage = ({isConfirmed, isAuthenticated, logout}) => (
  <div>
    {isAuthenticated ? (
      <Button onClick={() => logout()}>Logout</Button>) : (
      <div>
        <Link to='/login'>Login</Link> or <Link to='/signup'>Signup</Link>
      </div>
    )
    }
    {!isConfirmed && <ConfirmEmailMessage/>}
  </div>
)

DashboardPage.propTypes = {
  isConfirmed: PropType.bool.isRequired,
  isAuthenticated: PropType.bool.isRequired,
  logout: PropType.func.isRequired
}

function mapStateProps (state) {
  return {
    isConfirmed: !!state.user.confirmed,
    isAuthenticated: !!state.user.token
  }
}

export default connect(mapStateProps, {logout: actions.logout})(DashboardPage)