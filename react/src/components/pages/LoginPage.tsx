import * as React from "react";

import LoginForm from "../forms/LoginForm";

class LoginPage extends React.Component {
    public submit = (data) => {
        window.console.log(data);
    };

    public render = () =>
        (<div>
            <h1>Page de Login</ h1>
            <LoginForm submit={this.submit}/>
        </div>)

}

export default LoginPage;
