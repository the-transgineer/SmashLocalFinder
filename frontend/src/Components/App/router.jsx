import React from "react";
import {BrowserRouter, Route} from "react-router-dom";
import Regions from "../../Routes/Regions";
import Region from "../../Routes/Region";

const Router = () => (
    <BrowserRouter>
      <Route component={Regions} exact path="/regions"/>
      <Route component={Region} path="/regions/:region" />
    </BrowserRouter>
);

export default Router