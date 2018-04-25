import React from 'react'
import {connect} from 'react-redux'
import PropTypes from 'prop-types'
import ConfirmEmailMessage from '../messages/ConfirmEmailMessage'
import {allBooksSelector} from '../../reducers/books'

import {Button, Grid, Segment} from 'semantic-ui-react'
import {Link} from "react-router-dom"
import MatchsListCtA from "../ctas/MatchsListCtA";

class DashboardPage extends React.Component {
    state = {
        loading: false,
        leagues: [],
    };

    render() {
        const {isConfirmed} = this.props;
        const {leagues} = this.state;
        return (
            <div>
                {!isConfirmed && <ConfirmEmailMessage/>}

                <Grid columns={2}>
                    <Grid.Column>
                        <Segment><MatchsListCtA/></Segment>
                    </Grid.Column>
                    <Grid.Column>
                        {leagues.length === 0 && (
                            <Segment>
                                <h1>Pas de ligues actuellement</h1>
                                <Button primary as={Link} to="/league/new">Ajouter une ligue</Button>
                            </Segment>
                        )}
                    </Grid.Column>
                </Grid>
            </div>
        )
    }
}

DashboardPage.propTypes = {
    isConfirmed: PropTypes.bool.isRequired,
    books: PropTypes.arrayOf(PropTypes.shape({
        title: PropTypes.string.isRequired
    }).isRequired).isRequired
}

function mapStateProps(state) {
    return {
        isConfirmed: !!state.user.confirmed,
        books: allBooksSelector(state)
    }
}

export default connect(mapStateProps)(DashboardPage)