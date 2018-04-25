import React from 'react'
import {Redirect, Route} from 'react-router-dom'
import PropTypes from "prop-types"

import {connect} from "react-redux"

const AdminRoute = ({isAuthenticated, isAdmin, component: Component, ...rest}) => (
    <Route
        {...rest}
        render={props =>
            (isAuthenticated && isAdmin) ? <Component {...props} /> : <Redirect to="/"/>
        }
    />
);

AdminRoute.propTypes = {
    component: PropTypes.func.isRequired,
    isAuthenticated: PropTypes.bool.isRequired,
    isAdmin: PropTypes.bool.isRequired
};

function mapStateToProps(state) {
    return {
        isAuthenticated: !!state.user.email,
        isAdmin: !!state.user.admin
    }
};


export default connect(mapStateToProps)(AdminRoute)