import React from 'react'
import {Dropdown, Image, Menu, Segment} from 'semantic-ui-react'
import {Link} from 'react-router-dom'
import PropTypes from 'prop-types'
import gravatarUrl from 'gravatar-url'
import * as actions from '../../actions/auth'

import {connect} from 'react-redux'

const TopNavigation = ({user, logout}) => (
    <Segment as={Menu} inverted>
        <Menu.Item as={Link} to='/dashboard'>
            Dashboard
        </Menu.Item>
        <Menu.Item as={Link} to='/league/new'>
            Nouvelle ligue
        </Menu.Item>
        {user.admin && (
            <Menu.Item as={Link} to='/matchs/new'>Ajouter des matchs</Menu.Item>
        )}
        {user.admin && (
            <Menu.Item as={Link} to='/teams'>Ajouter des équipes</Menu.Item>
        )}
        <Menu.Menu position="right">
            <Dropdown trigger={<Image avatar src={gravatarUrl(user.email)}/>}>
                <Dropdown.Menu>
                    <Dropdown.Item>{user.email}</Dropdown.Item>
                    <Dropdown.Item onClick={() => logout()}>
                        Logout
                    </Dropdown.Item>
                </Dropdown.Menu>
            </Dropdown>
        </Menu.Menu>
    </Segment>
);

TopNavigation.propTypes = {
    user: PropTypes.shape({
        email: PropTypes.string.isRequired,
        admin: PropTypes.bool
    }).isRequired,
    logout: PropTypes.func.isRequired
};

function mapStateToProps(state) {
    return {
        user: {
            email: state.user.email,
            admin: !!state.user.admin
        }
    }
}

export default connect(mapStateToProps, {logout: actions.logout})(TopNavigation)