import React from "react";
import "../global.css";
import { Get, GetAPILink } from "../_Utils/ApiConnector";
import { UserRow } from "./UserRow";

export class SelectAll extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      users: [],
    };
  }

  componentDidMount() {
    GetAPILink().then((link) => {
      Get(link + "/users").then((data) => {
        this.setState({
          users: data,
        });
      });
    });
  }

  render() {
    return (
      <div className="AppContent">
        <h1>Select All Users</h1>
        <table className="UserTable">
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Secret Key</th>
            </tr>
          </thead>
          <tbody>
            {this.state.users !== undefined &&
              this.state.users.map((user, index) => {
                return (
                  <UserRow
                    key={index}
                    id={user.id}
                    name={user.name}
                    secretKey={user.secret_key}
                  />
                );
              })}
          </tbody>
        </table>
      </div>
    );
  }
}
