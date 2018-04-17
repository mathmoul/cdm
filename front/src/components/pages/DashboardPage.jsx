import React from 'react'
import { connect } from 'react-redux'
import PropType from 'prop-types'
import ConfirmEmailMessage from '../messages/ConfirmEmailMessage'

const DashboardPage = ({isConfirmed}) => (
  <div>
    {!isConfirmed && <ConfirmEmailMessage/>}
  </div>
)

DashboardPage.propTypes = {
  isConfirmed: PropType.bool.isRequired
}

function mapStateProps (state) {
  return {
    isConfirmed: !!state.user.confirmed
  }
}

export default connect(mapStateProps)(DashboardPage)