import React from 'react'
import {connect} from 'react-redux'
import PropTypes from 'prop-types'

import {Button, Segment} from 'semantic-ui-react'
import {fetchAllTeams} from "../../actions/teams";

import TeamForm from "../forms/TeamForm"

class TeamsPage extends React.Component {
    state = {
        loading: false,
        addTeam: false,
        teams: []
    };

    componentDidMount() {
        this.onInit(this.props)
    }

    onInit = (props) => {
        this.setState({loading: true});
        props.fetchAllTeams()
            .then(teams => this.setState({teams, loading: false}))
            .catch(error => console.log(error))
    }

    onTeamSubmit = (data) => {

    }

    render() {
        console.log(this.state)
        const {loading, teams, addTeam} = this.state;
        return (
            <div>
                <Button onClick={() => this.setState({addTeam: !addTeam})}>Ajouter une equipe</Button>
                {addTeam && <Segment><TeamForm onSubmit={this.onTeamSubmit}/></Segment>}
                <Segment loading={loading}>
                    {
                        !teams ?
                            <h1>Pas encore d'equipes</h1> :
                            <div>{teams.map((team, index) => <Segment index={index}>{team}</Segment>)}</div>
                    }

                </Segment>

            </div>
        )
    }
}

TeamsPage.propTypes = {
    isConfirmed: PropTypes.bool.isRequired,
    isAdmin: PropTypes.bool.isRequired,
    fetchAllTeams: PropTypes.func.isRequired,
};

function mapStateToProps(state) {
    return {
        isConfirmed: !!state.user.confirmed,
        isAdmin: !!state.user.admin
    }
}

export default connect(mapStateToProps, {fetchAllTeams})(TeamsPage)