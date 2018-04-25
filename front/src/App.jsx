import * as React from "react";
import {Route} from "react-router-dom";
import PropTypes from "prop-types";
import {connect} from "react-redux";

import HomePage from "./components/pages/HomePage";
import LoginPage from "./components/pages/LoginPage";
import DashboardPage from "./components/pages/DashboardPage";
import SignupPage from "./components/pages/SignupPage";
import ConfirmationPage from "./components/pages/ConfirmationPage";
import NewBookPage from "./components/pages/NewBookPage";

import GuestRoute from "./components/routes/GuestRoute";
import UserRoute from "./components/routes/UserRoute";
import AdminRoute from './components/routes/AdminRoute';

import ForgotPasswordPage from "./components/pages/ForgotPasswordPage";
import ResetPasswordPage from "./components/pages/ResetPasswordPage";
import TopNavigation from "./components/navigation/TopNavigation";
import NewMatchsPage from "./components/pages/NewMatchsPage";
import TeamsPage from "./components/pages/TeamsPage";


class App extends React.Component {
    state = {
        visible: false
    };

    render() {
        const {location, isAuthenticated} = this.props;
        return (
            <div className="ui container">
                {isAuthenticated && <TopNavigation/>}
                <Route location={location} path="/" exact component={HomePage}/>
                <Route
                    location={location}
                    path="/confirmation/:token"
                    exact
                    component={ConfirmationPage}
                />
                <GuestRoute
                    location={location}
                    path="/login"
                    exact
                    component={LoginPage}
                />
                <GuestRoute
                    location={location}
                    path="/signup"
                    exact
                    component={SignupPage}
                />
                <GuestRoute
                    location={location}
                    path="/forgot_password"
                    exact
                    component={ForgotPasswordPage}
                />
                <GuestRoute
                    location={location}
                    path="/reset_password/:token"
                    exact
                    component={ResetPasswordPage}
                />
                <UserRoute
                    location={location}
                    path="/dashboard"
                    exact
                    component={DashboardPage}
                />
                <UserRoute
                    location={location}
                    path="/books/new"
                    exact
                    component={NewBookPage}
                />
                <AdminRoute
                    location={location}
                    path="/matchs/new"
                    exact
                    component={NewMatchsPage}
                />
                <AdminRoute
                    location={location}
                    path="/teams"
                    exact
                    component={TeamsPage}
                />
            </div>
        )
    }

}

App.propTypes = {
    location: PropTypes.shape({
        pathname: PropTypes.string.isRequired
    }).isRequired,
    isAuthenticated: PropTypes.bool.isRequired
};

function mapStateToProps(state) {
    return {
        isAuthenticated: !!state.user.email
    };
}

export default connect(mapStateToProps)(App);
