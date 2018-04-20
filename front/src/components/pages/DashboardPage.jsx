import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import ConfirmEmailMessage from '../messages/ConfirmEmailMessage'

import AddBookCtA from '../ctas/AddBookCtA'
import { allBooksSelector } from '../../reducers/books'

const DashboardPage = ({isConfirmed, books}) => (
  <div>
    {!isConfirmed && <ConfirmEmailMessage/>}

    {books.length === 0 && <AddBookCtA/>}
  </div>
)

DashboardPage.propTypes = {
  isConfirmed: PropTypes.bool.isRequired,
  books: PropTypes.arrayOf(PropTypes.shape({
    title: PropTypes.string.isRequired
  }).isRequired).isRequired
}

function mapStateProps (state) {
  return {
    isConfirmed: !!state.user.confirmed,
    books: allBooksSelector(state)
  }
}

export default connect(mapStateProps)(DashboardPage)