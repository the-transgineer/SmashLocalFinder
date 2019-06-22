import React from "react";
import {BrowserRouter, Route} from "react-router-dom";
import Regions from "../../Routes/Regions";

const RegionRoute = ({match}) => {
    console.log(match.params.region);
    return <div></div>
}

const Router = () => (
    <BrowserRouter>
      <Route component={Regions} exact path="/regions"/>
      <Route component={RegionRoute} path="/regions/:region" />
    </BrowserRouter>
);

export default Router