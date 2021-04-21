import React from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import { Delete } from "../Delete/Delete";
import "../global.css";
import { Home } from "../Home/Home";
import { Select } from "../Select/Select";
import { SelectAll } from "../SelectAll/SelectAll";
import { Upsert } from "../Upsert/Upsert";

export class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <Router>
        <div className="AppHeader">
          <h2>GDRP Demo: User App</h2>
          <nav>
            {" "}
            {"> "}
            <Link to={"/"} className="Link">
              Home
            </Link>{" "}
            {" | "}
            <Link to={"/select-all"} className="Link">
              Select All
            </Link>{" "}
            {" | "}
            <Link to={"/select"} className="Link">
              Select
            </Link>{" "}
            {" | "}
            <Link to={"/upsert"} className="Link">
              Upsert
            </Link>{" "}
            {" | "}
            <Link to={"/delete"} className="Link">
              Delete
            </Link>
          </nav>
        </div>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route path="/select-all" component={SelectAll} />
          <Route path="/select" component={Select} />
          <Route path="/upsert" component={Upsert} />
          <Route path="/delete" component={Delete} />
        </Switch>
      </Router>
    );
  }
}
