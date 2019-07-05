import React from "react";
import {BrowserRouter, Route} from "react-router-dom";
import Regions from "../../Routes/Regions";
import Region from "../../Routes/Region";
import Admin from "../../Routes/Admin";
import Create from "../../Routes/Create";

const Router = () => (
    <BrowserRouter>
      <Route component={Regions} exact path="/regions"/>
      <Route component={Region} path="/regions/:region" />
      <Route component={Admin} path="/admin/"/>
      <Route component={Create} path="/admin/create"/>
    </BrowserRouter>
);

export default Router