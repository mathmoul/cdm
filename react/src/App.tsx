import * as React from "react";
import {Route} from "react-router-dom"

import HomePage from "./components/HomePage";
import LoginPage from "./components/LoginPage";

const App = function() {
    return <div>
        <Route path="/" exact component={HomePage} />
        <Route path="/login" exact component={LoginPage} />
   </div>
}

export default App;